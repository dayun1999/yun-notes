# 排序代码汇总

- <a href="#merge_sort">归并排序</a>
- <a href="#quick_sort">快速排序</a>
- <a href="#heap_sort">堆排序</a>



## <a name="merge_sort">归并排序</a>

```go
func MergeSort(arr []int) []int {
	// write code here
	// 既然是递归就要有结束条件
	if len(arr) <= 1 {
		return arr
	}
	// 找到中点
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	return merge(left, right)
}

// 归并排序
func merge(l1, l2 []int) []int {
	res := []int{}
	i, j := 0, 0
	for i < len(l1) && j < len(l2) {
		if l1[i] < l2[j] {
			res = append(res, l1[i])
			i++
		} else {
			res = append(res, l2[j])
			j++
		}
	}
	for i < len(l1) {
		res = append(res, l1[i])
		i++
	}
	for j < len(l2) {
		res = append(res, l2[j])
		j++
	}
	return res
}
```



## <a name="quick_sort">快速排序</a>

```go
func quickSort(a []int, low, high int) {
    if low < high {
        // 找到基准
        index := partition(a, low, high, low)
        quickSort(a, low, index-1)
        quickSort(a, index+1, high)
    }
}

func partition(a []int, low, high, pivotIndex int) int {
    pivotValue := a[pivotIndex]
    // 将轴值和最后一个交换
    a[high], a[pivotIndex] = a[pivotIndex], a[high]
    storeIndex := low

    for i := low; i <= high-1; i++ {
        if a[i] <= pivotValue {
            a[i], a[storeIndex] = a[storeIndex], a[i]
            storeIndex++
        }
    }
    a[high], a[storeIndex] = a[storeIndex], a[high]
    return storeIndex
}
```



## <a name="heap_sort">堆排序</a>

```go

```

