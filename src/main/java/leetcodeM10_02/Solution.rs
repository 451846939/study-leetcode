/*
编写一种方法，对字符串数组进行排序，将所有变位词组合在一起。变位词是指字母相同，但排列不同的字符串。

注意：本题相对原题稍作修改

示例:

输入: ["eat", "tea", "tan", "ate", "nat", "bat"],
输出:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
说明：

所有输入均为小写字母。
不考虑答案输出的顺序。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/group-anagrams-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
impl Solution {
    pub fn group_anagrams(strs: Vec<String>) -> Vec<Vec<String>> {
        let mut ans = vec![];
        let mut map = std::collections::HashMap::new();
        for s in &strs {
            let mut cnts = vec![0; 26];
            for x in s.as_bytes() {
                cnts[((x)-b'a') as usize]+=1;
            }
            let mut str="".to_string();
            for i in cnts.iter() {
                str+=(i.to_string().as_str().to_owned()+"_").as_str()
            }
            let key = str.to_string();
            let default = map.entry(key).or_insert(vec![]);
            default.push(s.clone());

        }
        for list in map.values() {
            ans.push(list.clone())
        }
        ans
    }
}