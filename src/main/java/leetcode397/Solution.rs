/*
给定一个正整数n ，你可以做如下操作：

如果n是偶数，则用n / 2替换n 。
如果n是奇数，则可以用n + 1或n - 1替换n 。
n变为 1 所需的最小替换次数是多少？



示例 1：

输入：n = 8
输出：3
解释：8 -> 4 -> 2 -> 1
示例 2：

输入：n = 7
输出：4
解释：7 -> 8 -> 4 -> 2 -> 1
或 7 -> 6 -> 3 -> 2 -> 1
示例 3：

输入：n = 4
输出：2


提示：

1 <= n <= 231 - 1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/integer-replacement
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
impl Solution {
    pub fn integer_replacement(n: i32) -> i32 {
        let mut n = n as i64;
        let mut ans = 0;
        while n != 1 {
            if n & 1 == 0 {
                n = n >> 1;
            } else {
                if n != 3 && ((n >> 1) & 1 == 1) {
                    n+=1;
                }else {
                    n-=1;
                }
            }
            ans += 1;
        }
        ans
    }
}