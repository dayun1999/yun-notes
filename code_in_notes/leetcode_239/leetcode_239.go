// 滑动窗口的最大值
package leetcode_239

func maxSlidingWindow(nums []int, k int) []int {
	N := len(nums)
	if N < 2 {
		return nums
	}
	// 单调队列, 存放的是索引
	deque := []int{}
	// 结果数组
	// res := make([]int, N-k+1)
	res := []int{}
	// 遍历nums数组
	for i := 0; i < N; i++ {
		// 这里需要保证新加入的数前面没有比其更小的数了
		for len(deque) != 0 && nums[deque[len(deque)-1]] <= nums[i] {
			// 将小的数移除队列尾部
			deque = deque[:len(deque)-1]
		}
		// 将该数加入队尾
		deque = append(deque, i)
		// 需要判断队首中的值是否有效, 队首是最大值的索引, 但是窗口是移动的
		// 每次移动的时候队首就不一定在下一个窗口中了
		if deque[0] <= i-k {
			// 将失效的数据从队首删除
			deque = deque[1:]
		}

		// 当窗口长度为k时, 保存当前窗口中的最大值
		if i+1 >= k {
			// res[i+1-k] = nums[deque[0]]
			res = append(res, nums[deque[0]])
		}
	}
	return res
}
