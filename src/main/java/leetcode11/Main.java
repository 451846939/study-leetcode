package leetcode11;

/**
 * 给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。
 * 在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，
 * 使得它们与 x 轴共同构成的容器可以容纳最多的水。
 *
 * 说明：你不能倾斜容器，且 n 的值至少为 2。
 *
 *
 *
 * 图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
 *
 * https://aliyun-lc-upload.oss-cn-hangzhou.aliyuncs.com/aliyun-lc-upload/uploads/2018/07/25/question_11.jpg
 *
 * 示例:
 *
 * 输入: [1,8,6,2,5,4,8,3,7]
 * 输出: 49
 *
 * 来源：力扣（LeetCode）
 * 链接：https://leetcode-cn.com/problems/container-with-most-water
 * 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
public class Main {
    public static void main(String[] args) {
         int[] a ={1,8,6,2,5,4,8,3,7};
        System.out.println(maxArea(a));
    }

    public static int maxArea(int[] height) {
//        满足height[i]*(j-i)最大
        int i=0;
        int j=height.length-1;
        int temp=0;
        while (i<j){
            if (height[i]>height[j]){
                if (height[j]*(j-i)>temp){
                    temp=height[j]*(j-i);
                }
                j--;
            }else {
                if (height[i]*(j-i)>temp){
                    temp=height[i]*(j-i);
                }
                i++;
            }
        }
        return temp;
    }
}
