/*

有 n 个城市，其中一些彼此相连，另一些没有相连。如果城市 a 与城市 b 直接相连，且城市 b 与城市 c 直接相连，那么城市 a 与城市 c 间接相连。

省份 是一组直接或间接相连的城市，组内不含其他没有相连的城市。

给你一个 n x n 的矩阵 isConnected ，其中 isConnected[i][j] = 1 表示第 i 个城市和第 j 个城市直接相连，而 isConnected[i][j] = 0 表示二者不直接相连。

返回矩阵中 省份 的数量。

 

示例 1：
https://assets.leetcode.com/uploads/2020/12/24/graph1.jpg

输入：isConnected = [[1,1,0],[1,1,0],[0,0,1]]
输出：2
示例 2：
https://assets.leetcode.com/uploads/2020/12/24/graph2.jpg

输入：isConnected = [[1,0,0],[0,1,0],[0,0,1]]
输出：3
 

提示：

1 <= n <= 200
n == isConnected.length
n == isConnected[i].length
isConnected[i][j] 为 1 或 0
isConnected[i][i] == 1
isConnected[i][j] == isConnected[j][i]


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-provinces
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

impl Solution {
    pub fn find_circle_num(is_connected: Vec<Vec<i32>>) -> i32 {
        let n = is_connected.len();
        let mut visited = vec![false; n];
        let mut count=0;
        fn dfs(is_connected: &Vec<Vec<i32>>,visited: &mut Vec<bool>,i: usize){
            for j in 0..is_connected.len() {
                if is_connected[i][j]==1&&!visited[j] {
                    visited[j]=true;
                    dfs(is_connected,visited,j);
                }
            }
        }

        for i in 0..n {
            if !visited[i] {
                dfs(&is_connected, &mut visited, i);
                count+=1;
            }
        }

        count
    }
    pub fn find_circle_num2(is_connected: Vec<Vec<i32>>) -> i32 {
        // 并查集
        let p=&mut (0..is_connected.len()).collect();
        fn f(x:usize, p:&mut Vec<usize>) -> usize {
            if p[x]!=x {
                p[x]=f(p[x],p);
            }
            p[x]
        }
        let mut ans=is_connected.len();
        for i in 0..is_connected.len() {
            for j in 0..is_connected.len() {
                if is_connected[i][j]==1 {
                    let (pi,pj)=(f(i,p),f(j,p));
                    if pi!=pj {
                        p[pj]=pi;
                        ans-=1;
                    }
                }
            }
        }
        ans as i32
    }

}