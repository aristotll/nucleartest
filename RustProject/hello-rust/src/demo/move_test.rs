pub fn __main__() {
    let mut a = 1;
    let mut b = a;
    b = 100;
    println!("a: {}, b: {}", a, b);
    func1(a);
    println!("{}", a);

	func1_(&mut a);
	println!("{}", a);

    let s = String::from("111");
    //func2(s);
    println!("{}", s);
}

// i 的所有权不会转移
fn func1(i: i64) {
    let mut _b = i;
    _b = 10086;
	println!("func1: {}", _b);
}

fn func1_(i: &mut i64) {
	let mut _i = i;
	*_i = 10086;
}

// s 的所有权会转移
#[allow(dead_code)]
fn func2(s: String) {
    let mut _ss = s;
    _ss = String::from("666");
}
