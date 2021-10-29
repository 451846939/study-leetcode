/*
有一个长度为arrLen的数组，开始有一个指针在索引0 处。

每一步操作中，你可以将指针向左或向右移动 1 步，或者停在原地（指针不能被移动到数组范围外）。

给你两个整数steps 和arrLen ，请你计算并返回：在恰好执行steps次操作以后，指针仍然指向索引0 处的方案数。

由于答案可能会很大，请返回方案数 模10^9 + 7 后的结果。



示例 1：

输入：steps = 3, arrLen = 2
输出：4
解释：3 步后，总共有 4 种不同的方法可以停在索引 0 处。
向右，向左，不动
不动，向右，向左
向右，不动，向左
不动，不动，不动
示例 2：

输入：steps = 2, arrLen = 4
输出：2
解释：2 步后，总共有 2 种不同的方法可以停在索引 0 处。
向右，向左
不动，不动
示例 3：

输入：steps = 4, arrLen = 2
输出：8


提示：

1 <= steps <= 500
1 <= arrLen<= 10^6

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-ways-to-stay-in-the-same-place-after-some-steps
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

impl Solution {
    pub fn num_ways(steps: i32, arr_len: i32) -> i32 {
        let MODULO=1000000007;
        let max_column = steps.min(arr_len-1);
        let mut dp = vec![vec![0; (max_column+1) as usize]; (steps + 1) as usize];
        dp[0][0]=1;
        for i in 1..=steps as usize {
            for j in 0..=max_column as usize {
                dp[i][j]=dp[i-1][j];
                if j>=1 {
                    dp[i][j]=(dp[i][j]+dp[i-1][j-1])%MODULO;
                }
                if j+1<= max_column as usize {
                    dp[i][j]=(dp[i][j]+dp[i-1][j+1])%MODULO;
                }
            }
        }
        dp[steps as usize][0]
    }
}