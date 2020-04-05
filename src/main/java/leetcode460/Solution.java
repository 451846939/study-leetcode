package leetcode460;

import java.util.*;

/**
 * @author lin
 * Email 451846939@qq.com
 * @version V1.0
 * @description
 * @date 2020/4/5 5:24 下午
 */
public class Solution {



    class LFUCache {
        class Node{
            int key;
            int value;
            int freq = 1;

            public Node(int key, int value) {
                this.key = key;
                this.value = value;
            }
        }
        int size;
        int capacity;
        int min; // 存储当前最小频次
        private Map<Integer, List<Node>> freqMap=new HashMap<>();
        private Map<Integer,Node> cache=new HashMap<>();
        public LFUCache(int capacity) {
            this.capacity = capacity;
        }
        public int get(int key) {
            Node node = cache.get(key);
            if (node == null) {
                return -1;
            }
            freqInc(node);
            return node.value;
        }

        public void put(int key, int value) {
            if (capacity == 0) {
                return;
            }
            Node node = cache.get(key);
            if (node != null) {
                node.value = value;
                freqInc(node);
            } else {
                if (size == capacity) {
                    Node deadNode = removeNode();
                    cache.remove(deadNode.key);
                    size--;
                }
                Node newNode = new Node(key, value);
                cache.put(key, newNode);
                addNode(newNode);
                size++;
            }
        }

        void freqInc(Node node) {
            // 从原freq对应的链表里移除, 并更新min
            int freq = node.freq;
            List<Node> set = freqMap.get(freq);
            set.remove(node);
            if (freq == min && set.size() == 0) {
                min = freq + 1;
            }
            // 加入新freq对应的链表
            node.freq++;
            List<Node> newSet = freqMap.get(freq + 1);
            if (newSet == null) {
                newSet = new LinkedList<>();
                freqMap.put(freq + 1, newSet);
            }
            newSet.add(node);
        }

        void addNode(Node node) {
            List<Node> set = freqMap.get(1);
            if (set == null) {
                set = new LinkedList<>();
                freqMap.put(1, set);
            }
            set.add(node);
            min = 1;
        }

        Node removeNode() {
            List<Node> set = freqMap.get(min);
            Node deadNode = set.iterator().next();
            set.remove(deadNode);
            return deadNode;
        }

    }

/**
 * Your LFUCache object will be instantiated and called as such:
 * LFUCache obj = new LFUCache(capacity);
 * int param_1 = obj.get(key);
 * obj.put(key,value);
 */
}
