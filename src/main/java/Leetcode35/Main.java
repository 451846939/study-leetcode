package Leetcode35;

/**
 * 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
 *
 * 你可以假设数组中无重复元素。
 *
 * 示例 1:
 *
 * 输入: [1,3,5,6], 5
 * 输出: 2
 * 示例 2:
 *
 * 输入: [1,3,5,6], 2
 * 输出: 1
 * 示例 3:
 *
 * 输入: [1,3,5,6], 7
 * 输出: 4
 * 示例 4:
 *
 * 输入: [1,3,5,6], 0
 * 输出: 0
 *
 * 来源：力扣（LeetCode）
 * 链接：https://leetcode-cn.com/problems/search-insert-position
 * 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
public class Main {
    public static void main(String[] args) {
        System.out.println(searchInsert(new int[]{1},1));
//        System.out.println(3/2);
    }
    public static int searchInsert(int[] nums, int target) {
        if (nums.length==0){
            return 0;
        }
        int n=nums.length;
        int l = 0;
        int i=0;
        if (nums[0]>=target){
            return 0;
        }
        if (nums[nums.length-1]<target){
            return nums.length;
        }
        while (i<=nums.length/2){
            i++;
            int tmp = (n + l) / 2;
            if (nums[tmp]==target){
                return tmp;
            }else if (nums[tmp]>target){
                if (i==nums.length/2){
                    return tmp;
                }
                n=tmp;
            }else {
                if (i==nums.length/2){
                    return tmp+1;
                }
                l=tmp;
            }

        }
        return i;
    }

}
