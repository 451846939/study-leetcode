package leetcode1711

class Solution {
    fun countPairs(deliciousness: IntArray): Int {
        val mod = 1000000007
        var maxVal = deliciousness.maxOrNull()
        val maxSum = maxVal?.times(2)
        val map = mutableMapOf<Int, Int>()
        var ans=0
        deliciousness.forEach {
            var sum = 1
            while (sum <= maxSum!!) {
                val count = map[sum - it] ?: 0
                ans=(ans+count)%mod
                sum=sum.shl(1)
            }
            map[it]=(map[it]?:0)+1
        }
        return ans
    }
}