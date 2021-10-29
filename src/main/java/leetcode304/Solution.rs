/*
给定一个二维矩阵，计算其子矩形范围内元素的总和，该子矩阵的左上角为 (row1,col1) ，右下角为 (row2,col2) 。
https://assets.leetcode-cn.com/aliyun-lc-upload/images/304.png

上图子矩阵左上角(row1, col1) = (2, 1)，右下角(row2, col2) = (4, 3)，该子矩形内元素的总和为 8。



示例：

给定 matrix = [
  [3, 0, 1, 4, 2],
  [5, 6, 3, 2, 1],
  [1, 2, 0, 1, 5],
  [4, 1, 0, 1, 7],
  [1, 0, 3, 0, 5]
]

sumRegion(2, 1, 4, 3) -> 8
sumRegion(1, 1, 2, 2) -> 11
sumRegion(1, 2, 2, 4) -> 12


提示：

你可以假设矩阵不可变。
会多次调用sumRegion方法。
你可以假设row1 ≤ row2 且col1 ≤ col2 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/range-sum-query-2d-immutable
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
struct NumMatrix {
    sums:Vec<Vec<i32>>
}


/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl NumMatrix {

    fn new(matrix: Vec<Vec<i32>>) -> Self {
        if matrix.len()>0 {
            let mut sums = vec![vec![0; matrix[0].len()+1]; matrix.len()+1];
            for i in 0..matrix.len() {
                for j in 0..matrix[0].len() {
                    // 多加了一次所以减去sums[i][j]
                    sums[i+1][j+1]=sums[i+1][j]+sums[i][j+1]-sums[i][j]+matrix[i][j];
                }
            }
            NumMatrix{sums}
        }else {
            NumMatrix{ sums: vec![] }
        }
    }

    fn sum_region(&self, row1: i32, col1: i32, row2: i32, col2: i32) -> i32 {
        let sums = &self.sums;
        sums[row2 as usize + 1][col2 as usize + 1] - sums[row1 as usize][col2 as usize + 1] - sums[row2 as usize + 1][col1 as usize] + sums[row1 as usize][col1 as usize]
    }
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * let obj = NumMatrix::new(matrix);
 * let ret_1: i32 = obj.sum_region(row1, col1, row2, col2);
 */