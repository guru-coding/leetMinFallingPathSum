package leetMinFallingPathSum

import "slices"

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	if n == 1 {
		return matrix[0][0]
	}
	dp := make([]int, n)
	mem := make([]int, n)
	for i := 0; i < n; i++ {
		dp[0] = min(mem[0], mem[1]) + matrix[i][0]
		dp[n-1] = min(mem[n-2], mem[n-1]) + matrix[i][n-1]
		for j := 1; j < n-1; j++ {
			dp[j] = min(mem[j-1], mem[j], mem[j+1]) + matrix[i][j]
		}
		dp, mem = mem, dp
	}
	return slices.Min(mem)
}
