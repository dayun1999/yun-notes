// 力扣169 多数元素
package leetcode169

// 超过n/2的数最多只能有1个
func majorityElement(nums []int) int {
	count, candidates := 0, 0
	for _, v := range nums {
		if count == 0 {
			candidates = v
		}
		if candidates == v {
			count++
		} else {
			count--
		}
	}
	return candidates
}
