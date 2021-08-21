# 笔记：git的使用

##  <font color="blue">操作GitHub的一般流程</font>

```python
#前提：本地安装git && 完成git配置(你的uername和email)

#1.在GitHub网站create新的仓库(reposity)
#2.在本地(Windows)的一个目标目录打开git(右键git bash)
#3.git clone 你的仓库地址(HTTPS 或者 SSH)
#4.克隆下来的一个空项目进行编写
#5.git add -A,将全部(包括你的项目配置文件)添加到暂存区
#6.git commit -m "输入你想输入的信息"

#7.git remote add <远程名字> <你的仓库地址> 注意：如果是头一次使用GitHub，那么远程名字可以写origin,
不同的项目地址使用不同的远程名字
比如：
git remote add algorithms https://github.com/code4EE/Algorithms-4th-Edition-Golang.git

#8.将我们暂存区的所有内容提交(push)到GitHub
#命令如下： git push <远程名字> <分支名字>
如： git push algorithms master
分支默认是master，程序员可以新建分支，从而实现多个人多个分支，共同编写项目，最后合并(merge)分支

```

## <font color="red"> 如何正确pull request</font>

```python
#1.先fork别人的仓库
点击github页面的图标--fork 即可
```

```python
#2.克隆(clone)别人的仓库
$ git clone <fork后在你的仓库里面的地址>
```

```python
#3.进入克隆的仓库
$ cd <仓库名>
```

```python
#4.创建一个新的分支(branch)，这是一个好习惯
$ git checkout -b <分支的名字>
```

```python
#5.修改你想修改的地方并且提交(commit),步骤依次如下：
$ git status 目的：查看修改的文件有哪些(红色的那部分)
$ git add .   目的：将所有修改的文件添加到暂存区
$ git commit -m "你想输入的信息" 目的：提交你所做的更改
```

```python
#6.将更改推送(push)到github,步骤依次如下：
$ git remote  目的：识别远程名(remote's name),比如输入这个命令之后，出现origin,那么origin就是remote name
$ git push <执行git remote命令之后出现的远程名> <前面我们新建的分支的名字> 目的：安全地推送更改到github

```

```python
#7.创建pull request
到我们自己的仓库，会发现有一个按钮<Compare & pull request>,点击即可，
然后输入你做的改变，比如你修复了哪个错误(也就是告诉作者你做了怎样的更改)，然后提交(submit),如果你的pull request被作者认可了，你会收到相应的邮件
```

```python
#8.同步你fork过的主分支
--第一步，检查你在哪一个分支
$ git branch 
--第二步，切换到master分支(因为上面我们自己建了一个新分支，所以要切换回主分支)
$ git checkout master
--第三步，添加原始仓库作为一个upstream仓库
$ git remote add upstream <别人的HTTP仓库地址>
--第四步，拿到原始仓库(fetch the repositry)
$ git fetch upstream
--第五步，合并
$ git merge upstream/master
--第六步，将更改推送到github
$ git push origin master

注意： 在同步我们fork过的仓库的主分支之后，我们也可以删掉自己的那个远程名(这个例子里面是upstream),但如果你未来可能会再次更新/同步，所以保留就好
#删除的命令如下
$ git remote
$ git remote rm upstream
```

```python
#9.删除不必要的分支
#由于本次为了修改代码我们新建了一个分支，代码修改完，提交完之后可以删除这个分支
$ git branch -d <我们新建的分支的名字>

#同时也可以删除github上面的那个分支
$ git push origin --deleted <我们新建的分支的名字>
```



## <font color="red">如何创建新的分支</font>

```python
#1. 创建新的分支
git branch <新的分支名字>
```

```python
#2. 切换分支
git checkout <新的分支名字>
```

```python
#3. 与远程分支相关联
git remote add origin https://github.com/... 
```

```python
#4. 将分支上传
git push origin 分支名字
```





详细请点击[这里](https://github.com/GarageGames/Torque2D/wiki/Cloning-the-repo-and-working-with-Git)

#### <font color="orange">只删除远程仓库的文件，不影响本地文件</font>

```vim
git rm --cached 文件名
#删除文件夹需要加 -r
git commit -m "delete some files or directories"
git push
```

#### <font color="orange">删除本地和远程仓库的文件</font>

```vim
git rm [-r] <文件名>
git commit -m "deleted some files"
git push
```

#### <font color="orange">只将修改的文件add到缓存区</font>

```python
git add -u
#或者将所有文件(包括untarcked files)添加
git add -A (也即 git add . && git add -u)
```

#### <font color="green">提交到暂存区（index）后又想删了重新提交</font>

``` python
#直接删除index区里的所有文件
git rm [-r] --cached <文件> 
#-r参数在删除文件夹的时候需要
```

#### 如何取消 git add 操作

```python
$git reset <file> 或者  git reset 取消所有
```

#### 如何取消 git init 操作

```python
$rm -rf .git
```

#### 如何取消对本地文件的修改

```python
$git reset --hard  #Discard all local changes to all files permanently
$git checkout -- <file> #Discarding local changes (permanently) to a file
$git stash #Discard all local changes, but save them for possible re-use later
```



#### <font color="orange">如何查看并更改远程名字</font>

```python
#查看当前使用的remote name
git remote -v
#修改当前的远程名字
git remote rename <当前的远程名字> <想要修改成的远程名字>
```



## 常见问题记录

#### 【问题】Your branch is ahead of 'origin/main' by 3 commits.

```python
#最直接的，添加--hard参数后，会回到上次commit的状态，也就是说从上次commit之后的的修改都将被重置，换句话说这些数据都丢失了，所以要谨慎操作
git reset --hard origin/main
然后再获取远程仓库的最新数据:
git pull origin main
```





