package Leetcode120;

import java.util.List;
import java.util.stream.Collectors;
import java.util.stream.Stream;

/**
 * 给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
 *
 * 例如，给定三角形：
 *
 * [
 *      [2],
 *     [3,4],
 *    [6,5,7],
 *   [4,1,8,3]
 * ]
 * 自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
 *
 * 说明：
 *
 * 如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。
 *
 * 来源：力扣（LeetCode）
 * 链接：https://leetcode-cn.com/problems/triangle
 * 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
public class Main {
    public static void main(String[] args) {
//        ,,,
//        System.out.println(minimumTotal(Stream.of(Stream.of(2).collect(Collectors.toList()),Stream.of(3,4).collect(Collectors.toList()),Stream.of(6,5,7).collect(Collectors.toList()),Stream.of(4,1,8,3).collect(Collectors.toList())).collect(Collectors.toList())));
//        [  [-1],
        //  [2, 3],
        // [1,-1,-3]]
        System.out.println(minimumTotal(Stream.of(Stream.of(-1).collect(Collectors.toList()),Stream.of(2,3).collect(Collectors.toList()),Stream.of(1,-1,-3).collect(Collectors.toList())).collect(Collectors.toList())));
    }

    public static int minimumTotal(List<List<Integer>> triangle) {
    //需要自下向顶思考，因为从下向上，每一层都是确定的，从上向下每下面一层都无法确定
        int n = triangle.size();
        int res[]=new int[n+1];
       //res[j] 使用的时候默认是上一层的，赋值之后变成当前层
        for (int i = n-1; i >=0 ; i--) {
            List<Integer> integers = triangle.get(i);
            for (int j = 0; j < integers.size(); j++) {
                res[j]=Math.min(res[j],res[j+1])+integers.get(j);
            }
        }
        return res[0];
    }

}
