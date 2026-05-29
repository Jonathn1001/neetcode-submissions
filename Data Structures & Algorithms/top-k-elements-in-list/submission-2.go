// Solution 1: Sorted Frequency O(nlogn)
// func topKFrequent(nums []int, k int) []int {
//     freqs := make(map[int]int) 
//     result := make([]int, 0, k)
//     // step 1: count frequency
//     for _, v := range nums {
//         freqs[v]++
//     }

//     // step 2: map -> slice then sort with descending frequency
//     type pair struct {
//         num int
//         freq int
//     }
//     pairs := make([]pair,0, len(freqs))
//     for num, freq  := range freqs {
//         pairs = append(pairs,pair{num, freq})
//     }
//     sort.Slice(pairs, func(i, j int) bool {
//         return pairs[i].freq > pairs[j].freq
//     })

//     // Step 3: get k elements frorm start
//     for i := 0; i < k; i++ {
//         result = append(result, pairs[i].num)
//     }

//     return result;
// }

// Solution 2; Bucket Sort
func topKFrequent(nums []int, k int) []int {    
    // Step 1: count frequency
    freqs := make(map[int]int)
    for _, n := range nums {
        freqs[n]++
    }

    // Step 2: create buckets[i] = slice contains frequency = i
    // need [n+1] buckets due to the max freq = n (index 0 -> n)
    buckets := make([][]int, len(nums)+1)
    for num, freq := range freqs {
        buckets[freq] = append(buckets[freq], num)
    }

    // Step 3: gather the top k elements
    result := make([]int, 0, k)
    for i := len(buckets) - 1; i >= 1 && len(result) < k; i-- {
        for _, num := range buckets[i] {
            result = append(result, num)
            if len(result) == k {
                return result
            }
        }
    }
    return result
} 