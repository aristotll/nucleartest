package nio;

import java.io.IOException;
import java.net.InetSocketAddress;
import java.nio.ByteBuffer;
import java.nio.channels.SelectionKey;
import java.nio.channels.Selector;
import java.nio.channels.ServerSocketChannel;
import java.nio.channels.SocketChannel;

public class NIOEchoServer {
    public static void main(String[] args) throws IOException {
        var port = 8080;
        var serverSocketChannel = ServerSocketChannel.open();
        serverSocketChannel.socket().bind(new InetSocketAddress(port));
        System.out.println("listen in " + port);
        serverSocketChannel.configureBlocking(false); // 非阻塞
        var selector = Selector.open(); // 创建一个 selector
        // 将 serverSocketChannel 注册到 selector
        serverSocketChannel.register(selector, SelectionKey.OP_ACCEPT);
        while (selector.select() > 0) { // 产生事件了
            var selectionKeys = selector.selectedKeys();
            var iterator = selectionKeys.iterator();
            while (iterator.hasNext()) {    // 遍历所有事件进行处理
                var next = iterator.next();
                if (next.isAcceptable()) { // 如果是连接事件
                    var conn = serverSocketChannel.accept();
                    System.out.println("a accept event is coming, addr: "
                            + conn.getRemoteAddress());
                    conn.configureBlocking(false);
                    // 将该连接注册到 selector，关注事件为读事件
                    conn.register(selector, SelectionKey.OP_READ, ByteBuffer.allocate(1024));
                } else if (next.isReadable()) { // 如果是可读事件
                    // 通过 key 反向获取到对应的 channel
                    var channel = (SocketChannel) next.channel();
                    System.out.println("a read event is coming, addr: "
                            + channel.getRemoteAddress());
                    var buffer = (ByteBuffer) next.attachment();
                    try {
                        while (channel.read(buffer) != -1) {
                            // 不加貌似也没影响，应该是写模式切换到读模式时才需要调用
                            // 更新：必须要加，不然用 nc 连接会无限收到消息
                            // flip 会把当前指针置 0，limit 置之前的当前指针（也就是写多少读多少）
                            buffer.flip();
                            channel.write(buffer);
                            buffer.clear();
                        }
                    } catch (IOException err) {
                        err.printStackTrace();
                        System.out.println(err.getMessage());
                        channel.close();
                    }
                }
            }
            iterator.remove();
        }
    }
}
