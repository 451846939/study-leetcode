/*
几乎每一个人都用乘法表。但是你能在乘法表中快速找到第k小的数字吗？

给定高度m、宽度n 的一张m * n的乘法表，以及正整数k，你需要返回表中第k小的数字。

例1：

输入: m = 3, n = 3, k = 5
输出: 3
解释: 
乘法表:
1	2	3
2	4	6
3	6	9

第5小的数字是 3 (1, 2, 2, 3, 3).
例 2：

输入: m = 2, n = 3, k = 6
输出: 6
解释: 
乘法表:
1	2	3
2	4	6

第6小的数字是 6 (1, 2, 2, 3, 4, 6).
注意：

m 和n的范围在 [1, 30000] 之间。
k 的范围在 [1, m * n] 之间。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/kth-smallest-number-in-multiplication-table
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

impl Solution {
    pub fn find_kth_number(m: i32, n: i32, k: i32) -> i32 {
        let mut left = 1;
        let mut right = m * n;
        fn count(m: i32, n: i32,k: i32) -> i32{
            let mut  res =0;
            for i in 1..=m {
                res+=(k/i).min(n);
            }
            res
        }
        while left < right{
            let mid = left + ((right - left) >> 1);
            if count(m,n,mid)>=k {
                right = mid;
            }else {
                left = mid + 1;
            }
        }
        left
    }
}