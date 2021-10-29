/*
给你一个 m 行 n 列的矩阵matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。



示例 1：
https://assets.leetcode.com/uploads/2020/11/13/spiral1.jpg

输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]
示例 2：
https://assets.leetcode.com/uploads/2020/11/13/spiral.jpg

输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]


提示：

m == matrix.length
n == matrix[i].length
1 <= m, n <= 10
-100 <= matrix[i][j] <= 100


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/spiral-matrix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
impl Solution {
    pub fn spiral_order(matrix: Vec<Vec<i32>>) -> Vec<i32> {
        if matrix.is_empty() || matrix[0].is_empty() { return vec![]; }

        let (mut row_start, mut row_end) = (0, matrix.len() - 1);
        let (mut col_start, mut col_end) = (0, matrix[0].len() - 1);
        let mut res = vec![];

        while row_start <= row_end && col_start <= col_end {
            (col_start..=col_end).for_each(|i| res.push(matrix[row_start][i]));
            (row_start + 1..row_end).for_each(|i| res.push(matrix[i][col_end]));
            if row_start != row_end {
                (col_start..=col_end).rev().for_each(|i| res.push(matrix[row_end][i]));
            }
            if col_start != col_end {
                (row_start + 1..row_end).rev().for_each(|i| res.push(matrix[i][col_start]))
            }

            row_start += 1;
            row_end = row_end.saturating_sub(1);
            col_start += 1;
            col_end = col_end.saturating_sub(1);
        }

        res
    }
}