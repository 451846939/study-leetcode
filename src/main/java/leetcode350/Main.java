package leetcode350;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;

/**
 * 给定两个数组，编写一个函数来计算它们的交集。
 * <p>
 * 示例 1:
 * <p>
 * 输入: nums1 = [1,2,2,1], nums2 = [2,2]
 * 输出: [2,2]
 * 示例 2:
 * <p>
 * 输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
 * 输出: [4,9]
 * 说明：
 * <p>
 * 输出结果中每个元素出现的次数，应与元素在两个数组中出现的次数一致。
 * 我们可以不考虑输出结果的顺序。
 * 进阶:
 * <p>
 * 如果给定的数组已经排好序呢？你将如何优化你的算法？
 * 如果 nums1 的大小比 nums2 小很多，哪种方法更优？
 * 如果 nums2 的元素存储在磁盘上，磁盘内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？
 * <p>
 * 来源：力扣（LeetCode）
 * 链接：https://leetcode-cn.com/problems/intersection-of-two-arrays-ii
 * 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
public class Main {
    public static void main(String[] args) {
//        System.out.println(Arrays.toString(intersect(new int[]{4,9,5}, new int[]{9,4,9,8,4})));
        System.out.println(Arrays.toString(intersect(new int[]{1,2,2,1}, new int[]{2,2})));
    }

    public static int[] intersect(int[] nums1, int[] nums2) {
        int num = Math.max(nums1.length, nums2.length);
        ArrayList<Integer> objects = new ArrayList<>();
        HashMap<Integer,Integer> hash = new HashMap(num+1);
        for (int i = 0; i < nums1.length; i++) {
            Integer orDefault = hash.getOrDefault(nums1[i], 0);
            hash.put(nums1[i],orDefault+1);
        }
        for (int i = 0; i < nums2.length; i++) {
            int now = hash.getOrDefault(nums2[i],0);
            if (now!=0){
                hash.put(nums2[i],now-1);
                objects.add(nums2[i]);
            }
        }
        return objects.stream().mapToInt(Integer::intValue).toArray();
    }
}
