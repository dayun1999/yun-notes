# 背包问题的套路

[一篇文章吃透背包问题](https://leetcode-cn.com/problems/coin-change/solution/yi-pian-wen-zhang-chi-tou-bei-bao-wen-ti-sq9n/)

## 常见背包类型

常见的背包类型主要有以下几种：
1、0/1背包问题：每个元素最多选取一次
2、完全背包问题：每个元素可以重复选择
3、组合背包问题：背包中的物品要考虑顺序
4、分组背包问题：不止一个背包，需要遍历每个背包

## 背包分类的模板：

**1、0/1背包：  外循环nums, 内循环target, target倒序且 target>=nums[i];**
**2、完全背包：外循环nums, 内循环target, target正序且 target>=nums[i];**
**3、组合背包(考虑顺序)：外循环target,内循环nums,target正序且target>=nums[i];**
**4、分组背包：这个比较特殊，需要三重循环：外循环背包bags,内部两层循环根据题目的要求转化为1,2,3三种背包类型的模板**

## 问题分类的模板：

1、最值问题:   `dp[i] = max/min(dp[i], dp[i-nums]+1)`或`dp[i] = max/min(dp[i], dp[i-num]+nums)`;
2、存在问题：`dp[i]=dp[i]||dp[i-num]`;
3、组合问题：`dp[i] += dp[i-num]`

这样遇到问题将两个模板往上一套大部分问题就可以迎刃而解

