func groupAnagrams(strs []string) [][]string {
    // map[key]list_of_anagrams
    // Tương đương với new Map() trong JS
    groups := make(map[string][]string)
    
    for _, s := range strs {
        // Bước 1: tạo canonical key bằng cách sort các ký tự
        bytes := []byte(s)
        sort.Slice(bytes, func(i, j int) bool {
            return bytes[i] < bytes[j]
        })
        key := string(bytes)
        
        // Bước 2: append chuỗi gốc vào group tương ứng
        groups[key] = append(groups[key], s)
    }
    
    // Bước 3: convert map values thành [][]string để trả về
    result := make([][]string, 0, len(groups))
    for _, group := range groups {
        result = append(result, group)
    }
    return result
}