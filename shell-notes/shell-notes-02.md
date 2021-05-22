# shell-notes-02

## 内容大纲

- *find命令*
- *玩转xargs*
- *tr命令*

## 1.find命令

## 2.玩转xargs

> `xargs` 接收stdin作为主要的数据源, 将从stdin读取的数据作为指定命令的参数并执行命令

### 需要留意的事

- `xargs` 命令应该紧跟在管道操作符('|')之后\
- `xargs`默认执行echo命令 \

### 举例

#### 1、在一组txt文件中搜索指定字符

```bash
# 在一组txt文件中搜索指定字符
>>> ls *.txt | xargs grep hello
```

#### 2、`-n` 选项可以限制每次调用命令时用到的参数个数

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

#### 3、`-I`选项可以替换字符串

```bash
>>> cat args.txt | xargs -I {} ./cecho.sh -p {} -1

-p arg1 -1#
-p arg2 -1#
-p arg3 -1#
```

#### 4、`-d`选项可以指定分隔参数

```bash
>>> echo "helloXhelloXhelloX" | xargs -d X
hello hello hello 

```

> find命令里面的`-print0` 选项可以指定用0(NULL)来分隔查找到的参数, 然后再用xargs里面对应的`-0`选项来进行解析,具体例子出现在下文的代码中

#### 5、与find结合使用

```bash
# 删除当前目录下所有txt文件
>>> find . -type f -name "*.txt" -print0 | xargs -0 rm -f
```

#### 6、统计源代码目录中所有C程序文件的行数

```bash
# 统计要搜索的目录下的所有C语言文件的代码行数
>>> find <source_code_dir_path> -type f -iname "*.c" -print0 | xargs -0 wc -l
```

## 3. tr命令

### 需要留意的事情

- 命令格式: `tr [options] set1 set2`, 意思是从输入接收到的字符会按照对应位置从set1映射到set2
- tr只能通过stdin接收输入，不能通过命令行参数接收

#### 1、将输入的字符中的大写全部转为小写

```bash
>>> echo "HELLO WORLD" | tr 'A-Z' 'a-z'
 hello world

```

#### 2、ROT13加密算法

```bash
>>> echo "tr came, tr saw, tr conquered." | tr 'a-zA-Z' 'n-za-mN-ZA-M' 
ge pnzr, ge fnj, ge pbadhrerq.

```

#### 3、`-d`选项用来删除字符

```bash
>>> echo "Hello 124 world 456, wdy001" | tr -d '0-9'
Hello  world , wdy

```

#### 4、`-c`选项与字符组补集

```bash
>>> echo "Hello 1 C 2 world 3" | tr -d -c '0-9\n' 
123

```

#### 5、`-s`选项与字符压缩

```bash
# 1.压缩空格
>>> echo "Hello          world." | tr -s ' '
Hello world.

# 2.删除多余的换行符
>>> cat multi_blanks.txt | tr -s '\n'
line 1
line 2
line 3
line 4

# 3.将文件中的数字列表相加
>>> cat sum.txt 
1
2
3
4
5
# 下面代码中$[ oprtation ]执行算术运算,echo $[ $(tr '\n' '+') 0 ]相当于echo $[ 1+2+3+4+5+0 ]
>>> cat sum.txt | echo $[ $(tr '\n' '+') 0 ]
15

# 4.在包含数字和字母的文件中计算数字之和
>>> cat test.txt | tr -d 'a-z' | echo "total: $[ $(tr ' ' '+') ]"
total: 6

```

### tr 中会用到的字符类

|名称|含义|
|:--|:--|
|alnum|字母和数字|
|alpha|字母|
|cntrl|控制字符|
|digit|数字|
|graph|图形字符|
|lower|小写字母|
|upper|大写字母|
|print|可打印符号|
|space|空白字符|
|xdigit|十六进制字符|
|punct|标点符号|

> 使用如下: tr '[:lower:]' '[:upper:]'
