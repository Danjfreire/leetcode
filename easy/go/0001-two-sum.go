// O(n)
func twoSum(nums []int, target int) []int {
   diffMap:= make(map[int]int) 

   for i := 0; i < len(nums); i++ {
	diff := target - nums[i]

	value, found := diffMap[diff]

	if(found) {
		return []int{i,value}
	} else {
		diffMap[nums[i]] = i
	}
   }

   return []int{}
}