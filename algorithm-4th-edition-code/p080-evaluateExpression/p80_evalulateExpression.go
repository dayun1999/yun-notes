/*
Dijkstra的双栈算术表达式求值(算法第4版-P80)
流程:
1.将操作数压入"操作数栈"
2.将运算符压入"运算符栈"
3.忽略左括号
4.遇到右括号, 弹出一个运算符, 弹出所需数量的操作数, 并将计算结果压入操作数栈
*/
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/code4EE/algorithm-4th-edition/stack"
)

func main() {
	ops := stack.New()
	vals := stack.New()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入你要计算的表达式:")

	for {
		s, _ := reader.ReadString('\n')
		s = strings.TrimSuffix(s, "\n")
		if s != "" {
			if s == "(" {
				// 忽略左括号
			} else if s == "+" || s == "-" || s == "*" || s == "/" || s == "sqrt" {
				// 如果是运算符就压入操作符栈里面
				ops.Push(s)
			} else if s == ")" {
				// 处理右括号
				var op string = ops.Pop().(string)
				var v float64 = vals.Pop().(float64)
				if op == "+" {
					v = vals.Pop().(float64) + v
				} else if op == "-" {
					v = vals.Pop().(float64) - v
				} else if op == "*" {
					v = vals.Pop().(float64) * v
				} else if op == "/" {
					v = vals.Pop().(float64) / v
				} else if op == "sqrt" {
					v = math.Sqrt(v)
				}
				vals.Push(v)
			} else {
				// 如果不是括号也不是运算符,就将它压入操作数栈里面
				val, _ := strconv.ParseFloat(s, 64)
				vals.Push(val)
			}
		} else {
			break
		}
	}
	fmt.Println(vals.Pop())
}
