/*
给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是回文。

返回符合要求的 最少分割次数 。



示例 1：

输入：s = "aab"
输出：1
解释：只需一次分割就可将s 分割成 ["aa","b"] 这样两个回文子串。
示例 2：

输入：s = "a"
输出：0
示例 3：

输入：s = "ab"
输出：1


提示：

1 <= s.length <= 2000
s 仅由小写英文字母组成

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-partitioning-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

impl Solution {
    pub fn min_cut(s: String) -> i32 {
        let s=s.as_bytes();
        let mut dp1=vec![vec![true;s.len()];s.len()];
        for i in (0..s.len()).rev() {
            for j in i + 1..s.len() {
                dp1[i][j]=(s[i]==s[j])&dp1[i+1][j-1]
            }
        }
        let mut dp2 =vec![std::i32::MAX; s.len()];
        for i in 0..s.len() {
            if dp1[0][i] {
                dp2[i]=0
            }else {
                for j in 0..i {
                    if dp1[j+1][i] {
                        dp2[i]=dp2[i].min(dp2[j]+1)
                    }
                }
            }
        }
        dp2[s.len()-1]
    }
}