/*
数组中占比超过一半的元素称之为主要元素。给你一个 整数 数组，找出其中的主要元素。若没有，返回 -1 。请设计时间复杂度为 O(N) 、空间复杂度为 O(1) 的解决方案。



示例 1：

输入：[1,2,5,9,5,9,5,5,5]
输出：5
示例 2：

输入：[3,2]
输出：-1
示例 3：

输入：[2,2,1,1,1,2,2]
输出：2

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-majority-element-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
impl Solution {
    pub fn majority_element(nums: Vec<i32>) -> i32 {
        let mut ans =0;
        let mut count=0;
        for x in &nums {
            if count==0 {
                count+=1;
                ans=*x
            }else {
                if ans==*x {
                    count+=1;
                }else {
                    count-=1;
                }
            }
        }
        let mut verification=0;
        for x in &nums {
            if *x==ans {
                verification+=1
            }
        }
        if verification<=nums.len()>>1 {
            ans=-1
        }
        ans
    }
}