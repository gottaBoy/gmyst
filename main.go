package main

/*
curl "http://localhost:9999"
Hello gmyst
curl "http://localhost:9999/panic"
{"message":"Internal Server Error"}

2021/06/14 16:11:32 Route  GET - /
2021/06/14 16:11:32 Route  GET - /panic
2021/06/14 16:11:53 [200] / in 23.791µs
2021/06/14 16:11:54 [404] /favicon.ico in 13.848µs
2021/06/14 16:13:17 [200] / in 4.781µs
2021/06/14 16:13:51 runtime error: index out of range [100] with length 1
Traceback:
        /usr/local/Cellar/go/1.16.3/libexec/src/runtime/panic.go:965
        /usr/local/Cellar/go/1.16.3/libexec/src/runtime/panic.go:88
        /Users/minyi/go/src/gmyst/main.go:45
        /Users/minyi/go/src/gmyst/gmyst/context.go:41
        /Users/minyi/go/src/gmyst/gmyst/recovery.go:37
        /Users/minyi/go/src/gmyst/gmyst/context.go:41
        /Users/minyi/go/src/gmyst/gmyst/logger.go:15
        /Users/minyi/go/src/gmyst/gmyst/context.go:41
        /Users/minyi/go/src/gmyst/gmyst/router.go:99
        /Users/minyi/go/src/gmyst/gmyst/gmyst.go:130
        /usr/local/Cellar/go/1.16.3/libexec/src/net/http/server.go:2888
        /usr/local/Cellar/go/1.16.3/libexec/src/net/http/server.go:1953
        /usr/local/Cellar/go/1.16.3/libexec/src/runtime/asm_amd64.s:1372

2021/06/14 16:13:51 [500] /panic in 186.146µs

*/

import (
	"fmt"
	"gmyst/gmystcache"
	"log"
	"net/http"
)

//func main() {
//	r := gmyst.Default()
//	r.GET("/", func(c *gmyst.Context) {
//		c.String(http.StatusOK, "Hello gmyst\n")
//	})
//	// index out of range for testing Recovery()
//	r.GET("/panic", func(c *gmyst.Context) {
//		names := []string{"gmyst"}
//		c.String(http.StatusOK, names[100])
//	})
//
//	r.Run(":9999")
//}

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

// curl http://localhost:9999/_gmystcache/scores/Tom
func main() {
	gmystcache.NewGroup("scores", 2<<10, gmystcache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))

	addr := "localhost:9999"
	peers := gmystcache.NewHTTPPool(addr)
	log.Println("gmystcache is running at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}
