package main

import (
	"container/heap"
	"container/list"
	"fmt"
	"math"
)

/*


设计一个简化版的推特(Twitter)，可以让用户实现发送推文，关注/取消关注其他用户，能够看见关注人（包括自己）的最近十条推文。你的设计需要支持以下的几个功能：

postTweet(userId, tweetId): 创建一条新的推文
getNewsFeed(userId): 检索最近的十条推文。每个推文都必须是由此用户关注的人或者是用户自己发出的。推文必须按照时间顺序由最近的开始排序。
follow(followerId, followeeId): 关注一个用户
unfollow(followerId, followeeId): 取消关注一个用户
示例:

Twitter twitter = new Twitter();

// 用户1发送了一条新推文 (用户id = 1, 推文id = 5).
twitter.postTweet(1, 5);

// 用户1的获取推文应当返回一个列表，其中包含一个id为5的推文.
twitter.getNewsFeed(1);

// 用户1关注了用户2.
twitter.follow(1, 2);

// 用户2发送了一个新推文 (推文id = 6).
twitter.postTweet(2, 6);

// 用户1的获取推文应当返回一个列表，其中包含两个推文，id分别为 -> [6, 5].
// 推文id6应当在推文id5之前，因为它是在5之后发送的.
twitter.getNewsFeed(1);

// 用户1取消关注了用户2.
twitter.unfollow(1, 2);

// 用户1的获取推文应当返回一个列表，其中包含一个id为5的推文.
// 因为用户1已经不再关注用户2.
twitter.getNewsFeed(1);

https://leetcode-cn.com/problems/design-twitter/
*/
func main() {
	//["Twitter","postTweet","postTweet","postTweet","postTweet","postTweet","postTweet","postTweet","postTweet","postTweet","postTweet","postTweet","getNewsFeed"]
	//[[],[1,5],[1,3],[1,101],[1,13],[1,10],[1,2],[1,94],[1,505],[1,333],[1,22],[1,11],[1]]
	//t1()
	obj := Constructor()
	obj.PostTweet(1, 5)
	obj.PostTweet(1, 3)
	obj.PostTweet(1, 101)
	obj.PostTweet(1, 13)
	obj.PostTweet(1, 10)
	obj.PostTweet(1, 2)
	obj.PostTweet(1, 94)
	obj.PostTweet(1, 505)
	obj.PostTweet(1, 333)
	obj.PostTweet(1, 22)
	obj.PostTweet(1, 11)
	fmt.Println(obj.GetNewsFeed(1))
	//obj.Follow(1, 2)
	//obj.PostTweet(2, 6)
	//fmt.Println(obj.GetNewsFeed(1))
	//obj.Unfollow(1, 2)
	//fmt.Println(obj.GetNewsFeed(1))
}

func t1() {
	obj := Constructor()
	obj.PostTweet(1, 5)
	fmt.Println(obj.GetNewsFeed(1))
	obj.Follow(1, 2)
	obj.PostTweet(2, 6)
	fmt.Println(obj.GetNewsFeed(1))
	obj.Unfollow(1, 2)
	fmt.Println(obj.GetNewsFeed(1))
}

type Twitter struct {
	msgs         map[int]int
	users        map[int]user
	time, maxLen int
}
type msg struct {
	time int
	tid  int
}
type user struct {
	tids    []int
	follows *list.List
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{make(map[int]int), make(map[int]user), 0, 10}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	msgs := this.msgs
	users := this.users
	this.time += 1
	msgs[tweetId] = this.time
	userInfo, found := users[userId]
	if !found {
		userInfo = user{make([]int, 0), list.New()}
		userInfo.follows.PushBack(userId)
		users[userId] = userInfo
	}
	userInfo.tids = append(userInfo.tids, tweetId)
	users[userId] = userInfo
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	msgs := this.msgs
	users := this.users
	userInfo, found := users[userId]
	if !found {
		userInfo = user{make([]int, 0), list.New()}
		userInfo.follows.PushBack(userId)
		users[userId] = userInfo
		return []int{}
	}
	tids := userInfo.tids
	myheap := make(IntHeap, 0, 10)
	heap.Init(&myheap)
	for i := range tids {
		//if myheap.Len()>=this.maxLen {
		//	heap.Pop(&myheap)
		//}
		heap.Push(&myheap, msg{msgs[tids[i]], tids[i]})
	}
	follows := userInfo.follows
	fFront := follows.Front()
	for ; fFront != nil; fFront = fFront.Next() {
		uid := fFront.Value
		if uid == userId {
			continue
		}
		m := uid.(int)
		u := users[m]
		for i := range u.tids {
			//if myheap.Len()>=this.maxLen {
			//	heap.Pop(&myheap)
			//}
			heap.Push(&myheap, msg{msgs[u.tids[i]], u.tids[i]})
		}
	}
	mi := math.Min(float64(myheap.Len()), float64(this.maxLen))
	for myheap.Len() > this.maxLen {
		heap.Pop(&myheap)
	}
	min := int(mi)
	res := make([]int, min)
	for i := min - 1; i >= 0; i-- {
		pop := heap.Pop(&myheap)
		m := pop.(msg)
		res[i] = m.tid
	}
	return res
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	users := this.users
	userInfo, found := users[followerId]
	if !found {
		userInfo = user{make([]int, 0), list.New()}
		userInfo.follows.PushBack(followerId)
		users[followerId] = userInfo
	}
	front := userInfo.follows.Front()
	have := false
	for front != nil {
		if front.Value == followeeId {
			have = true
			break
		}
		front = front.Next()
	}
	if !have {
		userInfo.follows.PushBack(followeeId)
	}
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	users := this.users
	userInfo, found := users[followerId]
	if found {
		front := userInfo.follows.Front()
		have := false
		for front != nil {
			if front.Value == followeeId {
				have = true
				break
			}
			front = front.Next()
		}
		if have {
			userInfo.follows.Remove(front)
		}
	}
}

type IntHeap []msg // 定义一个类型

func (h IntHeap) Len() int { return len(h) } // 绑定len方法,返回长度
func (h IntHeap) Less(i, j int) bool { // 绑定less方法
	return h[i].time < h[j].time // 如果h[i]<h[j]生成的就是小根堆，如果h[i]>h[j]生成的就是大根堆
}
func (h IntHeap) Swap(i, j int) { // 绑定swap方法，交换两个元素位置
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Pop() interface{} { // 绑定pop方法，从最后拿出一个元素并返回
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap) Push(x interface{}) { // 绑定push方法，插入新元素
	m := x.(msg)
	*h = append(*h, m)
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
