package main

func main() {

}
func merge(A []int, m int, B []int, n int) {
	i := 0
	j := 0
	tmp := make([]int, m)
	for k := 0; k < m; k++ {
		tmp[k] = A[k]
	}
	for k := 0; k < len(A); k++ {
		if i < m && j < n {
			if tmp[i] <= B[j] {
				A[k] = tmp[i]
				i++
			} else {
				A[k] = B[j]
				j++
			}
		} else if i == m {
			A[k] = B[j]
			j++
		} else {
			A[k] = tmp[i]
			i++
		}
	}
}
