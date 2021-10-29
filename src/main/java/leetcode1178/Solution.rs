/*
外国友人仿照中国字谜设计了一个英文版猜字谜小游戏，请你来猜猜看吧。

字谜的迷面puzzle 按字符串形式给出，如果一个单词word符合下面两个条件，那么它就可以算作谜底：

单词word中包含谜面puzzle的第一个字母。
单词word中的每一个字母都可以在谜面puzzle中找到。
例如，如果字谜的谜面是 "abcdefg"，那么可以作为谜底的单词有 "faced", "cabbage", 和 "baggage"；而 "beefed"（不含字母 "a"）以及"based"（其中的 "s" 没有出现在谜面中）。
返回一个答案数组answer，数组中的每个元素answer[i]是在给出的单词列表 words 中可以作为字谜迷面puzzles[i]所对应的谜底的单词数目。



示例：

输入：
words = ["aaaa","asas","able","ability","actt","actor","access"], 
puzzles = ["aboveyz","abrodyz","abslute","absoryz","actresz","gaswxyz"]
输出：[1,1,3,2,4,0]
解释：
1 个单词可以作为 "aboveyz" 的谜底 : "aaaa" 
1 个单词可以作为 "abrodyz" 的谜底 : "aaaa"
3 个单词可以作为 "abslute" 的谜底 : "aaaa", "asas", "able"
2 个单词可以作为"absoryz" 的谜底 : "aaaa", "asas"
4 个单词可以作为"actresz" 的谜底 : "aaaa", "asas", "actt", "access"
没有单词可以作为"gaswxyz" 的谜底，因为列表中的单词都不含字母 'g'。


提示：

1 <= words.length <= 10^5
4 <= words[i].length <= 50
1 <= puzzles.length <= 10^4
puzzles[i].length == 7
words[i][j], puzzles[i][j]都是小写英文字母。
每个puzzles[i]所包含的字符都不重复。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-valid-words-for-each-puzzle
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
impl Solution {
    pub fn find_num_of_valid_words(words: Vec<String>, puzzles: Vec<String>) -> Vec<i32> {
        let mut result = Vec::new();

        let mut words_set = vec![0; (2 as usize).pow(27)-1];
        for word in words {
            let mut posts = 0;
            for character in word.chars() {
                posts |= (1 << (character as i32 - 'a' as i32));
            }
            words_set[posts] += 1;
        }

        for puzzle in puzzles {
            let mut count = 0;

            let mut posts = 0;
            let mut characters = puzzle.chars();
            let head = characters.nth(0).unwrap();
            for character in characters {
                posts |= (1 << (character as i32 - 'a' as i32))
            }
            let mut sub = posts;
            loop {
                sub = (sub - 1) & posts;
                count += words_set[sub + (1 << (head as i32 - 'a' as i32))];
                if sub == posts { break; }
            }
            result.push(count);
        }

        result
    }
}