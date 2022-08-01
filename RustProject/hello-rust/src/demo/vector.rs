fn main() {
    let mut v = vec![1, 2, 3, 4];
    let i = &mut v[0];
    println!("{}", i);

    *i = 100;
    println!("{}", i);
    println!("{:?}", v);

    let index = 100;
    let res = v.get(index);
    match res {
        Some(val) => {
            println!("index {} is: {}", index, val);
        },
        None => println!("index {} is overflow", index),
    }

    for val in &v {
        print!("{} ", val);
    }
    println!();

    for (index, val) in v.iter().enumerate() {
        print!("[{}]:{} ", index, val);
    }
    println!();

    println!("{:?}", v);
}
