/*
给定一个非空且只包含非负数的整数数组nums，数组的度的定义是指数组里任一元素出现频数的最大值。

你的任务是在 nums 中找到与nums拥有相同大小的度的最短连续子数组，返回其长度。



示例 1：

输入：[1, 2, 2, 3, 1]
输出：2
解释：
输入数组的度是2，因为元素1和2的出现频数最大，均为2.
连续子数组里面拥有相同度的有如下所示:
[1, 2, 2, 3, 1], [1, 2, 2, 3], [2, 2, 3, 1], [1, 2, 2], [2, 2, 3], [2, 2]
最短连续子数组[2, 2]的长度为2，所以返回2.
示例 2：

输入：[1,2,2,3,1,4,2]
输出：6


提示：

nums.length在1到 50,000 区间范围内。
nums[i]是一个在 0 到 49,999 范围内的整数。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/degree-of-an-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
use std::collections::HashMap;

impl Solution {
    pub fn find_shortest_sub_array(nums: Vec<i32>) -> i32 {
        let mut map = HashMap::new();
        let mut degree=0;
        for i in 0..nums.len() {
            let call = map.entry(nums[i]).or_insert((0, i, 1));
            call.0+=1;
            call.2=i-call.1+1;
            degree=degree.max(call.0);
        }
        map.values().filter(|e| e.0==degree).map(|e|e.2).min().unwrap() as i32
    }
}