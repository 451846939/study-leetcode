/*
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。

 

示例：

s = "leetcode"
返回 0

s = "loveleetcode"
返回 2

提示：你可以假定该字符串只包含小写字母。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/first-unique-character-in-a-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

impl Solution {
    pub fn first_uniq_char(s: String) -> i32 {
        let mut count=vec![0;26];
        let bytes = s.as_bytes();
        for x in bytes {
            count[(x-b'a') as usize]+=1;
        }
        for i in 0..bytes.len() {
            if count[(bytes[i]-b'a') as usize]==1 {
                return i as i32;
            }
        }
        -1
    }
}