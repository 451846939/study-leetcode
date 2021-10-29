package leetcode930

class Solution {
    fun numSubarraysWithSum(nums: IntArray, goal: Int): Int {
        var sum=0
        val map = mutableMapOf<Int, Int>()
        var ret=0
        nums.forEach {
            map[sum]=(map[sum]?:0)+1
            sum+=it
            ret+=map[sum-goal]?:0
        }
        return ret
    }
}