/*
实现一个基本的计算器来计算一个简单的字符串表达式 s 的值。



示例 1：

输入：s = "1 + 1"
输出：2
示例 2：

输入：s = " 2-1 + 2 "
输出：3
示例 3：

输入：s = "(1+(4+5+2)-3)+(6+8)"
输出：23


提示：

1 <= s.length <= 3* 105
s 由数字、'+'、'-'、'('、')'、和 ' ' 组成
s 表示一个有效的表达式

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/basic-calculator
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
impl Solution {
    pub fn calculate(s: String) -> i32 {
        let mut stack=vec![];
        stack.push(1);
        let mut sign=1;
        let mut s =s.chars().peekable();
        let mut ret=0;
        while let Some(c) = s.next() {
            match c {
                ' '=>{continue},
                '+'=>{
                    sign= *stack.last().unwrap();
                },
                '-'=>{
                    sign= -*stack.last().unwrap();
                },
                '('=>{
                    stack.push(sign)
                },
                ')'=>{
                    stack.pop();
                },
                _ => {
                    let mut n = (c as u8 - b'0') as i32;

                    while let Some(&x) = s.peek() {
                        if x.is_ascii_digit() {
                            n = n * 10 + (x as u8 - b'0') as i32;
                            s.next();
                        } else {
                            break;
                        }
                    }
                    ret+=n*sign;
                }
            }
        }
        ret
    }
}