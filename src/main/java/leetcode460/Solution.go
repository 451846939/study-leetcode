package main

import (
	"fmt"
)

/*设计并实现最不经常使用（LFU）缓存的数据结构。它应该支持以下操作：get 和 put。

get(key) - 如果键存在于缓存中，则获取键的值（总是正数），否则返回 -1。
put(key, value) - 如果键不存在，请设置或插入值。当缓存达到其容量时，它应该在插入新项目之前，使最不经常使用的项目无效。在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，最近最少使用的键将被去除。

进阶：
你是否可以在 O(1) 时间复杂度内执行两项操作？

示例：

LFUCache countTable = new LFUCache( 2  capacity (缓存容量)  );

countTable.put(1, 1);
countTable.put(2, 2);
countTable.get(1);       // 返回 1
countTable.put(3, 3);    // 去除 key 2
countTable.get(2);       // 返回 -1 (未找到key 2)
countTable.get(3);       // 返回 3
countTable.put(4, 4);    // 去除 key 1
countTable.get(1);       // 返回 -1 (未找到 key 1)
countTable.get(3);       // 返回 3
countTable.get(4);       // 返回 4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/lfu-cache
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
func main() {
	//testone()
	//t2()
	//t3()
	//t4()
	obj := Constructor(3)
	//param_1 := obj.Get(key)
	obj.Put(3, 1)
	obj.Put(2, 1)
	obj.Put(2, 1)
	obj.Put(2, 2)
	obj.Put(3, 1)
	obj.Put(4, 4)
	fmt.Println(obj.Get(2))
	fmt.Println(obj.Get(2))
	fmt.Println(obj.Get(4))
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(2))
	fmt.Println(obj.Get(2))

	//["LFUCache","put","put","put","put","put","get","put","get","get","put","get","put","put","put","get","put","get","get","get","get","put","put","get","get","get","put","put","get","put","get","put","get","get","get","put","put","put","get","put","get","get","put","put","get","put","put","put","put","get","put","put","get","put","put","get","put","put","put","put","put","get","put","put","get","put","get","get","get","put","get","get","put","put","put","put","get","put","put","put","put","get","get","get","put","put","put","get","put","put","put","get","put","put","put","get","get","get","put","put","put","put","get","put","put","put","put","put","put","put"]
	//[[10],[10,13],[3,17],[6,11],[10,5],[9,10],[13],[2,19],[2],[3],[5,25],[8],[9,22],[5,5],[1,30],[11],[9,12],[7],[5],[8],[9],[4,30],[9,3],[9],[10],[10],[6,14],[3,1],[3],[10,11],[8],[2,14],[1],[5],[4],[11,4],[12,24],[5,18],[13],[7,23],[8],[12],[3,27],[2,12],[5],[2,9],[13,4],[8,18],[1,7],[6],[9,29],[8,21],[5],[6,30],[1,12],[10],[4,15],[7,22],[11,26],[8,17],[9,29],[5],[3,4],[11,30],[12],[4,29],[3],[9],[6],[3,4],[1],[10],[3,29],[10,28],[1,20],[11,13],[3],[3,12],[3,8],[10,9],[3,26],[8],[7],[5],[13,17],[2,27],[11,15],[12],[9,19],[2,15],[3,16],[1],[12,17],[9,1],[6,19],[4],[5],[5],[8,1],[11,7],[5,2],[9,28],[1],[2,2],[7,4],[4,22],[7,24],[9,26],[13,28],[11,26]]
}

func t4() {
	//["LFUCache","put","put","put","put","get"]
	//[[2],[3,1],[2,1],[2,2],[4,4],[2]]

	obj := Constructor(2)
	//param_1 := obj.Get(key)
	obj.Put(3, 1)
	obj.Put(2, 1)
	obj.Put(2, 2)
	obj.Put(4, 4)
	fmt.Println(obj.Get(2))
}

func t3() {
	//["LFUCache","put","put","put","put","get","get"]
	//[[2],[2,1],[1,1],[2,3],[4,1],[1],[2]]
	obj := Constructor(2)
	//param_1 := obj.Get(key)
	obj.Put(2, 1)
	obj.Put(1, 1)
	obj.Put(2, 3)
	obj.Put(4, 1)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(2))
}

func t2() {
	//	["LFUCache","put","put","get","get","get","put","put","get","get","get","get"]
	//[[3],        [2,2],[1,1],[2],[1],[2],[3,3],[4,4],[3],[2],[1],[4]]
	obj := Constructor(3)
	//param_1 := obj.Get(key)
	obj.Put(2, 2)
	obj.Put(1, 1)
	fmt.Println(obj.Get(2))
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(2))
	obj.Put(3, 3)
	obj.Put(4, 4)
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(2))
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(4))
}

func testone() {
	obj := Constructor(2)
	//param_1 := obj.Get(key)
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1))
	obj.Put(3, 3)
	fmt.Println(obj.Get(2))
	fmt.Println(obj.Get(3))
	obj.Put(4, 4)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(4))
}

type node struct {
	k, v, count int
}
type keyNode struct {
	count, index int
}
type LFUCache struct {
	maxLength, nowlen int
	countTable        map[int][]node
	keyTable          map[int]keyNode
}

func Constructor(capacity int) LFUCache {
	return LFUCache{capacity, 0, make(map[int][]node, capacity), make(map[int]keyNode)}
}

func (this *LFUCache) Get(key int) int {
	countTable := this.countTable
	keyTable := this.keyTable
	keyN, have := keyTable[key]
	if have {
		res := countTable[keyN.count][keyN.index]
		//更新index和count以及countmap
		if len(countTable[keyN.count]) == keyN.index+1 {
			countTable[keyN.count] = countTable[keyN.count][:keyN.index]
		} else {
			//中间元素
			countTable[keyN.count] = append(countTable[keyN.count][:keyN.index], countTable[keyN.count][keyN.index+1:]...)
			for i := keyN.index; i < len(countTable[keyN.count]); i++ {
				temp := keyTable[countTable[keyN.count][i].k]
				temp.index--
				keyTable[countTable[keyN.count][i].k] = temp
			}
		}
		keyN.count++
		res.count++
		countTable[keyN.count] = append(countTable[keyN.count], res)
		keyN.index = len(countTable[keyN.count]) - 1
		keyTable[key] = keyN
		return res.v
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
	//为0直接返回
	if this.maxLength <= 0 {
		return
	}
	newNode := node{key, value, 0}
	countTable := this.countTable
	keyTable := this.keyTable
	//先查找keyTable中是否存在这个键，如果存在就更新
	keyN, have := keyTable[key]
	if have {
		if len(countTable[keyN.count]) == keyN.index+1 {
			countTable[keyN.count] = countTable[keyN.count][:keyN.index]
		} else {
			//中间元素
			countTable[keyN.count] = append(countTable[keyN.count][:keyN.index], countTable[keyN.count][keyN.index+1:]...)
			for i := keyN.index; i < len(countTable[keyN.count]); i++ {
				temp := keyTable[countTable[keyN.count][i].k]
				temp.index--
				keyTable[countTable[keyN.count][i].k] = temp
			}
		}
		keyN.count++
		newNode.count = keyN.count
		countTable[keyN.count] = append(countTable[keyN.count], newNode)
		keyN.index = len(countTable[keyN.count]) - 1
		keyTable[key] = keyN
		return
	}
	if this.nowlen >= this.maxLength {
		//	找到最小的使用并且删除
		minCount := 0
		for true {
			if len(countTable[minCount]) != 0 {
				break
			}
			minCount++
		}
		delete(keyTable, countTable[minCount][0].k)
		countTable[minCount] = countTable[minCount][1:]
		for _, e := range countTable[minCount] {
			temp := keyTable[e.k]
			temp.index--
			keyTable[e.k] = temp
		}
		this.nowlen--
	}
	nodes := countTable[0]
	nodes = append(nodes, newNode)
	countTable[0] = nodes
	keyTable[key] = keyNode{0, len(nodes) - 1}
	this.nowlen++
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
