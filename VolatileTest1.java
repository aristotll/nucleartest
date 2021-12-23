public class VolatileTest1 {
    private int count = 0;

    public static void main(String[] args) {
        while (true) {
            final VolatileTest1 test = new VolatileTest1();
            new Thread(() -> {
                test.count = 1;
            }).start();

            new Thread(() -> {
                System.out.println(test.count);
            }).start();

            try {
                Thread.sleep(500);
            } catch (Exception e) {
                System.out.print(e);
            }
        }
    }
}
