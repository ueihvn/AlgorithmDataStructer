package main

func longestCommonPrefixTwoString(str1, str2 string) string {
	var res string
	var lenSmallestString int

	if len(str1) > len(str2) {
		lenSmallestString = len(str2)
	} else {
		lenSmallestString = len(str1)
	}

	for i := 0; i < lenSmallestString; i++ {
		if str1[i] == str2[i] {
			res += string(str1[i])
		} else {
			break
		}
	}

	return res

}

func longestCommonPrefix(strs []string) string {
	res := strs[0]
	for i := 1; i < len(strs); i++ {
		lcp := longestCommonPrefixTwoString(strs[i], res)
		if lcp != res {
			res = lcp
		}
	}
	return res
}
