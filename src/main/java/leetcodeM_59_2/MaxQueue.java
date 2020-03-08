package leetcodeM_59_2;

import java.util.Deque;
import java.util.LinkedList;
import java.util.Queue;

class MaxQueue {
    private int size;
    private Queue<Integer> queue;
    private Deque<Integer> maxQueue;

    public MaxQueue() {
        size = 0;
        queue = new LinkedList<>();
        maxQueue = new LinkedList<>();
    }

    public int max_value() {
        if (size == 0) {
            return -1;
        }
        return maxQueue.getFirst();
    }

    public void push_back(int value) {
        while (!maxQueue.isEmpty() && maxQueue.getLast() < value) {
            maxQueue.removeLast();
        }
        maxQueue.add(value);
        queue.add(value);
        size++;
    }

    public int pop_front() {
        if (size == 0) {
            maxQueue.clear();
            return -1;
        }
        Integer remove = queue.remove();
        if (maxQueue.getFirst().equals(remove)) {
            maxQueue.removeFirst();
        }
        size--;
        return remove;
    }

    public static void main(String[] args) {
        MaxQueue obj = new MaxQueue();
        int param_1 = obj.max_value();
        obj.push_back(1);
        obj.push_back(3);
        obj.push_back(3);
        obj.push_back(2);
        int param_3 = obj.pop_front();
        obj.pop_front();
        obj.pop_front();
        obj.pop_front();
    }
}