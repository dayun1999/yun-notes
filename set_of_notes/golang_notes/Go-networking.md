# Go网络编程知识点总结

## 目录

- #### <a href="#io_model">IO模型</a>
  
  - <a href="#"></a>
- #### <a href="#how_to_find_bug">线上如何排查问题</a>
  
  - <a href="#"></a>
  - <a href="#"></a>
- <a href="#"></a>
- <a href="#"></a>
- <a href="#"></a>



## <a name="io_model">IO模型</a>

建议阅读`Unix Network Programming volume I`

- 阻塞式IO
- 非阻塞式IO
- IO multiplexing(`select`和`poll`)
- 信号驱动式IO
- 异步IO



## <a name="hot_to_find_bug">如何线上排查问题</a>

Linxu性能分析工具全家桶： [Linux eBPF Tracing Tools (brendangregg.com)](https://www.brendangregg.com/ebpf.html)

### 1. top命令

具体可看[(34条消息) top命令详解_wisgood的专栏-CSDN博客](https://blog.csdn.net/wisgood/article/details/38959881)

- 按1——[TOP命令 详解CPU 查看多个核心的利用率按1 - AmilyAmily - 博客园 (cnblogs.com)](https://www.cnblogs.com/AmilyWilly/p/7016319.html)



### 2. 如何做`si`的均衡

- 富人方案: 网卡多队列(多核心)
- 穷人方案: 单队列的CPU均衡(单核心)

[MYSQL数据库网卡软中断不平衡问题及解决方案 | 系统技术非业余研究 (yufeng.info)](https://blog.yufeng.info/archives/2037)

#### 视频中提及

- `ByPass kernel`
- `DPDK`



### 3. `nmon`命令【推荐】

```bash
yum install nmon
```

[性能测试之nmon对linux服务器的监控 - 第二个卿老师 - 博客园 (cnblogs.com)](https://www.cnblogs.com/qgc1995/p/7523786.html)

- 按`k`， 查看上下文切换和中断数
- 按`n`，查看网络的包



### 4. `nload`命令

查看网卡的实时流量非常方便



### 5. `tcpflow`命令



### 6. `ifconfig`命令

[每天一个linux命令（52）：ifconfig命令 - peida - 博客园 (cnblogs.com)](https://www.cnblogs.com/peida/archive/2013/02/27/2934525.html)

主要为了排除网卡或者网卡驱动是否有问题，核心看有没有丢包

```bash
RX packets 254  bytes 22418 (21.8 KiB)
RX errors 0  dropped 0  overruns 0  frame 0 // 比如看看接收到多少错误
```

### 7. `dmesg` 命令或者`vim /var/log/messages`

查看系统有没有什么硬件问题



### 8. `netstat -S`

`netstat`其他的命令不要乱敲,会把机器搞死机



### 9. `ss`命令

```bash
$ ss -s
Total: 200 (kernel 228)
TCP:   17 (estab 2, closed 1, orphaned 0, synrecv 0, timewait 1/0), ports 0

Transport Total     IP        IPv6
*	  228       -         -        
RAW	  0         0         0        
UDP	  5         4         1        
TCP	  16        14        2        
INET	  21        18        3        
FRAG	  0         0         0   
```

`timewait` 很重要, 重点需要看一篇文章



### 10. `vmstat`命令



### 11. `iostat`命令

[Linux iostat命令详解 - 小a玖拾柒 - 博客园 (cnblogs.com)](https://www.cnblogs.com/ftl1012/p/iostat.html)

```bash
$ iostat -x 1
Device:         rrqm/s   wrqm/s     r/s     w/s    rkB/s    wkB/s avgrq-sz avgqu-sz   await r_await w_await  svctm  %util
vda               0.00     0.00    0.00    0.00     0.00     0.00     0.00     0.00    0.00    0.00    0.00   0.00   0.00

```



### 12. `iotop`命令

查看io被谁吃掉了



### 13. `lsof -p $pid`

查看进程打开了哪些文件



### 14. `perf`命令【重要】

[系统级性能分析工具perf的介绍与使用 - ArnoldLu - 博客园 (cnblogs.com)](https://www.cnblogs.com/arnoldlu/p/6241297.html)

```bash
$ perf top
```



### 15. `free`命令



### 16. `ethtool`命令

可以修改MTU



conntrack是个大坑




### 其他

#### hlist

[Linux 内核 hlist 详解 - 怀想天空_2013 - 博客园 (cnblogs.com)](https://www.cnblogs.com/cyyljw/p/10722709.html)

#### 单机连接端口数耗尽解决

虚拟网卡、多开几个IP



## 粘包











