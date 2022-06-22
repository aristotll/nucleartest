struct User {
    username: String,
    password: String,
}

fn main() {
    let u = &User{
        username: String::from("zhang3"),
        password: String::from("123456"),
    };
    let x: Option<&User> = Some(u);
    match x {
        Some(_) => {
            println!("x is not null");
        }
        None => println!("x is null")
    }
}