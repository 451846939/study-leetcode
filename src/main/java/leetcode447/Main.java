package leetcode447;

import java.util.HashMap;

/**
 * 给定平面上 n 对不同的点，“回旋镖” 是由点表示的元组 (i, j, k) ，其中 i 和 j 之间的距离和 i 和 k 之间的距离相等（需要考虑元组的顺序）。
 *
 * 找到所有回旋镖的数量。你可以假设 n 最大为 500，所有点的坐标在闭区间 [-10000, 10000] 中。
 *
 * 示例:
 *
 * 输入:
 * [[0,0],[1,0],[2,0]]
 *
 * 输出:
 * 2
 *
 * 解释:
 * 两个回旋镖为 [[1,0],[0,0],[2,0]] 和 [[1,0],[2,0],[0,0]]
 *
 * 来源：力扣（LeetCode）
 * 链接：https://leetcode-cn.com/problems/number-of-boomerangs
 * 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
public class Main {
    public static void main(String[] args) {
//        System.out.println(numberOfBoomerangs(new int[][]{{0,0},{1,0},{2,0}}));
        System.out.println(numberOfBoomerangs(new int[][]{{1,1},{2,2},{3,3}}));
    }

    public static int numberOfBoomerangs(int[][] points) {
        int res=0;
        for (int i = 0; i < points.length; i++) {
            HashMap<Integer, Integer> map = new HashMap<>();
            for (int j = 0; j < points.length; j++) {
                if (i!=j){
                    Integer orDefault = map.getOrDefault(dic(points[i], points[j]), 0);
                    map.put(dic(points[i], points[j]),orDefault+1);
                }
            }
            int[] values = map.values().stream().mapToInt(Integer::intValue).toArray();
            for (int k = 0; k < values.length; k++) {
                res+=values[k]*(values[k]-1);
            }
        }
        return res;
    }

    public static int dic(int[] point1,int[] point2){
        return (int) (Math.pow(point1[0]-point2[0],2)+Math.pow(point1[1]-point2[1],2));
    }
}
