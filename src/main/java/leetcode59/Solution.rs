/*

给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。



示例 1：
https://assets.leetcode.com/uploads/2020/11/13/spiraln.jpg

输入：n = 3
输出：[[1,2,3],[8,9,4],[7,6,5]]
示例 2：

输入：n = 1
输出：[[1]]


提示：

1 <= n <= 20

*/
impl Solution {
    pub fn generate_matrix(n: i32) -> Vec<Vec<i32>> {
        let n=n as usize;
        let mut res=vec![vec![0;n];n];
        let directions=vec![(0,1),(1,0),(0,-1),(-1,0)];
        let (mut row,mut col)=(0,0);
        let mut dir_index=0;
        for i in 1..=n * n {
            res[row][col]=i as i32;
            let next_row=row as i32+directions[dir_index].0;
            let next_col=col as i32+directions[dir_index].1;
            if next_row<0||next_row>= n as i32 ||next_col<0||next_col>=n as i32||res[next_row as usize][next_col as usize]!=0 {
                dir_index=(dir_index+1)%4
            }
            row = (row as i32 + directions[dir_index].0) as usize;
            col= (col as i32 + directions[dir_index].1) as usize;
        }
        res
    }
}