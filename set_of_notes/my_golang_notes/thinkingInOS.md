# 操作系统的一些思考

## 目录

- #### <a href="#thread_safe">线程安全</a>

  - <a href="#what_is_thread_safe">1. 何为线程安全</a>
  - <a href="#how_to_ensure_thread_safe">2. 如何保障线程安全</a>

- #####  <a href="#user_and_kernel_space">用户态和内核态</a>

  - <a href="#what_are_user_and_kernel_sapce">是什么</a>
  - <a href="#why_distinguish_user_and_kernel_space">为什么区分用户态和内核态</a>
  - <a href="#"></a>

- <a href="#"></a>

  - <a href="#"></a>
  - <a href="#"></a>
  - <a href="#"></a>

- <a href="#"></a>

  - <a href="#"></a>
  - <a href="#"></a>

- <a href="#"></a>

  - <a href="#"></a>
  - <a href="#"></a>

- <a href="#"></a>

- 



## <a name="thread_safe">线程安全</a>

### <a name="what_is_thread_safe">1. 何为线程安全</a>

【推荐阅读》》》】[线程安全](https://www.cnblogs.com/lixinjie/p/a-answer-about-thread-safety-in-a-interview.html)

“线程安全”不是指线程的安全，而是指内存的安全

在每个进程的内存空间中都会有一块特殊的公共区域，通常称为堆（内存）。进程内的所有线程都可以访问到该区域，这就是造成问题的潜在原因。

假设某个线程把数据处理到一半，觉得很累，就去休息了一会，回来准备接着处理，却发现数据已经被修改了，不是自己离开时的样子了。可能被其它线程修改了。

**所以线程安全指的是，在堆内存中的数据由于可以被任何线程访问到，在没有限制的情况下存在被意外修改的风险**。

即堆内存空间在没有保护机制的情况下，对多线程来说是不安全的地方，因为你放进去的数据，可能被别的线程“破坏”。

### <a name="how_to_ensure_thread_safe">2. 如何保障线程安全</a>

- 根据操作系统会为每个线程分配属于它自己的内存空间，通常称为栈内存，其它线程无权访问。将资源设置为局部变量，它们都会被分配在线程栈内存中；
- 将资源设置为常量
- 加锁
  - 悲观锁，持悲观态度，就是假设我的数据一定会被意外修改，那干脆直接加锁得了。
  - 乐观锁，持乐观态度，就是假设我的数据不会被意外修改，如果修改了，就放弃，比如利用`CAS（Compare And Swap）`。就是在并发很小的情况下，数据被意外修改的概率很低，但是又存在这种可能性，此时就用`CAS`。



## <a name="user_and_kernel_space">用户态和内核态</a>

### <a name="what_are_user_and_kernel_sapce">1. 是什么</a>

**当进程运行在内核空间时就处于内核态，而进程运行在用户空间时则处于用户态**

#### 1. 用户空间和内核空间

操作系统(32bit)将虚拟地址空间划分为两部分，一部分为内核空间，另一部分为用户空间。针对 Linux 操作系统而言，最高的 1G 字节(从虚拟地址 0xC0000000 到 0xFFFFFFFF)由内核使用，称为内核空间。而较低的 3G 字节(从虚拟地址 0x00000000 到 0xBFFFFFFF)由各个进程使用，称为用户空间。

对上面这段内容我们可以这样理解：**「每个进程的 4G 地址空间中，最高 1G 都是一样的，即内核空间。只有剩余的 3G 才归进程自己使用。」**

**「换句话说就是， 最高 1G 的内核空间是被所有进程共享的！**

![](C:\Users\26646\Desktop\github.com\yun-notes\set_of_notes\my_golang_notes\pictures\user_space_kernel_space.jpg)

<img src="C:\Users\26646\Desktop\github.com\yun-notes\set_of_notes\my_golang_notes\pictures\userspace.png" style="zoom:50%;" />

### <a name="why_distinguish_user_and_kernel_space">2. 为什么区分用户态和内核态</a>

首先，**不是所有“操作系统”都分“用户态”和“内核态”**

在 CPU 的所有指令中，有些指令是非常危险的，如果错用，将导致系统崩溃，比如清内存、设置时钟等。如果允许所有的程序都可以使用这些指令，那么系统崩溃的概率将大大增加。

所以，CPU 将指令分为特权指令和非特权指令，对于那些危险的指令，只允许操作系统及其相关模块使用，普通应用程序只能使用那些不会造成灾难的指令。

比如 Intel 的 **CPU 将特权等级分为 4 个级别**：Ring0~Ring3。其实 Linux 系统只使用了 Ring0 和 Ring3 两个运行级别(Windows 系统也是一样的)。

当进程运行在 Ring3 级别时被称为运行在用户态，而运行在 Ring0 级别时被称为运行在内核态。



### <a name=""></a>

### <a name=""></a>

### <a name=""></a>

### <a name=""></a>

### <a name=""></a>

### <a name=""></a>

### <a name=""></a>

































