/*
编写一个高效的算法来判断m x n矩阵中，是否存在一个目标值。该矩阵具有如下特性：

每行中的整数从左到右按升序排列。
每行的第一个整数大于前一行的最后一个整数。


示例 1：
https://assets.leetcode.com/uploads/2020/10/05/mat.jpg

输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
输出：true
示例 2：
https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/11/25/mat2.jpg

输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
输出：false


提示：

m == matrix.length
n == matrix[i].length
1 <= m, n <= 100
-104 <= matrix[i][j], target <= 104

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/search-a-2d-matrix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
impl Solution {
    pub fn search_matrix(matrix: Vec<Vec<i32>>, target: i32) -> bool {
        let m=matrix.len();
        let n=matrix[0].len();
        let mut l =0;
        let mut r =m*n-1;
        while l <= r {
            let mid=l+((r-l)>>1);
            let x = matrix[mid / n][mid % n];
            if x==target {
                return true
            }else if x>target {
                if mid==0 {
                    return false
                }
                r=mid-1
            }else {
                l=mid+1
            }
        }
        false
    }
}