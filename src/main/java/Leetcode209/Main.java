package Leetcode209;

/**
 * 给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的连续子数组。如果不存在符合条件的连续子数组，返回 0。
 *
 * 示例: 
 *
 * 输入: s = 7, nums = [2,3,1,2,4,3]
 * 输出: 2
 * 解释: 子数组 [4,3] 是该条件下的长度最小的连续子数组。
 * 进阶:
 *
 * 如果你已经完成了O(n) 时间复杂度的解法, 请尝试 O(n log n) 时间复杂度的解法。
 *
 * 来源：力扣（LeetCode）
 * 链接：https://leetcode-cn.com/problems/minimum-size-subarray-sum
 * 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
public class Main {
    public static void main(String[] args) {
//        int s = 7;
//        int[] nums = {2,3,1,2,4,3};
        int s = 15;
        int[] nums = {1,2,3,4,5};
        System.out.println(minSubArrayLen(s,nums));
    }

    public static int minSubArrayLen(int s, int[] nums) {
        int l=0;
        int r=0;
        int count=0;
        while (l<nums.length){
            int temp=0;
            for (int i = 0; i < r - l; i++) {
                temp=temp+nums[l+i];
            }
            if (temp>=s){
                if (r-l>0&&count>r-l||count==0){
                    count=r-l;
                }
                l++;
            }else if (r<nums.length){
                r++;
            }else {
                l++;
            }
        }
        return count;
    }
}
