package main

import (
	"fmt"
	"strings"
)

func uniqueChars(str string) bool {
	str = strings.ToLower(str)
	mp := make(map[rune]struct{})
	// со string'ом range работает как с рунами
	for _, v := range str {
		if _, ok := mp[v]; ok {
			return false
		}
		mp[v] = struct{}{}
	}
	return true
}

func main() {
	str1 := "abcd"
	str2 := "abcdD"
	str3 := "abcda"
	fmt.Println(uniqueChars(str1)) //true
	fmt.Println(uniqueChars(str2)) //false
	fmt.Println(uniqueChars(str3)) //false
}
