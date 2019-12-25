package leetcode1041;

/**
 * 在无限的平面上，机器人最初位于 (0, 0) 处，面朝北方。机器人可以接受下列三条指令之一：
 *
 * "G"：直走 1 个单位
 * "L"：左转 90 度
 * "R"：右转 90 度
 * 机器人按顺序执行指令 instructions，并一直重复它们。
 *
 * 只有在平面中存在环使得机器人永远无法离开时，返回 true。否则，返回 false。
 *
 *  
 *
 * 示例 1：
 *
 * 输入："GGLLGG"
 * 输出：true
 * 解释：
 * 机器人从 (0,0) 移动到 (0,2)，转 180 度，然后回到 (0,0)。
 * 重复这些指令，机器人将保持在以原点为中心，2 为半径的环中进行移动。
 * 示例 2：
 *
 * 输入："GG"
 * 输出：false
 * 解释：
 * 机器人无限向北移动。
 * 示例 3：
 *
 * 输入："GL"
 * 输出：true
 * 解释：
 * 机器人按 (0, 0) -> (0, 1) -> (-1, 1) -> (-1, 0) -> (0, 0) -> ... 进行移动。
 *  
 *
 * 提示：
 *
 * 1 <= instructions.length <= 100
 * instructions[i] 在 {'G', 'L', 'R'} 中
 *
 * 来源：力扣（LeetCode）
 * 链接：https://leetcode-cn.com/problems/robot-bounded-in-circle
 * 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
public class Main{
    public static void main(String[] args) {
        System.out.println(new Solution().isRobotBounded("GL"));
    }
        }

/**
 * 1、若终点和起点坐标一致，则返回true；
 * 2、若终点不一致，此时判断起点与终点的方向关系：
 * （1）不一致，则一定可以在有限次执行指令后回到原点，返回true；
 * （2）一致，则无限远离起点，返回false；
 *只要走完一轮后，方向改变，即不是直走的话，最后无论再走多少轮总有一轮会走回终点。
 */
class Solution {
    public boolean isRobotBounded(String instructions) {
        Point p = new Point(0,0,0);
        for(int i=0; i<instructions.length(); i++) {
            p.move(instructions.charAt(i));
        }
        if (p.getX()==0 && p.getY()==0) {
            return true;
        }else if (p.getDerec() == 0) {
            return false;
        }
        return true;
    }
}

class Point {
    private int x;
    private int y;
    private int derec;//方向
    public Point(int x, int y, int derec) {
        this.x = x;
        this.y = y;
        this.derec = derec;
    }
    public void setX(int x) {
        this.x = x;
    }

    public void setY(int y) {
        this.y = y;
    }
    public void setDerec(int derec) {
        this.derec = derec;
    }
    public int getX() {
        return this.x;
    }
    public int getY() {
        return this.y;
    }
    public int getDerec() {
        return this.derec;
    }
    public void move(char c) {
        if (c == 'L') {
            derec=(derec+1)%4;//左转
        }else if (c == 'R') {
            derec = (derec+3)%4;//右转
        }else if (c == 'G') {
            switch (derec) {//在不同方向下前进一步的坐标变化
                case 0: this.x++;break;
                case 1: this.y++;break;
                case 2: this.x--;break;
                case 3: this.y--;break;
            }
        }
    }
}

