use std::fs;
use std::io::{Read, Write};
use std::net::{TcpListener, TcpStream};

pub fn start_tcp_server() {
    let lis = TcpListener::bind("127.0.0.1:7788");
    match lis {
        Ok(tcp_lis) => {
            println!("listen in {}", tcp_lis.local_addr().unwrap());
            loop {
                let (conn, addr) = tcp_lis.accept().unwrap();
                //echo_handler(conn);
                http_handler(conn);
            }
        }
        Err(err) => println!("bind err: {}", err),
    }
}

fn echo_handler(mut stream: TcpStream) {
    let mut buf = [0; 1024];
    loop {
        // &mut 代表借用 buf，并且可以对其进行写入操作
        stream.read(&mut buf).unwrap();
        println!(
            "[{}] recv: {}",
            stream.local_addr().unwrap(),
            String::from_utf8_lossy(&buf[..])
        );
        stream.write(&buf[..]).unwrap();
    }
}

fn http_handler(mut stream: TcpStream) {
    let mut buf = [0; 1024];
    loop {
        // &mut 代表借用 buf，并且可以对其进行写入操作
        stream.read(&mut buf).unwrap();

        let file_src = "~/pj/justtest/RustProject/hello-rust/src/demo/hello.html";
        let contents = fs::read_to_string(file_src).unwrap();
        let response = format!(
            "HTTP/1.1 200 OK\r\nContent-Length: {}\r\n\r\n{}",
            contents.len(),
            contents
        );
        stream.write(response.as_bytes()).unwrap();
        stream.flush().unwrap();
    }
}
