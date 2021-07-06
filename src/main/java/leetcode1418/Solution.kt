package leetcode1418

class Solution {
    fun displayTable(orders: List<List<String>>): List<List<String>> {
        val menu = HashSet<String>(orders.map { it[2] }).sorted()
        val tables = HashSet<String>(orders.map { it[1] })
        val map = HashMap<Pair<String, String>, Int>()
        orders.forEach { map[it[1] to it[2]] = (map[it[1] to it[2]] ?: 0) + 1 }
        val head = mutableListOf("Table")
        head.addAll(menu)
        val ans = mutableListOf(head)

        for (t in tables.sortedBy { it.toInt() }) {
            val ms = mutableListOf(t)
            ms.addAll(menu.map { "${map[t to it] ?: 0}" })
            ans.add(ms)
        }

        return ans
    }
}