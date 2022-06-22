import io.netty.bootstrap.ServerBootstrap;
import io.netty.buffer.ByteBuf;
import io.netty.buffer.Unpooled;
import io.netty.channel.*;
import io.netty.channel.nio.NioEventLoopGroup;
import io.netty.channel.socket.SocketChannel;
import io.netty.channel.socket.nio.NioServerSocketChannel;

import java.net.InetSocketAddress;

public class EchoServer {
    private final int port;

    public EchoServer(int port) {
        this.port = port;
    }

    public static void main(String[] args) throws Exception {
        if (args.length != 1) {
            System.out.println("usage: " +
                    EchoServer.class.getSimpleName() + " <port> ");
        }
        var port = Integer.parseInt(args[0]);
        new EchoServer(port).start();
    }

    public void start() throws Exception {
        final var serverHandler = new EchoServerHandler();
        var group = new NioEventLoopGroup(); // 创建事件循环
        try {
            var b = new ServerBootstrap(); // 用来初始化 server
            b.group(group)
                    .channel(NioServerSocketChannel.class)  // 指定 channel 类型
                    .localAddress(new InetSocketAddress(this.port)) // 指定端口
                    .childHandler(
                            // ChannelInitializer 当一个新的连接被接受时，一个新的子 Channel 将会被创建，
                            // 而 ChannelInitializer 将会把一个你的 EchoServerHandler 的实例添加到该
                            // Channel 的 ChannelPipeline 中
                            new ChannelInitializer<SocketChannel>() {
                                @Override
                                public void initChannel(SocketChannel ch) throws Exception {
                                    //  添加一个 EchoServerHandler 到子 Channel 的
                                    //  ChannelPipeline
                                    ch.pipeline().addLast(serverHandler);
                                }
                            });
            // 等待绑定完成。对 sync()方法的调用将导致当前 Thread 阻塞，一直到绑定操作完成为止
            var f = b.bind().sync();
            //
            f.channel().closeFuture().sync();
        } finally {
            group.shutdownGracefully().sync();
        }
    }

    @ChannelHandler.Sharable
    public class EchoServerHandler extends ChannelInboundHandlerAdapter {
        @Override
        public void exceptionCaught(ChannelHandlerContext ctx, Throwable cause) throws Exception {
            cause.printStackTrace();
            ctx.close();
        }

        @Override
        public void channelRead(ChannelHandlerContext ctx, Object msg) throws Exception {
            var in = (ByteBuf) msg;
            System.out.println("server received: " + in.toString());
            ctx.write(in);
        }

        @Override
        public void channelReadComplete(ChannelHandlerContext ctx) throws Exception {
            ctx.writeAndFlush(Unpooled.EMPTY_BUFFER)
                    .addListener(ChannelFutureListener.CLOSE);
        }
    }
}
