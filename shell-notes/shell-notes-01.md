# shell-notes-01

## 内容大纲

- 自定义linux提示符
- 别名
- tput命令和stty命令
- 命令行参数
- read命令的用法

## 1.自定义linux提示符

在$HOME目录下,找到`.bashrc`文件并编辑,修改(如果没有则添加)PS1如下

```bash
# 普通修改
PS1="\u \w >>>"
```

效果如图所示:
![图片](https://github.com/code4EE/images/blob/main/20210516190532.png)

```bash
# 添加颜色
PS1='❄\[\e[1;37m\][\[\e[1;36m\]\u \[\e[1;35m\]\w\[\e[1;37m\]]\[\e[1;31m\]>\[\e[1;33m\]>\[\e[1;32m\]>'
```

效果如图所示:
![图片](https://github.com/code4EE/images/blob/main/20210516182709.png)

### 解释说明各个符号的意义

- `❄`  没什么意义,只想演示编辑支持unicode字符
- `\e[1:31m` 修改字体颜色,其中数字31代表红色,可修改为其他颜色,黑色=30,绿色=32,黄色=33,蓝色=34,洋红=35,青色=36,白色=37
- `\[`和`\]`  不键入该两个字符你会发现输入长命令的时候不会自动换行而且linux prompt会被blackspace键删掉
[感兴趣的可以看这个](https://unix.stackexchange.com/questions/150492/backspace-deletes-bash-prompt)

**PS: 更多配置见图片**
![图片](https://github.com/code4EE/images/blob/main/20210516185059.png)

## 2.别名

- ### 使用别名

使用 `alias` 命令为长命令创建别名,实现便捷化 \
格式为

```bash
alias new_command='existing command'
```

举例:

```bash
# 利用alias命令解决长命令的简短化
alias install='sudo yum install -y'
# 使用新命令安装gcc工具链
install gcc
```

但是这种设置效果只是暂时的,可以将其写入 `~/.bashrc` 中实现一直使用

```bash
# >> 代表追加写入
echo 'alias new_command="existning_command"' >> ~/.bashrc
```

- ### 删除别名

从 `~/.bashrc` 删除即可,或者直接使用 `unalias` 取消别名,或者 `alias command=` 取消

```bash
# 举例
unalias install
alias install=
```

- #### 使用别名需要注意的事情

```bash
# 1.使用alias命令可以新建一个与原命令相同的别名,使用'\'可以使用原来的命令,这一点可以防止一些安全问题
\command

# 2.列举所有当前已经定义的别名,直接使用alias命令
alias
```

## 3. tput命令和stty命令

### tput命令的简单介绍

使用 `tput` 命令设置背景终端背景色

```bash
tput setb n # n=range(0-7)
```

|n的取值|终端背景|
|:--|:-----|
|0|黑色 rgb[0, 0, 0] #000000|
|1|蓝色 rgb[30, 144, 245] #1E90F5|
|2|绿色 rgb[0, 100, 0] #006400|
|3|青色? rgb[0,205, 205] #00CDCD|
|4|红色 rgb[187, 0, 0] #BB0000|
|5|紫色 rgb[187, 0, 187] #BB00BB|
|6|黄色 rgb[200, 175, 0] #C8AF00|
|7|白色 rgb[235, 235, 235] #EBEBEB|

使用 `tput` 命令设置背景终端前景色

```bash
tput setf n # n=range(0-7)
```

### stty命令的简单介绍

输入密码的时候脚本不应该显示密码内容,使用stty可以实现

```bash
#!/bin/bash
#Filename: passwd.sh
echo -e "Enter password: "
# 在读取密码之前禁止回显
stty -echo
read password
# 重新允许回显
stty echo
echo
echo Password read.
```

## 4.命令行参数

```bash
# 修改文件的权限
chmod 755 test.sh
# 传递参数
./test.sh arg1 arg2 arg3
```

|符号|意义|
|:--|:--|
|$0 |脚本名称|
|$1 |第一个参数,例子中的arg1|
|$2 |第二个参数,例子中的arg2|
|$@ |以列表的形式打印所有的参数|
|$* |类似于"$@",但是将所有参数视为一个整体|

举例:

```bash
#!/bin/bash
#Filename: args.sh
for i in `seq 1 $#`
do
    echo $i is $1
    shift
done
```

执行并展示运行结果:

```bash
❄[wdy ~/shell_learning]>>> ./args.sh a b c d e f

1 is a
2 is b
3 is c
4 is d
5 is e
6 is f
```

## 5.read命令的用法

```bash
# 1.读取n个字符并存入变量
# read -n number_of_character <变量名>
# 举例
❄[wdy ~/shell_learning]>>> read -n 2 var

# 2.无回显的方式读取密码
# read -s <变量名>
# 举例
❄[wdy ~/shell_learning]>>> read -s var

# 3.使用提示信息
# read -p "提示信息" <变量名>
# 举例
❄[wdy ~/shell_learning]>>> read -p "Enter input:" var

# 4.给定期限读取输入
# read -t timeout <变量名>
# 举例
❄[wdy ~/shell_learning]>>> read -t 3 var #单位:秒

# 5.给定分隔符结束输入
# read -d <delimiter> <变量名>
# 举例
❄[wdy ~/shell_learning]>>> read -d ":" var
hello:
```
