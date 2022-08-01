#[derive(Debug)]
struct User {
    name: String,
    age: i8,
}

fn __error1__() {
    let z;
    {
        let u = &User {
            name: "123".to_string(),
            age: 18,
        };
        z = u;
        println!("{:#?}", z);
    }
    //println!("{:#?}", z);
}

// 为什么 u = &User，之后在一个 {} 内输出，编译会报错？
// 去掉 & 可以正常编译，或者用 __true2__ 中的方法
// fn __error2__() {
//     let u;
//     let str1 = String::from("li");
//     {
//         let str = "zhang";
//         // 猜测：莫非是因为这里借用了一个临时变量，而这个临时变量因为没有绑定变量，
//         // 所以立马被释放了，导致 u 引用了一个悬空指针？
//         u = &User {
//             //name: str.to_string(),
//             name: str1,
//             age: 0,
//         };
//         // ^- temporary value is freed at the end of this statement
//         println!("{:#?}", u);
//     }
//     //println!("{:#?}", u);
// }
//
// // 为什么定义一个 uu 保存 &User，再赋值给 u，编译就不会报错？
// fn __true2__() {
//     let u: &User;
//     let str1 = String::from("li");
//     {
//         let str = "zhang";
//         let uu = &User {
//             //name: str.to_string(),
//             name: str1,
//             age: 0,
//         };
//         println!("{:#?}", uu);
//         //u = uu;
//         //println!("{:#?}", u);
//     }
// }

pub fn __main__() {
    //__error2__();
}
