package main

/*
输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。



示例 1：

输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]
示例 2：

输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]


限制：

0 <= matrix.length <= 100
0 <= matrix[i].length <= 100

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shun-shi-zhen-da-yin-ju-zhen-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {

}
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	step := 0
	size := len(matrix) * len(matrix[0])
	//定义四个方向的端点
	top, bottom, right, left := 0, len(matrix)-1, len(matrix[0])-1, 0
	nums := make([]int, size)
	for step < size {
		//从左到右
		for i := left; i <= right && step < size; i++ {
			nums[step] = matrix[top][i]
			step++
		}
		top++
		//从上到下
		for i := top; i <= bottom && step < size; i++ {
			nums[step] = matrix[i][right]
			step++
		}
		right--
		//从右到左
		for i := right; i >= left && step < size; i-- {
			nums[step] = matrix[bottom][i]
			step++
		}
		bottom--
		//从下到上
		for i := bottom; i >= top && step < size; i-- {
			nums[step] = matrix[i][left]
			step++
		}
		left++
	}
	return nums
}
