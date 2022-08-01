use lazy_static::lazy_static;
use std::collections::HashMap;

#[derive(Debug)]
pub struct User<'a> {
    #[allow(dead_code)]
    username: &'a str,
    #[allow(dead_code)]
    password: &'a str,
}

#[allow(dead_code)]
fn usage() {
    let u = &User {
        username: &String::from("zhang3"),
        password: &String::from("123456"),
    };
    let x: Option<&User> = Some(u);
    match x {
        Some(_) => {
            println!("x is not null");
        }
        None => println!("x is null"),
    }

    let uu = &User {
        username: &"123".to_string(),
        password: &"456".to_string(),
    };
    check_option_struct(Some(uu));
}

pub fn __main__() {
    // let v = search_from_db(&"10001".to_string());
    // match v {
    //     Some(v) => println!("{:#?}", v),
    //     None =>
    // }
}

//static data: HashMap<String, &User> = init_map();

lazy_static! {
    static ref DATA: HashMap<&'static str, &'static User<'static>> = {
        let mut m = HashMap::new();
        m.insert(
            "10001",
            &User {
                username: "abc",
                password: "123456",
            },
        );
        m.insert(
            "10002",
            &User {
                username: "root",
                password: "666666",
            },
        );
        m.insert(
            "10003",
            &User {
                username: "wwww",
                password: "aaaaaa",
            },
        );
        m
    };
}

//fn init_map() -> HashMap<String, &'static User> {
// let mut m: HashMap<String, &'static User> = HashMap::new();
// m.insert(
//     String::from("10001"),
//     &User {
//         username: String::from("abc"),
//         password: String::from("123456"),
//     },
// );
// m.insert(
//     String::from("10002"),
//     &User {
//         username: String::from("root"),
//         password: String::from("666666"),
//     },
// );
// m.insert(
//     String::from("10003"),
//     &User {
//         username: String::from("wwww"),
//         password: String::from("aaaaaa"),
//     },
// );
// m
//}

#[allow(dead_code)]
pub fn search_from_db(id: &str) -> Option<&&User> {
    DATA.get(id)
}

#[allow(dead_code)]
fn check_option_struct(opt: Option<&User>) {
    // match opt {
    //     Some(&User {
    //         username: "123".to_string(),
    //         password: "456".to_string(),
    //     }) => println!("uname = 123 pwd = 456"),
    //     None => println!("user is nil"),
    // }
}

#[allow(dead_code)]
fn check_option_str(opt: Option<&str>) {
    match opt {
        Some("123") => println!("number"),
        Some("abc") => println!("word"),
        None => println!("none"),
        _ => println!("_"),
    }
}
