package leetcode75;

import java.util.Arrays;

public class Main {
    public static void main(String[] args) {
        Solution solution = new Solution();
        int[] a={2,1,0,0,0,1,2};
        solution.sortColors(a);
        System.out.println(Arrays.toString(a));
    }
    static class Solution {
        public void sortColors(int[] nums) {
            for (int i = 0,one=-1,tow=nums.length; i < tow;) {
                if (nums[i]==0){
                    one++;
                    swap(nums,i,one);
                    i++;
                }else if (nums[i]==2){
                    tow--;
                    swap(nums,i,tow);
                }else {
                    i++;
                }
            }
        }
        public static void swap(int[] nums,int i,int j){
            int k=nums[i];
            nums[i]=nums[j];
            nums[j]=k;
        }
    }
}
