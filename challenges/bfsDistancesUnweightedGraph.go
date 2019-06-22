func bfsDistancesUnweightedGraph(matrix [][]bool, startVertex int) []int {
	n := len(matrix)
	d := make([]int, n)
	q := make([]int, n)
	h := 0
	t := 1
	q[0] = startVertex
	d[startVertex] = 1
	for h < t {
		currentVertex := q[h]
		for i := 0; i < n; i++ {
			if matrix[currentVertex][i] && d[i] == 0 {
				d[i] = d[currentVertex] + 1
				q[t] = i
				t += 1
			}
		}
		h += 1
	}
	for i := 0; i < n; i++ {
		d[i] -= 1
	}
	return d
}
