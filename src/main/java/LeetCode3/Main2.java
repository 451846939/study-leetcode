package LeetCode3;

/**
 * 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
 *
 * 示例 1:
 *
 * 输入: "abcabcbb"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 * 示例 2:
 *
 * 输入: "bbbbb"
 * 输出: 1
 * 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
 * 示例 3:
 *
 * 输入: "pwwkew"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
 *      请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
 *
 * int [26] 用于字母 ‘a’ - ‘z’ 或 ‘A’ - ‘Z’
 * int [128] 用于ASCII码
 * int [256] 用于扩展ASCII码
 * 来源：力扣（LeetCode）
 * 链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
 * 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
public class Main2 {


    public static void main(String[] args) {
        System.out.println(lengthOfLongestSubstring("dvdf"));
    }

    public static int lengthOfLongestSubstring(String s) {
        int[] ints = s.chars().toArray();
        int [] map =new int[256];
        int l=0;
        int r=-1;
        int count=0;
        while (l<ints.length){
            if (r+1<ints.length&&map[ints[r+1]]==0){
                r++;
                map[ints[r]]=1;
            }else {
                map[ints[l]]=0;
                l++;
            }
            count=Math.max(count,r-l+1);
        }
        return count;
    }
}
