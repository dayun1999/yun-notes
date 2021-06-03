# shell-notes-05-一团乱码?没这回事

## :blue_book:内容大纲

- *wget命令介绍*
- *cURL介绍*

## 1. `wget`命令介绍

wget是一个用于文件下载的命令行工具\
命令基本格式

```bash
wget <URL>
#或者一次性下载多个
wget <URL1> <URL2> ...
```

|选项|描述|举例|
|:--:|:--|:--|
|`-O`(大写)|指定输出文件名|`wget baidu.com -O baidu_index.html`|
|`-o`(小写)|指定一个日志文件, 下载信息就不会打印到stdout了|`wget baidu.com -o baidu.log`|
|`-t`|指定最多下载重试次数(由于网络不稳定会造成下载失败,所以要多试几次)|`wget -t 5 baidu.com`,最多尝试5次, 当`-t`选项设置为0的时候wget会不断进行重试|
|`--limite-rate`|限定下载的最大带宽,后面的单位可以为k(千字节)或m(兆字节)|`wget --limit-rate 1k https://geektutu.com/`|
|`--quota或者-Q`|指定最大下载配额|`wget -Q 10m https://geektutu.com/`|
|`-c`|断点续传||
|`--mirror`|复制整个网站(镜像)||
|`-r`|递归下载||
|`-l`|指定页面层级 (深度) ||

## 2. cURL介绍

|选项|描述|举例|
|:--:|:--|:--|
|`-C -`|断点续传,并且希望cURL自定推断出断点在哪，也可以自己指定断点偏移位置`-C offset`|curl -C - URL`|
|`--cookie-jar`|设置cookie, 指定cookie文件|`curl --cookie-jar cookir-file`|
|`--user-agent`|设置用户代理字符串|`curl URL --user-agent "Mozilla/5.0"`|
|`--limit-rate`|限定下载速率,和wget一样||
|`--max-filesize`|指定最大下载量,下载成功返回0|`curl URL --max-filesize bytes`|
|`-u`|完成HTTPS或者FTP认证|`curl -u user:pass URL`|
|`-I`或者`--head`|只打印HTTP头部信息,无须下载远程文件||
|`--silent`|不显示下载进度信息||
|`--progress`|显示下载进度条||
|`-O`|指明将下载数据写入文件,*注意URL必须是完整的,不能只是域名*|`curl www.knopper.net/index.html -O --progress`执行完后会生成index.html文件|
|`-o`|指定输出文件名||

### curl发送Web页面并读取响应

|选项|描述|举例|
|:--:|:--|:--|
|`-d`|||