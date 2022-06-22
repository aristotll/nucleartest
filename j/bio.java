import java.io.IOException;
import java.io.InputStream;
import java.io.OutputStream;
import java.net.ServerSocket;
import java.net.Socket;
import java.util.concurrent.Executor;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

public class bio {
    public static void main(String[] args) throws IOException {
        var port = 6666;
        var executor = Executors.newCachedThreadPool();
        var serverSocket = new ServerSocket(port);
        System.out.printf("start server in %d\n", port);
        while (true) {
            var conn = serverSocket.accept();
            System.out.println("a socket in, addr: " + conn.getRemoteSocketAddress());
            executor.execute(() -> {
                var ch = new EchoHandler(conn);
                ch.start();
            });
        }
    }
}

class EchoHandler extends Thread {
    private Socket socket;

    public EchoHandler(Socket socket) {
        this.socket = socket;
    }

    @Override
    public void run() {
        while (!Thread.currentThread().isInterrupted() &&
                !socket.isClosed()) {
            byte[] buf = new byte[4096];
            try {
                var inputStream = socket.getInputStream();
                int n = inputStream.read(buf);
                var outputStream = socket.getOutputStream();
                outputStream.write(buf);
            } catch (IOException e) {
                throw new RuntimeException(e);
            }
        }
    }
}