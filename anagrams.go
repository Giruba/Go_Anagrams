package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println("Are anagrams of each other?")
	fmt.Println("---------------------------")
	
        fmt.Println("Enter first string")
        var firstString string
	fmt.Scanln(&firstString)

	firstStringC := strings.Split(firstString, "")
    	sort.Strings(firstStringC)
    	firstString = strings.Join(firstStringC, "")

	fmt.Println("Enter second string")
	var secondString string
	fmt.Scanln(&secondString)
	
	
	secondStringC := strings.Split(secondString, "")
    	sort.Strings(secondStringC)
    	secondString = strings.Join(secondStringC, "")

	if firstString == secondString {
		fmt.Println("Entered strings are anagrams of each other", firstString, secondString)
	} else{
		fmt.Println("Entered strings are not anagrams of each other", firstString, secondString)
	}
}
