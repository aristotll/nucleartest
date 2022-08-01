fn func(mut s: String) {
	s = String::from("123");
}

fn main() {
	let s = String::from("1");
	func(s);
	println!("{}", s);
}
