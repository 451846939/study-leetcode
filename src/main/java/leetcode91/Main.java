//package leetcode91;
//
//import java.util.HashMap;
//
///**
// * 一条包含字母 A-Z 的消息通过以下方式进行了编码：
// * <p>
// * 'A' -> 1
// * 'B' -> 2
// * ...
// * 'Z' -> 26
// * 给定一个只包含数字的非空字符串，请计算解码方法的总数。
// * <p>
// * 示例 1:
// * <p>
// * 输入: "12"
// * 输出: 2
// * 解释: 它可以解码为 "AB"（1 2）或者 "L"（12）。
// * 示例 2:
// * <p>
// * 输入: "226"
// * 输出: 3
// * 解释: 它可以解码为 "BZ" (2 26), "VF" (22 6), 或者 "BBF" (2 2 6) 。
// * <p>
// * 来源：力扣（LeetCode）
// * 链接：https://leetcode-cn.com/problems/decode-ways
// * 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
// */
//public class Main {
//    public static void main(String[] args) {
//        System.out.println(numDecodings("230"));
//    }
//
//
//    public static int numDecodings(String s) {
////        1122
////       1  1  0  2  2
////todo 动态规划
//        if (s.length()==1&&Integer.valueOf(s.substring(0,1))==0){
//            return 0;
//        }
//        int n = s.length();
//        int[] ss = new int[n + 2];
//        for (int i = 1; i < n + 1; i++) {
//            ss[i] = Integer.valueOf(s.substring(i-1, i));// 1 2 3 4    12 3 4   1 23 4
//        }
//        int[] res = new int[n + 2];
////        1                res[1]=1
////        1 2              res[2]=res[1]+2
////        1 2 3            res[3]=res[2]+2
////        1 2 3 4          res[4]=res[]
//        HashMap<Object, Character> objectObjectHashMap = new HashMap<>();
//        objectObjectHashMap.getOrDefault("", Character.MIN_VALUE);
//        for (int i = 0; i+1 <= ss.length;) {
//            i++;
//            if (i+1==n+1){
//                break;
//            }
//            if (i == 1) {
//                if (ss[i] == 0)
//                    return 0;
//                if ( ss[i] > 6|| (ss[i + 1] > 6 || ss[i + 1] < 1)) {
//                    res[i] = 1;
//                }else {
//                    res[i] = 2;
//                }
//                continue;
//            }
//            if (n % 2 == 1 && (ss[i] == 0&&ss[i+1]==0 )||( ss[i] > 6 && ss[i + 1] > 6)) {
//                return 0;
//            }
//            if (i % 2 == 1) {
//                if (ss[i] <= 6 && ss[i] > 0 && ss[i + 1] >= 1 && ss[i + 1] <= 6) {
//                    res[i] = res[i - 1] + 1;
//                } else {
//                    res[i] = res[i - 1]-1;
//                }
//            } else {
//                if (ss[i] <= 6 && ss[i] > 0 && ss[i + 1] >= 1 && ss[i + 1] <= 6) {
//                    res[i] = res[i - 1] + res[i - 2];
//                } else {
//                    res[i] = res[i - 1]-1;
//                }
//            }
//
//        }
//        if (n%2==1&&Integer.valueOf(s.substring(n-1,n))==0&&Integer.valueOf(s.substring(n-2,n))>26){
//            return res[n-1];
//        }
//        if (n%2==1){
//            return res[n-1]+1;
//        }else {
//            return res[n-1];
//        }
//
//
////        if (ss.length==2&&ss[1]>=1){
////            return 1;
////        }
////        if (ss.length>3&&ss[ss.length-1]==0&&ss[ss.length-2]==0){
////            return 0;
////        }
////
////        for (int i = 1; i <= n; i++) {
////            if (i==1&&ss[i]!=0){
////                if (ss[i+1]==0){
////                    i++;
////                    res[i]=2;
////                }else {
////                    res[i]=1;
////                }
////                continue;
////            }else if (i == 1){
////                break;
////            }
////            if (i==2){
////                if (ss[i]<=6&&ss[i]>0&&ss[i-1]>=1&&ss[i-1]<=6){
////                    res[i]=res[i-1]+2;
////                }else {
////                    res[i]=res[i-1]+1;
////                }
////                continue;
////            }
////
////            if (i%2==1){
////                if (ss[i]<=6&&ss[i]>0&&ss[i-1]>=1&&ss[i-1]<=6){
////                    res[i]=res[i-1]+1;
////                }else {
////                    res[i]=res[i-1];
////                }
////            }else{
////                if (ss[i]<=6&&ss[i]>0&&ss[i-1]>=1&&ss[i-1]<=6){
////                    res[i]=res[i-1]+res[i-2];
////                }else {
////                    res[i]=res[i-1];
////                }
////            }
////
////        }
////        if ( res[n]-1>=0){
////            return res[n]-1;
////        }else {
////            return 0;
////        }
//    }
//}
