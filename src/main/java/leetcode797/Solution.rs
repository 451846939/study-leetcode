/*
给你一个有n个节点的 有向无环图（DAG），请你找出所有从节点 0到节点 n-1的路径并输出（不要求按特定顺序）

二维数组的第 i 个数组中的单元都表示有向图中 i 号节点所能到达的下一些节点，空就是没有下一个结点了。

译者注：有向图是有方向的，即规定了 a→b 你就不能从 b→a 。



示例 1：
https://assets.leetcode.com/uploads/2020/09/28/all_1.jpg


输入：graph = [[1,2],[3],[3],[]]
输出：[[0,1,3],[0,2,3]]
解释：有两条路径 0 -> 1 -> 3 和 0 -> 2 -> 3
示例 2：
https://assets.leetcode.com/uploads/2020/09/28/all_2.jpg


输入：graph = [[4,3,1],[3,2,4],[3],[4],[]]
输出：[[0,4],[0,3,4],[0,1,3,4],[0,1,2,3,4],[0,1,4]]
示例 3：

输入：graph = [[1],[]]
输出：[[0,1]]
示例 4：

输入：graph = [[1,2,3],[2],[3],[]]
输出：[[0,1,2,3],[0,2,3],[0,3]]
示例 5：

输入：graph = [[1,3],[2],[3],[]]
输出：[[0,1,2,3],[0,3]]


提示：

n == graph.length
2 <= n <= 15
0 <= graph[i][j] < n
graph[i][j] != i（即，不存在自环）
graph[i] 中的所有元素 互不相同
保证输入为 有向无环图（DAG）

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/all-paths-from-source-to-target
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

impl Solution {
    pub fn all_paths_source_target(graph: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        // 队列
        let mut queue = std::collections::VecDeque::new();
        // 初始节点
        queue.push_back(vec![0]);
        // 一个 vector 用于保存结果的
        let mut result = vec![];
        while let Some(path) = queue.pop_front() {
            let end = path[path.len()-1];
            if end == (graph.len()-1) as i32 {
                result.push(path);
            } else {
                for next in &graph[end as usize] {
                    let mut path1 = path.clone();
                    path1.push(*next);
                    queue.push_back(path1);
                }
            }
        }
        result
    }
}