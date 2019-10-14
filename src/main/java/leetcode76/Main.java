package leetcode76;

import java.util.HashMap;
import java.util.Map;

/**
 * 给你一个字符串 S、一个字符串 T，请在字符串 S 里面找出：包含 T 所有字母的最小子串。
 * <p>
 * 示例：
 * <p>
 * 输入: S = "ADOBECODEBANC", T = "ABC"
 * 输出: "BANC"
 * 说明：
 * <p>
 * 如果 S 中不存这样的子串，则返回空字符串 ""。
 * 如果 S 中存在这样的子串，我们保证它是唯一的答案。
 * <p>
 * 来源：力扣（LeetCode）
 * 链接：https://leetcode-cn.com/problems/minimum-window-substring
 * 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
public class Main {
    public static void main(String[] args) {
        System.out.println(minWindow("ADOBECODEBANC","ABC"));
        System.out.println(minWindow("a","a"));
        System.out.println(minWindow("ab","b"));
        System.out.println(minWindow("aa","aa"));
        System.out.println(minWindow("abc","aabc"));
        System.out.println(minWindow("abcc","aabc"));
        System.out.println(minWindow("aas","s"));
        System.out.println(minWindow("bdab","ab"));
//        String s="avbs";
//        System.out.println(s.substring(2,3+1));
    }
    public static String minWindow(String s, String t) {
        if (s==null||"".equals(s)){
            return "";
        }
        Map<Character, Integer> ts = new HashMap<>(t.length());
        for (int r = 0; r < t.length(); r++) {
            Integer orDefault = ts.getOrDefault(t.charAt(r), 0);
            ts.put(t.charAt(r), orDefault+1);
        }
        int r = -1, l = 0;
        int mindlenth=s.length()+1;
        int havelenth=0;
        int ansL=l;
        int ansR=r+1;
        Map<Character, Integer> ss = new HashMap<>(s.length());
        while (l<s.length()){
            if (r+1<s.length()&&havelenth<t.length()){
                ss.put(s.charAt(r+1),ss.getOrDefault(s.charAt(r+1),0)+1);
                if (ts.containsKey(s.charAt(r+1))&&ss.get(s.charAt(r+1)).intValue()<=ts.get(s.charAt(r+1)).intValue()){
                    havelenth++;
                }
                r++;
            }else {
                if (havelenth==t.length()&&r-l+1<mindlenth){
                    mindlenth=r-l+1;
                    ansL=l;
                    ansR=r;
                }
                ss.put(s.charAt(l),ss.get(s.charAt(l))-1);
                if (ts.containsKey(s.charAt(l))&&ss.get(s.charAt(l)).intValue()<ts.get(s.charAt(l)).intValue()){
                    havelenth--;
                }
                l++;
            }

        }
        if (mindlenth==s.length()+1){
            return "";
        }
        return s.substring(ansL,ansR+1);
    }
}
