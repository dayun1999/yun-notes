# 并查集

**Created By: 蜜雪冰熊**

## 内容大纲

- **<a href="#并查集的基本介绍">1.并查集的基本介绍</a>**
- **<a href="#Quick Find的并查集">2.Quick Find的并查集</a>**
- **<a href="#Quick Union的并查集">3.Quick Union的并查集</a>**
- **<a href="#按秩合并的并查集">4.按秩合并的并查集</a>**
- **<a href="#路径压缩优化的并查集">5.路径压缩优化的并查集</a>**
- **<a href="#基于路径压缩的按秩合并优化的并查集">6.基于路径压缩的按秩合并优化的并查集</a>**





## <a name="并查集的基本介绍">1.并查集的基本介绍</a>

### 思考

如果给你一些顶点，并且告诉你每个顶点的连接关系，你如何才能快速的找出两个顶点是否具有连通性呢？如下图「图 5. 连通性问题」，该图给出了顶点与顶点之间的连接关系，那么，我们如何让计算机快速定位 (0, 3) , (1, 5), (7, 8) 是否相连呢？此时我们就需要机智的「并查集」数据结构了。很多地方也会称「并查集」为算法，在本 Leetbook 中，我们将其视为数据结构。

![](C:\Users\26646\Desktop\LeetCode刷题专题笔记\图\连通性问题.png)

### 并查集的常用术语

- 父节点——顶点的**直接**父亲节点, 比如上图的3的父节点是1, 顶点9的父节点是其本身
- 根节点——没有父节点的节点，3的根节点是0, 0的根节点是0



### 并查集的基本思想

略

### 并查集的编程思想

- 并查集的`find`函数
- 并查集的`union`函数

### 并查集的两个重要函数

- find
- union







## <a name="Quick Find的并查集">2.Quick Find的并查集</a>

<img src="C:\Users\26646\Desktop\LeetCode刷题专题笔记\图\quick_find原理图.png" style="zoom:80%;" />

`Quick Find`是将数组中的**元素对应了顶点的根节点**

实现代码如下:

```go
func main() {
	fmt.Println("Hello World")
	uf := New(10)
	uf.Union(1, 2)
	uf.Union(2, 5)
	uf.Union(5, 6)
	uf.Union(6, 7)
	uf.Union(3, 8)
	uf.Union(8, 9)
	fmt.Println(uf.connected(1, 5)) //
	fmt.Println(uf.connected(5, 7))
	fmt.Println(uf.connected(4, 9))

	uf.Union(9, 4)
	fmt.Println(uf.connected(4, 9))
}

// implement Quick-Find
type UnionFind struct {
	root []int
}

func New(size int) *UnionFind {
	root := make([]int, size, size)
	for i := range root {
		root[i] = i
	}
	return &UnionFind{
		root: root,
	}
}

func (uf *UnionFind) Find(x int) int {
	return uf.root[x]
}

func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	// 如果根节点不同
	if rootX != rootY {
		// 遍历整个数组
		for i := range uf.root {
			if uf.root[i] == rootY {
				uf.root[i] = rootX
			}
		}
	}
}

func (uf *UnionFind) connected(x, y int) bool {
	return uf.root[x] == uf.root[y]
}
```

### 时间复杂度

||`UnionFind`构造函数|`find`函数|`union`函数|`connnected`函数|
|:--|:--|:--|:--|:--|
|时间复杂度|O(N)|O(1)|O(N)|O(1)|
其中N=顶点个数








## <a name="Quick Union的并查集">3.Quick Union的并查集</a>

<img src="C:\Users\26646\Desktop\LeetCode刷题专题笔记\图\quick_union原理图.png" style="zoom:80%;" />

`Quick Union`将数组中的**元素对应了顶点的父节点/父顶点**

### 代码实现

```go
// implement Quick-Union
type UnionFind struct {
	root []int
}

func New(size int) *UnionFind {
	root := make([]int, size, size)
	for i := range root {
		root[i] = i
	}
	return &UnionFind{
		root: root,
	}
}

func (uf *UnionFind) Find(x int) int {
	for x != uf.root[x] {
		x = uf.root[x]
	}
	return x
}

func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	// 如果根节点不同
	if rootX != rootY {
		// 遍历整个数组
		uf.root[rootY] = rootX
	}
}

func (uf *UnionFind) connected(x, y int) bool {
	return uf.root[x] == uf.root[y]
}

```

### 时间复杂度

|| `UnionFind`构造函数 | `find`函数 |`union`函数|`connected`函数|
|:--|:--|:-- |:--|:--|
|时间复杂度|O(N)|O(H)|O(H)|O(H)|

N=顶点个数, H为树的高度



### 谁更好？`Quick Find` VS `Quick Union`

总体来说Quick Union更好， 但是两者都有很大的缺点！！！







## <a name="按秩合并的并查集">4.按秩合并的并查集</a>

> 【这里是对Quick Union里面的`Union`函数的优化】

小伙伴看到这里的时候，我们其实已经实现了 2 种「并查集」。但它们都有一个很大的缺点，这个缺点就是通过 union 函数连接顶点之后，可能所有顶点连成一条线形成「图 5. 一条线的图」，这就是我们 find 函数在最坏的情况下的样子。那么我们有办法解决吗？

当然，伟大的科学家已经给出了解决方案，就是按秩合并。这里的「秩」可以理解为「秩序」。之前我们在 union 的时候，我们是随机选择 x 和 y 中的一个根节点/父节点作为另一个顶点的根节点。但是在「按秩合并」中，我们是按照「某种秩序」选择一个父节点。

这里的「秩」指的是每个顶点所处的高度。我们每次 union 两个顶点的时候，选择根节点的时候不是随机的选择某个顶点的根节点，而是将「秩」大的那个根节点作为两个顶点的根节点，换句话说，我们将低的树合并到高的树之下，将高的树的根节点作为两个顶点的根节点。这样，我们就避免了所有的顶点连成一条线，这就是按秩合并优化的「并查集」。

<img src="C:\Users\26646\Desktop\LeetCode刷题专题笔记\图\按秩合并的并查集.png" style="zoom: 50%;" />

【总结】就是在Union的操作中确保合并的时候, 高度小的树服从高度大的树的管理，这里就是将小的树的根节点的父节点设置为高度大的树的根节点`root[rootOfMin] = rootOfMax`, 其中高度是存放在`rank`数组里面的

### 实现代码

```go
func main() {
	fmt.Println("Hello World")
	uf := New(10)
	uf.Union(1, 2)
	uf.Union(2, 5)
	uf.Union(5, 6)
	uf.Union(6, 7)
	uf.Union(3, 8)
	uf.Union(8, 9)
	fmt.Println(uf.Connected(1, 5)) //
	fmt.Println(uf.Connected(5, 7))
	fmt.Println(uf.Connected(4, 9))

	uf.Union(9, 4)
	fmt.Println(uf.Connected(4, 9))
}

// implementing 按秩合并并查集
type UnionFind struct {
	root []int // 存放顶点的数组
	rank []int // 存放顶点高度的数组
}

func New(size int) *UnionFind {
	root := make([]int, size, size)
	rank := make([]int, size, size)
	for i := range root {
		root[i] = i
		rank[i] = 1
	}
	return &UnionFind{root: root, rank: rank}
}

func (uf *UnionFind) Find(x int) int {
	// 只要索引不等于其值,就继续寻找
	for uf.root[x] != x {
		x = uf.root[x]
	}
	return x
}

func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	// 如果父节点不相同
	if rootX != rootY {
		// 按秩合并, 谁大谁做主
		if uf.rank[rootX] > uf.rank[rootY] {
			uf.root[rootY] = rootX
		} else if uf.rank[rootX] < uf.rank[rootY] {
			uf.root[rootY] = rootY
		} else {
			uf.root[rootY] = rootX
			uf.rank[rootX]++
		}

	}
}

func (uf *UnionFind) Connected(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}
```

### 时间复杂度

|| `UnionFind`构造函数 | `find`函数 |`union`函数|`connected`函数|
|:--|:--|:-- |:--|:--|
|时间复杂度|O(N)|O(logN)|O(logN)|O(logN)|







## <a name="路径压缩优化的并查集">5.路径压缩优化的并查集</a>

> 【这里是对Quick Union里面的`Find`函数的优化】

从前面的「并查集」实现方式中，我们不难看出，要想找到一个元素的根节点，需要沿着它的父亲节点的足迹一直遍历下去，直到找到它的根节点为止。如果下次再查找同一个元素的根节点，我们还是要做相同的操作。那我们有没有什么办法将它升级优化下呢？

答案是可以的！如果我们**在找到根节点之后，将所有遍历过的元素的父节点都改成根节点**，那么我们下次再查询到相同元素的时候，我们就仅仅只需要遍历两个元素就可以找到它的根节点了，这是非常高效的实现方式。那么问题来了，我们如何将所有遍历过的元素的父节点都改成根节点呢？这里就要拿出「递归」算法了。这种优化我们称之为「路径压缩」优化，它是对 find 函数的一种优化。

<img src="C:\Users\26646\Desktop\LeetCode刷题专题笔记\图\路径压缩优化的并查集2.png" style="zoom:80%;" />

### 代码实现

```go
type UnionFind struct {
	root []int
}

func New(size int) *UnionFind {
	root := make([]int, size, size)
	for i := range root {
		root[i] = i
	}
	return &UnionFind{root: root}
}

// 递归
func (uf *UnionFind) Find(x int) int {
	if uf.root[x] == x {
		return x
	}
	uf.root[x] = uf.Find(uf.root[x])
	return uf.root[x]
}

func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	if rootY != rootX {
		uf.root[rootY] = x
	}
}

func (uf *UnionFind) Connected(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}
```

### 时间复杂度

|| `UnionFind`构造函数 | `find`函数 |`union`函数|`connected`函数|
|:--|:--|:-- |:--|:--|
|时间复杂度|O(N)|O(logN)|O(logN)|O(logN)|







## <a name="基于路径压缩的按秩合并优化的并查集">6.基于路径压缩的按秩合并优化的并查集 </a>

这个优化就是将「路径压缩优化」和「按秩合并优化」合并后形成的「并查集」的实现方式。

```go
func main() {
	fmt.Println("Hello World")
	uf := New(10)
	uf.Union(1, 2)
	uf.Union(2, 5)
	uf.Union(5, 6)
	uf.Union(6, 7)
	uf.Union(3, 8)
	uf.Union(8, 9)
	fmt.Println(uf.Connected(1, 5)) //
	fmt.Println(uf.Connected(5, 7))
	fmt.Println(uf.Connected(4, 9))

	uf.Union(9, 4)
	fmt.Println(uf.Connected(4, 9))
}

type UnionFind struct {
	root []int
	rank []int
}

func New(size int) *UnionFind {
	root := make([]int, size, size)
	rank := make([]int, size, size)
	for i := range root {
		root[i] = i
		rank[i] = 1
	}
	return &UnionFind{root: root, rank: rank}
}

// 此处的find函数与路径压缩优化版本的find函数一样
func (uf *UnionFind) Find(x int) int {
	// 结束条件: 找到根节点
	if uf.root[x] == x {
		return x
	}
	uf.root[x] = uf.Find(uf.root[x])
	return uf.root[x]
}

// 按秩合并优化的union函数
func (uf *UnionFind) Union(x, y int) {
	//先找到根节点
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	// 判断两个树的高度, 矮的树服从高的树
	if rootX != rootY {
		if uf.rank[rootX] > uf.rank[rootY] {
			uf.root[rootY] = rootX
		} else if uf.rank[rootX] < uf.rank[rootY] {
			uf.root[rootX] = rootY
		} else {
			uf.root[rootY] = rootX
			uf.rank[rootX]++
		}
	}
}

func (uf *UnionFind) Connected(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}
```

### 时间复杂度

|| `UnionFind`构造函数 | `find`函数 |`union`函数|`connected`函数|
|:--|:--|:-- |:--|:--|
|时间复杂度|O(N)|O(α(N))|O(α(N))|O(α(N))|



## <a name="并查集数据结构总结">并查集数据结构总结</a>

在「并查集」数据结构中，其中心思想是将所有连接的顶点，无论是直接连接还是间接连接，都将他们指向同一个父节点或者根节点。此时，如果要判断两个顶点是否具有连通性，只要判断它们的根节点是否为同一个节点即可。

在「并查集」数据结构中，它的两个灵魂函数，分别是 find和 union。find 函数是为了找出给定顶点的根节点。 union 函数是通过更改顶点根节点的方式，将两个原本不相连接的顶点表示为两个连接的顶点。对于「并查集」来说，它还有一个重要的功能性函数 connected。它最主要的作用就是检查两个顶点的「连通性」。find 和 union 函数是「并查集」中必不可少的函数。connected 函数则需要根据题目的意思来决定是否需要。

### 并查集的刷题小技巧

「并查集」的代码是高度模版化的。所以作者建议大家熟记「并查集」的实现代码，这样小伙伴们在遇到「并查集」的算法题目的时候，就可以淡定的应对了。作者推荐大家在理解的前题下，请熟记「基于路径压缩+按秩合并的并查集」的实现代码。



## <a name="">练习题：省份数量</a>

【并查集、DFS、BFS】

![](C:\Users\26646\Desktop\LeetCode刷题专题笔记\图\练习题-547-省份数量.png)

### 代码实现

```go
func findCircleNum(isConnected [][]int) int {
    size := len(isConnected)
    res := 0
    uf := New(size)
    for i := 0; i < size; i++ {
    	for j := 0; j < size; j++ {
			if isConnected[i][j] == 1 && {
				uf.Union(i, j)
			}
		}
    }
    // 找到所有不同的根节点, 即为所求
    // 下面这个解法就是错的, 因为写的代码是基于Quick Union的
    // 而Quick Union的数组元素代表的是父节点, 不是Quick Find的根节点
	//    m := make(map[int]int)
	//    for _, v := range uf.root {
	//    	if _, ok := m[v]; !ok {
	//			m[v]++
	//		}
	//    }
	for i, p := range uf.root {
		if i == p {
			res++
		}
	}
    return res
}

type UnionFind struct {
	root []int
	rank []int
}

func New(size int) *UnionFind {
	root := make([]int, size, size)
	rank := make([]int, size, size)
	for i := range root {
		root[i] = i
		rank[i] = 1
	}
	return &UnionFind{root: root, rank: rank}
}

// 此处的find函数与路径压缩优化版本的find函数一样
func (uf *UnionFind) Find(x int) int {
	// 结束条件: 找到根节点
	if uf.root[x] == x {
		return x
	}
	uf.root[x] = uf.Find(uf.root[x])
	return uf.root[x]
}

// 按秩合并优化的union函数
func (uf *UnionFind) Union(x, y int) {
	//先找到根节点
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	// 判断两个树的高度, 矮的树服从高的树
	if rootX != rootY {
		if uf.rank[rootX] > uf.rank[rootY] {
			uf.root[rootY] = rootX
		} else if uf.rank[rootX] < uf.rank[rootY] {
			uf.root[rootX] = rootY
		} else {
			uf.root[rootY] = rootX
			uf.rank[rootX]++
		}
	}
}
```


