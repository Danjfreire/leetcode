package main

func maximumWealth(accounts [][]int) int {
	maxWealth := 0

	for i := 0; i < len(accounts); i++ {
		wealth := 0
		for j := 0; j < len(accounts[i]); j++ {
			wealth += accounts[i][j]
		}

		if wealth > maxWealth {
			maxWealth = wealth
		}
	}

	return maxWealth
}
