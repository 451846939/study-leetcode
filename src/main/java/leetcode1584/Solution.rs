/*
给你一个points数组，表示 2D 平面上的一些点，其中points[i] = [xi, yi]。

连接点[xi, yi] 和点[xj, yj]的费用为它们之间的 曼哈顿距离：|xi - xj| + |yi - yj|，其中|val|表示val的绝对值。

请你返回将所有点连接的最小总费用。只有任意两点之间 有且仅有一条简单路径时，才认为所有点都已连接。



示例 1：

https://assets.leetcode.com/uploads/2020/08/26/d.png

输入：points = [[0,0],[2,2],[3,10],[5,2],[7,0]]
输出：20
解释：
https://assets.leetcode.com/uploads/2020/08/26/c.png
我们可以按照上图所示连接所有点得到最小总费用，总费用为 20 。
注意到任意两个点之间只有唯一一条路径互相到达。
示例 2：

输入：points = [[3,12],[-2,5],[-4,1]]
输出：18
示例 3：

输入：points = [[0,0],[1,1],[1,0],[-1,1]]
输出：4
示例 4：

输入：points = [[-1000000,-1000000],[1000000,1000000]]
输出：4000000
示例 5：

输入：points = [[0,0]]
输出：0

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/min-cost-to-connect-all-points
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
impl Solution {
    pub fn min_cost_connect_points(points: Vec<Vec<i32>>) -> i32 {
        prim(&points,0)
    }
}


pub fn prim(points:&Vec<Vec<i32>>,start:usize)->i32{
    /*
链接：https://leetcode-cn.com/problems/min-cost-to-connect-all-points/solution/prim-and-kruskal-by-yexiso-c500/
Part1. 解题思路
抽象（假想:假设存在，但不存在）出两个集合，集合V 和集合Vnew

最开始，所以的图节点都在集合V中

如果一个节点加入到了最小生成树中，则将该节点加入到Vnew(即Vnew保存的是最小生成树中的节点)

说明： Vnew即最小生成树

Part2. 数据结构
Prim算法主要维护2个数组

lowcost 数组，表示V中的节点，保存V中每个节点离Vnew中所有节点的最短距离。如果节点已经加入到了Vnew中，则置为-1

v 数组，表示V中节点的访问情况，最开始全部为0,表示未加入到Vnew中，若某节点加入到了Vnew中， 则将其置为-1

Part3. 步骤：
随机选择一个起点，并将其加入到Vnew中。同时，更新此时的lowcost和v
遍历lowcost，寻找lowcost中的最小值min（假设索引为 j ），将与索引 j 相对应的节点加入到Vnew中，更新lowcost[j]和v[j]。
找到最小值 j 后，此时lowcost中的所有节点都要更新，因为此时Vnew节点增加了，V集合中的节点离Vnew集合的距离可能会缩短。
此时已更新完所有的lowcost。
重复步骤2,直到访问了所有的节点。
很明显，最后需要计算的最小生成树各节点之间的距离和便是每一步lowcost中的最小值min的和

Part4. 举例：
https://pic.leetcode-cn.com/1611023745-EuGMdh-image.png
*/
    fn abs(i:i32)->i32{
        if i<0 { return -i; }
        i
    }
    let n = points.len();
    let mut res=0;
    // 1. 将points转化成邻接矩阵
    let mut g = vec![vec![0; n]; n];
    for i in 0..n {
        for j in 0..n {
            let dist = abs(points[i][0] - points[j][0]) + abs(points[i][1] - points[j][1]);
            g[i][j]=dist;
            g[j][i]=dist
        }
    }
    // 记录V[i]到Vnew的最近距离
    let mut low_cost = vec![0x7fffffff; n];
    // 记录V[i]是否加入到了Vnew
    let mut v = vec![-1;n];
    // 2. 先将start加入到Vnew
    v[start]=0;
    for i in 0..n {
        if i==start {
            continue
        }
        low_cost[i]=g[i][start]
    }
    // 3. 剩余n - 1个节点未加入到Vnew，遍历
    for i in 1..n {
        // 找出此时V中，离Vnew最近的点
        let mut min_index=0;
        let mut min_val=0x7fffffff;
        for j in 0..n {
            if v[j]==0 {
                continue
            }
            if low_cost[j]<min_val {
                min_index=j;
                min_val=low_cost[j];
            }
        }
        // 将该点加入Vnew，更新lowcost和v
        res+=min_val;
        v[min_index]=0;
        low_cost[min_index]=-1;

        // 更新集合V中所有点的lowcost
        for j in 0..n {
            if v[j]==-1 &&g[j][min_index]<low_cost[j]{
                low_cost[j]=g[j][min_index];
            }
        }
    }
    res
}