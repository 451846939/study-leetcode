/*
给出一个字符串s（仅含有小写英文字母和括号）。

请你按照从括号内到外的顺序，逐层反转每对匹配括号中的字符串，并返回最终的结果。

注意，您的结果中 不应 包含任何括号。



示例 1：

输入：s = "(abcd)"
输出："dcba"
示例 2：

输入：s = "(u(love)i)"
输出："iloveu"
示例 3：

输入：s = "(ed(et(oc))el)"
输出："leetcode"
示例 4：

输入：s = "a(bcdefghijkl(mno)p)q"
输出："apmnolkjihgfedcbq"


提示：

0 <= s.length <= 2000
s 中只有小写英文字母和括号
我们确保所有括号都是成对出现的

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-substrings-between-each-pair-of-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
impl Solution {
    pub fn reverse_parentheses(s: String) -> String {
        let s = s.as_bytes();
        let n = s.len();
        let mut pair = vec![0; n];
        let mut stack = vec![];
        for i in 0..n {
            if s[i]==b'(' {
                stack.push(i);
            }else if s[i]==b')' {
                let j = stack.pop().unwrap();
                pair[i]=j;
                pair[j]=i;
            }
        }
        let mut str = String::new();
        let mut index:i32=0;
        let mut step=1;
        while index < n as i32 {
            if s[index as usize]==b'('||s[index as usize]==b')' {
                index= pair[index as usize] as i32;
                step=-step;
            }else {
                str.push(s[index as usize] as char)
            }
            index+=step
        }
        str.to_string()
    }
}