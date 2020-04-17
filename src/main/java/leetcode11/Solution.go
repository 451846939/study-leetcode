package main

func main() {

}
func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	tmp := 0
	for l < r {
		if height[l] > height[r] {
			tmp = max(height[r]*(r-l), tmp)
			r--
		} else {
			tmp = max(height[l]*(r-l), tmp)
			l++
		}
	}
	return tmp
}
func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
