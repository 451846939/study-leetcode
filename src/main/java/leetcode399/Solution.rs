/*

给你一个变量对数组 equations 和一个实数值数组 values 作为已知条件，其中 equations[i] = [Ai, Bi] 和 values[i] 共同表示等式 Ai / Bi = values[i] 。每个 Ai 或 Bi 是一个表示单个变量的字符串。

另有一些以数组 queries 表示的问题，其中 queries[j] = [Cj, Dj] 表示第 j 个问题，请你根据已知条件找出 Cj / Dj = ? 的结果作为答案。

返回 所有问题的答案 。如果存在某个无法确定的答案，则用 -1.0 替代这个答案。



注意：输入总是有效的。你可以假设除法运算中不会出现除数为 0 的情况，且不存在任何矛盾的结果。



示例 1：

输入：equations = [["a","b"],["b","c"]], values = [2.0,3.0], queries = [["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]
输出：[6.00000,0.50000,-1.00000,1.00000,-1.00000]
解释：
条件：a / b = 2.0, b / c = 3.0
问题：a / c = ?, b / a = ?, a / e = ?, a / a = ?, x / x = ?
结果：[6.0, 0.5, -1.0, 1.0, -1.0 ]
示例 2：

输入：equations = [["a","b"],["b","c"],["bc","cd"]], values = [1.5,2.5,5.0], queries = [["a","c"],["c","b"],["bc","cd"],["cd","bc"]]
输出：[3.75000,0.40000,5.00000,0.20000]
示例 3：

输入：equations = [["a","b"]], values = [0.5], queries = [["a","b"],["b","a"],["a","c"],["x","y"]]
输出：[0.50000,2.00000,-1.00000,-1.00000]


提示：

1 <= equations.length <= 20
equations[i].length == 2
1 <= Ai.length, Bi.length <= 5
values.length == equations.length
0.0 < values[i] <= 20.0
1 <= queries.length <= 20
queries[i].length == 2
1 <= Cj.length, Dj.length <= 5
Ai, Bi, Cj, Dj 由小写英文字母与数字组成
*/

impl Solution {
    pub fn calc_equation(equations: Vec<Vec<String>>, values: Vec<f64>, queries: Vec<Vec<String>>) -> Vec<f64> {
        /*
         * graph1: 将图转化为邻接矩阵
         * graph2: 每个变量的分组情况，不同组之间没有通路
         * graph3: 每个变量的虚拟值（方便计算）
         * unchecked: 尚未分组的变量数
         * secnum: 组别数量
         */
        let mut graph1: HashMap<String, HashMap<String, f64>> = HashMap::new();
        let n = equations.len();
        for i in 0..n {
            graph1.entry(equations[i][0].clone()).or_insert(HashMap::new()).entry(equations[i][1].clone()).or_insert(values[i]);
            graph1.entry(equations[i][1].clone()).or_insert(HashMap::new()).entry(equations[i][0].clone()).or_insert(1f64 / values[i]);
        }
        let mut graph2 = HashMap::new();
        for key in graph1.keys() {
            graph2.insert(key, 0);
        }
        let mut unchecked = graph2.len();
        let mut secnum = 0;
        let mut graph3 = HashMap::new();
        for key in graph1.keys() {
            if graph2.get(key) == Some(&0) {
                graph3.entry(key).or_insert(1f64);
                secnum += 1;
                graph2.insert(key, secnum);
                let mut keys = vec![key];
                while keys.len() > 0 {
                    let mut keys_n = vec![];
                    for item in keys.clone() {
                        for key_n in graph1.get(item).unwrap().keys() {
                            if graph2.get(&key_n) == Some(&0) {
                                graph2.insert(key_n, secnum);
                                unchecked -= 1;
                                let r = graph3.get(item).unwrap() * graph1.get(key_n).unwrap().get(item).unwrap();
                                graph3.entry(key_n).or_insert(r);
                                keys_n.push(key_n);
                            }
                        }
                    }
                    keys = keys_n;
                }
            }
        }
        let mut ret = vec![];
        for query in queries.iter() {
            if graph2.contains_key(&query[0]) && graph2.contains_key(&query[1]) && graph2.get(&query[0]) == graph2.get(&query[1]) {
                ret.push(graph3.get(&query[0]).unwrap() / graph3.get(&query[1]).unwrap());
            } else {
                ret.push(-1f64);
            }
        }
        ret
    }
}