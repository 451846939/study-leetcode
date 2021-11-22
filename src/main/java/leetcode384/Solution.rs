/*

给你一个整数数组 nums ，设计算法来打乱一个没有重复元素的数组。

实现 Solution class:

Solution(int[] nums) 使用整数数组 nums 初始化对象
int[] reset() 重设数组到它的初始状态并返回
int[] shuffle() 返回数组随机打乱后的结果


示例：

输入
["Solution", "shuffle", "reset", "shuffle"]
[[[1, 2, 3]], [], [], []]
输出
[null, [3, 1, 2], [1, 2, 3], [1, 3, 2]]

解释
Solution solution = new Solution([1, 2, 3]);
solution.shuffle();    // 打乱数组 [1,2,3] 并返回结果。任何 [1,2,3]的排列返回的概率应该相同。例如，返回 [3, 1, 2]
solution.reset();      // 重设数组到它的初始状态 [1, 2, 3] 。返回 [1, 2, 3]
solution.shuffle();    // 随机返回数组 [1, 2, 3] 打乱后的结果。例如，返回 [1, 3, 2]


提示：

1 <= nums.length <= 200
-106 <= nums[i] <= 106
nums 中的所有元素都是 唯一的
最多可以调用 5 * 104 次 reset 和 shuffle
*/



/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
use std::time::{SystemTime, UNIX_EPOCH};
struct Solution {
    nums: Vec<i32>
}
impl Solution {
    fn new(nums: Vec<i32>) -> Self {
        Self { nums }
    }
    fn reset(&self) -> Vec<i32> {
        self.nums.clone()
    }
    fn shuffle(&self) -> Vec<i32> {
        let mut nums = self.nums.clone();
        let mut n = nums.len() as u128;
        while n > 1 {
            let m = SystemTime::now().duration_since(UNIX_EPOCH).unwrap().as_nanos() % n;
            nums.swap(m as usize, n as usize - 1);
            n -= 1;
        }
        nums
    }
}


/**
 * Your Solution object will be instantiated and called as such:
 * let obj = Solution::new(nums);
 * let ret_1: Vec<i32> = obj.reset();
 * let ret_2: Vec<i32> = obj.shuffle();
 */