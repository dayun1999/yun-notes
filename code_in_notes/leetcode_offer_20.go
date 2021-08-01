// 剑指offer第20题
package leetcode_offer_20

// 表示数值的字符串
// 0 起始的blank
// 1 e之前的sign
// 2 dot之前的数字digit
// 3 dot之后的数字digit
// 4 dot之前为空时, dot之后的数字digit
// 5 e
// 6 e之后的sign
// 7 e之后的digit
// 8 尾部的blank
func isNumber(s string) bool {
	var states = []map[byte]int{
		{' ': 0, 's': 1, 'd': 2, '.': 4},
		{'d': 2, '.': 4},
		{'d': 2, '.': 3, 'e': 5, ' ': 8},
		{'d': 3, 'e': 5, ' ': 8},
		{'d': 3},
		{'s': 6, 'd': 7},
		{'d': 7},
		{'d': 7, ' ': 8},
		{' ': 8},
	}
	p := 0
	for i := 0; i < len(s); i++ {
		state := getType(s[i])
		if _, ok := states[p][state]; !ok {
			return false
		}
		p = states[p][state]
	}
	return p == 2 || p == 3 || p == 7 || p == 8
}

func getType(c byte) byte {
	var res byte
	if '0' <= c && c <= '9' {
		res = 'd'
	} else if c == '+' || c == '-' {
		res = 's'
	} else if c == 'e' || c == 'E' {
		res = 'e'
	} else if c == '.' || c == ' ' {
		res = c
	} else {
		res = '?'
	}
	return res
}
