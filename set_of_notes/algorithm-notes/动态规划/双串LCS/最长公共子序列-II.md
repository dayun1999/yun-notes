## 题目

[最长公共子序列-II_](https://www.nowcoder.com/practice/6d29638c85bb4ffd80c020fe244baf11?tpId=117&&tqId=37798&&companyId=665&rp=1&ru=/company/home/code/665&qru=/ta/job-code-high/question-ranking)

## 分析

和**最长公共子序列**类似, 这不过这道题返回的是字符串,所以怎么解决呢？ 

将dp数组的元素设置为`[]byte`, 当找到字符串的时候就append右上角, 换汤不换药

![](C:\Users\26646\Desktop\牛客网刷题笔记\Pictures\最长公共子序列.png)

## 解答

```go
// NC92: 最长公共子序列II
func LCS(str1 string, str2 string) string {
	// write code here
	S1 := len(str1)
	S2 := len(str2)
    // 将dp设置为[]byte数组
	dp := make([][][]byte, S1+1)
	for i := range dp {
		dp[i] = make([][]byte, S2+1)
	}
	for i := 1; i <= S1; i++ {
		for j := 1; j <= S2; j++ {
			// 当两个字符相等
			if str1[i-1] == str2[j-1] {
				dp[i][j] = append(dp[i][j], dp[i-1][j-1]...)
				dp[i][j] = append(dp[i][j], str1[i-1])
			} else {
				if len(dp[i-1][j]) > len(dp[i][j-1]) {
					dp[i][j] = append(dp[i][j], dp[i-1][j]...)
				} else {
					dp[i][j] = append(dp[i][j], dp[i][j-1]...)
				}
			}
		}
	}
	if len(dp[S1][S2]) == 0 {
		return "-1"
	}
	return string(dp[S1][S2])
}
```

