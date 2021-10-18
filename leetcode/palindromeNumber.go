package main

import (
	"strconv"
)

func isPalinromeNumberString(x int) bool {
	if x < 0 {
		return false
	}

	strNum := strconv.Itoa(x)
	len := len(strNum)
	for i := 0; i < len/2; i++ {
		if strNum[i] != strNum[len-1-i] {
			return false
		}

	}
	return true
}

func isPalinromeNumberRevert(x int) bool {
	if x < 0 {
		return false
	} else if x == 0 {
		return true
	}

	var revertX int
	x2 := x
	for x2 != 0 {
		revertX *= 10
		revertX += x2 % 10
		x2 = x2 / 10
	}

	if revertX != x {
		return false
	}
	return true
}
