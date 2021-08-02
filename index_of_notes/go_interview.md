# Go语言常问题目

**Created By: 蜜雪冰熊**

## 常见语法题

### 1

下面代码会按预期输出吗？

```go
type query func(string) string

func exec(name string, vs ...query) string {
	ch := make(chan string)
	fn := func(i int) {
		ch <- vs[i](name)
	}
	for i, _ := range vs {
		go fn(i)
	}
	return <-ch
}

func main() {
	ret := exec("111",
		func(n string) string {
			return n + "func1"
		},
		func(n string) string {
			return n + "func2"
		},
		func(n string) string {
			return n + "func3"
		},
		func(n string) string {
			return n + "func4"
		})
	fmt.Println(ret)
}
```

#### 解答

只会输出一行`111func1`或者`111func4`或者...也就是四个函数不确定执行哪一个,但是exec只会有一个输出

----

### 2

下面代码有什么问题？

```go
type People struct {
	Name string
}

func (p *People) String() string {
	return fmt.Sprintf("print: %v", p)
}

func main() {
	p := &People{}
	p.String()
}
```

#### 解答

这里设计到Go语言的打印的接口`Stringer`
```go
type Stringer interface {
    String() string
}
```
本身`fmt.Sprintf()`就已经实现了Stringer接口, 这里结构体`People`显然也实现了`Stringer`结构, 而main函数中调用`p.String()`就产生了循环调用, 最终的结果是:
```go
runtime: sp=0xc0201615b8 stack=[0xc020160000, 0xc040160000]
fatal error: stack overflow

runtime stack:
runtime.throw({0x405cb9, 0x48c8c0})
	C:/Go/src/runtime/panic.go:1198 +0x76
runtime.newstack()
	C:/Go/src/runtime/stack.go:1086 +0x5bb
runtime.morestack()
	C:/Go/src/runtime/asm_amd64.s:461 +0x93
```

----

### 3

找出下面代码的问题所在

```go

```

----

### 4


----

### 5


----

### 6

