/*
给你一个 n个点的带权无向连通图，节点编号为 0到 n-1，同时还有一个数组 edges，其中 edges[i] = [fromi, toi, weighti]表示在fromi和toi节点之间有一条带权无向边。最小生成树(MST) 是给定图中边的一个子集，它连接了所有节点且没有环，而且这些边的权值和最小。

请你找到给定图中最小生成树的所有关键边和伪关键边。如果从图中删去某条边，会导致最小生成树的权值和增加，那么我们就说它是一条关键边。伪关键边则是可能会出现在某些最小生成树中但不会出现在所有最小生成树中的边。

请注意，你可以分别以任意顺序返回关键边的下标和伪关键边的下标。



示例 1：
https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/06/21/ex1.png


输入：n = 5, edges = [[0,1,1],[1,2,1],[2,3,2],[0,3,2],[0,4,3],[3,4,3],[1,4,6]]
输出：[[0,1],[2,3,4,5]]
解释：上图描述了给定图。
下图是所有的最小生成树。
https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/06/21/msts.png
注意到第 0 条边和第 1 条边出现在了所有最小生成树中，所以它们是关键边，我们将这两个下标作为输出的第一个列表。
边 2，3，4 和 5 是所有 MST 的剩余边，所以它们是伪关键边。我们将它们作为输出的第二个列表。
示例 2 ：

https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/06/21/ex2.png

输入：n = 4, edges = [[0,1,1],[1,2,1],[2,3,1],[0,3,1]]
输出：[[],[0,1,2,3]]
解释：可以观察到 4 条边都有相同的权值，任选它们中的 3 条可以形成一棵 MST 。所以 4 条边都是伪关键边。


提示：

2 <= n <= 100
1 <= edges.length <= min(200, n * (n - 1) / 2)
edges[i].length == 3
0 <= fromi < toi < n
1 <= weighti<= 1000
所有 (fromi, toi)数对都是互不相同的。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-critical-and-pseudo-critical-edges-in-minimum-spanning-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
use std::cell::{RefCell, Ref};
use std::mem::swap;
impl Solution {
    pub fn find_critical_and_pseudo_critical_edges(n: i32, mut edges: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let m = edges.len();
        for i in 0..m {
            edges[i].push(i as i32);
        }
        edges.sort_by(|u,v|{
            let tmp = u[2] - v[2];
            return if tmp == 0 {
                std::cmp::Ordering::Equal
            } else if tmp < 0 {
                std::cmp::Ordering::Less
            } else {
                std::cmp::Ordering::Greater
            }
        });
        let mut uf_std = UnionFind::new(n as usize);
        let mut value=0;
        for i in 0..m {
            if uf_std.unite(edges[i][0] as usize,edges[i][1] as usize) {
                value+=edges[i][2];
            }
        }
        let mut ans = vec![vec![]; 2];
        for i in 0..m {
            // 判断是否是关键边
            let mut uf = UnionFind::new(n as usize);
            let mut v=0;
            for j in 0..m {
                if i != j && uf.unite(edges[j][0] as usize, edges[j][1] as usize) {
                    v += edges[j][2];
                }
            }
            if uf.set_count != 1 || (uf.set_count == 1 && v > value) {
                ans[0].push(edges[i][3]);
                continue;
            }
            // 判断是否是伪关键边
            let mut uf = UnionFind::new(n as usize);
            uf.unite(edges[i][0] as usize, edges[i][1] as usize);
            v = edges[i][2];
            for j in 0..m {
                if i != j && uf.unite(edges[j][0] as usize, edges[j][1] as usize) {
                    v += edges[j][2];
                }
            }
            if v == value {
                ans[1].push(edges[i][3]);
            }
        }
        ans
    }
}

pub struct UnionFind{
    parent:Vec<usize>,
    size:Vec<usize>,
    n:usize,
    set_count:i32
}

impl UnionFind {
    pub fn new( n: usize) -> Self {
        UnionFind { parent: (0..n).collect(),size: vec![1;n], n, set_count:n as i32 }
    }
    fn findset(&mut self, x:usize) ->usize{
        let mut vec = &mut self.parent;
        UnionFind::find(x, &mut vec)
    }

    fn find( x: usize, vec: &mut Vec<usize>) -> usize {
        if vec[x] != x {
            vec[x] = UnionFind::find( vec[x],vec);
        }
        vec[x]
    }

    fn unite(&mut self,mut x:usize,mut y:usize)->bool{
        x = self.findset(x);
        y = self.findset(y);
        if x==y {
            return false;
        }
        if self.size[x]<self.size[y] {
            swap(&mut x,&mut y)
        }
        self.parent[y]=x;
        self.size[x]+=self.size[y];
        self.set_count-=1;
        true
    }

    fn connected(&mut self, x:usize, y:usize) ->bool{
        self.findset(x)==self.findset(y)
    }
}