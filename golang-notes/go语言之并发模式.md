# go 语言之并发模式

推荐直接阅读[官方博客](https://blog.golang.org/pipelines#TOC_2.)

## 内容大纲

- _go 并发模式-Pipelines and cancellation_
- _go 并发模式-Context_

## Pipelines and cancellation

### 何为 pipeline?

没有明确定义, 暂时理解为: 一个 pipeline 就是一系列被 channel 连接起来的阶段(stage), 其中每个阶段可以被定义为一组运行着同一个函数的 goroutine \
每个阶段, goroutine 做的事情如下:

- 通过入 channel(inbound channel)从上游获取数据
- 对数据进行处理(一般都会因此产生新的值)
- 通过出 channel(outbound channel)向下游发送数据

### 官方例子: Squareing numbers

上面说过 pipeline 就是一系列的阶段(stage), 那么现在有一个任务: \
_将一个整数列表中的每一个元素进行平方_，这个任务可以被分隔为三个阶段，且这三个阶段被一个 channel 连接:

- _阶段一: 定义一个函数`gen`, 将整个列表中的数据挨个发送给一个 channel_
- _阶段二: 定义平方函数`sq`, 从 channel 中获取数据，平方后发送给一个 channel_
- _阶段三: main 函数接收从第二阶段发送过来的值并且打印_

```go
package main

import (
    "fmt"
)

// first stage
func gen(nums ...int) <-chan int {
    out := make(chan int)
    go func(){
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}

// second stage
func sq(in <-chan int) <-chan int {
    out := make(chan int)
    go func(){
        for n := range in {
            out <- n*n
        }
        close(out)
    }()
    return out
}

// third stage
func main() {
    // Set up the pipeline
    c := gen(2, 3, 4, 5)
    out := sq(c)

    // Consume the output
    for n := range out {
        fmt.Println(n)
    }
}
```

### 扇入、扇出

- 扇出(Fan-out): 多个函数从同一个channel中读取数据(直到chnnel关闭)
- 扇入(Fan-in ): 一个函数从多个channel读取数据, 将这些输入channel复用到一个通道，直到这个单个的channel被关闭

接着上面的这个例子来理解fan-in, 引入一个函数`merge`, 来将`sq`函数的两个实例结合到一起

```go
func main() {
    in := gen(9, 10)

    // 将平方这项工作分发给两个goroutine, 让它们都从 'in' 中读取数据
    c1 := sq(in)
    c2 := sq(in)

    // 处理最终从c1和c2中合并的结果
    for n := range merge(c1, c2) {
        fmt.Println(n)
    }
}

func merge(cs ...<-chan int) <-chan int {
    // 引入WaitGroup来保证同步
    var wg sync.WaitGroup
    // out就是那个被复用的channel
    out := make(chan int)

    output := func(in <-chan int) {
        for n := range in {
            out <- n
        }
        wg.Done()
    }

    wg.Add(len(cs))

    // 对于每个input channel,都启动一个goroutine
    for _, c := range cs {
        go output(c)
    }

    // 另外启动一个goroutine来负责关闭out channel
    go func() {
        wg.Wait()
        close(out)
    }()

    return out
}
```

### goroutine的取消

思考，上面的程序有个特点:

- *所有发送操作结束之后下游的channel才关闭*
- *下游一直接收上游发来的值直到上游的所有channel关闭*

问题在于现实程序中我并不一定需要channel里面所有的值, 这就需要*明确取消channel*的操作了

```go
package main

import (
    "sync"
    "fmt"
)

// first stage
func gen(done <-chan struct{}, nums ...int) <-chan int {
    out := make(chan int)
    go func(){
        defer close(out)
        for _, n := range nums {
            select {
            case out <- n:
            case <-done:
                return
            }
        }
        // close(out)
    }()
    return out
}

// second stage
// 修改后,增加参数done <-chan struct{}
func sq(done <-chan struct{}, in <-chan int) <-chan int {
    out := make(chan int)
    go func(){
        // defer确保关闭out channel
        defer close(out)
        for n := range in {
            select{
            case out <- n*n:
            case <-done:
                return
            }
        }
        // close(out)
    }()
    return out
}

func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
    var wg sync.WaitGroup
    out := make(chan int)

    // 为每个输入channel启动一个输出goroutine
    output := func(c <-chan int) {
        // 利用defer确保执行
        defer wg.Done()
        for n := range c {
            select {
            case out <- n:
            case <-done:
                return
            }

        }
        // wg.Done()
    }
    wg.Add(len(cs))
    for _, c := range cs {
        go output(c)
    }

    // 另外启动一个goroutine来关闭out channel一旦所有的output goroutine都结束的时候
    go func(){
        wg.Wait()
        close(out)
    }()

    return out
}

// third stage
func main() {
    // 设置一个共享整个pipeline的channel--done
    // 并且当这个pipeline退出的时候关闭这个channel, 作为一个广播信号
    // 告诉所有goroutine我们要退出了
    done := make(chan struct{})
    defer close(done)

    // Set up the pipeline
    // 将done作为参数传递进去
    in := gen(done, 9, 10, 11, 12)
    c1 := sq(done, in)
    c2 := sq(done, in)

    // Consume the first value from output
    out := merge(done, c1, c2)
    fmt.Println(<-out)

}
```

### 计算一个目录所有文件的MD5值

```go
package main

import(
    "fmt"
    "path/filepath"
    "io/ioutil"
    "crypto/md5"
    "os"
    "sort"
)
// serial
func MD5All_Serial(root string) (map[string][md5.Size]byte, error) {
    m := make(map[string][md5.Size]byte)
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.Mode().IsRegular() {
            return nil
        }
        data ,err := ioutil.ReadFile(path)
        if err != nil {
            return err
        }
        m[path] = md5.Sum(data)
        return nil
    })

    if err != nil {
        return nil, err   
    }

    return m, nil
}

// concurrency
//func MD5All_Concurrency(root string) {

//}

func main() {
    // read the dir path from command line
    // calculate all the md5 value of all files in the given path
    m, err := MD5All_Serial(os.Args[1])
    if err != nil {
        fmt.Println(err)
        return
    }
    var paths []string
    for path := range m {
        paths = append(paths, path)
    }
    // sort asending
    sort.Strings(paths)
    for _, path := range paths {
        fmt.Printf("%x %s\n", m[path], path)
    }
}
```