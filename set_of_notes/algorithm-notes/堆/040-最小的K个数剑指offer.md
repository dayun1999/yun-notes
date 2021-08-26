## 题目

#### [剑指 Offer 40. 最小的k个数](https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/)

输入整数数组 `arr` ，找出其中最小的 `k` 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4

```go
输入：arr = [3,2,1], k = 2
输出：[1,2] 或者 [2,1]
```



## 分析




**时间复杂度: ** `O()`

**空间复杂度:  **`O()`

## 解答

```go
func getLeastNumbers(nums []int, k int) []int {
    N := len(nums)
    if k == 0 || N == 0 {
        return []int{}
    }
    h := &IntHeap{}
    for i := 0; i < k; i++ {
        heap.Push(h, nums[i])
    }

    for i := k; i <N; i++ {
        if nums[i] <= h.Peek().(int) {
            heap.Pop(h)
            heap.Push(h, nums[i])
        }
    }
    res := []int{}
    for h.Len() > 0 {
        res = append(res, heap.Pop(h).(int))
    }
    return res
}

// 最大堆的实现
type IntHeap []int
func (h IntHeap) Len() int { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Pop() (x interface{}) {
    old := *h
    n := len(old)
    x = old[n-1]
    *h = old[0:n-1]
    return
}

func (h *IntHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *IntHeap) Peek() (x interface{}) {
    return (*h)[0]
}
```

