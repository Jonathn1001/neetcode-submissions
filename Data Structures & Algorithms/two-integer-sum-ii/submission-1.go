func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers) - 1
	for i < j  {
		currentSum:= numbers[i] + numbers[j]

		if currentSum == target {
			return []int{i+1,j+1}
		} else if currentSum > target {
			j--
		} else {
			i++
		}
	}
	return nil
}
