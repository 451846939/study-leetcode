///给你一个数组pairs ，其中pairs[i] = [xi, yi]，并且满足：
/// 
/// pairs中没有重复元素
/// xi < yi
/// 令ways为满足下面条件的有根树的方案数：
/// 
/// 树所包含的所有节点值都在 pairs中。
/// 一个数对[xi, yi] 出现在pairs中当且仅当xi是yi的祖先或者yi是xi的祖先。
/// 注意：构造出来的树不一定是二叉树。
/// 两棵树被视为不同的方案当存在至少一个节点在两棵树中有不同的父节点。
/// 
/// 请你返回：
/// 
/// 如果ways == 0，返回0。
/// 如果ways == 1，返回 1。
/// 如果ways > 1，返回2。
/// 一棵 有根树指的是只有一个根节点的树，所有边都是从根往外的方向。
/// 
/// 我们称从根到一个节点路径上的任意一个节点（除去节点本身）都是该节点的 祖先。根节点没有祖先。
/// 
/// 
/// 
/// 示例 1：
/// https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2021/01/09/trees2.png
/// 
/// 输入：pairs = [[1,2],[2,3]]
/// 输出：1
/// 解释：如上图所示，有且只有一个符合规定的有根树。
/// 示例 2：
/// https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2021/01/09/tree.png
/// 
/// 输入：pairs = [[1,2],[2,3],[1,3]]
/// 输出：2
/// 解释：有多个符合规定的有根树，其中三个如上图所示。
/// 示例 3：
/// 
/// 输入：pairs = [[1,2],[2,3],[2,4],[1,5]]
/// 输出：0
/// 解释：没有符合规定的有根树。
/// 
/// 
/// 提示：
/// 
/// 1 <= pairs.length <= 105
/// 1 <= xi < yi <= 500
/// pairs中的元素互不相同。
/// 
/// 来源：力扣（LeetCode）
/// 链接：https:///leetcode-cn.com/problems/number-of-ways-to-reconstruct-a-tree
/// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

impl Solution {
    pub fn check_ways(pairs: Vec<Vec<i32>>) -> i32 {
        let mut  adj:std::collections::HashMap<i32,Vec<i32>> = std::collections::HashMap::new();
        for p in pairs {
            let entry1 = adj.entry(p[0]);
            entry1.or_default().push(p[1]);
            let entry2 = adj.entry(p[1]);
            entry2.or_default().push(p[0]);
        }
        let mut root=-1;
        let adj_len = adj.len();
        for (k, v) in &adj {
            let node=k;
            if v.len()== adj_len -1 {
                root=*node;
            }
        }
        if root==-1 {
            return 0;
        }
        let mut res=1;
        for (k, v) in adj.iter() {
            let node=k;
            let neighbours=v;
            if *node==root {
                continue
            }
            let curr_degree = neighbours.len() as i32;
            let mut parent=-1;
            let mut parent_degree =std::i32::MAX;
            for neighbour in neighbours.iter() {
                if adj.get(&neighbour).unwrap().len()< parent_degree as usize &&adj.get(&neighbour).unwrap().len()>= curr_degree as usize {
                    parent=*neighbour;
                    parent_degree= adj.get(&neighbour).unwrap().len() as i32;
                }
            }
            if parent==-1 {
                return 0;
            }
            for neighbour in neighbours.iter() {
                if *neighbour==parent {
                    continue
                }
                if !adj.get(&parent).unwrap().contains(&neighbour) {
                    return 0;
                }
            }
            if parent_degree== curr_degree {
                res=2;
            }
        }
        res
    }
}