package main

func romanToInt(s string) int {
	var res int
	romanInt := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for i := 0; i < len(s)-1; i++ {
		if romanInt[string(s[i])] < romanInt[string(s[i+1])] {
			res -= romanInt[string(s[i])]
		}

		res += romanInt[string(s[i])]

	}
	res += romanInt[string(s[len(s)-1])]
	return res
}
