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

# ORM 框架
数据库	面向对象的编程语言
表(table)	类(class/struct)
记录(record, row)	对象 (object)
字段(field, column)	对象属性(attribute)

reflect.ValueOf() 获取指针对应的反射值。
reflect.Indirect() 获取指针指向的对象的反射值。
(reflect.Type).Name() 返回类名(字符串)。
(reflect.Type).Field(i) 获取第 i 个成员变量。

## Go 语言中使用比较广泛 ORM 框架是 gorm 和 xorm
https://github.com/jinzhu/gorm
https://github.com/go-xorm/xorm

### 第一天：database/sql 基础
安装数据库 https://geektutu.com/post/cheat-sheet-sqlite.html
```
brew install sqlite3
sqlite3 gmyst.db
CREATE TABLE User(Name text, Age integer);
INSERT INTO User(Name, Age) VALUES ("Tom", 18), ("Jack", 25);
.head on
# 查找 `Age > 20` 的记录；
SELECT * FROM User WHERE Age > 20;
# 统计记录个数。
SELECT COUNT(*) FROM User;
# .table 查看当前数据库中所有的表(table)，执行 .schema <table> 查看建表的 SQL 语句
.table
.schema User
```
### 第二天：对象表结构映射
### 第三天：记录新增和查询
### 第四天：链式操作与更新删除
### 第五天：实现钩子(Hooks)
### 第六天：支持事务(Transaction)
### 第七天：数据库迁移(Migrate)

# GmystRPC
### 第一天 - 服务端与消息编码
消息的序列化与反序列化

第二天 - 支持并发与异步的客户端
第三天 - 服务注册(service register)
第四天 - 超时处理(timeout)
第五天 - 支持HTTP协议
第六天 - 负载均衡(load balance)
第七天 - 服务发现与注册中心(registry)