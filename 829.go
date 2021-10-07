package main

func consecutiveNumbersSum(n int) int {
	res := 0
	count := 1
	for i := 1; i*i/2 <= n; i++ {
		if i%2 == 0 {
			if n%i == count {
				res++
			}
			count++
		} else {
			if n%i == 0 {
				res++
			}
		}
	}

	return res
}

// func main() {
// 	fmt.Println(consecutiveNumbersSum(3))
// }
