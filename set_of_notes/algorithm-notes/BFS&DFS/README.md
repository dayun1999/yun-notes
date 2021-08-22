# BFS & DFS



## 两种算法的比对以及使用场景

### 数据结构的对比

- `DFS`采用栈，`BFS`采用队列

### 效率的对比

- 深度优先搜素算法：不全部保留结点，占用空间少；有回溯操作(即有入栈、出栈操作)，**运行速度慢**
-  广度优先搜索算法：保留全部结点， 占用空间大； 无回溯操作(即无入栈、出栈操作)，**运行速度快**

 

`DFS`不全部保留结点，扩展完的结点从数据库中弹出删去，这样，一般在数据库中存储的结点数就是深度值，因此它占用空间较少。

所以，当搜索树的结点较多，用其它方法易产生内存溢出时，深度优先搜索不失为一种有效的求解方法。 　

` BFS`，一般需存储产生的所有结点，占用的存储空间要比深度优先搜索大得多，因此，程序设计中，必须考虑溢出和节省内存空间的问题。

但广度优先搜索法一般无回溯操作，即入栈和出栈的操作，所以运行速度比深度优先搜索要快些