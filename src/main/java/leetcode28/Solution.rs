/*
实现strStr()函数。

给你两个字符串haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回 -1 。



说明：

当needle是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。

对于本题而言，当needle是空字符串时我们应当返回 0 。这与 C 语言的strstr()以及 Java 的indexOf()定义相符。



示例 1：

输入：haystack = "hello", needle = "ll"
输出：2
示例 2：

输入：haystack = "aaaaa", needle = "bba"
输出：-1
示例 3：

输入：haystack = "", needle = ""
输出：0


提示：

0 <= haystack.length, needle.length <= 5 * 104
haystack 和 needle 仅由小写英文字符组成

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/implement-strstr
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
impl Solution {
    pub fn str_str(haystack: String, needle: String) -> i32 {
        // KMP
        let haystack = haystack.as_bytes();
        let needle = needle.as_bytes();
        let n = haystack.len();
        let m = needle.len();
        if m==0 {
            return 0;
        }
        // 构建next
        let mut next=vec![0;m];
        let mut j=0;
        for i in 1..m {
            while j > 0 && needle[i] != needle[j] {
                j=next[j-1];
            }
            if needle[i]==needle[j] {
                j+=1;
            }
            next[i]=j;
        }
        let mut j=0;
        for i in 0..n {
            while j > 0 && haystack[i] != needle[j] {
                j=next[j-1];
            }
            if haystack[i]==needle[j] {
                j+=1;
            }
            if j==m {
                return (i - m + 1) as i32;
            }
        }
        -1
    }
    pub fn str_str2(haystack: String, needle: String) -> i32 {
        let haystack = haystack.as_bytes();
        let needle = needle.as_bytes();
        let n = haystack.len();
        let m = needle.len();
        if m==0 {
            return 0;
        }
        for i in 0..n {
            if i+m>n {
                break;
            }
            let mut flag=true;
            for j in 0..m {
                if haystack[i+j]!=needle[j] {
                    flag=false;
                    break;
                }
            }
            if flag {
                return i as i32;
            }
        }
        -1
    }
}