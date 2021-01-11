/*
给你一个字符串 s，以及该字符串中的一些「索引对」数组 pairs，其中 pairs[i] = [a, b] 表示字符串中的两个索引（编号从 0 开始）。

你可以 任意多次交换 在 pairs 中任意一对索引处的字符。

返回在经过若干次交换后，s 可以变成的按字典序最小的字符串。

 

示例 1:

输入：s = "dcab", pairs = [[0,3],[1,2]]
输出："bacd"
解释：
交换 s[0] 和 s[3], s = "bcad"
交换 s[1] 和 s[2], s = "bacd"
示例 2：

输入：s = "dcab", pairs = [[0,3],[1,2],[0,2]]
输出："abcd"
解释：
交换 s[0] 和 s[3], s = "bcad"
交换 s[0] 和 s[2], s = "acbd"
交换 s[1] 和 s[2], s = "abcd"
示例 3：

输入：s = "cba", pairs = [[0,1],[1,2]]
输出："abc"
解释：
交换 s[0] 和 s[1], s = "bca"
交换 s[1] 和 s[2], s = "bac"
交换 s[0] 和 s[1], s = "abc"
 

提示：

1 <= s.length <= 10^5
0 <= pairs.length <= 10^5
0 <= pairs[i][0], pairs[i][1] < s.length
s 中只含有小写英文字母

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/smallest-string-with-swaps
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

impl Solution {
    pub fn smallest_string_with_swaps(s: String, pairs: Vec<Vec<i32>>) -> String {
        if pairs.is_empty() {
            return s;
        }

        let len = s.len();

        let mut parent = (0..len).collect::<Vec<usize>>();
        let mut rank = vec![1usize; len];
        fn find(parent: &mut Vec<usize>,x :usize)->usize{
            if parent[x]!=x {
                parent[x]=find(parent,parent[x]);
            }
            parent[x]
        }
        fn union(parent:&mut Vec<usize>,rank:&mut Vec<usize>,x:usize,y:usize){
            let root_x = find(parent, x);
            let root_y = find(parent, y);
            if root_x == root_y {
                return;
            }
            if rank[root_x] == rank[root_y] {
                parent[root_x] = root_y;
                rank[root_y] += 1;
            } else if rank[root_x] < rank[root_y] {
                parent[root_x] = parent[root_y]
            } else {
                parent[root_y] = parent[root_x]
            }
        }

        for pair in pairs {
            union(&mut parent,&mut rank,pair[0] as usize,pair[1] as usize);
        }

        let s = s.into_bytes();
        let mut map = std::collections::HashMap::<usize, std::collections::BinaryHeap<std::cmp::Reverse<u8>>>::new();
        for i in 0..len {
            let root = find(&mut parent, i);
            if map.contains_key(&root) {
                map.get_mut(&root).unwrap().push(std::cmp::Reverse(s[i]));
            }else {
                let mut min_heap = std::collections::BinaryHeap::<std::cmp::Reverse<u8>>::new();
                min_heap.push(std::cmp::Reverse(s[i]));
                map.insert(root,min_heap);
            }
        }

        let mut ret = Vec::<u8>::with_capacity(len);
        for i in 0..len {
            let root = find(&mut parent, i);
            let c = map.get_mut(&root).unwrap().pop().unwrap().0;
            ret.push(c);
        }

        String::from_utf8(ret).unwrap()
    }
}