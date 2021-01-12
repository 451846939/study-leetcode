/*

公司共有 n 个项目和  m 个小组，每个项目要不无人接手，要不就由 m 个小组之一负责。

group[i] 表示第 i 个项目所属的小组，如果这个项目目前无人接手，那么 group[i] 就等于 -1。（项目和小组都是从零开始编号的）小组可能存在没有接手任何项目的情况。

请你帮忙按要求安排这些项目的进度，并返回排序后的项目列表：

同一小组的项目，排序后在列表中彼此相邻。
项目之间存在一定的依赖关系，我们用一个列表 beforeItems 来表示，其中 beforeItems[i] 表示在进行第 i 个项目前（位于第 i 个项目左侧）应该完成的所有项目。
如果存在多个解决方案，只需要返回其中任意一个即可。如果没有合适的解决方案，就请返回一个 空列表 。



示例 1：



输入：n = 8, m = 2, group = [-1,-1,1,0,0,1,0,-1], beforeItems = [[],[6],[5],[6],[3,6],[],[],[]]
输出：[6,3,4,1,5,2,0,7]
示例 2：

输入：n = 8, m = 2, group = [-1,-1,1,0,0,1,0,-1], beforeItems = [[],[6],[5],[6],[3],[],[4],[]]
输出：[]
解释：与示例 1 大致相同，但是在排序后的列表中，4 必须放在 6 的前面。


提示：

1 <= m <= n <= 3 * 104
group.length == beforeItems.length == n
-1 <= group[i] <= m - 1
0 <= beforeItems[i].length <= n - 1
0 <= beforeItems[i][j] <= n - 1
i != beforeItems[i][j]
beforeItems[i] 不含重复元素
*/

impl Solution {
    pub fn sort_items(n: i32, m: i32, group: Vec<i32>, before_items: Vec<Vec<i32>>) -> Vec<i32> {
        pub fn sort_items(n: i32, m: i32, group: Vec<i32>, before_items: Vec<Vec<i32>>) -> Vec<i32> {
            let mut ans = vec![];
            let mut graph = vec![vec![]; (n + m) as usize];
            let mut in_degree = vec![0; (n + m) as usize];


            for i in 0..group.len() {
                if group[i] == -1 {
                    continue;
                }
                graph[(n + group[i]) as usize].push(i as i32);
                in_degree[i] += 1;
            }
            for i in 0..before_items.len() {
                for &item in &before_items[i] {
                    let rep_before_item = if group[item as usize] == -1 {
                        item
                    } else {
                        n + group[item as usize]
                    };
                    let rep_current_item = if group[i] == -1 {
                        i as i32
                    } else {
                        n + group[i]
                    };

                    if rep_before_item == rep_current_item {
                        graph[item as usize].push(i as i32);
                        in_degree[i] += 1;
                    } else {
                        graph[rep_before_item as usize].push(rep_current_item);
                        in_degree[rep_current_item as usize] += 1;
                    }
                }
            }


            fn dfs(graph: &mut Vec<Vec<i32>>, in_degree: &mut Vec<i32>, cur: i32, n: i32, res: &mut Vec<i32>) {
                if cur < n {
                    res.push(cur);
                }
                in_degree[cur as usize] -= 1;

                for &child in &graph[cur as usize].clone() {
                    in_degree[child as usize] -= 1;
                    if in_degree[child as usize] == 0 {
                        dfs(graph, in_degree, child, n, res);
                    }
                }
            }

            for i in 0..n + m {
                if in_degree[i as usize] == 0 {
                    dfs(&mut graph, &mut in_degree, i, n, &mut ans);
                }
            }
            if ans.len() as i32 == n {
                ans
            } else {
                vec![]
            }
        }
    }
}