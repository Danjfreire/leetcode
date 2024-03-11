package main

func findMaxAverage(nums []int, k int) float64 {
 currentSum := 0 

 for idx := 0; idx < k; idx ++ {
	currentSum += nums[idx]
 }
 
 topSum := currentSum 

 for idx := k; idx < len(nums); idx++ {
	currentSum -= nums[idx - k]
	currentSum += nums[idx]

	if currentSum > topSum {
		topSum = currentSum
	}
 }

 return float64(topSum) / float64(k)
}