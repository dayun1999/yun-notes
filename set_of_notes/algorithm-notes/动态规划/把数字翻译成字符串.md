## 题目

#### [剑指 Offer 46. 把数字翻译成字符串](https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/)

```go
输入: 12258
输出: 5
解释: 12258有5种不同的翻译，分别是"bccfi", "bwfi", "bczi", "mcfi"和"mzi"
```



## 分析

仔细分析之后, 发现这跟斐波那契数列是类似的

## 解答

```go
// 动态规划--斐波那契数列的条件版
func translateNum(num int) int {
	if num < 10 {
		 return 1
	}
	str := strconv.Itoa(num)
	N := len(str)
	dp := make([]int, N+1)
	dp[0] = 1 // 这步需要反推
	dp[1] = 1
	
	for i := 2; i <= N; i++ {
		// 把前面两个数字合起来判断能否直接翻译(10-25之间都是能直接翻译的)
		temp := string(str[i-2]) + string(str[i-1]) 
		if temp >= "10" && temp <= "25" {
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	} 
	return dp[N]
}
```

