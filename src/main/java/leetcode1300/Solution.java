package leetcode1300;

import java.util.Arrays;

/**
 * @author lin
 * Email 451846939@qq.com
 * @version V1.0
 * @description
 * @date 2020/6/14 6:45 下午
 */
public class Solution {
    public int findBestValue(int[] arr, int target) {
        if (arr==null){
            return 0;
        }
        Arrays.sort(arr);
        int sum=0;
        int length = arr.length;
        for (int i = 0; i < length; i++) {
            int x = (target - sum) /  (length - i);
            if (x<arr[i]){
                double t = (double)(target - sum) / (double)(length - i);
                if (t-x>0.5){
                    return x+1;
                }else {
                    return x;
                }
            }
            sum+=arr[i];
        }
        return arr[length-1];
    }
}