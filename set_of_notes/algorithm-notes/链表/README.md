# 链表的一些操作

## 头插法和尾插法

头插法

```go
nodeNew.Next = cur.Next
cur.Next = nodeNew
```

尾插法--将元素插在末尾

```go
cur.Next = nodeNew
cur = cur.Next
cur.Next = nil
```



## 如何找中点

### 1. `1->2->3->4->nil`中找到3

```go
// 找中点-如果是偶数, 那么就是后半部分的第一个节点[1 2 3 4]里面的3
func getMiddle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
    // 关键在for循环怎么写
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
```



### 2.`1->2->3->4->nil`中找到2

```go
func getMiddle1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
    // 关键在for循环怎么写
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
```

