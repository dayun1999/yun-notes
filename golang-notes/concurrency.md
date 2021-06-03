# Concurrency

推荐直接阅读:point_right:[Effective go](https://golang.org/doc/effective_go#concurrency)

并发优秀代码收集:point_right:[点这](https://github.com/code4EE/yun-notes/golang-notes/blob/main/golang-notes/concurrency_code_collections.md)

## 内容大纲

- *并发还是并行:question:*
- *关于并发的谚语*
- *那些可以优化的并发的代码*

## 并发还是并行?

- ### 并发的思想: 将一个程序构建成(structure)多个独立的执行模块

- ### 并行的思想: 为了获得效率从而在多个物理CPU上并行地执行计算,并行只有在有多个物理CPU核心的情况下才可能存在，是真真正正的同一时刻运行多个程序

## 关于并发的谚语

1. Concurrency is parallelism.

2. Keep yourself busy or do the work yourself.

3. Leave concurrency to the caller.

4. Never start a goroutine without knowing when it will stop.

5. Only use log.Fatal from main.main or init functions.

## 那些可以优化的并发的代码

### 缓冲Channel的使用

- #### bad

```go
var sem = make(chan int, MaxOutstanding)

func handle(r *Request) {
    sem <- 1    //等待缓冲channel耗尽
    process(r)  //需要一定时间处理
    <-sem       //完成;可以让下一个request进来了
}

func Serve(queue chan *Request) {
    for {
        req := <-queue
        go handle(req) //不用等到handle函数完成
    }
}
```

##### 上述代码的问题就在于,来若干个请求(request)都会主动创建对应若干个goroutine,即使他们中最多只有`MaxOutstanding`个可以随时运行,剩下的goroutine只能等待，这样做大量地创建goroutine会导致无限制的消耗大量的资源

- #### good

```go
func hande(queue chan *Request) {
    for r := range queue {
        process(r)
    }
}

func Serve(clientRequests chan *Request, quit chan bool) {
    for i := 0; i < MaxOutstanding; i++ {
        go handle(clientRequests)
    }
    <-quit  //一直等待,直到被告知要退出
}
```
