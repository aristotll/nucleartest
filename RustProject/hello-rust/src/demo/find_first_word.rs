sfn find_first_word(s: &String) -> String {
	let mut ss: String = String::from("");
	for c in s.chars() {
		if c == ' ' {
			break;		
		}
		ss.push(c);
	}
	return ss;
}

fn main() {
	let str = String::from("my name");
	let ret = find_first_word(&str);
	println!("{}", ret);
}
