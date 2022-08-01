pub fn __main__() {
    let a = &mut 123;
    *a = 50;
    println!("{}", a);

    let mut b = 1;  // b 也必须是 mut 的
    // 如果引用加上了 mut（&mut），那么引用的本身（b）也必须是 mut 的
    let bb = &mut b;
}