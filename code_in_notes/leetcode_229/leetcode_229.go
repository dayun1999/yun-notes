// 229 求众数II
// 多数投票升级版：
// 超过n/3的数最多只能有两个。先选出两个候选人A,B。 遍历数组，分三种情况：

// 1.如果投A（当前元素等于A），则A的票数++;
// 2.如果投B（当前元素等于B），B的票数++；
// 3.如果A,B都不投（即当前与A，B都不相等）,那么检查此时A或B的票数是否减为0：

// 遍历结束后选出了两个候选人，但是这两个候选人是否满足>n/3，还需要再遍历一遍数组，找出两个候选人的具体票数。
package leetcode229

func majorityElement(nums []int) []int {
	N := len(nums)
	if N == 0 {
		return nil
	}
	res := []int{}

	// 记住: 超过n/3的候选人最多有两个
	// 记为A和B
	candidateA, candidateB := nums[0], nums[0]
	countA, countB := 0, 0

	// 遍历数组
	for _, num := range nums {
		// 情况1-如果当前的候选人也是A
		if num == candidateA {
			countA++
			continue
		}
		// 情况2-如果当前的候选人也是B
		if num == candidateB {
			countB++
			continue
		}
		// 情况3-如果都不是, 判断countA和countB谁为0
		if countA == 0 {
			candidateA = num
			countA++
			continue
		}
		if countB == 0 {
			candidateB = num
			countB++
			continue
		}
		// 情况3.1-如果两个候选人的票数不为0, 则都减去1
		if countA != 0 && countB != 0 {
			countA--
			countB--
		}
	}

	// 得到两个候选人了, 检查一遍是否他们的和大于n/3
	countA, countB = 0, 0
	for _, num := range nums {
		if num == candidateA {
			countA++
		} else if num == candidateB {
			countB++
		}
	}
	if countA > N/3 {
		res = append(res, candidateA)
	}
	if countB > N/3 {
		res = append(res, candidateB)
	}
	return res
}
