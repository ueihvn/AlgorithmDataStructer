package main

import "fmt"

func lengthOfLastWord(s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == " " {
			continue
		}
		res := 1
		for j := i - 1; j >= 0; j-- {
			if string(s[j]) != " " {
				res++
			} else {
				break
			}
		}
		return res

	}
	return 0
}

func testLengthOfLastWord() {
	arrS := []string{
		"Hello World",
		"   fly me   to   the moon  ",
		"luffy is still joyboy",
		"Hello",
		"",
	}

	for i := range arrS {
		fmt.Println(arrS[i])
		fmt.Println(lengthOfLastWord(arrS[i]))
	}
}
