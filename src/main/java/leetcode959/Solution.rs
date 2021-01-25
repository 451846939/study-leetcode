/*
在由 1 x 1 方格组成的 N x N 网格grid 中，每个 1 x 1方块由 /、\ 或空格构成。这些字符会将方块划分为一些共边的区域。

（请注意，反斜杠字符是转义的，因此 \ 用 "\\"表示。）。

返回区域的数目。



示例 1：

输入：
[
 " /",
 "/ "
]
输出：2
解释：2x2 网格如下：
https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/15/1.png
示例 2：

输入：
[
 " /",
 "  "
]
输出：1
解释：2x2 网格如下：
https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/15/2.png
示例 3：

输入：
[
 "\\/",
 "/\\"
]
输出：4
解释：（回想一下，因为 \ 字符是转义的，所以 "\\/" 表示 \/，而 "/\\" 表示 /\。）
2x2 网格如下：
https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/15/3.png
示例 4：

输入：
[
 "/\\",
 "\\/"
]
输出：5
解释：（回想一下，因为 \ 字符是转义的，所以 "/\\" 表示 /\，而 "\\/" 表示 \/。）
2x2 网格如下：
https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/15/4.png
示例 5：

输入：
[
 "//",
 "/ "
]
输出：3
解释：2x2 网格如下：
https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/15/5.png


提示：

1 <= grid.length == grid[0].length <= 30
grid[i][j] 是'/'、'\'、或' '。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/regions-cut-by-slashes
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
impl Solution {
    pub fn regions_by_slashes(grid: Vec<String>) -> i32 {
        let n=grid.len();
        let size=n*n*4;
        let mut union_find = UnionFind::new(size);
        for i in 0..n {
            let row = grid[i].as_str().chars().collect::<Vec<char>>();
            for j in 0..n {
                let index=4*(i*n+j);
                let c = row[j];
                if c=='/' {
                    union_find.unite(index,index+3);
                    union_find.unite(index+1,index+2);
                }else if c=='\\' {
                    union_find.unite(index,index+1);
                    union_find.unite(index+2,index+3);
                }else {
                    union_find.unite(index,index+1);
                    union_find.unite(index+1,index+2);
                    union_find.unite(index+2,index+3);
                }
                if j+1<n {
                    union_find.unite(index+1,4*(i*n+j+1)+3);
                }
                if i+1<n {
                    union_find.unite(index+2,4*((i+1)*n+j));
                }
            }
        }
        union_find.set_count
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
            std::mem::swap(&mut x,&mut y)
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
