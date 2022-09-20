use hyper::service::{make_service_fn, service_fn};
use hyper::{Body, Request, Response, Server};
use hyper::{Method, StatusCode};
use serde::Deserialize;
use std::convert::Infallible;
use std::net::SocketAddr;


#[derive(Deserialize, Debug)]
struct User {
    id: i64,
    uname: String,
}

async fn hello_world(req: Request<Body>) -> Result<Response<Body>, hyper::Error> {
    let mut response = Response::new(Body::empty());
    match (req.method(), req.uri().path()) {
        (&Method::GET, "/") => {
            *response.body_mut() = Body::from("Try POSTing data to /echo");
        }
        (&Method::POST, "/") => {
            let b = hyper::body::to_bytes(req).await?;
            let bb = b.iter().as_slice();
            let u: User = serde_json::from_slice(bb).unwrap();
            println!("{:?}\n", u);
        }
        _ => {
            *response.status_mut() = StatusCode::NOT_FOUND;
        }
    }
    Ok(response)
}


#[tokio::main]
async fn main() {
    let addr = SocketAddr::from(([127, 0, 0, 1], 8080));
    let make_svc = make_service_fn(|conn| async { Ok::<_, Infallible>(service_fn(hello_world)) });

    let server = Server::bind(&addr).serve(make_svc);
    // 运行 server
    if let Err(e) = server.await {
        eprintln!("server error: {}", e);
    }
}
