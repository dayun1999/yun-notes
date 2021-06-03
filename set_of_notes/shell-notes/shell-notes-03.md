# shell-notes-03-以文件之名

## 内容大纲

- *生成任意大小的文件*
- *文件权限、所有权与粘滞位*

## 1. 生成任意大小的文件

## 2. 文件权限、所有权与粘滞位

### 2.1 `chmod`命令设置权限

- #### 设置文件权限

```bash
# u代表用户权限
# g代表用户组权限
# o代表其他用户权限
>>> chmod u=rwx, g=rw, o=r filename

#或者使用八进制
# r=4
# w=2
# x=1
>>> chmod 755 filename
```

- #### 增加或删除权限

```bash
# 增加权限
# 给所有(a代表all)权限类别添加可执行权限
>>> chmod a+x filename

# 删除权限
>>> chmod a-x filename
```

### 2.2 `chown`命令更改所有权

格式: `chown user:group filename`

```bash
>>> chown wdy:root hello.txt
```

### 2.3 设置粘滞位(sticky bit)

> 粘滞位是 *目录* 的一个特殊权限\
如果目录没有设置执行权限(x),粘滞位使用`T`;\
如果目录有设置执行权限(x), 粘滞位使用`t`;\
粘滞位可以应用于目录, 设置之后只有目录的所有者才能删除目录, 即使他人拥有写权限也无法执行删除操作

```bash
>>> chmod a+t dir_name
```

### 2.4 递归方式设置文件权限和文件所有权

`-R`选项即可

```bash
# .代表当前目录, 可替换
>>> chmod 777 . -R

>>> chmod user:group . -R
```

### 2.5 以不同的身份运行可执行文件(setuid)

`setuid`只能用于ELF格式的可执行文件

```bash
>>> chmod +s executable_file
>>> chmod root:root executable_file
>>> chmod +s executable_file
>>> ./executable_file
# 现在无论是谁发起调用, 该文件都是以root的身份运行
```

### 2.6 `chattr`将文件设置为不可修改

设置之后文件不可删除, 除非撤销不可修改属性

```bash
# 添加不可修改属性
>>> sudo chattr +i number.txt

# 试图删除文件
>>> rm -f number.txt 
rm: cannot remove ‘number.txt’: Operation not permitted

# 移除不可修改属性
>>> sudo chattr -i number.txt 

```
