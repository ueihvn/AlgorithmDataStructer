package main

import "fmt"

func addBinary(a string, b string) string {
	i, j, increment := len(a)-1, len(b)-1, 0
	res := ""
	for !(i < 0 && j < 0) {
		if i >= 0 && j >= 0 {
			eleA := fromStringToInt(string(a[i]))
			eleB := fromStringToInt(string(b[j]))
			temp := eleA + eleB + increment

			increment = temp / 2
			res = fromIntToString(temp%2) + res
			i--
			j--
		} else if i >= 0 {
			eleA := fromStringToInt(string(a[i]))
			temp := eleA + increment

			increment = temp / 2
			res = fromIntToString(temp%2) + res
			i--
		} else if j >= 0 {
			eleB := fromStringToInt(string(b[j]))
			temp := eleB + increment

			increment = temp / 2
			res = fromIntToString(temp%2) + res
			j--
		}
	}

	if increment == 1 {
		return fmt.Sprintf("1%s", res)
	}

	return res
}

func fromStringToInt(str string) int {
	switch str {
	case "0":
		return 0
	case "1":
		return 1
	}
	return -1
}

func fromIntToString(i int) string {
	switch i {
	case 0:
		return "0"
	case 1:
		return "1"
	}
	return "-1"
}

func testAddBinary() {
	arrInput := [][]string{
		{"11", "1"},
		{"1010", "1011"},
	}

	for _, v := range arrInput {
		fmt.Println(v)
		fmt.Println(addBinary(v[0], v[1]))
	}
}
