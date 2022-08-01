fn ret() -> i64 {
	return 5;
}

fn ret1() -> i64 {
	5
}

fn ret2() -> (i64, i64) {
	(1, 2)
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
	
	let (a, b) = ret2();
	println!("{}, {}", a, b);
}
