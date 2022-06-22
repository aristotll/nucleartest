fn ret() -> i64 {
	return 5;
}

fn ret1() -> i64 {
	5
}

fn main() {
	let y = {
		let x = 3;
		let x = x + 1;
		x + 1
	};
	println!("y: {}", y);
	println!("{}", ret());
	println!("{}", ret1());
}
