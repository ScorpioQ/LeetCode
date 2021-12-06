package main

import "fmt"

func detectCapitalUse(word string) bool {
	if len(word) == 1 {
		return true
	}

	CheckStatus := 0
	for _, item := range word {
		switch CheckStatus {
		case 0: // first
			if item <= 'Z' {
				CheckStatus = 1
			} else {
				CheckStatus = 3
			}
		case 1: // first Big
			if item <= 'Z' {
				CheckStatus = 2
			} else {
				CheckStatus = 3
			}
		case 2: // need all big
			if item > 'Z' {
				return false
			}

		case 3: // need all small
			if item <= 'Z' {
				return false
			}
		}
	}

	return true
}

func main() {
	fmt.Println(detectCapitalUse("USA"))
	fmt.Println(detectCapitalUse("leetcode"))
	fmt.Println(detectCapitalUse("Google"))
	fmt.Println(detectCapitalUse("FlaG"))
}
