func maxArea(heights []int) int {
	 left := 0
	 right := len(heights) - 1
	 maxWater := 0
	 
	 for left < right {
		// 1. Calculate the width between 2 pointers
		width := right - left
		// 2: Find the lower bar
		var currentHeight = min(heights[left], heights[right])
		// Step 3: Calculate the current area & update maxWater
		currentWater := width * currentHeight
		if currentWater > maxWater {
			maxWater = currentWater
		}
		// Step 3; Move the lower bar's pointer
		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	 }

	 return maxWater
}
