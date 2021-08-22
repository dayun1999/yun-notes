## 题目



## 分析



## 解答

```go
func myPow(x float64, n int) float64 {
    if n < 0 {
        return 1.0/ pow(x, -n)
    }
    return pow(x, n)
}

func pow(x float64, n int) (res float64) {
    res = 1.0
    x_contribute := x
    for n > 0 {
        if n&1 == 1 {
            res *= x_contribute
        }
        x_contribute *= x_contribute
        n >>= 1
    }
    return
}
```

