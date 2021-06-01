/*
给你一个下标从 0 开始的正整数数组candiesCount，其中candiesCount[i]表示你拥有的第i类糖果的数目。同时给你一个二维数组queries，其中queries[i] = [favoriteTypei, favoriteDayi, dailyCapi]。

你按照如下规则进行一场游戏：

你从第0天开始吃糖果。
你在吃完 所有第 i - 1类糖果之前，不能吃任何一颗第 i类糖果。
在吃完所有糖果之前，你必须每天 至少吃 一颗糖果。
请你构建一个布尔型数组answer，满足answer.length == queries.length 。answer[i]为true的条件是：在每天吃 不超过 dailyCapi颗糖果的前提下，你可以在第favoriteDayi天吃到第favoriteTypei类糖果；否则 answer[i]为 false。注意，只要满足上面 3 条规则中的第二条规则，你就可以在同一天吃不同类型的糖果。

请你返回得到的数组answer。



示例 1：

输入：candiesCount = [7,4,5,3,8], queries = [[0,2,2],[4,2,4],[2,13,1000000000]]
输出：[true,false,true]
提示：
1- 在第 0 天吃 2 颗糖果(类型 0），第 1 天吃 2 颗糖果（类型 0），第 2 天你可以吃到类型 0 的糖果。
2- 每天你最多吃 4 颗糖果。即使第 0 天吃 4 颗糖果（类型 0），第 1 天吃 4 颗糖果（类型 0 和类型 1），你也没办法在第 2 天吃到类型 4 的糖果。换言之，你没法在每天吃 4 颗糖果的限制下在第 2 天吃到第 4 类糖果。
3- 如果你每天吃 1 颗糖果，你可以在第 13 天吃到类型 2 的糖果。
示例 2：

输入：candiesCount = [5,2,6,4,1], queries = [[3,1,2],[4,10,3],[3,10,100],[4,100,30],[1,3,1]]
输出：[false,true,true,false,false]


提示：

1 <= candiesCount.length <= 105
1 <= candiesCount[i] <= 105
1 <= queries.length <= 105
queries[i].length == 3
0 <= favoriteTypei < candiesCount.length
0 <= favoriteDayi <= 109
1 <= dailyCapi <= 109

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/can-you-eat-your-favorite-candy-on-your-favorite-day
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
impl Solution {
    pub fn can_eat(candies_count: Vec<i32>, queries: Vec<Vec<i32>>) -> Vec<bool> {
        let pre_sum = candies_count.iter().enumerate().fold(Vec::new(), |mut ps, (idx, &c)| {
            if idx == 0 {
                ps.push(0);
                ps.push(c as i64);
            } else {
                ps.push(c as i64 + ps[ps.len() - 1]);
            }
            return ps;
        });
        queries.iter().map(|q| {
            let favorite_type = q[0] as i64 + 1;
            let favorite_day = q[1] as i64 + 1;
            let daily_cap = q[2] as i64;
            if (favorite_day - 1) >= pre_sum[favorite_type as usize] {
                return false;
            } else if daily_cap * favorite_day <= pre_sum[favorite_type as usize - 1] {
                return false;
            }
            return true;
        }).collect()
    }
}