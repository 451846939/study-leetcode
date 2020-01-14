package leetcode994;


import com.eclipsesource.json.JsonArray;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.LinkedList;

/**
 * 在给定的网格中，每个单元格可以有以下三个值之一：
 *
 * 值 0 代表空单元格；
 * 值 1 代表新鲜橘子；
 * 值 2 代表腐烂的橘子。
 * 每分钟，任何与腐烂的橘子（在 4 个正方向上）相邻的新鲜橘子都会腐烂。
 *
 * 返回直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1。
 *https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2019/02/16/oranges.png
 *  
 *
 * 示例 1：
 *
 *
 *
 * 输入：[[2,1,1],[1,1,0],[0,1,1]]
 * 输出：4
 * 示例 2：
 *
 * 输入：[[2,1,1],[0,1,1],[1,0,1]]
 * 输出：-1
 * 解释：左下角的橘子（第 2 行， 第 0 列）永远不会腐烂，因为腐烂只会发生在 4 个正向上。
 * 示例 3：
 *
 * 输入：[[0,2]]
 * 输出：0
 * 解释：因为 0 分钟时已经没有新鲜橘子了，所以答案就是 0 。
 *  
 *
 * 提示：
 *
 * 1 <= grid.length <= 10
 * 1 <= grid[0].length <= 10
 * grid[i][j] 仅为 0、1 或 2
 *
 * 来源：力扣（LeetCode）
 * 链接：https://leetcode-cn.com/problems/rotting-oranges
 * 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
class Solution {
    LinkedList<Point> bedPoint = new LinkedList<>();
    int[][] points;
    class Point{
        public int x;
        public int y;
        public int count;
        @Override
        public String toString() {
            return x+":"+y;
        }
    }

    private boolean next(Point point){
        boolean flag=false;
//        left
        if (point.x-1>=0&&point.x-1< points.length&& points[point.x-1][point.y]==1){
//            count++;
            Point newPoint = new Point();
            newPoint.x=point.x-1;
            newPoint.y=point.y;
            points[newPoint.x][newPoint.y]=2;
            newPoint.count=point.count+1;
            bedPoint.add(newPoint);
            flag=true;
        }
//        up
        if (point.y-1>=0&&point.y-1< points[0].length&& points[point.x][point.y-1]==1){
//            count++;
            Point newPoint = new Point();
            newPoint.x=point.x;
            newPoint.y=point.y-1;
            points[newPoint.x][newPoint.y]=2;
            newPoint.count=point.count+1;
            bedPoint.add(newPoint);
            flag=true;
        }
//        right
        if (point.x+1>=0&&point.x+1< points.length&& points[point.x+1][point.y]==1){
//            count++;
            Point newPoint = new Point();
            newPoint.x=point.x+1;
            newPoint.y=point.y;
            points[newPoint.x][newPoint.y]=2;
            newPoint.count=point.count+1;
            bedPoint.add(newPoint);
            flag=true;
        }
        //        down
        if (point.y+1>=0&&point.y+1< points[0].length&& points[point.x][point.y+1]==1){
//            count++;
            Point newPoint = new Point();
            newPoint.x=point.x;
            newPoint.y=point.y+1;
            points[newPoint.x][newPoint.y]=2;
            newPoint.count=point.count+1;
            bedPoint.add(newPoint);
            flag=true;
        }
        return flag;
    }
    public int orangesRotting(int[][] grid) {
        int count=0;
        bedPoint.clear();
        this.points =grid;
        for (int i = 0; i < points.length; i++) {
            for (int j = 0; j < points[i].length; j++) {
                if (points[i][j]==2){
                    Point point = new Point();
                    point.x=i;
                    point.y=j;
                    point.count=0;
                    bedPoint.add(point);
                }
            }
        }
        while (!bedPoint.isEmpty()){
            Point remove = bedPoint.remove();
            next(remove);
            count=remove.count;
        }
        for (int[] row: points)
            for (int v: row)
                if (v == 1)
                    return -1;
        return count;
    }

}

public class MainClass {
    public static int[] stringToIntegerArray(String input) {
        input = input.trim();
        input = input.substring(1, input.length() - 1);
        if (input.length() == 0) {
          return new int[0];
        }

        String[] parts = input.split(",");
        int[] output = new int[parts.length];
        for(int index = 0; index < parts.length; index++) {
            String part = parts[index].trim();
            output[index] = Integer.parseInt(part);
        }
        return output;
    }

    public static int[][] stringToInt2dArray(String input) {
        JsonArray jsonArray = JsonArray.readFrom(input);
        if (jsonArray.size() == 0) {
          return new int[0][0];
        }

        int[][] arr = new int[jsonArray.size()][];
        for (int i = 0; i < arr.length; i++) {
          JsonArray cols = jsonArray.get(i).asArray();
          arr[i] = stringToIntegerArray(cols.toString());
        }
        return arr;
    }

    public static void main(String[] args) throws IOException {
        BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
        String line;
        while ((line = in.readLine()) != null) {
            int[][] grid = stringToInt2dArray(line);

            int ret = new Solution().orangesRotting(grid);

            String out = String.valueOf(ret);

            System.out.print(out);
        }
    }
}