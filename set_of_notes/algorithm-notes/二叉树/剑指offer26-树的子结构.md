## 题目

#### [剑指 Offer 26. 树的子结构](https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof/)

```go
输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)

B是A的子结构， 即 A中有出现和B相同的结构和节点值。

例如:
给定的树 A:

     3
    / \
   4   5
  / \
 1   2
给定的树 B：

   4 
  /
 1
返回 true，因为 B 与 A 的一个子树拥有相同的结构和节点值。
```

## 分析

递归

## 解答[已修改]

```go
func isSubStructure(A *TreeNode, B *TreeNode) bool {
    if A == nil || B == nil { return false }
    return isSub(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func isSub(A *TreeNode, B *TreeNode) bool {
    if B == nil {
        return true
    }
    if A == nil || B == nil {
        return false
    }
    // 这边已经修改
    if A.Val != B.Val  {
        return false
    }
    return isSub(A.Left, B.Left) && isSub(A.Right, B.Right)
}
```

