/*
给你一个仅由大写英文字母组成的字符串，你可以将任意位置上的字符替换成另外的字符，总共可最多替换k次。在执行上述操作后，找到包含重复字母的最长子串的长度。

注意：字符串长度 和 k 不会超过104。



示例 1：

输入：s = "ABAB", k = 2
输出：4
解释：用两个'A'替换为两个'B',反之亦然。
示例 2：

输入：s = "AABABBA", k = 1
输出：4
解释：
将中间的一个'A'替换为'B',字符串变为 "AABBBBA"。
子串 "BBBB" 有最长重复字母, 答案为 4。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-repeating-character-replacement
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

impl Solution {
    pub fn character_replacement(s: String, k: i32) -> i32 {
        let mut map = vec![0; 26];
        let (mut left, mut right)=(0, 0);
        let x = s.as_bytes();
        let mut max=0;
        while right < x.len() {
            map[(x[right]-b'A') as usize]+=1;
            max = std::cmp::max(max, map[(x[right]-b'A') as usize]) as usize;
            if right-left+1-max> k as usize {
                map[(x[left]-b'A')as usize]-=1;
                left+=1;
            }
            right+=1;
        }
        (right-left) as i32
    }
}