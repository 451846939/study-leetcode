/*

Alice 和 Bob 共有一个无向图，其中包含 n 个节点和 3  种类型的边：

类型 1：只能由 Alice 遍历。
类型 2：只能由 Bob 遍历。
类型 3：Alice 和 Bob 都可以遍历。
给你一个数组 edges ，其中 edges[i] = [typei, ui, vi] 表示节点 ui 和 vi 之间存在类型为 typei 的双向边。请你在保证图仍能够被 Alice和 Bob 完全遍历的前提下，找出可以删除的最大边数。如果从任何节点开始，Alice 和 Bob 都可以到达所有其他节点，则认为图是可以完全遍历的。

返回可以删除的最大边数，如果 Alice 和 Bob 无法完全遍历图，则返回 -1 。



示例 1：
https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/09/06/5510ex1.png


输入：n = 4, edges = [[3,1,2],[3,2,3],[1,1,3],[1,2,4],[1,1,2],[2,3,4]]
输出：2
解释：如果删除 [1,1,2] 和 [1,1,3] 这两条边，Alice 和 Bob 仍然可以完全遍历这个图。再删除任何其他的边都无法保证图可以完全遍历。所以可以删除的最大边数是 2 。
示例 2：

https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/09/06/5510ex2.png

输入：n = 4, edges = [[3,1,2],[3,2,3],[1,1,4],[2,1,4]]
输出：0
解释：注意，删除任何一条边都会使 Alice 和 Bob 无法完全遍历这个图。
示例 3：

https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/09/06/5510ex3.png

输入：n = 4, edges = [[3,2,3],[1,1,2],[2,3,4]]
输出：-1
解释：在当前图中，Alice 无法从其他节点到达节点 4 。类似地，Bob 也不能达到节点 1 。因此，图无法完全遍历。


提示：

1 <= n <= 10^5
1 <= edges.length <= min(10^5, 3 * n * (n-1) / 2)
edges[i].length == 3
1 <= edges[i][0] <= 3
1 <= edges[i][1] < edges[i][2] <= n
所有元组 (typei, ui, vi) 互不相同
*/

impl Solution {
    pub fn max_num_edges_to_remove(n: i32, mut edges: Vec<Vec<i32>>) -> i32 {
        let mut edges = &mut edges;
        let mut ufa = UnionFind::new(n as usize);
        let mut ufb = UnionFind::new(n as usize);
        let mut ans = 0;
        for x in edges.into_iter() {
            x[1] -= 1;
            x[2] -= 1;
        }
        // 公共边
        for x in edges.into_iter() {
            if x[0] == 3 {
                if !ufa.unite(x[1] as usize, x[2] as usize) {
                    ans += 1;
                } else {
                    ufb.unite(x[1] as usize, x[2] as usize);
                }
            }
        }
        // 独占边
        for x in edges.into_iter() {
            if x[0] == 1 {
                if !ufa.unite(x[1] as usize, x[2] as usize) {
                    ans += 1;
                }
            } else if x[0] == 2 {
                if !ufb.unite(x[1] as usize, x[2] as usize) {
                    ans += 1;
                }
            }
        }
        if ufa.set_count != 1 || ufb.set_count != 1 {
            return -1;
        }
        ans
    }
}

pub struct UnionFind {
    parent: Vec<usize>,
    size: Vec<usize>,
    n: usize,
    set_count: i32,
}

impl UnionFind {
    pub fn new(n: usize) -> Self {
        UnionFind { parent: (0..n).collect(), size: vec![1; n], n, set_count: n as i32 }
    }
    fn findset(&mut self, x: usize) -> usize {
        let mut vec = &mut self.parent;
        UnionFind::find(x, &mut vec)
    }

    fn find(x: usize, vec: &mut Vec<usize>) -> usize {
        if vec[x] != x {
            vec[x] = UnionFind::find(vec[x], vec);
        }
        vec[x]
    }

    fn unite(&mut self, mut x: usize, mut y: usize) -> bool {
        x = self.findset(x);
        y = self.findset(y);
        if x == y {
            return false;
        }
        if self.size[x] < self.size[y] {
            std::mem::swap(&mut x, &mut y)
        }
        self.parent[y] = x;
        self.size[x] += self.size[y];
        self.set_count -= 1;
        true
    }

    fn connected(&mut self, x: usize, y: usize) -> bool {
        self.findset(x) == self.findset(y)
    }
}