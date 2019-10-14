package leetcode88;

/**
 * 给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。
 * <p>
 * 说明:
 * <p>
 * 初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
 * 你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
 * 示例:
 * <p>
 * 输入:
 * nums1 = [1,2,3,0,0,0], m = 3
 * nums2 = [2,5,6],       n = 3
 * <p>
 * 输出: [1,2,2,3,5,6]
 * 来源：力扣（LeetCode）
 * 链接：https://leetcode-cn.com/problems/merge-sorted-array
 * 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
public class Main {
    public static void main(String[] args) {
        int []a={1,2,3,0,0,0};
        int []b={        2,5,6};
        merge(a,3,b,3);
    }

    public static void merge(int[] nums1, int m, int[] nums2, int n) {
        int[] tmp = new int[m];
        for (int i = 0; i < m; i++) {
            tmp[i] = nums1[i];
            nums1[i] = 0;
        }
        int tmpf = 0;
        int nums2f = 0;
        for (int i = 0; i < nums1.length; i++) {
            if (tmpf == tmp.length && nums2f < nums2.length) {
                nums1[i] = nums2[nums2f];
                nums2f++;
            } else if (nums2f == nums2.length && tmpf < tmp.length) {
                nums1[i] = tmp[tmpf];
                tmpf++;
            } else if (tmp[tmpf] <= nums2[nums2f]) {
                nums1[i] = tmp[tmpf];
                tmpf++;
            } else if (tmp[tmpf] >= nums2[nums2f]) {
                nums1[i] = nums2[nums2f];
                nums2f++;
            }
        }
    }
}
