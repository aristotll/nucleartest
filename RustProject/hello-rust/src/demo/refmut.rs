fn main() {
	let mut a = 10;
	let b = &mut a;
	//let d = &mut b;
	println!("{}", b);
	let c = &mut a;
	println!("{}", c);
}
