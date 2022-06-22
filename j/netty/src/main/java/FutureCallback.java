import io.netty.buffer.Unpooled;
import io.netty.channel.Channel;
import io.netty.channel.ChannelFuture;
import io.netty.channel.socket.SocketChannel;
import io.netty.channel.socket.nio.NioServerSocketChannel;

import java.net.InetSocketAddress;

public class FutureCallback {
    public static void main(String[] args) {
        var ch = new NioServerSocketChannel();
        var cf = ch.bind(new InetSocketAddress(8080));
        cf.addListener((future) -> {
            if (future.isSuccess()) {
                System.out.println("监听 8080 端口成功");
            } else {
                System.out.println("监听 8080 端口失败");
                Throwable cause = future.cause();
                cause.printStackTrace();
            }
        });
    }
}
