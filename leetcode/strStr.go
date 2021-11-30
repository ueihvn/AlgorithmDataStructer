package main

func strStr(haystack string, needle string) int {
	lenNeedle := len(needle)
	lenHaystack := len(haystack)
	if lenNeedle == 0 {
		return 0
	}
	i := 0
Loopi:
	for i <= lenHaystack-lenNeedle {
		if haystack[i] == needle[0] {
			for j := 1; j < lenNeedle; j++ {
				if haystack[i+j] != needle[j] {
					i++
					continue Loopi
				}
			}
			return i
		}
		i++
	}

	return -1
}
