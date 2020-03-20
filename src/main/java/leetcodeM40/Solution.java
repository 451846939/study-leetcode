package leetcodeM40;

import java.util.Arrays;
import java.util.Comparator;
import java.util.PriorityQueue;

/**
 * @author chenxu.lin
 * @description
 * @date 2020/3/20 12:47
 */
public class Solution {
    public static void main(String[] args) {
        System.out.println(Arrays.toString(getLeastNumbers2(new int[]{3, 2, 1}, 2)));
    }
    public static int[] getLeastNumbers(int[] arr, int k) {
        PriorityQueue<Integer> maxheap = new PriorityQueue<>(Comparator.reverseOrder());
        for (int i = 0; i < arr.length; i++) {
            if (maxheap.size()<k){
                maxheap.offer(arr[i]);
            }else if (maxheap.peek()!=null&&arr[i]<maxheap.peek()){
                maxheap.poll();
                maxheap.offer(arr[i]);
            }
        }
        int[] res = new int[maxheap.size()];
        for (int i=maxheap.size()-1;i>=0;i--) {
            if (!maxheap.isEmpty())
            res[i] = maxheap.poll();
        }
        return res;
    }

    public static int[] getLeastNumbers2(int[] arr, int k) {
        return Arrays.stream(arr).sorted().limit(k).toArray();
    }
}
