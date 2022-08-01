#[derive(Debug)]
struct Rectangle {
    width: u32,
    height: u32,
}

fn main() {
    let rect1 = Rectangle {
        width: 30,
        height: 50,
    };
    
    // rect1 is Rectangle {
    // width: 30,
    // height: 50,
    // }   
    println!("rect1 is {:#?}", rect1);

    // Alice, this is Bob. Bob, this is Alice
    println!("{0}, this is {1}. {1}, this is {0}", "Alice", "Bob");
    // 1, 2 
    println!("{}, {}", 1, 2);

    let st = "wang5";
    // z3, l4, wang5
    println!("{a}, {b}, {c}", a="z3", b="l4", c=st);
    
}
