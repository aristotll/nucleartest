package nio;

import java.io.IOException;
import java.net.InetSocketAddress;
import java.nio.ByteBuffer;
import java.nio.channels.SocketChannel;

public class NIOEchoClient {
    public static void main(String[] args) throws IOException {
        var socketChannel = SocketChannel.open();
        socketChannel.configureBlocking(false);
        var addr = new InetSocketAddress("127.0.0.1", 8080);
        if (!socketChannel.connect(addr)) {
            while (!socketChannel.finishConnect()) {
                System.out.println("连接需要时间,客户端不会阻塞...先去吃个宵夜");
            }
        }
        var msg = "Hello";
        var buf = ByteBuffer.wrap(msg.getBytes());
        socketChannel.write(buf);
        //buf.reset();
        socketChannel.read(buf);
        System.out.println("read from conn: " + new String(buf.array(), 0, buf.limit()));
        socketChannel.close();
        System.out.println("客户端退出");
    }
}
