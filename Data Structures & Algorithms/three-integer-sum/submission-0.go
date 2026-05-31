// nums[i] + nums[j] + nums[k] = 0 => num[j] + num[k] = -nums[i]
func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums) // Step 1: Sort the array
	// Step 2: Iterate through the array
	for i := 0; i < len(nums)-2; i++ {	
		// Step 3: Skip duplicate anchors to avoid identical triplets
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// If the anchor is greater than 0, 
		// since the array is sorted, no 3 positive numbers can ever sum to 0.
		if nums[i] > 0 {
			break
		}

		// Step 4: Two-pointer setup for the rest of  the array
		left := i+1
		right := len(nums)-1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// Step 5: Skip duplicate values for left and right pointer
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				// Move pointers inward after finding a match
				left++
				right--
			} else if sum < 0 {
				left++ // Sum is too small, move to a larger number
			} else {
				right-- // Sum is too large, move to a smaller number
			}
		}
	}
	return result
}
