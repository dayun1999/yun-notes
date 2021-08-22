## 题目



## 分析



## 解答

```go
type Trie struct {
	children [26]*Trie
	word string // 可以理解为一个分支的结束符
}

func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.word = word
}

// 注意前缀树的根节点啥也没有
func buildTree(words []string) *Trie {
	root := &Trie{}
	for _, w := range words {
		root.Insert(w)
	}
	return root
}

func findWords(board [][]byte, words []string) []string {
	M, N := len(board), len(board[0])
	res := []string{}
	
	var dfs func(int, int, *Trie)
	dfs = func(i, j int, node *Trie) {
		// 结束条件
		// 1.越界
		if !(i >= 0 && i < M && j >= 0 && j < N) {
			return
		}
		ch := board[i][j]
        // 2. 已经被访问过
        if ch == '#' {
            return
        }
		// 3.节点为空
		if node.children[ch-'a'] == nil {
            return 
        }
		node = node.children[ch-'a']
		// 4.找到节点
		if node.word != "" {
			res = append(res, node.word)
			node.word = "" // 去重
            // return // 错误, 为什么不能return，因为ab是一种结果, abc又是另一个结果, 所以return就检查不到abc这个单词的存在了
 		}

         // 标记当前格子被访问过
         board[i][j] = '#'
         defer func() { board[i][j] = ch } ()
		
		// 四个方向遍历
		dfs(i-1, j, node)
		dfs(i+1, j, node)
		dfs(i, j-1, node)
		dfs(i, j+1, node)
 	}
	
    // 构造前缀树
    root := buildTree(words)
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			dfs(i, j, root)
		}
	}
	return res
}
```

