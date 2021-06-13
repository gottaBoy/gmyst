package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)
/**
 * find . -name "*_test.go" | xargs cat | wc -l
 * find . -name "*.go" | xargs cat | wc -l
 */
func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, gmyst")
	})
	// 解析路径参数
	// 匹配 /user/gmyst
	r.GET("/user/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK, "Hello %s", name)
	})
	// 获取Query参数
	// 匹配users?name=xxx&role=xxx，role可选
	// curl "http://localhost:9999/users?name=Tom&role=student"
	r.GET("/users", func(context *gin.Context) {
		name := context.Query("name")
		role := context.Query("role")
		context.String(http.StatusOK, "%s is a %s", name, role)
	})
	// 获取POST参数
	// curl http://localhost:9999/form  -X POST -d 'username=gmyst&password=1234'
	r.POST("/form", func(context *gin.Context) {
		username := context.PostForm("username")
		password := context.DefaultPostForm("password", "123456")
		context.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})
	// Query和POST混合参数
	// curl "http://localhost:9999/posts?id=9876&page=7"  -X POST -d 'username=gmyst&password=1234'
	r.POST("/posts", func(context *gin.Context) {
		id := context.Query("id")
		page := context.DefaultQuery("page", "0")
		username := context.PostForm("username")
		password := context.DefaultPostForm("password", "123456")
		context.JSON(http.StatusOK, gin.H{
			"id":       id,
			"page":     page,
			"username": username,
			"password": password,
		})
	})
	// Map参数(字典参数)
	// curl -g "http://localhost:9999/post?ids[Jack]=001&ids[Tom]=002" -X POST -d 'names[a]=Sam&names[b]=David'
	r.POST("/post", func(context *gin.Context) {
		ids := context.QueryMap("ids")
		names := context.PostFormMap("names")
		context.JSON(http.StatusOK, gin.H{
			"ids":       ids,
			"names":     names,
		})
	})

	// 重定向(Redirect)
	// curl -i http://localhost:9999/redirect
	// curl "http://localhost:9999/goindex"
	r.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/index")
	})
	r.GET("/goindex", func(c *gin.Context) {
		c.Request.URL.Path = "/"
		r.HandleContext(c)
	})

	// 分组路由(Grouping Routes)
	// group routes 分组路由
	defaultHandler := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"path": c.FullPath(),
		})
	}
	// group: v1
	// curl http://localhost:9999/v1/posts
	v1 := r.Group("/v1")
	{
		v1.GET("/posts", defaultHandler)
		v1.GET("/series", defaultHandler)
	}
	// group: v2
	// curl http://localhost:9999/v2/posts
	v2 := r.Group("/v2")
	{
		v2.GET("/posts", defaultHandler)
		v2.GET("/series", defaultHandler)
	}

	// 上传文件
	// 单个文件
	r.POST("/upload1", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		// c.SaveUploadedFile(file, dst)
		c.String(http.StatusOK, "%s uploaded!", file.Filename)
	})
	// 多个文件
	r.POST("/upload2", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)
			// c.SaveUploadedFile(file, dst)
		}
		c.String(http.StatusOK, "%d files uploaded!", len(files))
	})

	// HTML模板(Template)
	type student struct {
		Name string
		Age  int8
	}
	r.LoadHTMLGlob("templates/*")
	stu1 := &student{Name: "gmyst", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}
	// curl http://localhost:9999/arr
	r.GET("/arr", func(c *gin.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gin.H{
			"title":  "Gin",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	//中间件(Middleware)
	// 作用于全局
	//r.Use(gin.Logger())
	r.Use(Logger())
	r.Use(gin.Recovery())

	// 作用于单个路由
	//r.GET("/benchmark", MyBenchLogger(), benchEndpoint)

	// 作用于某个组
	//authorized := r.Group("/")
	//authorized.Use(AuthRequired())
	//{
	//	authorized.POST("/login", loginEndpoint)
	//	authorized.POST("/submit", submitEndpoint)
	//}

	// 热加载调试 Hot Reload
	//github.com/codegangsta/gin
	//github.com/pilu/fresh

	r.Run(":9999") // listen and serve on 0.0.0.0:8080
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// 给Context实例设置一个值
		c.Set("gmyst", "1111")
		// 请求前
		c.Next()
		// 请求后
		latency := time.Since(t)
		log.Print(latency)
	}
}