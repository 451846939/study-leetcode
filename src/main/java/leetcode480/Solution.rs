/*
中位数是有序序列最中间的那个数。如果序列的大小是偶数，则没有最中间的数；此时中位数是最中间的两个数的平均数。

例如：

[2,3,4]，中位数是3
[2,3]，中位数是 (2 + 3) / 2 = 2.5
给你一个数组 nums，有一个大小为 k 的窗口从最左端滑动到最右端。窗口中有 k 个数，每次窗口向右移动 1 位。你的任务是找出每次窗口移动后得到的新窗口中元素的中位数，并输出由它们组成的数组。



示例：

给出nums = [1,3,-1,-3,5,3,6,7]，以及k = 3。

窗口位置                      中位数
---------------               -----
[1  3  -1] -3  5  3  6  7       1
 1 [3  -1  -3] 5  3  6  7      -1
 1  3 [-1  -3  5] 3  6  7      -1
 1  3  -1 [-3  5  3] 6  7       3
 1  3  -1  -3 [5  3  6] 7       5
 1  3  -1  -3  5 [3  6  7]      6
因此，返回该滑动窗口的中位数数组[1,-1,-1,3,5,6]。



提示：

你可以假设k始终有效，即：k 始终小于输入的非空数组的元素个数。
与真实值误差在 10 ^ -5 以内的答案将被视作正确答案。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sliding-window-median
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
use std::collections::{BinaryHeap, HashMap};
use std::cmp::Reverse;

impl Solution {
    pub fn median_sliding_window(nums: Vec<i32>, k: i32) -> Vec<f64> {
        let mut ret: Vec<f64> = vec![];
        let mut left_heap_max = BinaryHeap::new();
        let mut right_heap_min: BinaryHeap<Reverse<i32>> = BinaryHeap::new();
        let flag = if k % 2 == 0 { true } else { false };  // if flag is false, left has 1 more
        let mut balance = k % 2;   // maintain 2 binary heap is to keep balance 0
        let mut map: HashMap<i32, i32> = HashMap::new();

        // init heaps
        for i in 0..k as usize {
            left_heap_max.push(nums[i]);
            balance -= 1;
        }
        Solution::balance_heaps(&mut left_heap_max, &mut right_heap_min, &mut balance, &mut map);
        ret.push(Solution::get_median(&left_heap_max, &right_heap_min, flag));

        // iter
        for i in 1..nums.len()-k as usize +1 {
            Solution::proc_outdate(nums[i-1], &mut left_heap_max, &mut right_heap_min, &mut map, &mut balance);
            Solution::push_new(nums[i+k as usize -1], &mut left_heap_max, &mut right_heap_min, &mut balance);
            Solution::balance_heaps(&mut left_heap_max, &mut right_heap_min, &mut balance, &mut map);
            ret.push(Solution::get_median(&left_heap_max, &right_heap_min, flag));
        }
        ret
    }

    fn balance_heaps(left: &mut BinaryHeap<i32>, right: &mut BinaryHeap<Reverse<i32>>, balance: &mut i32,
                     map: &mut HashMap<i32, i32>) {
        if *balance == 0 { return; }
        if *balance < 0 {
            while *balance != 0 {
                right.push(Reverse(left.pop().unwrap()));
                Solution::check_hashmap_left(left, map);
                *balance += 2;
            }
        } else {
            while *balance != 0  {
                left.push(right.pop().unwrap().0);
                Solution::check_hashmap_right(right, map);
                *balance -= 2;
            }
        }
    }

    fn get_median(left: &BinaryHeap<i32>, right: &BinaryHeap<Reverse<i32>>, flag: bool) -> f64 {
        if flag {
            (*left.peek().unwrap_or(&0) as f64 + right.peek().unwrap_or(&Reverse(0)).0 as f64) / 2.0
        } else {
            *left.peek().unwrap() as f64
        }
    }

    fn proc_outdate(num: i32, left: &mut BinaryHeap<i32>, right: &mut BinaryHeap<Reverse<i32>>,
                    map: &mut HashMap<i32, i32>, balance: &mut i32) {
        if let Some(x) = left.peek() {
            if num == *x {
                left.pop();
                *balance += 1;
                Solution::check_hashmap_left(left, map);
                return;
            } else if num < *x {
                let entry = map.entry(num).or_insert(0);
                *entry += 1;
                *balance += 1;
                return;
            }
        }

        if let Some(y) = right.peek() {
            if num == y.0 {
                right.pop();
                *balance -= 1;
                Solution::check_hashmap_right(right, map);
            } else if num > y.0 {
                let entry = map.entry(num).or_insert(0);
                *entry += 1;
                *balance -= 1;
            }
        }
    }

    fn check_hashmap_left(heap: &mut BinaryHeap<i32>, map: &mut HashMap<i32, i32>) {
        while let Some(x) = heap.peek() {
            let key = *x;
            if Solution::decrease_record(key, map) {
                heap.pop();
            } else {
                break;
            }
        }
    }

    fn check_hashmap_right(heap: &mut BinaryHeap<Reverse<i32>>, map: &mut HashMap<i32, i32>) {
        while let Some(x) = heap.peek() {
            let key = x.0;
            if Solution::decrease_record(key, map) {
                heap.pop();
            } else {
                break;
            }
        }
    }

    fn decrease_record(key: i32, map: &mut HashMap<i32, i32>) -> bool {
        if map.contains_key(&key) {
            if map[&key] > 1 {
                *map.entry(key).or_insert(0) -= 1;
            } else if map[&key] == 1 {
                map.remove(&key);
            }
            return true;
        }
        return false;
    }

    fn push_new(num: i32, left: &mut BinaryHeap<i32>, right: &mut BinaryHeap<Reverse<i32>>, balance: &mut i32) {
        if let Some(x) = left.peek() {
            if num <= *x {
                left.push(num);
                *balance -= 1;
                return;
            }
        }
        right.push(Reverse(num));
        *balance += 1;
    }
}