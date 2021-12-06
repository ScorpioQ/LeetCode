package main

import "fmt"

func removeDuplicates(s string, k int) string {
	res := ""
	for i := 0; i < len(s); i++ {
		res = res + string(s[i])
		lenRes := len(res)
		needDel := true
		for j := 1; j < k; j++ {
			if lenRes-1-j < 0 {
				needDel = false
				break
			}
			if res[lenRes-1-j] != res[lenRes-j] {
				needDel = false
				break
			}
		}
		if needDel {
			res = res[:len(res)-k]
		}
	}

	return res
}

func main() {
	fmt.Println(removeDuplicates("pbbcggttciiippooaais", 2))
}
