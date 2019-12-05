import java.util.HashSet;
import java.util.Set;

public class Test {
    public static void main(String[] args) {
        int a=2;
        Set<String> tags=new HashSet<>();
        boolean hitFlag=false;
        if (a>>1==1){
            hitFlag=true;
            tags.add("测试");
        }
//        。。。。。。省略一堆函数
        if (hitFlag){
//           高级别
            System.out.println("定级高");
        }
    }
}
