package main

func twoSum(nums []int, target int) []int {
	tmpMap := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		tmpMap[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		leftValue := target - nums[i]
		if leftIndex, ok := tmpMap[leftValue]; ok && leftIndex > i {
			return []int{i, leftIndex}
		}
	}
	return nil
}
