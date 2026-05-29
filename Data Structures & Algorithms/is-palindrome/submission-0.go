func isPalindrome(s string) bool {
	// trimmed := strings.ReplaceAll(s, " ", "")
	// Compile regex: matches anything NOT in a-z, A-Z, or 0-9
	str := strings.ToLower(s)
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		panic(err)
	}
	// Replace all non-alphanumeric matches with empty string
	trimmed := reg.ReplaceAllString(str, "")
	i, j := 0, len(trimmed) - 1

	for i < j {
		if trimmed[i] != trimmed[j] {
			return false
		}
		i++
		j--
	}

	return true
}
