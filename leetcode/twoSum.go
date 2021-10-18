package main

// use map for store nums was travel
// check if target - num[i] has in map?
// yes -> return i and map[target-num[i]]
// no -> map[num[i]] = i
func twoSum(nums []int, target int) []int {
	mapNums := make(map[int]int)

	for i, v := range nums {
		j := target - v
		if v, ok := mapNums[j]; ok {
			return []int{v, i}
		}
		mapNums[v] = i
	}
	return nil
}
