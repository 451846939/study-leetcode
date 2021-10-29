/*
你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。

给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，能够偷窃到的最高金额。



示例1：

输入：nums = [2,3,2]
输出：3
解释：你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。
示例 2：

输入：nums = [1,2,3,1]
输出：4
解释：你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
    偷窃到的最高金额 = 1 + 3 = 4 。
示例 3：

输入：nums = [0]
输出：0


提示：

1 <= nums.length <= 100
0 <= nums[i] <= 1000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/house-robber-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
impl Solution {
    pub fn rob(nums: Vec<i32>) -> i32 {
        let mut rob_fst=(0,0,nums[0]);
        let mut rob_lst=(0,0,0);
        for i in 1..nums.len(){
            if i!=nums.len()-1 {
                rob_fst=(rob_fst.1,rob_fst.2,nums[i]+rob_fst.0.max(rob_fst.1));
            }
            rob_lst=(rob_lst.1,rob_lst.2,nums[i]+rob_lst.0.max(rob_lst.1));
        }
        rob_fst.1.max(rob_fst.2).max(rob_lst.1).max(rob_lst.2)
    }
    pub fn rob2(nums: Vec<i32>) -> i32 {
        let n = nums.len();
        if n==1 {
            return nums[0];
        }else if n==2 {
            return nums[0].max(nums[1]);
        }
        fn _rob(robs:&[i32])->i32{
            let mut sum=robs[0].max(robs[1]);
            let mut first=robs[0];
            for v in robs.split_at(2).1 {
                let tmp=sum;
                sum=sum.max(first+v);
                first=tmp;
            }
            sum
        }
        _rob(nums.split_at(n-1).0).max(_rob(nums.split_at(1).1))
    }
}