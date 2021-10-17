// import java.lang.system;

public class FuncStaticVar {
    public static void test() {
        int x = 0;
        // static int x = 0; error
        System.out.println(x);
    }
    
    
    public static void main(String[] args) {
        for (int i = 0; i < 5; i++) {
            test();
        }
    }
}

