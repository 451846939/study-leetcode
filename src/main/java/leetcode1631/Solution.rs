/*
你准备参加一场远足活动。给你一个二维rows x columns的地图heights，其中heights[row][col]表示格子(row, col)的高度。一开始你在最左上角的格子(0, 0)，且你希望去最右下角的格子(rows-1, columns-1)（注意下标从 0 开始编号）。你每次可以往 上，下，左，右四个方向之一移动，你想要找到耗费 体力 最小的一条路径。

一条路径耗费的 体力值是路径上相邻格子之间 高度差绝对值的 最大值决定的。

请你返回从左上角走到右下角的最小体力消耗值。



示例 1：
https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/10/25/ex1.png


输入：heights = [[1,2,2],[3,8,2],[5,3,5]]
输出：2
解释：路径 [1,3,5,3,5] 连续格子的差值绝对值最大为 2 。
这条路径比路径 [1,2,2,2,5] 更优，因为另一条路径差值最大值为 3 。
示例 2：
https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/10/25/ex2.png


输入：heights = [[1,2,3],[3,8,4],[5,3,5]]
输出：1
解释：路径 [1,2,3,4,5] 的相邻格子差值绝对值最大为 1 ，比路径 [1,3,5,3,5] 更优。
示例 3：
https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/10/25/ex3.png

输入：heights = [[1,2,1,1,1],[1,2,1,2,1],[1,2,1,2,1],[1,2,1,2,1],[1,1,1,2,1]]
输出：0
解释：上图所示路径不需要消耗任何体力。


提示：

rows == heights.length
columns == heights[i].length
1 <= rows, columns <= 100
1 <= heights[i][j] <= 106

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/path-with-minimum-effort
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
use std::cmp::Ordering;
use std::collections::BinaryHeap;
#[derive(Debug)]
pub struct Item(usize, usize, i32);

impl Eq for Item {}
impl PartialEq for Item {
    fn eq(&self, other: &Self) -> bool {
        self.2 == other.2
    }
}

impl PartialOrd for Item {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        other.2.partial_cmp(&self.2)
    }
}

impl Ord for Item {
    fn cmp(&self, other: &Item) -> Ordering {
        // self.cnt.cmp(&other.cnt)
        other.2.cmp(&self.2)
    }
}

impl Solution {
    pub fn minimum_effort_path(heights: Vec<Vec<i32>>) -> i32 {
        let dirs: Vec<Vec<i32>> = vec![vec![-1, 0], vec![1, 0], vec![0, -1], vec![0, 1]];
        let rows = heights.len();
        let cols = heights[0].len();
        let mut heap:BinaryHeap<Item> = std::collections::BinaryHeap::new();
        heap.push(Item(0,0,0));

        let mut dist = vec![vec![std::i32::MAX; cols]; rows];
        let mut seen = vec![vec![false; cols]; rows];
        dist[0][0] = 0;
        while let Some(Item(x,y,d)) = heap.pop() {
            if seen[x][y] {
                continue;
            }
            if x==rows-1&&y==cols-1 {
                break;
            }
            seen[x][y]=true;
            for dir in &dirs {
                let nx = dir[0]+x as i32;
                let ny = dir[1]+y as i32;
                if nx>=0&&nx<rows as i32&&ny>=0&&ny<cols as i32&&std::cmp::max(d,(heights[x][y] - heights[nx as usize][ny as usize]).abs())<dist[nx as usize][ny as usize] {
                    dist[nx as usize][ny as usize]=std::cmp::max(d,(heights[x][y] - heights[nx as usize][ny as usize]).abs());
                    heap.push(Item(nx as usize,ny as usize,dist[nx as usize][ny as usize]));
                }

            }
        }
        dist[rows-1][cols-1]
    }
}