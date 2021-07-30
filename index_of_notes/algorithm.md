# 数据结构设计与实现

**Created By: 蜜雪冰熊**
## 大纲

- **<a href="#data_structure">数据结构设计与实现</a>**
  - <a href="#lru">LRU缓存</a>
  - <a href="#trieTree">前缀树Trie Tree</a>

## <a name="data_structure">数据结构设计与实现</a>

#### <a name="lru">1.LRU缓存</a>

```go
// LRU缓存
type LRUCache struct {
	size       int                       // 当前的缓存大小
	capacity   int                       // 缓存容量
	cache      map[int]*DoubleLinkedNode // 哈希链表
	head, tail *DoubleLinkedNode         // 链表的头部和尾部
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache:    make(map[int]*DoubleLinkedNode),
		head:     initDoubleLinkedNode(0, 0),
		tail:     initDoubleLinkedNode(0, 0),
		capacity: capacity,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

// 双向链表
type DoubleLinkedNode struct {
	key, value int
	prev, next *DoubleLinkedNode
}

func initDoubleLinkedNode(key, value int) *DoubleLinkedNode {
	return &DoubleLinkedNode{
		key:   key,
		value: value,
	}
}

// 查询操作
func (lru *LRUCache) Get(key int) int {
	// 检查key在不
	node, ok := lru.cache[key]
	if ok {
		lru.moveToHead(node)
		return node.value
	}
	return -1
}

// 插入/更新操作
func (lru *LRUCache) Put(key int, value int) {
	// 先检查缓存里面有没有同一个key
	if _, ok := lru.cache[key]; !ok {
		node := initDoubleLinkedNode(key, value)
		lru.cache[key] = node
		lru.addToHead(node) // 新来的总是最近最常使用, 所以放到链表头部
		lru.size++
		// 检查是否大于缓存了
		if lru.size > lru.capacity {
			// 去掉最不常使用的数据
			removed := lru.removeTail()
			delete(lru.cache, removed.key)
			lru.size--
		}
	} else {
		// 如果有, 就更新值
		node := lru.cache[key] // 先取出key所在的节点
		node.value = value     // 更新
		lru.moveToHead(node)   // 放到开头, 这是常用的数据
	}
}

// 双向链表的操作
// 添加头结点
func (lru *LRUCache) addToHead(node *DoubleLinkedNode) {
	node.prev = lru.head
	node.next = lru.head.next
	lru.head.next.prev = node
	lru.head.next = node
}

// 删除节点
func (lru *LRUCache) removeNode(node *DoubleLinkedNode) {
	node.prev.next = node.next // 前一节点指向我的下一节点
	node.next.prev = node.prev // 后一节点指向我的前一节点
    // 清除指针的内容, 可写可不写
	// node.next = nil
	// node.prev = nil
}

// 移到头结点
func (lru *LRUCache) moveToHead(node *DoubleLinkedNode) {
    lru.removeNode(node)
	lru.addToHead(node)
}

// 移除尾结点
func (lru *LRUCache) removeTail() *DoubleLinkedNode {
	// lru其实已经保存尾结点了
	node := lru.tail.prev
	lru.removeNode(node)
	return node
}
```

#### 2. <a name="trieTree">前缀树</a>

```go
// 数组实现前缀树
type Trie struct {
    children [26]*Trie	// 二叉树有两个子树,这里是多叉树
    isEnd bool	// 是否到了一个单词的结尾
}


/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}


/** Inserts a word into the trie. */
func (t *Trie) Insert(word string)  {
	node := t
	for  _, ch := range word {
		// 根据字母找到索引
		ch -= 'a' 
		// 如果孩子为空, 就新建孩子节点
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}


/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	node := t.SearchPrefix(word)
	return node != nil && node.isEnd
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	return t.SearchPrefix(prefix) != nil
}

func (t *Trie) SearchPrefix(prefix string) *Trie {
	node := t
	for _, ch := range prefix {
		ch -= 'a'
		if node.children[ch] == nil {
			return nil
		}
		node = node.children[ch]
	}
	return node
}
```