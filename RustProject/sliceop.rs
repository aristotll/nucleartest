fn main() {
	let s = String::from("12345");
	let mut ss = &s[0..3];
	println!("s=12345, s[0..3]={}", ss);
	ss = "456";
	println!("s={}", s);	
}
