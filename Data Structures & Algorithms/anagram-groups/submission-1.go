func groupAnagrams(strs []string) [][]string {
    // English letters through a-z that they have corresponding ASCII code from 97 - 122
    // map[K]v : K = [26]int, v = []string
    groups := make(map[[26]int][]string) 

    for _, s := range strs {
        var count [26]int // init = [0,0,...,0]
        for i := 0; i < len(s); i++ {
            // a = 97 -> every char from a-z minus a will has the value from 0 - 25.
            count[s[i] - 'a']++  // create a key is an array has 26 element from 0 - 25.
        }
        groups[count] = append(groups[count], s) 
    }

    result := make([][]string, 0, len(groups))
    for _, group := range groups {
        result = append(result, group)
    }
    return result
}
