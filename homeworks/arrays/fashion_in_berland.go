package arrays

import (
	"fmt"
	"strconv"
)

func fashionInBerland() {
	var strN string
	fmt.Scanln(&strN)
	n, _ := strconv.Atoi(strN)
	var temp string
	var nums string
	for i := 0; i < n; i++ {
		fmt.Scanf("%s", &temp)
		nums += temp
	}

	if len(nums) == 1 && string(nums[0]) == "1" {
		fmt.Println("YES")
		return
	} else if len(nums) == 1 && string(nums[0]) != "1" {
		fmt.Println("NO")
		return
	}
	var existZero bool
	for _, v := range nums {
		if string(v) == "0" && existZero {
			fmt.Println("NO")
			return
		} else if string(v) == "0" && !existZero {
			existZero = true
		}
	}
	if existZero {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
	return
}
