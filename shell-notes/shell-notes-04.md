# shell-notes-04-让文本飞

## 内容大纲

- *正则表达式*
- *`grep`命令搜索文本*
- *`cut`命令切分文本*
- *`sed`命令替换文本*
- *`awk`命令进行高级文本处理*

## 1. 正则表达式

|正则表达式|描述|示例|
|:--|:--|:--|
|^|指定想要匹配的文本的*首部*|`^he`能够匹配以he起始的行|
|$|指定想要匹配的文本的*尾部*|`he$`能够匹配以he结尾的行|
|.|匹配任意*一个*字符|he.只能匹配her和hey, 不能匹配hery, 只匹配单个字符|
|[]|匹配括号[]中的任意*一个*字符|he[ry]只能匹配her或者hey|
|[^]|匹配不在括号[]中的任意一个字符|9[^01]只能匹配92和93, 不能匹配90和91|
|?|匹配之前的项1次或0次|colou?r只能匹配colour或者color|
|+|匹配之前的项1次或多次|略|
|*|匹配之前的项0次或者多次|略|
|{n}|匹配n次|略|
|{n,}|匹配至少n次|略|
|{n,m}|匹配的最少次数为n, 最大次数为m|略略略|
|()|括号中的内容视为一个整体|ma(tri)?x能够匹配max或者matrix|
|`\|`|选择结构,可以匹配`\|`两边的任意一项|Oct (1st \| 2nd)能够匹配Oct 1st或者Oct 2nd|
|`\`|转义字符,将特殊字符的特殊意义去掉|原本a.b能够匹配acb或者afb之类的(因为`.`代表匹配任意一个字符)，但是`a\.b`只能匹配a.b|

感兴趣的戳这`:point_right:`[可视化正则表达式网站](https://regexper.com/)

## 2. `grep`命令搜索文本

|常用选项|描述|举例|
|:--|:--|:--|
|`--color`|输出行中着重标记匹配到的模式,`--color`的值只有auto,never,always|`grep --color=never wdy hello.txt`|
|`-E`|可以使grep使用扩展正则表达式|`grep -E " [a-z]{3} " hello.txt`|
|`-o`|只输出匹配到的文本||
|`-v`|打印不匹配的所有行,`v`指代invert||
|`-c`|统计出匹配模式的文本行数(注意不是匹配次数,是行数)|`grep -c "text" filename`|
|`-n`|打印出匹配内容所在行数||
|`-b`|打印出匹配出现在行中的偏移,配合`-o`可以打印出匹配所在的字符或者字节偏移||
|`-l`|列出匹配模式所在的文件|`grep hello . -Rl`在当前目录搜索存在hello字符串的所有文件|
|`-L`|和`-l`相反||
|`-R`或者`-r`|递归搜索|`grep "text" . -R -n`|
|`-i`|忽略大小写||
|`-e`|匹配多个模式|`grep -e "pattern1" -e "pattern2"`|
|`--include`|搜索指定文件||
|`--exclude`|排除指定文件||
|`--exclude-dir`|排除目录||
|`-f`|当匹配模式多了就可以定义在一个文件中, `-f` 选项后面跟的就是pattern file|`grep -f pattern_file source_filename`|
|`-q`|静默输出,就是不看打印的搜索结果||
|`-A`|打印匹配文本之前的行|`seq 10 \| grep 5 -A 3`打印匹配到的5的前三行内容|
|`-B`|打印匹配文本之后的行|`seq 10 \| grep 5 -B 3`打印匹配到的5的后三行内容|
|`-C`|打印匹配的内容以及其前三行和后三行的内容||

## 3. `cut`命令切分文本

cut命令可以按列切分文件, 对于处理固定宽度字段的文件、CSV文件或是由空格分隔的文件都十分方便

|常用选项|描述|例子|
|:--|:--|:--|
|`-f`|指定要提取的字段,`-f1,3`显示第一列和第三列|`cut -f field_list filename`|
|`--complement`|显示没有被`-f`指定的那些字段,等同于补集||
|`-d`|设置分隔符|`cut -f2 -d ";" data.txt`|
|`-c N-M`|将切分字段指定为N-M范围内的字符|`cut -c2-5 range_fields.txt`表示打印第2-5个字符|
|`-c -N`|将切分字段指定为前N个字符|`cut -c -2 range_fields.txt`表示打印前2个字符|
|`-c N-`|将切分字段指定为后N个字符|`cut -c 10- range_fields.txt`|
|`-b`|和`-c`类似, 只不过这里是表示字节(byte)||
|`--output-delimiter`|指定输出分隔符|`cut -c10-20,24-31 range_fields.txt --output-delimiter "--"`|

## 4. `sed`命令替换文本

sed是`stream editor`的缩写,常用与替换文本

```bash
# 使用语法
# pattern表示原文本中要被替换的文本模式
# replace_string表示替换的最终结果
sed 's/pattern/replace_string/' file

# 举例
>>> echo thisthisthis | sed 's/this/THIS/'
THISthisthis
```

|常用标记|描述|示例|
|:--|:--|:--|
|`-i`|该选项会用修改后的数据替换原文件|`sed -i 's/text/replace/' file`|
|`-i.bak`|该选项会用修改后的数据替换原文件但是同时会备份原文件,文件名为file.bak|`sed -i.bak 's/text/replace/' file`|
|`g`|该标记会进行全局替换|`sed 's/pattern/replace_string/g' file`|
|`/#g`|`#`代表数字,表示替换第N次出现的匹配|`sed 's/pattern/replace_string/2g' file`|
|`d`|表示不进行替换,而是将匹配到的行直接删除|`sed '/^$/d' file`表示移除文件中的空行|
|`&`|对应已经匹配到的字符串|`echo this is apple \| sed 's/\w\+/[&]/g'`|
|`\#`|`#`代表数字,比如\1代表匹配到的第一个字符串|`echo this is digit 7 \| sed 's/digit \([0-9]\)/\1/'`输出结果为`this is 7`|
|`-e`|组合多个表达式|`echo abc \| sed -e 's/a/A/' -e 's/c/C/'`也可以不适用`-e`选项,等同于`echo abc \| sed 's/a/A/;s/c/C/'`|

### sed命令的向后引用

```bash

```

### 补充

sed命令中表达式通常使用`''`单引号来引用,但是当想在表达式中使用变量的时候就可以使用双引号

```bash
❄[wdy ~/shell_learning]>>> echo hello world | sed "s/$text/HELLO/"
HELLO world
```

## 5. `awk`命令进行高级文本处理

awk可以处理数据流, 它支持关联数组、递归函数、条件语句等功能\
awk是以逐行的形式来处理文件的

```bash
# awk的脚本结构如下
awk 'BEGIN{ print "start" } pattern { commands } END{ print "end" }' file
```

1. 首先执行`BEGIN { commands }`,
2. 接着读取文件中的一行，如果匹配pattern，则执行随后的commands,
3. 当读取至流末尾的时候, 执行END { commands}语句块

模式是可选的, 不提供模式则认为所有行都是匹配的;\
和模式关联的语句块也是可选的, 如果不提供语句块, 则默认执行`{ print }`,即打印每一行

|特殊变量/选项|描述|示例|
|:--|:--|:--|
|`NR`|表示记录编号, 当awk将行作为记录时, NR相当于当前行号||
|`NF`|表示字段数量, 在处理当前记录时, 相当于字段数量；默认字段分隔符是空格||
|`$0`|当前记录(当前行)的文本内容||
|`$1`|第一字段的文本内容||
|`$2`|第二字段的文本内容||
|`-v`|借助`-v`可以将外部变量传递给awk|`VAR=100; echo \| awk -v variable=$VAR '{ print variable }'`|
|`-F`|设置字段分隔符|`awk -F: '{ print $NF }' /etc/passwd`  等同于`awk 'BEGIN { FS=":" } { print $NF }' /etc/passwd`|
|`OFS`|设置输出字段分隔符|`OFS="delimiter"`|

### 使用getline读取行

语法为: `getline var`,变量var包含了特定行;\
如果调用getline不含参数, 则可以用`$0、$1、$2`访问文本内容

```bash
❄[wdy ~/shell_learning]>>> seq 10 | awk 'BEGIN{ getline; print "Read ahead 1st line", $0 } { print $0 }'
Read ahead 1st line 1
2
3
4
5
6
7
8
9
10
```

### 使用过滤模式对awk处理的行进行过滤

```bash
# 行号小于5的行
>>> seq 10 | awk 'NR < 5'
1
2
3
4

# 行号在2-8之间的行
>>> seq 10 | awk 'NR==2, NR==8'
2
3
4
5
6
7
8

# 包含模式为linux的行
>>> echo -e "line1 unix\nline2 linux" | awk '/linux/'
line2 linux

# 不包含模式linux的行
>>> echo -e "line1 unix\nline2 linux" | awk '!/linux/'
line1 unix
```

### 使用循环

```bash
for(i=0;i<10;i++) { print $i ; }
# 还支持列表形式的for循环
for(i in array) { print array[i]; }
```

### awk内建的字符串处理函数

|函数名称|描述|
|:--|:--|
|`length(string)`|返回字符串string的长度|
|`index(string, search_string)`|返回search_string在字符串string中出现的位置|
|`split(string, array, delimiter)`|以delimiter为分隔符,分割string, 将生成的字符串存入数组array|
|`substr(regex, start_position, end-position)`|返回子串|
