/*
统计一个数字在排序数组中出现的次数。



示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: 2
示例2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: 0


限制：

0 <= 数组长度 <= 50000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
impl Solution {
    pub fn search(nums: Vec<i32>, target: i32) -> i32 {
        let result = nums.binary_search(&target);
        let mut res=0;
        match result {
            Ok(index) => {
                res+=1;
                for i in index+1..nums.len() {
                    if nums[i]==nums[i-1]{
                        res+=1;
                    }else {
                        break
                    }
                }
                for i in (0..=index).rev() {
                    if i>0&&nums[i]==nums[i-1]{
                        res+=1;
                    }else {
                        break
                    }
                }
            }
            Err(_) => {return 0}
        }
        res
    }
}