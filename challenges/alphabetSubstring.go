func alphabetSubstring(s string) bool {
	for i := 1; i < len(s); i++ {
		if int(s[i])-int(s[i-1]) != 1 {
			return false
		}
	}
	return true
}
