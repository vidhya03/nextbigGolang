package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println("hellow")
	// arr := []string{"some", "values", "list"}
	// fmt.Println(len(arr))
	// fmt.Println(arr[0])
	// fmt.Println(arr[1])
	// fmt.Println(arr[2])
	// fmt.Println(strings.Index("Rajesh", "Raj"))
	// fmt.Println(strings.Index("Sheldon", "Shely"))
	// fmt.Println(longestCommonPrefix([]string{"some", "some", "some"}))
	fmt.Println(longestCommonPrefix([]string{"flower","flow","flight"}))
	// fmt.Println(strings.Index("flower", "fl"))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var prefix string = strs[0]
	for i := 1; i < len(strs); i++ {		
		for strings.Index(strs[i],prefix) != 0 {
			prefix = prefix[:len(prefix)-1]
			if len(prefix) == 0 {
				return ""
			}
		}
	}

	return prefix
}
