func quasifactorial(n int) int {
	if n == 1 {
		return 1
	}
	return quasifactorial(n-1)*n - 1
}
