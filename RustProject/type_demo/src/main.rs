fn main() {
    let tup: (i32, f64, u8) = (500, 6.4, 1);
    let (x, y, z) = tup;

    println!("the value of y is: {}", y);

    let x1 = tup.0;
    let y1 = tup.1;
    let z1 = tup.2;

    println!("x1: {}, y1: {}, z1: {}", x1, y1, z1);
}
