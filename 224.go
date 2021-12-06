package main

import "fmt"

func calculate(s string) int {
	sign := []int{1}
	cur_sign := 1
	res := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ' ':
			continue

		case '+':
			cur_sign = sign[len(sign)-1]

		case '-':
			cur_sign = -sign[len(sign)-1]

		case '(':
			sign = append(sign, cur_sign)

		case ')':
			sign = sign[:len(sign)-1]

		default:
			num := 0
			for ; i < len(s) && '0' <= s[i] && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			i--
			res += num * cur_sign
		}
	}
	return res
}

func main() {
	fmt.Println(calculate("1 + 1"))
	fmt.Println(calculate(" 2-1 + 2 "))
	fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)"))
}
