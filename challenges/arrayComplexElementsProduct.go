func arrayComplexElementsProduct(real []int, imag []int) []int {
	r := 1
	i := 0
	for k := 0; k < len(real); k++ {
		rr := r*real[k] - i*imag[k]
		ii := r*imag[k] + i*real[k]
		r = rr
		i = ii
	}
	return []int{r, i}
}
