## 题目

#### [50. Pow(x, n) （中等）](https://leetcode-cn.com/problems/powx-n/)

## 分析

快速幂相乘

**时间复杂度: ** `O(lgN)`

**空间复杂度:  **`O(1)`

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

