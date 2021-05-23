# shell-notes-02

## 内容大纲

- *find命令*
- *玩转xargs*
- *tr命令*
- *校验和 (checksum)*
- *加密工具与散列*
- *行排序*

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

## 4. 校验和 (checksum)

校验和用来检查文件是否发生改变，有助于核实下载文件的完整性等等, 常见的校验和算法: `MD5`和`SHA-1`

### 4.1、MD5算法计算校验和-- md5sum 命令

#### PS: md5sum是一个 *32* 个字符的十六进制串

```bash
# MD5算法对应命令--md5sum
# 1. 对单个文件计算校验和
>>> md5sum cecho.sh 
5fe499be53b0d7c2b1ba0d3cec9bf653  cecho.sh

# 2. 对多个文件分别计算校验和
>>> md5sum cecho.sh hello.txt out.txt 
5fe499be53b0d7c2b1ba0d3cec9bf653  cecho.sh
5ebc7480a6da0d09d93879da9b71d707  hello.txt
dc8a89c05151ecdac50718932cd371c8  out.txt

# 3. 使用-c选项核实数据的完整性
>>> md5sum wdy.sh > file_sum.md5
# 中途修改wdy.sh的内容
# 然后核原来的校验和校验,发现不一样
>>> md5sum -c file_sum.md5 
wdy.sh: FAILED
md5sum: WARNING: 1 computed checksum did NOT match

```

### 4.2、SHA-1计算校验和-- sha1sum 命令

#### PS: sha1sum是一个 *40* 个字符的十六进制串

命令格式和`md5sum`一样

```bash
# SHA-1算法对应命令--sha1sum
# 对单个文件计算校验和
>>> sha1sum cecho.sh 
9370ff4e182174cb9a076226b50f8df4497a4c00  cecho.sh

# 对多个文件分别计算校验和
>>> sha1sum cecho.sh hello.txt out.txt 
9370ff4e182174cb9a076226b50f8df4497a4c00  cecho.sh
ba6f7a639cbac7e9c6918d03a4feefc1c40e62e5  hello.txt
b2bff929f414e8763495ad5834de8afdaade73d1  out.txt

```

### 4.3、 对目录进行校验

也即对目录中的所有文件计算校验和

#### `md5deep`命令的使用(使用之前先安装)

```bash
# -r代表递归
# -l代表使用相对路径,md5deep默认使用绝对路径
>>> sudo md5deep shell_learning/ -rl > directory.md5
>>> cat directory.md5
892020bedc3b7110c475bf43246b64ad  shell_learning/all_sh_files.txt
02e47d97d9071aa664e083fee26dde2a  shell_learning/wdy.sh
3a0666c07bad1a857d8359a70dedb685  shell_learning/a3
cc5e9216b37b3f82dfe398ed311d16e6  shell_learning/a2
...
...
d41d8cd98f00b204e9800998ecf8427e  shell_learning/*.txt
dc8a89c05151ecdac50718932cd371c8  shell_learning/out.txt

```

### 4.4、MD5和SHA-1已不再安全

MD5和SHA-1都是单向散列算法，均无法推出原始数据,由于计算能力的提升，上述加密算法已经不再安全，目前更推荐使用`bcrypt` 和 `sha512sum`这样的工具进行加密的 \
推荐阅读: [加密、散列和加盐的区别](https://www.bisend.cn/blog/difference-encryption-hashing-salting)

## 5. 加密工具与散列

加密算法和校验和算法不一样的是加密算法可以无损的重构数据 \
Linux中常见的加密算法有`bcrypt` `base64` `gpg`

## 6. 行排序

对文本进行按行排序

### 6.1、`sort`命令和`uniq`命令

- #### 排序一组文件

```bash
>>> sort file1.txt file2.txt > sorted.txt

# 也可写成
>>> sort file1.txt file2.txt -o sorted.txt
```

- #### 按照数字排序

```bash
>>> sort -n file.txt
```

- #### 逆序排序

```bash
>>> sort -r file.txt
```

- #### 更多用法见`man sort`
  
