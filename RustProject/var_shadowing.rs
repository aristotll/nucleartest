fn main() {
	let x = 5;
	{
		let x = x * 2;
		println!("The value of x in the inner scope is: {}", x);
	}
	println!("The value of x is: {}", x);
}

// func main() {
//        x := 5
//        {
//                x := x * 2
//                fmt.Println(x)
//        }
//        fmt.Println(x)
// }
