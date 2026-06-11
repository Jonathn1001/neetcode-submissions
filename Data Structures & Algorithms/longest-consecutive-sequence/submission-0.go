func longestConsecutive(nums []int) int {
    numSet := make(map[int]struct{}, len(nums))
    for _, num := range nums {
        numSet[num] = struct{}{}
    }

    longest := 0
    for num := range numSet {
        if _, exists := numSet[num-1]; !exists {
            currentNum := num
            currentLength := 1

            for {
                if _, ok := numSet[currentNum+1]; !ok {
                    break
                }
                currentNum++
                currentLength++
            }

            if currentLength > longest {
                longest = currentLength
            }
        }
    }

    return longest
}
