/*
原书109页ThreeSum程序go语言版本, 为了演示时间复杂度
*/
package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/code4EE/algorithm-4th-edition/utils/stopwatch"
)

// time complexitity is O(N^3)
func count(a []int) int {
	n := len(a)
	cnt := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if a[i]+a[j]+a[k] == 0 {
					cnt++
				}
			}
		}
	}
	return cnt
}

func main() {
	n, err := strconv.ParseInt(os.Args[1], 10, 32)
	fmt.Println("n是",n)
	if err != nil {
		log.Fatal(err)
	}
	a := make([]int, n)
	rand.Seed(time.Now().UnixNano())
	min, max := -1000000, 1000000
	for i := 0; i < int(n); i++ {
		a[i] = rand.Intn(max-min) + min
	}
	stopwatcher := stopwatch.New()
	cnt := count(a)
	t := stopwatcher.Elapsed()
	fmt.Println(cnt, " triples", t)
}
