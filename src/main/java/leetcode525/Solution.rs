/*

给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。



示例 1:

输入: nums = [0,1]
输出: 2
说明: [0, 1] 是具有相同数量0和1的最长连续子数组。
示例 2:

输入: nums = [0,1,0]
输出: 2
说明: [0, 1] (或 [1, 0]) 是具有相同数量0和1的最长连续子数组。


提示：

1 <= nums.length <= 105
nums[i] 不是 0 就是 1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/contiguous-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

impl Solution {
    pub fn find_max_length(nums: Vec<i32>) -> i32 {
        let mut max_len = 0;
        let mut map = std::collections::HashMap::new();
        let mut counter = 0;
        map.insert(0, -1);
        let n = nums.len();
        for i in 0..n {
            let num = nums[i];
            if num == 1 {
                counter += 1;
            } else {
                counter -= 1;
            }
            let prev_index = map.entry(counter).or_insert(i as i32);
            max_len = max_len.max(i as i32 - *prev_index)
        }
        max_len
    }
}