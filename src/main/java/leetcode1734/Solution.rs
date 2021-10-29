/*
给你一个整数数组perm，它是前n个正整数的排列，且n是个 奇数。

它被加密成另一个长度为 n - 1的整数数组encoded，满足encoded[i] = perm[i] XOR perm[i + 1]。比方说，如果perm = [1,3,2]，那么encoded = [2,1]。

给你encoded数组，请你返回原始数组perm。题目保证答案存在且唯一。



示例 1：

输入：encoded = [3,1]
输出：[1,2,3]
解释：如果 perm = [1,2,3] ，那么 encoded = [1 XOR 2,2 XOR 3] = [3,1]
示例 2：

输入：encoded = [6,5,4,6]
输出：[2,4,1,5,3]


提示：

3 <= n <105
n是奇数。
encoded.length == n - 1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/decode-xored-permutation
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
impl Solution {
    pub fn decode(encoded: Vec<i32>) -> Vec<i32> {
        let n=encoded.len();
        let mut total=0;
        for i in 1..=n+1 {
            total^=i as i32;
        }
        let mut odd=0;
        for i in (1..n).step_by(2) {
            odd^=encoded[i];
        }
        let mut perm = vec![0; n+1];
        perm[0]=total^odd;
        for i in 0..n  {
            perm[i+1]=perm[i]^encoded[i]
        }
        perm
    }
}