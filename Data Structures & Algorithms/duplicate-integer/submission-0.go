func hasDuplicate(nums []int) bool {
    appearances := make(map[int]int)

	for _, v := range nums {
		if appearances[v] >= 1 {
			return true
		} 
		appearances[v]++
	}
	
	return false
}
