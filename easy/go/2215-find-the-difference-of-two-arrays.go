package main

func findDifference(nums1 []int, nums2 []int) [][]int {
 nums1Map := getMap(nums1)
 nums2Map := getMap(nums2)

 onlyInNums1 := []int{}
 onlyInNums2 := []int{}

 for k, _ := range nums1Map {
	_, ok := nums2Map[k]
	if !ok {
		onlyInNums1 = append(onlyInNums1, k)
	}
 }

 for k, _ := range nums2Map {
	_, ok := nums1Map[k]
	if !ok {
		onlyInNums2 = append(onlyInNums2, k)
	}
 }

 return [][]int{onlyInNums1, onlyInNums2}
}

func getMap(nums []int) map[int]bool {
	numMap := make(map[int]bool);
	for i := 0; i < len(nums); i++ {
		_, ok := numMap[nums[i]]
		if !ok {
			numMap[nums[i]] = true	
		}
	}

	return numMap
}