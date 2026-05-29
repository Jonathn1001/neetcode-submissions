func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)
	for i, v := range nums {
		remaining := target - v
		if ri, ok := seen[remaining]; ok {
			return []int{ri, i}
		}
		seen[v] = i
	}
	return nil
}
