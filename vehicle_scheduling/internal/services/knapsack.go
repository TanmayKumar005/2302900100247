package services

import "github.com/TanmayKumar005/2302900100247/vehicle_scheduling/internal/models"

func OptimizeTasks(tasks []models.Vehicle, capacity int) ([]models.Vehicle, int) {
	n := len(tasks)

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	for i := 1; i <= n; i++ {
		for w := 0; w <= capacity; w++ {
			dp[i][w] = dp[i-1][w]

			if tasks[i-1].Duration <= w {

				candidate := dp[i-1][w-tasks[i-1].Duration] + tasks[i-1].Impact

				if candidate > dp[i][w] {
					dp[i][w] = candidate
				}
			}
		}
	}
	var selected []models.Vehicle

	w := capacity

	for i := n; i > 0; i-- {
		if dp[i][w] != dp[i-1][w] {
			selected = append(selected, tasks[i-1])
			w -= tasks[i-1].Duration
		}
	}
	return selected, dp[n][capacity]
}
