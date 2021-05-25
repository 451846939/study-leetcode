/*
给你一个整数数组 nums 和一个整数 k 。区间 [left, right]（left <= right）的 异或结果 是对下标位于left 和 right（包括 left 和 right ）之间所有元素进行 XOR 运算的结果：nums[left] XOR nums[left+1] XOR ... XOR nums[right] 。

返回数组中 要更改的最小元素数 ，以使所有长度为 k 的区间异或结果等于零。



示例 1：

输入：nums = [1,2,0,3,0], k = 1
输出：3
解释：将数组 [1,2,0,3,0] 修改为 [0,0,0,0,0]
示例 2：

输入：nums = [3,4,5,2,1,7,3,4,7], k = 3
输出：3
解释：将数组 [3,4,5,2,1,7,3,4,7] 修改为 [3,4,7,3,4,7,3,4,7]
示例 3：

输入：nums = [1,2,4,1,2,5,1,2,6], k = 3
输出：3
解释：将数组[1,2,4,1,2,5,1,2,6] 修改为 [1,2,3,1,2,3,1,2,3]


提示：

1 <= k <= nums.length <= 2000
0 <= nums[i] < 210


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/make-the-xor-of-all-segments-equal-to-zero
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
impl Solution {
    pub fn min_changes(nums: Vec<i32>, k: i32) -> i32 {
        use std::collections::HashMap;
        // x 的范围为 [0, 2^10)
        let MAXX=1<<10;
        // 极大值，为了防止整数溢出选择 INT_MAX / 2
        let INFTY=i32::max_value()>>1;
        let n = nums.len();
        let mut dp = vec![INFTY; MAXX];
        // 边界条件 f(-1,0)=0
        dp[0]=0;
        let k=k as usize;
        for i in 0..k {
            // 第 i 个组的哈希映射
            let mut cnt:HashMap<i32,i32>=HashMap::new();
            let mut size=0;
            for j in (i..n).step_by(k) {
                let entry = cnt.entry(nums[j]).or_insert(0);
                *entry+=1;
                size+=1;
            }
            let t2_min=dp.iter().min().unwrap();
            let mut g:Vec<i32> = vec![*t2_min; MAXX];
            for mask in 0..MAXX {
                for (x, cnt_x) in cnt.iter() {
                    g[mask as usize]=g[mask as usize].min(dp[(mask^x.clone() as usize) as usize]-cnt_x)
                }
            }
            for j in g.iter_mut() {
                *j+=size;
            }
            dp=g;
        }
        dp[0]
    }
}