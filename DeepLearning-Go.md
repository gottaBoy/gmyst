# 深入学习GO语言
### 数组和切片
数组类型的值（以下简称数组）的长度是固定的，而切片类型的值（以下简称切片）是可变长的。
```go
s1 := make([]int, 5, 8)
```
怎样正确估算切片的长度和容量？
问题 1：怎样估算切片容量的增长？
问题 2：切片的底层数组什么时候会被替换？
如果有多个切片指向了同一个底层数组，那么你认为应该注意些什么？
怎样沿用“扩容”的思想对切片进行“缩容”？
container/list
List和Element
1. 问题：为什么链表可以做到开箱即用？
2. 问题：Ring与List的区别在哪儿？
3. container/ring包中的循环链表的适用场景都有哪些？
4. 你使用过container/heap包中的堆吗？它的适用场景又有哪些呢？

### 字典的操作和约束
```go
maps := map[string]int{}
```
#### 通道
```go
ch1 := make(chan int, 3)
```
问题 1：发送操作和接收操作在什么时候可能被长时间的阻塞？
- 缓冲通道
- 非缓冲通道
问题 2：发送操作和接收操作在什么时候会引发 panic？
- 通道的长度代表着什么？它在什么时候会通道的容量相同？
- 元素值在经过通道传递时会被复制，那么这个复制是浅表复制还是深层复制呢？
chan<- 发送通道
<-chan 接收通道
```go
select {
case： express 
default: 
```
#### 函数
```go
type Printer func(contents string) (n int, err error)
```
高阶函数可以满足下面的两个条件：
1. 接受其他的函数作为参数传入；
2. 把其他的函数作为结果返回。
问题 1：如何实现闭包？
问题 2：传入函数的那些参数值后来怎么样了？

#### 结构体类型
```go
// AnimalCategory 代表动物分类学中的基本分类法。
type AnimalCategory struct {
kingdom string // 界。
phylum string // 门。
class  string // 纲。
order  string // 目。
family string // 科。
genus  string // 属。
species string // 种。
}

func (ac AnimalCategory) String() string {
return fmt.Sprintf("%s%s%s%s%s%s%s",
ac.kingdom, ac.phylum, ac.class, ac.order,
ac.family, ac.genus, ac.species)
}
```
#### 接口类型
```go
type Pet interface {
  SetName(name string)
  Name() string
  Category() string
}
```
iface的实例会包含两个指针，一个是指向类型信息的指针，另一个是指向动态值的指针。
这里的类型信息是由另一个专用数据结构的实例承载的，
其中包含了动态值的类型，以及使它实现了接口的方法和调用它们的途径，
问题 1：接口变量的值在什么情况下才真正为nil？
问题 2：怎样实现接口之间的组合？

#### 指针类型
unsafe.Pointer可以表示任何指向可寻址的值的指针，同时它也是前面提到的指针值和uintptr值之间的桥梁。也就是说，通过它，我们可以在这两种值之上进行双向的转换。这里有一个很关键的词——可寻址的（addressable）。
不可变的
临时结果
不安全的

问题 1：不可寻址的值在使用上有哪些限制？
问题 2：怎样通过unsafe.Pointer操纵可寻址的值？
```go
dog := Dog{"little pig"}
dogP := &dog
dogPtr := uintptr(unsafe.Pointer(dogP))
```

#### 进程与线程
Go 语言的运行时（runtime）系统会帮助我们自动地创建和销毁系统级的线程。
这里的系统级线程指的就是我们刚刚说过的操作系统提供的线程。
而对应的用户级线程指的是架设在系统级线程之上的，由用户（或者说我们编写的程序）完全控制的代码执行流程。
用户级线程的创建、销毁、调度、状态变更以及其中的代码和数据都完全需要我们的程序自己去实现和处理。

##### GPM
Go 语言不但有着独特的并发编程模型，以及用户级线程 goroutine，还拥有强大的用于调度 goroutine、对接系统级线程的调度器。
Go语言运行时系统的重要组成部分，它主要负责统筹调配 Go 并发编程模型中的三个主要元素，即：G（goroutine 的缩写）、P（processor 的缩写）和 M（machine 的缩写）

问题 1：怎样才能让主 goroutine 等待其他 goroutine？

```go
func main() {
	num := 10
	sign := make(chan struct{}, num)

	for i := 0; i < num; i++ {
		go func() {
			fmt.Println(i)
			sign <- struct{}{}
		}()
	}

	// 办法1。
	//time.Sleep(time.Millisecond * 500)

	// 办法2。
	for j := 0; j < num; j++ {
		<-sign
	}
}
```
问题 2：怎样让我们启用的多个 goroutine 按照既定的顺序运行？
```go
func main() {
	var count uint32
	trigger := func(i uint32, fn func()) {
		for {
			if n := atomic.LoadUint32(&count); n == i {
				fn()
				atomic.AddUint32(&count, 1)
				break
			}
			time.Sleep(time.Nanosecond)
		}
	}
	for i := uint32(0); i < 10; i++ {
		go func(i uint32) {
			fn := func() {
				fmt.Println(i)
			}
			trigger(i, fn)
		}(i)
	}
	trigger(10, func() {})
}
```
#### Error
net.Error接口除了拥有error接口的Error方法之外，还有两个自己声明的方法：Timeout和Temporary。
net包中有很多错误类型都实现了net.Error接口，
比如：*net.OpError；*net.AddrError；net.UnknownNetworkError等等。

#### panic函数、recover函数以及defer语句
从 panic 被引发到程序终止运行的大致过程是什么？
问题 1：怎样让 panic 包含一个值，以及应该让它包含什么样的值？
问题 2：怎样施加应对 panic 的保护措施，从而避免程序崩溃？
问题 3：如果一个函数中有多条defer语句，那么那几个defer函数调用的执行顺序是怎样的？
```go
func main() {
	fmt.Println("Enter function main.")
	defer func() {
		fmt.Println("Enter defer function.")
		// recover函数的正确用法。
		if p := recover(); p != nil {
			fmt.Printf("panic: %s\n", p)
		}
		fmt.Println("Exit defer function.")
	}()
	// recover函数的错误用法。
	fmt.Printf("no panic: %v\n", recover())
	// 引发panic。
	panic(errors.New("something wrong"))
	// recover函数的错误用法。
	p := recover()
	fmt.Printf("panic: %s\n", p)

	fmt.Println("Exit function main.")
}
```

#### 测试用例
单元测试、API 测试、集成测试、灰度测试
Go 程序编写三类测试，
即：功能测试（test）、基准测试（benchmark，也称性能测试），以及示例测试（example）
*testing.T类型
*testing.B类型
Example
go test file
问题 1：怎样解释功能测试的测试结果？
t.Log方法和t.Fail方法
t.Fatal方法和t.Fatalf方法
问题 2：怎样解释性能测试的测试结果？
go test -bench=. -run=^$

性能测试函数的执行次数 = `-cpu`标记的值中正整数的个数 x `-count`标记的值 x 探索式执行中测试函数的实际执行次数
功能测试函数的执行次数 = `-cpu`标记的值中正整数的个数 x `-count`标记的值

问题 1：-parallel标记的作用是什么？
go test -v
go test -count=2 -v
b.RunParallel方法、b.SetParallelism方法和-cpu标记
问题 2：性能测试函数中的计时器是做什么用的？
testing.B类型有这么几个指针方法：StartTimer、StopTimer和ResetTimer
-benchmem标记和-benchtime标记的作用分别是什么？
怎样在测试的时候开启测试覆盖度分析？如果开启，会有什么副作用吗？

#### 竞态条件、临界区与同步工具
用通讯的方式共享数据
竞态条件（race condition）
同步的用途有两个，一个是避免多个线程在同一时刻操作同一个数据块，另一个是协调多个线程，以避免它们在同一时刻执行同一个代码块。
存储资源、计算资源、I/O 资源、网络资源
竞态条件、临界区与同步工具
```go
mu.Lock()
_, err := writer.Write([]byte(data))
if err != nil {
 log.Printf("error: %s [%d]", err, id)
}
mu.Unlock()
```
问题 1：读写锁与互斥锁有哪些异同？
读写锁是读 / 写互斥锁的简称。
在 Go 语言中，读写锁由sync.RWMutex类型的值代表。与sync.Mutex类型一样，这个类型也是开箱即用的。
读写锁是把对共享资源的“读操作”和“写操作”区别对待了。它可以对这两种操作施加不同程度的保护。
换句话说，相比于互斥锁，读写锁可以实现更加细腻的访问控制。一个读写锁中实际上包含了两个锁，
即：读锁和写锁。sync.RWMutex类型中的Lock方法和Unlock方法分别用于对写锁进行锁定和解锁，
而它的RLock方法和RUnlock方法则分别用于对读锁进行锁定和解锁。另外，对于同一个读写锁来说有如下规则。
在写锁已被锁定的情况下再试图锁定写锁，会阻塞当前的 goroutine。
在写锁已被锁定的情况下试图锁定读锁，也会阻塞当前的 goroutine。
在读锁已被锁定的情况下试图锁定写锁，同样会阻塞当前的 goroutine。
在读锁已被锁定的情况下再试图锁定读锁，并不会阻塞当前的 goroutine。

#### 条件变量sync.Cond
条件变量怎样与互斥锁配合使用？
条件变量的初始化离不开互斥锁，并且它的方法有的也是基于互斥锁的。
条件变量提供的方法有三个：等待通知（wait）、单发通知（signal）和广播通知（broadcast）。


### unicode与字符编码
#### ASCII
#### Unicode
- UTF-8 [1, 4]
#### 字符串
- 字符类型 `type rune = int32`
```go
func main() {
	str := "Go爱好者"
	fmt.Printf("The string: %q\n", str)
	fmt.Printf("  => runes(char): %q\n", []rune(str))
	fmt.Printf("  => runes(hex): %x\n", []rune(str))
	fmt.Printf("  => bytes(hex): [% x]\n", []byte(str))
}
```
问题 1：Go 语言是用嵌入字段实现了继承吗？
问题 2：值方法和指针方法都是什么意思，有什么区别？

- 输出结果
```go
The string: "Go爱好者"
  => runes(char): ['G' 'o' '爱' '好' '者']
  => runes(hex): [47 6f 7231 597d 8005]
  => bytes(hex): [47 6f e7 88 b1 e5 a5 bd e8 80 85]
```  

- 字符串索引
```go
func main() {
	str := "Go爱好者"
	for i, c := range str {
		fmt.Printf("%d: %q [% x]\n", i, c, []byte(string(c)))
	}
}
```
- 输出结果
```go
0: 'G' [47]
1: 'o' [6f]
2: '爱' [e7 88 b1]
5: '好' [e5 a5 bd]
8: '者' [e8 80 85]
```
### strings包与字符串操作
- 问题 1：strings.Builder类型在使用上有约束吗
#### Builder
- 结构体
```go
type Builder struct {
	addr *Builder // of receiver, to detect copies by value
	buf  []byte
}
```
- strings.Builder 使用
```go
var builder strings.Builder
builder.WriteString()
builder.WriteByte()
builder.Write()
builder.Grow()
builder.Reset()
```

#### string.Reader
```go
type Reader struct {
	s        string
	i        int64 // current reading index
	prevRune int   // index of previous rune; or < 0
}
```
- 方法
```go
func (r *Reader) Len() int
func (r *Reader) Size() int64
func (r *Reader) Read(b []byte) (n int, err error)
func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)
func (r *Reader) ReadByte() (byte, error)
func (r *Reader) UnreadByte() error
func (r *Reader) ReadRune() (ch rune, size int, err error)
func (r *Reader) UnreadRune() error
func (r *Reader) Seek(offset int64, whence int) (int64, error)
func (r *Reader) WriteTo(w io.Writer) (n int64, err error)
func (r *Reader) Reset(s string)
func NewReader(s string) *Reader
```

#### bytes.Buffer
- bytes.Buffer类型的值记录的已读计数，在其中起到了怎样的作用？
```go
type Buffer struct {
	buf      []byte // contents are the bytes buf[off : len(buf)]
	off      int    // read at &buf[off], write at &buf[len(buf)]
	lastRead readOp // last read operation, so that Unread* can work correctly.
}
```
- 接口类
```go
func main() {
	// 示例1。
	builder := new(strings.Builder)
	_ = interface{}(builder).(io.Writer)
	_ = interface{}(builder).(io.ByteWriter)
	_ = interface{}(builder).(fmt.Stringer)

	// 示例2。
	reader := strings.NewReader("")
	_ = interface{}(reader).(io.Reader)
	_ = interface{}(reader).(io.ReaderAt)
	_ = interface{}(reader).(io.ByteReader)
	_ = interface{}(reader).(io.RuneReader)
	_ = interface{}(reader).(io.Seeker)
	_ = interface{}(reader).(io.ByteScanner)
	_ = interface{}(reader).(io.RuneScanner)
	_ = interface{}(reader).(io.WriterTo)

	// 示例3。
	buffer := bytes.NewBuffer([]byte{})
	_ = interface{}(buffer).(io.Reader)
	_ = interface{}(buffer).(io.ByteReader)
	_ = interface{}(buffer).(io.RuneReader)
	_ = interface{}(buffer).(io.ByteScanner)
	_ = interface{}(buffer).(io.RuneScanner)
	_ = interface{}(buffer).(io.WriterTo)

	_ = interface{}(buffer).(io.Writer)
	_ = interface{}(buffer).(io.ByteWriter)
	_ = interface{}(buffer).(io.ReaderFrom)

	_ = interface{}(buffer).(fmt.Stringer)

	// 示例4。
	src := strings.NewReader(
		"CopyN copies n bytes (or until an error) from src to dst. " +
			"It returns the number of bytes copied and " +
			"the earliest error encountered while copying.")
	dst := new(strings.Builder)
	written, err := io.CopyN(dst, src, 58)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Printf("Written(%d): %q\n", written, dst.String())
	}
}
```
- 问题 1：bytes.Buffer的扩容策略是怎样的？
- 新容器的容量 =2* 原有容量 + 所需字节数
```go
b.buf = b.buf[:length+need]
cap(b.buf)/2 >= b.Len() + need
```
- 问题 2：bytes.Buffer中的哪些方法可能会造成内容的泄露？

#### bufio
bufio包中的数据类型主要有：Reader；Scanner；Writer和ReadWriter。
```go
type Reader struct {
	buf          []byte
	rd           io.Reader // reader provided by the client
	r, w         int       // buf read and write positions
	err          error
	lastByte     int // last byte read for UnreadByte; -1 means invalid
	lastRuneSize int // size of last rune read for UnreadRune; -1 means invalid
}
```
- Writer
```go
type Writer struct {
	err error
	buf []byte
	n   int
	wr  io.Writer
}
```
- ReadWriter
```go
type ReadWriter struct {
	*Reader
	*Writer
}
```
- Scanner
```go
type Scanner struct {
	r            io.Reader // The reader provided by the client.
	split        SplitFunc // The function to split the tokens.
	maxTokenSize int       // Maximum size of a token; modified by tests.
	token        []byte    // Last token returned by split.
	buf          []byte    // Buffer used as argument to split.
	start        int       // First non-processed byte in buf.
	end          int       // End of data in buf.
	err          error     // Sticky error.
	empties      int       // Count of successive empty tokens.
	scanCalled   bool      // Scan has been called; buffer is in use.
	done         bool      // Scan has finished.
}
```

### OS包
在os包中，有这样几个函数，即：Create、NewFile、Open和OpenFile。
### 网络

### 性能分析
go tool pprof和go tool trace
- runtime/pprof.Lookup函数的正确调用方式是什么？
goroutine、heap、allocs、threadcreate、block和mutex。
- 如何为基于 HTTP 协议的网络服务添加性能分析接口？
```go
import _ "net/http/pprof"
log.Println(http.ListenAndServe("localhost:8082", nil))
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=60
```
- 涉及包
```go
    runtime/pprof；
    net/http/pprof；
    runtime/trace;
```
#### CPU 概要文件（CPU Profile）
pprof.StartCPUProfile(f)
#### 内存概要文件（Mem Profile）
runtime.MemProfileRat(f)
#### 和阻塞概要文件（Block Profile）
runtime.SetBlockProfileRate(f)


 
