/*
给你一个由一些多米诺骨牌组成的列表dominoes。

如果其中某一张多米诺骨牌可以通过旋转 0度或 180 度得到另一张多米诺骨牌，我们就认为这两张牌是等价的。

形式上，dominoes[i] = [a, b]和dominoes[j] = [c, d]等价的前提是a==c且b==d，或是a==d 且b==c。

在0 <= i < j < dominoes.length的前提下，找出满足dominoes[i] 和dominoes[j]等价的骨牌对 (i, j) 的数量。



示例：

输入：dominoes = [[1,2],[2,1],[3,4],[5,6]]
输出：1


提示：

1 <= dominoes.length <= 40000
1 <= dominoes[i][j] <= 9

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-equivalent-domino-pairs
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

impl Solution {
    pub fn num_equiv_domino_pairs(mut dominoes: Vec<Vec<i32>>) -> i32 {
        let mut freq = vec![0;100];
        let mut count=0;
        for x in &dominoes {
            let mut i1 = x[0];
            let mut i2 = x[1];
            if i1 > i2 {
                std::mem::swap(&mut i1, &mut i2);
            }
            let num = i1 * 10 + i2;
            count+=freq[num as usize];
            freq[num as usize]+=1;
        }
        count
    }
}
