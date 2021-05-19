# shell-notes-02

## 内容大纲

- *find命令*
- *玩转xargs*

## 1.find命令

## 2.玩转xargs

> `xargs` 接收stdin作为主要的数据源, 将从stdin读取的数据作为指定命令的参数并执行命令

### 需要留意

`xargs` 命令应该紧跟在管道操作符('|')之后\
`xargs`默认执行echo命令 \

### 举例

#### 在一组txt文件中搜索指定字符

```bash
# 在一组txt文件中搜索指定字符
>>> ls *.txt | xargs grep hello
```

#### `-n` 选项可以限制每次调用命令时用到的参数个数

```bash
>>> cat examples.txt
1 2 3 4 5 6 7
8 9 10 11
12 13
14

>>> cat examples.txt | xargs -n 3
1 2 3
4 5 6
7 8 9
10 11 12
13 14
```
