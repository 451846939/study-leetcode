/*
给定一种规律 pattern 和一个字符串 str ，判断 str 是否遵循相同的规律。

这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 str 中的每个非空单词之间存在着双向连接的对应规律。

示例1:

输入: pattern = "abba", str = "dog cat cat dog"
输出: true
示例 2:

输入:pattern = "abba", str = "dog cat cat fish"
输出: false
示例 3:

输入: pattern = "aaaa", str = "dog cat cat dog"
输出: false
示例 4:

输入: pattern = "abba", str = "dog dog dog dog"
输出: false
说明:
你可以假设 pattern 只包含小写字母， str 包含了由单个空格分隔的小写字母。 

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/word-pattern
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
use std::collections::{VecDeque, HashMap};
impl Solution {
    pub fn word_pattern(pattern: String, s: String) -> bool {
        let mut p_ms:HashMap<String,u8> = HashMap::new();
        let mut s_mp:HashMap<u8,String> = HashMap::new();
        let split = s.split(" ");
        let cstr = pattern.into_bytes();

        let mut si=0;
        for x in split {
            if si>=cstr.len() {
                return false
            }
            let ss = p_ms.entry(x.into()).or_insert(cstr[si]);
            let pp = s_mp.entry(cstr[si]).or_insert(x.into());
            if !cstr[si].eq(ss)|| !x.eq(pp) { return false;}
            si+=1;
        }
        return si==cstr.len();
    }
    pub fn word_pattern2(pattern: String, s: String) -> bool {
        let mut p_ms:HashMap<String,u8> = HashMap::new();
        let mut s_mp:HashMap<u8,String> = HashMap::new();
        let split = s.split(" ").collect::<Vec<&str>>();
        let cstr = pattern.into_bytes();
        if split.len()!=cstr.len() {
            return false;
        }
        let mut si=0;
        for x in split {
            let ss = p_ms.entry(x.into()).or_insert(cstr[si]);
            let pp = s_mp.entry(cstr[si]).or_insert(x.into());
            if !cstr[si].eq(ss)|| !x.eq(pp) { return false;}
            si+=1;
        }
        return true;
    }
}