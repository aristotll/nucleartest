package nio;

import java.io.IOException;
import java.net.InetSocketAddress;
import java.nio.ByteBuffer;
import java.nio.channels.*;
import java.util.Iterator;

public class GroupChatServer {
    private Selector selector;
    private int port;

    public GroupChatServer(int port) {
        this.port = port;
    }

    public void server() throws IOException {
        System.out.println("listen in " + port);
        this.selector = Selector.open();
        System.out.println(selector == null);

        var listener = ServerSocketChannel.open();
        listener.socket().bind(new InetSocketAddress(port));
        listener.configureBlocking(false);
        listener.register(selector, SelectionKey.OP_ACCEPT);

        while (selector.select() > 0) {
            var iterator = selector.selectedKeys().iterator();
            while (iterator.hasNext()) {
                var event = iterator.next();
                if (event.isAcceptable()) {
                    var conn = listener.accept();
                    System.out.println("a accept event is coming, addr: "
                            + conn.getRemoteAddress());
                    conn.configureBlocking(false);
                    conn.register(selector, SelectionKey.OP_READ);
                    System.out.println("register read event success");
                } else if (event.isReadable()) {
                    SocketChannel channel = null;
                    try {
                        channel = (SocketChannel) event.channel();
                        var buf = ByteBuffer.allocate(4096);
                        if (channel.read(buf) != -1) {
                            // 将读到的消息转发给所有其他用户（除自己）
                            for (var c : selector.keys()) {
                                Channel ch = c.channel();
                                // 排除自身
                                if (ch instanceof SocketChannel && ch != channel) {
                                    // 转型
                                    SocketChannel dest = (SocketChannel) ch;
                                    buf.flip(); // 没有这句话会导致接收方读不到消息，但是消息已经成功发生
                                    // 将 buffer 中的数据写入通道
                                    dest.write(buf);
                                    System.out.println(channel.getRemoteAddress()
                                            + " -> " + dest.getRemoteAddress()
                                            + " msg: "
                                            + new String(buf.array(), 0, buf.limit()));
                                }
                            }
                        }
                    } catch (IOException e) {
                        e.printStackTrace();
                        channel.close();
                    }
                }
                iterator.remove(); // 必须加这句
            }
        }
    }

    public static void main(String[] args) throws IOException {
        var server = new GroupChatServer(8080);
        server.server();
    }
}
