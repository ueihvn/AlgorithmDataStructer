package main

func generate(numRows int) [][]int {
	dp := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		dp[i] = make([]int, i+1)
	}
	dp[0][0] = 1
	for y := 1; y < numRows; y++ {
		for x := 0; x <= y; x++ {
			if x == 0 || x == y {
				dp[y][x] = dp[0][0]
			} else {
				dp[y][x] = dp[y-1][x-1] + dp[y-1][x]
			}
		}
	}
	return dp
}

func getRow(rowIndex int) []int {
	numRows := rowIndex + 1
	dp := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		dp[i] = make([]int, i+1)
	}
	dp[0][0] = 1
	for y := 1; y < numRows; y++ {
		for x := 0; x <= y; x++ {
			if x == 0 || x == y {
				dp[y][x] = dp[0][0]
			} else {
				dp[y][x] = dp[y-1][x-1] + dp[y-1][x]
			}
		}
	}
	return dp[rowIndex]
}
