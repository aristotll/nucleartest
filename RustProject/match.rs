enum Coin {
    Penny,
    Nickel,
    Dime,
    Quarter,
}

fn value_in_cents(coin: Coin) -> u8 {
    match coin {
        Coin::Penny => 1,
        Coin::Nickel => 5,
        Coin::Dime => 10,
        Coin::Quarter => 25,
    }
}

// match 必须要穷举，也就是把每种可能性都罗列出来
// 像上面的 enum Coin 一共有 4 种情况，在函数 value_in_cents 中的 match 全部罗列出来了
// 但是如果像下面这个函数一样，match 的是一个 i64 类型的值，如果没有将所有值全部罗列出来，会报错
// patterns `i64::MIN..=0_i64` and `3_i64..=i64::MAX` not covered  ，告诉我们没有覆盖，
// 需要将 i64::MIN 到 i64::MAX 范围内的所有值都罗列出来，但这显然是不现实的，为了解决这种情况，
// 可以添加一个默认分支，类似于 switch 的 default 分支
fn match_int(i: i64) {
    match i {
    //   ^ patterns `i64::MIN..=0_i64` and `3_i64..=i64::MAX` not covered    
        1 => println!("1"),
        2 => println!("2"),
        _ => println!("!1 and !2")
    }
}

#[derive(Debug)]
struct Rectangle {
    width: u32,
    height: u32,
}

fn match_struct(r: &Rectangle) {
    match (r.width, r.height) {
    //   ^^^^^^^^^^^^^^^^^^^ patterns `(0_u32, _)` and `(2_u32..=u32::MAX, _)` not covered    
        (1, 2) => println!("1, 2"),
        (3, 4) => println!("3, 4"),
        _ => println!("oth"),    
    }
}

fn main() {
	println!("Hello, World!");
	println!("{}", value_in_cents(Coin::Penny));
    match_int(1);
    match_int(10086);
    match_struct(&Rectangle{width: 1, height: 2});
    match_struct(&Rectangle{width: 3, height: 4});
    match_struct(&Rectangle{width: 5, height: 6});
}