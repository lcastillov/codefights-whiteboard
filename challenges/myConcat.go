func myConcat(strings []string, separator string) string {
	answer := ""
	for _, s := range strings {
		answer += s + separator
	}
	return answer
}
