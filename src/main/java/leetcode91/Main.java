package leetcode91;

/**
 * 一条包含字母 A-Z 的消息通过以下方式进行了编码：
 * <p>
 * 'A' -> 1
 * 'B' -> 2
 * ...
 * 'Z' -> 26
 * 给定一个只包含数字的非空字符串，请计算解码方法的总数。
 * <p>
 * 示例 1:
 * <p>
 * 输入: "12"
 * 输出: 2
 * 解释: 它可以解码为 "AB"（1 2）或者 "L"（12）。
 * 示例 2:
 * <p>
 * 输入: "226"
 * 输出: 3
 * 解释: 它可以解码为 "BZ" (2 26), "VF" (22 6), 或者 "BBF" (2 2 6) 。
 * <p>
 * 来源：力扣（LeetCode）
 * 链接：https://leetcode-cn.com/problems/decode-ways
 * 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
public class Main {
    public static void main(String[] args) {
        System.out.println(numDecodings("110"));
    }


    public static int numDecodings(String s) {
//         动态规划
        if (s == null || s.length() == 0) {
            return 0;
        }
        int[] res = new int[s.length() + 1];
        res[s.length()]=1;
        if (s.charAt(s.length()-1)=='0'){
            res[s.length()-1]=0;
        }else {
            res[s.length()-1]=1;
        }
        for (int i = s.length()-2; i >=0; i--) {
            if (s.charAt(i) == '0') {
                res[i] = 0;
                continue;
            }
            if (Integer.parseInt(s.substring(i , i + 2)) > 26) {
                res[i] = res[i + 1];
            } else {

                res[i] = res[i + 1] + res[i + 2];

            }
        }
        return res[0];
    }
}
