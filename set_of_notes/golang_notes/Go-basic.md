# GO知识点总结

**Created By: 蜜雪冰熊**

## 目录

- <a href="#slice">切片</a>
- <a href="#string">字符串</a>
- <a href="#map">map</a>
- <a href="#function_call">函数调用</a>
- <a href="#interface">接口</a>
- <a href="#defer">defer</a>
- <a href="#panic_recover">panic 和 recover</a>
- <a href="#make&new">make 和 new</a>
- <a href="#uintptr_and_unsafe.Pointer">uintptr和unsafe.Pointer的区别</a>
- <a href="#concurrency-context">并发编程--上下文Context</a>
- <a href="#concurrency-primitives&mutex">并发编程--同步原语与锁</a>
- <a href="#concurrency-channel">并发编程--Channel</a>
- <a href="#golang_gc">GC</a>
- <a href="#"></a>
- <a href="#"></a>



## <a name="slice">切片</a>

#### 1. 切片的运行时表示

```go
type SliceHeader struct {
    Data	uintptr
    Len		int
    Cap		int
}
```



####  2. 切片的扩容机制

在分配内存空间之前需要先确定新的切片容量，运行时根据切片的当前容量选择不同的策略进行扩容：
- 如果期望容量大于当前容量的两倍就会使用期望容量；
- 如果当前切片的长度小于 1024 就会将容量翻倍；
- <font color="red">如果当前切片的长度大于 1024 就会每次增加 25% 的容量，直到新容量大于期望容量；</font>

[`runtime.growslice`](https://draveness.me/golang/tree/runtime.growslice) 函数最终会返回一个新的切片，其中包含了新的数组指针、大小和容量，这个返回的三元组最终会覆盖原切片

#### 举例说明

```go
var arr []int64
arr = append(arr, 1, 2, 3, 4, 5)
```

简单总结一下扩容的过程，当我们执行上述代码时，会触发 [`runtime.growslice`](https://draveness.me/golang/tree/runtime.growslice) 函数扩容 `arr` 切片并传入期望的新容量 5，这时期望分配的内存大小为 40 字节；不过因为切片中的元素大小等于 `sys.PtrSize`，所以运行时会调用 [`runtime.roundupsize`](https://draveness.me/golang/tree/runtime.roundupsize) 向上取整内存的大小到 48 字节，所以新切片的容量为 48 / 8 = 6



#### 3. nil 切片和空切片

<img src="C:\Users\26646\Desktop\Go面试系列\pictures\nil-slice-empty-slice.jpg" style="zoom:67%;" />



## <a name="string">字符串</a>

#### 1. 字符串的底层表示

```go
type StringHeader struct {
    Data 	uintptr
    Len 	int
}
```

字符串可以说是一个只读的切片类型

所有在字符串上的写入操作都是通过拷贝实现的。

#### 2. 字符串的拼接

##### 2.1  直接用`+`拼接

运行时会调用`copy`将输入的多个字符串`拷贝`到目标字符串所在的内存空间, 新的字符串是一片新的内存空间，与原来的字符串也没有任何关联, **当拼接的字符串非常大的时候, 拷贝会带来无法忽略的性能损失**

##### 2.2 其他高效的字符串拼接

- 使用`strings.Builder`

```go
// 使用 strings.Builder


```

- 使用`bytes.Buffer`

```go
// 使用bytes.Buffer

```



#### 3. `string`与`[]byte`的类型转换



#### 4. `str == ""` 和 `len(str) == 0`有什么区别

没有区别, 汇编代码都是一样的



## <a href="#map">map</a>

#### 1. key与value的限制

key的类型一定要可比较`可以理解为支持 == 的操作`

| 可比较类型      | 不可比较类型 |
| --------------- | ------------ |
| bool            | slice        |
| numeric         | func         |
| string          | map          |
| pointer         |              |
| channel         |              |
| interface       |              |
| array 和 struct |              |

#### 2. map的扩容机制

如果之前为2^n ，那么下一次扩容是2^(n+1),每次扩容都是之前的两倍。扩容后需要重新计算每一项在hash中的位置，新表为老的两倍，此时前文的oldbacket用上了，用来存同时存在的两个心就map，等数据迁移完毕就可以释放oldbacket了

好处: **均摊扩容时间，一定程度上缩短了扩容时间**

那么overLoadFactor函数中有一个常量6.5（loadFactorNum/loadFactorDen）来进行影响扩容时机。这个值的来源是测试取中的结果

#### 3.  map的gc回收机制

一句话: `delete`并不会立刻删除map中的内容

```go
var intMap map[int]int // 万不可放到main函数里面,否则直接分配到栈上去了,全局变量分配在堆上

func main() {
	printMemStats("初始化")

	//
	intMap = make(map[int]int, 100000)
	for i := 0; i < 100000; i++ {
		intMap[i] = i
	}
	// 手动进行GC
	runtime.GC()

	printMemStats("添加map")

	log.Println("删除前map长度", len(intMap))
	for i := 0; i < 100000; i++ {
		delete(intMap, i)
	}
	log.Println("删除后map长度", len(intMap))

	// 手动GC
	runtime.GC()
	printMemStats("删除数据后")

	// 设置nil
	intMap = nil
	runtime.GC()
	printMemStats("设置为nil后")

}

// 查看当前内存情况
func printMemStats(message string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Println(message)
	log.Printf("---------- 内部分配=%vKB\n", m.Alloc/1024)
	log.Printf("---------------------GC 次数=%v\n", m.NumGC)
}
```

```go
// 输出结果
2021/08/19 11:06:57 初始化
2021/08/19 11:06:57 ---------- 内部分配=151KB
2021/08/19 11:06:57 ---------------------GC 次数=0
2021/08/19 11:06:57 添加map
2021/08/19 11:06:57 ---------- 内部分配=2846KB
2021/08/19 11:06:57 ---------------------GC 次数=1
2021/08/19 11:06:57 删除前map长度 100000
2021/08/19 11:06:57 删除后map长度 0
2021/08/19 11:06:57 删除数据后
2021/08/19 11:06:57 ---------- 内部分配=2846KB
2021/08/19 11:06:57 ---------------------GC 次数=2
2021/08/19 11:06:57 设置为nil后
2021/08/19 11:06:57 ---------- 内部分配=148KB
2021/08/19 11:06:57 ---------------------GC 次数=3
```





## <a name="function_call">函数调用</a>

#### 1. 为什么Go能返回多个值而C语言只能返回一个

**Go使用栈传递参数和接收返回值**, 所以只需要在栈上多分配一些内存就可以返回多个值

**C语言同时使用寄存器和栈传递参数, 使用`eax`寄存器传递返回值**：

上面两种设计的优缺点:

- C 语言的方式能够极大地减少函数调用的额外开销，但是也增加了实现的复杂度；
  - CPU 访问栈的开销比访问寄存器高几十倍；
  - 需要单独处理函数参数过多的情况；
- Go 语言的方式能够降低实现的复杂度并支持多返回值，但是牺牲了函数调用的性能；
  - 不需要考虑超过寄存器数量的参数应该如何传递；
  - 不需要考虑不同架构上的寄存器差异；
  - 函数入参和出参的内存空间需要在栈上进行分配；

通过堆栈传递参数，入栈的顺序是从右到左，而参数的计算是从左到右



#### 2. 参数传递

> Go语言都是值传递, 都会对传递的参数进行拷贝

##### 2.1 传递基本类型和数组

向函数传递`map`或者`channel`，函数内部的修改对外部可见, 是因为`map`或者`channel`都是运行时类型的一个指针, 一个`map`只是一个`runtime.hmap`结构体的指针

```go
// 总结一下就是:
向函数中传递 map channel slice, 内部的修改都会改变原有的值, 而传递[2]int数组不会
```



##### 2.2 结构体和指针

- 传递结构体时：会拷贝结构体中的全部内容；
- 传递结构体指针时：**会拷贝结构体指针**；



## <a name="interface">接口</a>

#### 1. 接口的底层实现

Go有两种接口：`见源码src/runtime/runtime2.go`

- `eface`是不带方法的接口
- `iface`是带方法的接口

```go
type iface struct {
	tab  *itab
	data unsafe.Pointer
}

type eface struct {
	_type *_type  // _type是Go语言类型运行时的表示
	data  unsafe.Pointer
}
```




#### 2. 结构体和指针实现接口(`*T` 和 `T`)

|             | T    | *T |
| :--| :-- | :-- |
| `struct{}`  | 通过 | 不通过 |
| `&struct{}` | 通过 | 通过 |

也就是说指针类型实现接口只有指针类型能调用

#### 3. 类型断言

```go
func main() {
    switch c.(type) {
        case:
    }
}
```



## <a name="defer">defer</a>

`defer`经常被用于关闭文件描述符、关闭数据库连接以及解锁资源

#### 1.作用域问题

例子1

```go
func main() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

$ go run main.go
4
3
2
1
0
```

例子2

```go
func main() {
	{
		defer fmt.Println("defer function returns") // 3
		fmt.Println("block ends") // 1
	}
	fmt.Println("main function returns") // 2
}

$ go run main.go
block ends
main function returns
defer function returns
```



**总结: **`defer` 传入的函数不是在退出代码块的作用域时执行的，它只会在当前函数和方法返回之前被调用

#### 2. 预计算参数

调用 `defer` 关键字会立刻拷贝函数中引用的外部参数函，数创建新的延迟调用时就会立刻拷贝函数的参数，**函数的参数不会等到真正执行时计算**；

```go
func main() {
	started := time.Now()
	defer fmt.Println(time.Since(started))
	
	time.Sleep(time.Second)
}
// 输出:
0s
```

输出不符合预期, 

```go
func main() {
	started := time.Now()
	defer func() {
		fmt.Println(time.Since(started))
	}()
	
	time.Sleep(time.Second)
}
// 输出结果
1.0022446s
```

#### 3. `defer`的数据结构

```go
type _defer struct {
	siz       int32 // 参数和结果的大小
	started   bool
	openDefer bool
	sp        uintptr // 栈指针
	pc        uintptr // 调用方的程序计数器
	fn        *funcval // defer关键字传入的函数
	_panic    *_panic
	link      *_defer
}
```

延迟调用链:

![嘿嘿](https://img.draveness.me/2020-01-19-15794017184603-golang-defer-link.png)

![嘿嘿](https://img.draveness.me/2020-01-19-15794017184614-golang-new-defer.png)

`defer` 关键字的插入顺序是从后向前的，而 `defer` 关键字执行是从前向后的，这也是为什么后调用的 `defer` 会优先执行



#### 4. `defer`关键字的实现

堆分配、栈分配和开放编码是处理 `defer` 关键字的三种方法

- 堆上分配- `1.1 ~ 1.12`
- 栈上分配- `1.13`
- 开放编码- `1.14 ~ 现在`



#### 5. defer的一些有趣的面试题

```go
func main() {
	fmt.Println(func1())
	fmt.Println(func2())
}

// 这里用到了caller-save模式, i会被存放在main的栈空间,
func func1() (i int) {
	i = 100
	// 闭包中对外部变量i进行写入,所以这里是存的变量i的指针
	defer func() {
		i += 1 // 最后执行 5 + 1
	}()
	return 5 // 返回的时候将i修改为5
}

func func2() int {
	var i int // i声明在func2,所以会被放进func2的栈空间中,对比一下func1是放在main函数的栈空间
	defer func() {
		i += 1         // 修改的还是func2的栈空间的数,所以这里是101,
		fmt.Println(i) // 101
	}()
	i += 100
	return i // 将i=100返回到main的栈空间
}

// 预期输出:
// 101
// 6
// 实际输出
// 6
// 100
```



## <a name="panic_recover">panic 和 recover</a>

这两个关键字上一节提到的 `defer` 有紧密的联系，它们都是 Go 语言中的内置函数，也提供了互补的功能

- **`panic` 能够改变程序的控制流，调用 `panic` 后会立刻停止执行当前函数的剩余代码，并在当前 `Goroutine`中递归执行调用方的 `defer`；**
- **`recover` 可以中止 `panic` 造成的程序崩溃。它是一个只能在 `defer` 中发挥作用的函数，在其他作用域中调用不会发挥作用；**

### 1. 几个现象

- `panic` 只会触发当前 `Goroutine` 的 `defer`；
- `recover` 只有在 `defer` 中调用才会生效；
- `panic` 允许在 `defer` 中嵌套多次调用；

##### 1.1 panic跨协程失效

`panic`只会触发当前`goroutine`的延迟函数调用,  请看下面现象

```go
// 演示panic跨协程失效
func main() {
	defer println("in main") // 这个并不会执行
	go func() {
		defer println("in goroutine")
		panic("啊啊啊")
	}()

	time.Sleep(time.Second)
}

// 输出结果
in goroutine
panic: 啊啊啊

goroutine 5 [running]:
main.main.func1()
	C:/Users/26646/AppData/Local/liteide/liteide/goplay.go:13 +0x7b
created by main.main
	C:/Users/26646/AppData/Local/liteide/liteide/goplay.go:11 +0x76
exit status 2
```

###### 解释
前面我们曾经介绍过 `defer` 关键字对应的 [`runtime.deferproc`](https://draveness.me/golang/tree/runtime.deferproc) 会将延迟调用函数与调用方所在 `Goroutine` 进行关联。所以当程序发生崩溃时只会调用当前 `Goroutine` 的延迟调用函数也是非常合理的

![panic跨协程失效](https://img.draveness.me/2020-01-19-15794253176199-golang-panic-and-defers.png)

##### 1.2 失效的崩溃恢复

调用`recover`试图中止程序的崩溃, 但是失败了, 为什么呢？

```go
func main() {
	defer fmt.Println("in main")
	if err := recover(); err != nil {
		fmt.Println(err)
	}

	panic("unkonwn err")
}

// 输出结果
in main
panic: unkonwn err

goroutine 1 [running]:
main.main()
	C:/Users/26646/AppData/Local/liteide/liteide/goplay.go:13 +0x8d
exit status 2

```

`recover`只有在发生`panic`之后调用才会生效, 所以需要在`defer`中使用

```go
func main() {
	defer fmt.Println("in main")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("recover捕获错误...")
	}()
	panic("unkonwn err")
}

// 输出结果
unkonwn err
recover捕获错误...
in main
```



#### 1.3 嵌套崩溃

好像有点难理解

```go
func main() {
	defer fmt.Println("in main...") // 为什么这个先执行
	
	defer func() {
		defer func() {
			panic("【3】panic again and again...")
		}()
		panic("【2】panic twice")
	}()
	panic("【1】panic once")
}

// 输出结果
in main...
panic: 【1】panic once
	panic: 【2】panic twice
	panic: 【3】panic again and again...

goroutine 1 [running]:
main.main.func1.1()
	C:/Users/26646/AppData/Local/liteide/liteide/goplay.go:12 +0x27
panic({0x4c4b80, 0x4f6200})
......
```

所以，我们可以确定程序多次调用 `panic` 也不会影响 `defer` 函数的正常执行，所以使用 `defer` 进行收尾工作一般来说都是安全的

### 2. 数据结构

`src/runtime/runtime2.go`

```go
type _panic struct {
	argp      unsafe.Pointer // pointer to arguments of deferred call run during panic; cannot move - known to liblink
	arg       interface{}    // argument to panic
	link      *_panic        // link to earlier panic
	pc        uintptr        // where to return to in runtime if this panic is bypassed
	sp        unsafe.Pointer // where to return to in runtime if this panic is bypassed
	recovered bool           // whether this panic is over
	aborted   bool           // the panic was aborted
	goexit    bool
}
```



## <a name="make&new">make 和 new</a>

#### 1. 各个的作用

- `make` 的作用是初始化内置的数据结构，也就是我们在前面提到的切片、哈希表和 Channel;
- `new` 的作用是根据传入的类型**分配一片内存空间并返回指向这片内存空间的指针**;

#### 2. make

```go
slice := make([]int,2, 5)
hash := make(map[int]bool, 10)
ch := make(chan int, 5)
```

1. `slice` 是一个包含 `data`、`cap` 和 `len` 的结构体 [`reflect.SliceHeader`](https://draveness.me/golang/tree/reflect.SliceHeader)；
2. `hash` 是一个指向 [`runtime.hmap`](https://draveness.me/golang/tree/runtime.hmap) 结构体的指针；
3. `ch` 是一个指向 [`runtime.hchan`](https://draveness.me/golang/tree/runtime.hchan) 结构体的指针；

#### 3. new

直接理解为返回一个指针，更深层次的以后再补





## <a name="uintptr_and_unsafe.Pointer">uintptr和unsafe.Pointer的区别</a>

- unsafe.Pointer只是单纯的通用指针类型，用于转换不同类型指针，它不可以参与指针运算；
- 而uintptr是用于指针运算的，GC 不把 uintptr 当指针，也就是说 uintptr 无法持有对象， uintptr 类型的目标会被回收；
- unsafe.Pointer 可以和 普通指针 进行相互转换；
- unsafe.Pointer 可以和 uintptr 进行相互转换





## <a name="concurrency-context">并发编程--上下文Context</a>

#### 1. 是什么&什么用途

`context.Context`是Go语言中用来设置截止日期、同步信号, 传递请求相关值的结构体，与`goroutine`有着密切的联系；

```go
type Context interface {
    // 返回context被取消的时间, 也就是完成工作的截止日期,多次调用返回相同的值
    Deadline() (deadline time.Time, ok bool)
    
    // 返回一个channel, 该channel会在工作完成或者上下文被取消后关闭,多次调用返回相同的值
    Done() <-chan struct{}
    
    // 返回context结束的原因, 只会在Done方法对应的channel关闭时返回非空的值
    // 返回非空之后, 多次调用返回相同的值
    Err() error
    
    // 从context中获取键对应的值, 多次调用返回相同的值
    Value(key interface{}) interface{}
}
```



[`context`](https://github.com/golang/go/tree/master/src/context) 包中提供的 [`context.Background`](https://draveness.me/golang/tree/context.Background)、[`context.TODO`](https://draveness.me/golang/tree/context.TODO)、[`context.WithDeadline`](https://draveness.me/golang/tree/context.WithDeadline) 和 [`context.WithValue`](https://draveness.me/golang/tree/context.WithValue) 函数会返回实现该接口的私有结构体;



#### 2. 设计原理

Context与Goroutine树

![](https://img.draveness.me/golang-context-usage.png)

使用Context同步信号

```go
func main() {
	// 设置一个上下文
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	go handle(ctx, 500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("[main] ...")
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Printf("[handle] %v 时间到\n", duration)
	}
}

// 输出结果
[handle] 500ms 时间到
[main] ...
```



#### 3. 默认上下文

建议直接阅读`src/context/context.go`

- [`context.Background`](https://draveness.me/golang/tree/context.Background) 是上下文的默认值，所有其他的上下文都应该从它衍生出来；
- [`context.TODO`](https://draveness.me/golang/tree/context.TODO) 应该仅在不确定应该使用哪种上下文时使用；

在多数情况下，如果当前函数没有上下文作为入参，我们都会使用 [`context.Background`](https://draveness.me/golang/tree/context.Background) 作为起始的上下文向下传递

```go
// An emptyCtx is never canceled, has no values, and has no deadline.
type emptyCtx int

var (
	background = new(emptyCtx)
	todo       = new(emptyCtx)
)

func Background() Context {
	return background
}

func TODO() Context {
	return todo
}
```

![](https://img.draveness.me/golang-context-hierarchy.png)



#### 4. 取消信号

建议阅读[这篇文章](https://www.sohamkamani.com/golang/context-cancellation-and-values/)





## <a href="#concurrency-primitives&mutex">并发编程--同步原语与锁</a>



### 拓展原语

#### 2. semaphore

信号量是在并发编程中常见的一种同步机制，**在需要控制访问资源的进程数量时就会用到信号量**，它会保证持有的计数器在 0 到初始化的权重之间波动



#### 3. SingleFlight

`singleflight`能够在一个服务中抑制对下游的多次重复请求

常用的场景:

> 缓存击穿





## <a name="golang_gc">GC</a>

#### 1. GC的触发条件

- 超过内存大小阈值---`gcpercent`
- 达到定时时间--默认`2 min`

```go
const (
	// gcTriggerHeap indicates that a cycle should be started when
	// the heap size reaches the trigger heap size computed by the
	// controller.
	gcTriggerHeap gcTriggerKind = iota

	// gcTriggerTime indicates that a cycle should be started when
	// it's been more than forcegcperiod nanoseconds since the
	// previous GC cycle.
	gcTriggerTime

	// gcTriggerCycle indicates that a cycle should be started if
	// we have not yet started cycle number gcTrigger.n (relative
	// to work.cycles).
	gcTriggerCycle
)
```

