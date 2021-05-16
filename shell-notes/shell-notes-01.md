# shell-notes-01

## 学习时主要参考书籍《Linux脚本攻略》第三版

### 1.自定义linux提示符

在$HOME目录下,找到`.bashrc`文件并编辑,修改(如果没有则添加)PS1如下

```bash
PS1='❄\[\e[1;37m\][\[\e[1;36m\]\u \[\e[1;35m\]\w\[\e[1;37m\]]\[\e[1;31m\]>\[\e[1;33m\]>\[\e[1;32m\]>'
```

效果如图所示:
![图片](https://github.com/code4EE/images/blob/main/20210516182709.png)

#### 解释说明各个符号的意义

- `❄`  没什么意义,只想演示编辑支持unicode字符
- `\e[1:31m` 修改字体颜色,其中数字31代表红色,可修改为其他颜色,黑色=30,绿色=32,黄色=33,蓝色=34,洋红=35,青色=36,白色=37
- `\[`和`\]`  不键入该两个字符你会发现输入长命令的时候不会自动换行而且linux prompt会被blackspace键删掉,[感兴趣的看这个](https://unix.stackexchange.com/questions/150492/backspace-deletes-bash-prompt)

**PS: 更多配置见图片**
![图片](https://github.com/code4EE/images/blob/main/20210516185059.png)
