package main

func maxArea(height []int) int {
    left := 0
	right := len(height)-1
    maxArea := 0

    for left < right {
        area := right - left

        if height[left] < height[right]{
            area *= height[left]
            left++
        }else{
            area *= height[right]
            right--
        }

        if area > maxArea {
            maxArea = area
        }
    } 

    return maxArea
}