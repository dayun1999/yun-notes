# 算法技巧积累

**Created By: 蜜雪冰熊**

## 大纲

- <a href="#快速幂">快速幂求余</a>
- <a href="#快速幂相乘">快速幂相乘</a>
- <a href="#snowball">滚雪球(见题目《移动零》)</a>
- <a href="#automaton">有限自动机/状态机(《Atoi的实现》)</a>
- <a href="#239">单调队列的使用(见题目《滑动窗口的最大值》)</a>

<br>

## <a name="快速幂">快速幂求余</a>

为什么需要快速幂求余？\
比如求解`a^b % p`, 这里存在着明显的问题,如果a和b过大，很容易就会溢出

```go
func PowerMod(a, b, c int) int {
    res := 1
    a = a%c
    for b > 0 {
        if b&1 == 1 {
            res = (res*a) % c
        }
        b = b >> 1 // 除以2
        a = (a*a) % c
    }
    return res
    }
```
#### 相关题目

\
\
<br>
## <a name="快速幂相乘">快速幂相乘</a>

```go
// (77)10 = (1001101)2
// x^77 = x^(2^6) * x^(2^4) * x^(2^3) * x^(2^1)
func pow(x float64, N int) float64 {
    ans := 1.0
    x_contribute := x
    // 迭代法
    for N > 0 {
        // 如果当前位是1
        if N%2 == 1 {
            ans *= x_contribute
        }
        x_contribute *= x_contribute
        N = N>>1
    }
    return ans
}
```

\
\
<br>
## <a name="snowball">滚雪球</a>
#### 相关题目

|题目|实现代码|
|:--|:--|
|[Leetcode 283-移动零](https://leetcode-cn.com/problems/move-zeroes/)|[传送门]()|

\
\
<br>
## <a name="automaton">有限自动机/状态机</a>

#### 相关题目

|题目|实现代码|
|:--|:--|
|[剑指offer20: 表示数值的字符串](https://leetcode-cn.com/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof/)|[传送门](https://github.com/code4EE/yun-notes/blob/main/code_in_notes/leetcode_offer_20.go)|
|[Leetcode 008: Atoi的实现](https://leetcode-cn.com/problems/string-to-integer-atoi/)|[传送门](https://github.com/code4EE/yun-notes/blob/main/code_in_notes/leetcode_008.go)|

\
\
<br>
## <a name="239">单调队列的使用(见题目《滑动窗口的最大值》)</a>

### What?

一种队列，同时保持其元素顺序单调
### When to use?

1. 需要维护一个队列结构，支持元素的插入和删除
2. 在任意时刻访问队列中元素的最大值
3. 各操作的平均复杂度为O(1)


















