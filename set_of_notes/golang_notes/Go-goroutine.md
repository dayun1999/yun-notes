# Goroutine

**Created By: 蜜雪冰熊**

## 目录

- ##### <a href="#goroutine_self">goroutine本身</a>
  
  - <a href="#why_goroutine_samll">goroutine为什么如此轻量?</a>
  - <a href="#how_goroutine_implemented">goroutine是如何实现的</a>
  - <a href="#how_to_acquire_goroutine_id">如何获取goroutine的ID？</a>
  - <a href="#"></a>
- <a href="#"></a>

- ##### <a href="#goroutine_and_memory">goroutine与内存</a>
  
  - <a href="#what_is_out_of_memory">什么是内存溢出（OOM）</a>
  - <a href="#what_is_memory_leak">什么是内存泄漏</a>
  - <a href="#what_is_goroutine_leak">什么是goroutine泄漏</a>
  - <a href="#some_scenarios_causing_oom">一些可能的内存/goroutine泄漏场景</a>
  - <a href="#how_to_supervise_goroutine_leak">如何监控和排查goroutine泄漏</a>
  - <a href="#"></a>
  - <a href="#"></a>
- ##### <a href="#goroutine_and_thread">goroutine与线程</a>
  - <a href="#relationship_goroutine_thread">goroutine与线程的关系</a>
  - <a href="#compare_goroutine_and_thread_context_switch">goroutine切换与线程上下文切换对比</a>
  - <a href="#">如何查看切换成本</a>
  - <a href="#"></a>
  - <a href="#"></a>
  - <a href="#"></a>
  - <a href="#"></a>
- <a href="#"></a>
  - <a href="#"></a>
  - <a href="#"></a>
  - <a href="#"></a>
  - <a href="#"></a>
- <a href="#"></a>
  - <a href="#"></a>
- <a href="#"></a>
  - <a href="#"></a>
- <a href="#"></a>
  - <a href="#"></a>
  - <a href="#"></a>
  - <a href="#"></a>
- <a href="#linux_performance_supervise">Linux系统中如何进行性能排查</a>
  - <a href="#how_to_supervise_context_switch">如何查看上下文切换？</a>
  - <a href="#"></a>
  - <a href="#"></a>











## <a name="goroutine_self">goroutine本身</a>

### <a name="why_goroutine_samll">1. 为什么goroutine如此轻量?</a>

#### 1.1 跟谁比算得上轻量？

既然我们是寻找goroutine轻量的原因，那么就必须有一个参照，即：goroutine相对什么显得轻量？——答案：相对于操作系统的原生线程来说goroutine很轻量

原生线程的开销主要表现在:

1. CPU在上下文切换中的开销
2. 线程的栈空间

#### 1.2 上下文切换有多贵？

**CPU 寄存器和程序计数器就是 CPU 上下文**，因为它们都是 CPU 在运行任何任务前，必须的依赖环境。

上下文切换先把前一个任务的 CPU 上下文（也就是 CPU 寄存器和程序计数器）保存起来，然后加载新任务的上下文到这些寄存器和程序计数器，最后再跳转到程序计数器所指的新位置，运行新任务。

这其中单次上下文切换的时间在微秒的粒度，即使仅仅10微秒，那么每一秒钟也仅仅能切换10万次，换句话说，每个核心在不干其他事情的情况下，只做上下文切换也只能切换10万次。

这个和动不动就上百万，上千万的goroutine的数量是不能比的



#### 1.3 Go的调度模型

goroutine是被Go自己运行时调度的，属于用户级线程，有如下几个特点:

- 相比线程,其启动的代价很小, 以很小的栈空间启动(2kB)
- 能够动态地伸缩栈的大小, 最大可以支持Gb级别
- 工作在用户态，切换成本很小
- 与线程关系是多对多`n : m`

**补充: **启动一个线程的开销可以用以下命令查看

```bash
$ ulimit -as | grep stack
stack size              (kbytes, -s) 8192  ## 默认8MB
```



### <a name="how_goroutine_implemented">2. goroutine是如何实现的?</a>





### <a name="how_to_acquire_goroutine_id">3. 如何获取goroutine的ID？</a>













## <a name="goroutine_and_memory">goroutine与内存</a>

### <a name="what_is_out_of_memory">1. 什么是内存溢出（OOM）</a>

#### 定义

**内存溢出 out of memory**，是指程序在申请内存时，没有足够的内存空间供其使用，出现out of memory；比如系统现在只有1G的空间，但是你偏偏要2个G空间，这就叫内存溢出

### <a name="what_is_memory_leak">2. 什么是内存泄漏</a>

#### 定义

**内存泄漏**（英语：**Memory leak**）是指程序未能释放已经不再使用的内存。内存泄漏并非指内存在物理上的消失，而是应用程序分配某段内存后，由于设计错误，导致在释放该段内存之前就失去了对该段内存的控制，从而造成了内存的浪费。

#### 后果

内存泄漏会因为减少可用内存的数量从而降低计算机的性能。最终，在最糟糕的情况下，过多的可用内存被分配掉导致全部或部分设备停止正常工作，或者应用程序崩溃

memory leak会最终会导致out of memory！



### <a name="what_is_goroutine_leak">3. 什么是goroutine泄漏</a>

goroutine没有被关闭，或者没有添加超时控制，让goroutine一只处于阻塞状态，不能被GC

### <a name="some_scenarios_causing_oom">4. 一些可能的内存/goroutine泄漏场景</a>

##### 可能导致内存goroutine泄漏的原因的总结<重点>

- 因为协程被永久阻塞资源一直无法释放而造成的永久性内存泄漏
- 因为没有停止不再使用的`time.Ticker`值而造成的永久性内存泄漏
- 因为不正确地使用终结期`finalizer`而造成的永久性内存泄漏
- `channel`导致的泄漏
  - 发送不接收
  - 接收不发送
  - nil channel
- Goroutine 内的业务逻辑进入长时间等待，有不断新增的 Goroutine 进入等待
- Goroutine 内正在进行 channel/mutex 等读写操作，但由于逻辑问题，某些情况下会被一直阻塞
- 



下面具体介绍一下: 

#### 4.1 因为协程被永久阻塞而造成的永久性内存泄漏

Go运行时并不会将处于永久阻塞状态的协程杀掉, 因此永久处于阻塞的协程所占用的资源得不到释放;  原因有二:

1. <font color="red">有时候Go运行时很难分辨出一个处于阻塞状态的协程是永久阻塞还是暂时性阻塞；</font>
2. <font color="red">有时候我们可能故意永久阻塞某些协程；</font>

#### 4.2 因为没有停止不再使用的`time.Ticker`值而造成的永久性内存泄漏

当一个`time.Timer`值不再被使用，一段时间后它将被自动垃圾回收掉；

但是<font color="blue">对于不再被使用的`time.Ticker`值，我们必须调用它的`Stop`方法结束它, 否则它将永远不会得到回收</font>；

#### 4.3 因为不正确地使用终结期`finalizer`而造成的永久性内存泄漏



#### 4.4 Channel导致的泄漏

##### 4.4.1 发送不接收

```go
func gen(nums ...int) <-chan int {
	out := make(chan int) // 无缓冲channel
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func main() {
	defer func() {
		fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
	}()

	out := gen(2, 3)

	for n := range out {
		fmt.Println(n) // 只会打印2
		time.Sleep(3 * time.Second) // 模拟处理其他事情
        // 模拟异常退出
		if true {
			break
		}
	}
}

// 输出结果
2
the number of goroutines:  2
```

发现另外一个goroutine一直阻塞, 因为接收者停止工作, 发送者`out <- n`并不知道

如何改进？：`可以通过关闭channel向所有的接收者发送广播消息`

```go
func gen(done <-chan struct{}, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
			case <-done: // done关闭后,这边会接收到done的零值
				return
			}
		}
	}()
	return out
}

func main() {
	defer func() {
		time.Sleep(time.Second)
		fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
	}()
	done := make(chan struct{})
	defer close(done)
	out := gen(done, 2, 3)

	for n := range out {
		fmt.Println(n)
		time.Sleep(1 * time.Second)
		if true {
			break
		}
	}
}

```

##### 4.4.2 接收不发送

```GO
func main() {
	defer func() {
		time.Sleep(time.Second)
		fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
	}()

    ch := make(chan struct{})
	go func() {
		ch <- struct{}{}
	}()
}

// 输出
the number of goroutines:  2
```



##### 4.4.3 nil channel

```go
var ch chan int // 未初始化
....
```

channel 如果忘记初始化，那么无论你是读，还是写操作，都会造成阻塞, 所以要写成`ch := make(chan int)`，用make初始化



#### 4.5 经典场景--奇怪的慢等待

```go
func main() {
    for {
        go func() {
            _, err := http.Get("https://www.xxx.com/")
            if err != nil {
                fmt.Printf("http.Get err: %v\n", err)
            }
            // do something...
    	}()

        time.Sleep(time.Second * 1)
        fmt.Println("goroutines: ", runtime.NumGoroutine())
	}
}

// 运行结果
number of goroutines: 7
number of goroutines: 13
number of goroutines: 19
number of goroutines: 25
number of goroutines: 31
number of goroutines: 37
number of goroutines: 43
.......
```

在这个例子中，展示了一个 Go 语言中经典的事故场景。也就是一般我们会在应用程序中去调用第三方服务的接口。

但是第三方接口，有时候会很慢，久久不返回响应结果。恰好，Go 语言中默认的 `http.Client` 是没有设置超时时间的。

因此就会导致一直阻塞，一直阻塞就一直爽，Goroutine 自然也就持续暴涨，不断泄露，最终占满资源，导致事故。

在 Go 工程中，我们一般建议至少对 `http.Client` 设置超时时间：

```go
httpClient := http.Client{
    Timeout: time.Second * 15,
}
```

并且要做限流、熔断等措施，以防突发流量造成依赖崩塌，依然吃 P0



#### 4.6 锁的不正确使用导致 

##### 4.6.1 互斥锁忘记解锁

```go
func main() {
	total := 0
	defer func() {
		time.Sleep(time.Second)
		fmt.Println("total", total)
		fmt.Printf("number of goroutines: %d\n", runtime.NumGoroutine())
	}()

	var mutex sync.Mutex
	for i := 0; i < 10; i++ {
		go func() {
			mutex.Lock()
            // defer mutex.Unlock
			total += 1
		}()
	}
}

// 输出结果
total 1
number of goroutines: 10
```



##### 4.6.2 同步锁使用不当

```go
func handle(v int) {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < v; i++ {
		fmt.Println("wdy")
		wg.Done()
	}
	wg.Wait()
}

func main() {
	defer func() {
		fmt.Println("number of goroutines：", runtime.NumGoroutine())
	}()
	go handle(3)
	time.Sleep(time.Second)
}
// 输出结果
wdy
wdy
wdy
number of goroutines： 2
```

由于 `wg.Add` 的数量与 `wg.Done` 数量并不匹配，因此在调用 `wg.Wait` 方法后一直阻塞等待

```go
 var wg sync.WaitGroup
    for i := 0; i < v; i++ {
        wg.Add(1)
        defer wg.Done() // 这里
        fmt.Println("wdy")
    }
    wg.Wait()
```
















### <a name="how_to_supervise_goroutine_leak">5. 如何监控和排查goroutine泄漏</a>

#### 5.1 `runtime.NumGoroutine()`

使用`runtime.NumGoroutine()`查看goroutine数量;

如果 goroutine 随着时间增加，数量在不断上升，而基本没有下降，基本可以确定存在泄露



#### 5.2 pprof

【建议直接看: [煎鱼--pprof](https://eddycjy.com/posts/go/tools/2018-09-15-go-tool-pprof/) 】

想要进行性能优化，首先瞩目在 Go 自身提供的工具链来作为分析依据，本文将带你学习、使用 Go 后花园，涉及如下：

- **runtime/pprof**：采集程序（非 Server）的运行数据进行分析
- **net/http/pprof**：采集 HTTP Server 的运行时数据进行分析

```bash
# 通过web页面
访问 http://127.0.0.1:6060/debug/pprof/
# 通过交互式终端
$ go tool pprof http://localhost:6060/debug/pprof/profile?seconds=60
$ go tool pprof http://localhost:6060/debug/pprof/heap
$ go tool pprof http://localhost:6060/debug/pprof/block
$ go tool pprof http://localhost:6060/debug/pprof/mutex
```

建议还是看博客吧







## <a name="goroutine_and_thread">goroutine与线程</a>

### <a name="relationship_goroutine_thread">1. goroutine与线程的关系</a>





### <a name="compare_goroutine_and_thread_context_switch">2. goroutine切换与线程上下文切换对比</a>

- Goroutine 上下文切换只涉及到三个寄存器（PC / SP / DX）的值修改；

- 线程的上下文切换则需要涉及模式切换（从用户态切换到内核态）、以及 16 个寄存器、PC、SP…等寄存器的刷新



## <a name="#"></a>



## <a name="#"></a>



## <a name="#"></a>



## <a name="linux_performance_supervise">Linux系统中如何进行性能排查</a>

### <a name="how_to_supervise_context_switch">如何查看上下文切换？</a>

#### 1. vmstat

`vmstat`命令查看系统总体的上下文切换情况

```bash
$ vmstat
procs -----------memory---------- ---swap-- -----io---- -system-- ------cpu-----
 r  b   swpd   free   buff  cache   si   so    bi    bo   in   cs us sy id wa st
 5  0      0 286948  27304 877076    0    0    12     3    2    9  1  1 98  0  0

```

Procs（进程）:

- r: 运行队列中进程数量
- b: 等待IO的进程数量

Memory（内存）:

- swpd: 使用虚拟内存大小
- free: 可用内存大小
- buff: 用作缓冲的内存大小
- cache: 用作缓存的内存大小

Swap:

- si: 每秒从交换区写到内存的大小
- so: 每秒写入交换区的内存大小

IO：（现在的Linux版本块的大小为1024bytes）

- bi: 每秒读取的块数
- bo: 每秒写入的块数

system：

- in: 每秒中断数，包括时钟中断
- <font color="red">cs: 每秒上下文切换数</font>

CPU（以百分比表示）

- us: 用户进程执行时间(user time)
- sy: 系统进程执行时间(system time)
- id: 空闲时间(包括IO等待时间)
- wa: 等待IO时间



#### 2. pidstat

要想查看**每个进程**的详细情况，需要 pidstat

```bash
# -w 表示显示每个进程的上下文切换, 1 表示每秒输出一次数据
$ pidstat -w 1 
```

