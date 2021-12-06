package main

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	m := map[string]int64{}
	for _, item := range cpdomains {
		rawDataArr := strings.Split(item, " ")
		count, _ := strconv.ParseInt(rawDataArr[0], 10, 64)
		domain := rawDataArr[1]
		m[domain] += count

		base := 0
		startIdx := 0
		for startIdx != -1 {
			startIdx = strings.Index(domain[base:], ".")
			if startIdx != -1 {
				startIdx++
				m[domain[base+startIdx:]] += count
				base += startIdx
			}
		}
	}

	result := []string{}
	for k, v := range m {
		result = append(result, fmt.Sprintf("%d %s", v, k))
	}
	return result
}

func main() {
	fmt.Println(subdomainVisits([]string{"9001 discuss.leetcode.com"}))
	fmt.Println(subdomainVisits([]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}))
}
