import java.util.Deque;
import java.util.LinkedList;

public class Test {
    public static void main(String[] args) {
        System.out.println(isValid("([)]"));
    }

    public static boolean isValid(String s) {
        Deque<Character> deque=new LinkedList();
        if (s.length()==0){
            return true;
        }
        if (s.length() <= 1) {
            return false;
        }
        deque.add(s.charAt(0));
        for (int i = 1; i < s.length(); i++) {
            char c = s.charAt(i);
            if (!deque.isEmpty()){
                Character poll = deque.peekLast();
                if (poll=='('&&c==')'){
                    deque.pollLast();
                    continue;
                }
                if (poll=='['&&c==']'){
                    deque.pollLast();
                    continue;
                }
                if (poll=='{'&&c=='}'){
                    deque.pollLast();
                    continue;
                }
            }
            deque.add(c);

        }
        if (deque.isEmpty()){
            return true;
        }else {
            return false;
        }
    }
}
