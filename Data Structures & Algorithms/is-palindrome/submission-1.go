func isPalindrome(s string) bool {
	i, j := 0, len(s) - 1

	for i < j {
        // Get runes (or bytes if sticking to ASCII)
        rI := rune(s[i])
        rJ := rune(s[j])

        // 1. Skip non-alphanumeric characters from the left
        if !isAlphanumeric(rI) {
            i++
            continue
        }

        // 2. Skip non-alphanumeric characters from the right
        if !isAlphanumeric(rJ) {
            j--
            continue
        }

        // 3. Compared characters by converting to lowercase
		if unicode.ToLower(rI) != unicode.ToLower(rJ) {
			return false
		}

        // 4. Move both pointers closer
		i++
		j--
	}

	return true
}

// Helper function to check alphanumeric status without regex
func isAlphanumeric(r rune) bool {
    return (r >= 'a' && r <= 'z') || 
           (r >= 'A' && r <= 'Z') || 
           (r >= '0' && r <= '9')
}
