# gmyst

### 缓存雪崩、缓存击穿与缓存穿透
- 缓存雪崩：缓存在同一时刻全部失效，造成瞬时DB请求量大、压力骤增，引起雪崩。缓存雪崩通常因为缓存服务器宕机、缓存的 key 设置了相同的过期时间等引起。
- 缓存击穿：一个存在的key，在缓存过期的一刻，同时有大量的请求，这些请求都会击穿到 DB ，造成瞬时DB请求量大、压力骤增。
- 缓存穿透：查询一个不存在的数据，因为不存在则不会写到缓存中，所以每次都会去请求 DB，如果瞬间流量过大，穿透到 DB，导致宕机。

使用 singleflight 防止缓存击穿，实现与测试

### protobuf
protobuf 即 Protocol Buffers，Google 开发的一种数据描述语言，是一种轻便高效的结构化数据存储格式，与语言、平台无关，可扩展可序列化。protobuf 以二进制方式存储，占用空间小。

### protoc
https://geektutu.com/post/quick-go-protobuf.html
#### 下载安装包
$ wget https://github.com/protocolbuffers/protobuf/releases/download/v3.11.2/protoc-3.11.2-linux-x86_64.zip
#### 解压到 /usr/local 目录下
$ sudo 7z x protoc-3.11.2-linux-x86_64.zip -o/usr/local

### protoc-gen-go
brew install protobuf
go get -u github.com/golang/protobuf/protoc-gen-go
protoc --proto_path=. --go_out=. *.proto
protoc  -I=. --go_out=plugins=grpc:. --go_opt=paths=source_relative gmystcache/gmystcachepb/*.proto
protoc --proto_path=. --go_out=plugins=grpc,paths=source_relative:. *.proto

https://cloud.tencent.com/developer/article/1768249
https://zhuanlan.zhihu.com/p/368079374
https://jishuin.proginn.com/p/763bfbd4ef61
