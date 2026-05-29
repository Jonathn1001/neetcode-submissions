func topKFrequent(nums []int, k int) []int {
    freqs := make(map[int]int) 
    result := make([]int, 0, k)
    // step 1: count frequency
    for _, v := range nums {
        freqs[v]++
    }

    // step 2: map -> slice then sort with descending frequency
    type pair struct {
        num int
        freq int
    }
    pairs := make([]pair,0, len(freqs))
    for num, freq  := range freqs {
        pairs = append(pairs,pair{num, freq})
    }
    sort.Slice(pairs, func(i, j int) bool {
        return pairs[i].freq > pairs[j].freq
    })

    // Step 3: get k elements frorm start
    for i := 0; i < k; i++ {
        result = append(result, pairs[i].num)
    }

    return result;
}
