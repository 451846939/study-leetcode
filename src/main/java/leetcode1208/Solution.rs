/*
给你两个长度相同的字符串，s 和 t。

将 s中的第i个字符变到t中的第 i 个字符需要|s[i] - t[i]|的开销（开销可能为 0），也就是两个字符的 ASCII 码值的差的绝对值。

用于变更字符串的最大预算是maxCost。在转化字符串时，总开销应当小于等于该预算，这也意味着字符串的转化可能是不完全的。

如果你可以将 s 的子字符串转化为它在 t 中对应的子字符串，则返回可以转化的最大长度。

如果 s 中没有子字符串可以转化成 t 中对应的子字符串，则返回 0。



示例 1：

输入：s = "abcd", t = "bcdf", cost = 3
输出：3
解释：s 中的 "abc" 可以变为 "bcd"。开销为 3，所以最大长度为 3。
示例 2：

输入：s = "abcd", t = "cdef", cost = 3
输出：1
解释：s 中的任一字符要想变成 t 中对应的字符，其开销都是 2。因此，最大长度为 1。
示例 3：

输入：s = "abcd", t = "acde", cost = 0
输出：1
解释：你无法作出任何改动，所以最大长度为 1。


提示：

1 <= s.length, t.length <= 10^5
0 <= maxCost <= 10^6
s 和t都只含小写英文字母。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/get-equal-substrings-within-budget
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

impl Solution {
    pub fn equal_substring(s: String, t: String, max_cost: i32) -> i32 {
        let s = s.as_bytes();
        let t = t.as_bytes();
        let mut len =0;
        let mut left = 0;   // 窗口左边界
        let mut cost = 0;   // 当前窗口消耗
        // i作为窗口右边界
        for i in 0..s.len() {
            cost +=abs(s, t, i);
            // 如果当前窗口消耗大于总开销，则左边界++，缩减窗口
            if cost as i32 > max_cost {
                cost -= abs(s,t,left);
                left += 1;
            }
            len = std::cmp::max(len, i - left + 1)
        }
        return len as i32;
    }
}
fn abs(s: &[u8], t: &[u8], i: usize) -> i32 {
    (s[i] as i32 - t[i] as i32).abs()
}