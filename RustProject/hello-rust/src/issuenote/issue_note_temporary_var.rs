pub fn __main__() {}

#[derive(Debug)]
struct User {
    name: String,
    age: i8,
}

fn __error3__() {
    let u;
    let str1 = String::from("li");
    u = &User { name: str1, age: 0 };
    // ^^^^^^^^^^^^^^^^^^^^^^^^^^^- temporary value is freed at the end of this statement
    println!("{:#?}", u);
    //                - borrow later used here
}

fn __true3__() {
    let u: &User;
    let str1 = String::from("li");
    {
        // 这里的写法和 error3 一模一样，只不过放到了一个单独的作用域内，为什么就不会报错了？
        // 可能的原因：
        // 貌似不是因为作用域的关系，而是这里用 let 定义了一个新的变量，而在 __error3__ 中
        // 是先定义了一个 u，再将 &User 赋值给 u，比如下面的 __true33__ 定义了一个新的变量
        // let u = &User{}，此时就不会产生错误
        // 二者的区别（猜测）：
        // 1. let u; u = &User{} 代表的是：u 借用了这个 User{} 变量
        // 2. let u = &User{} 代表的是：将 User 这个变量绑定到 u
        // 1 中的 User{} 因为没有绑定变量，所以成了临时变量，刚创建完生命周期就结束了，
        // 而 u 又借用了这个已经被释放的变量，所以会发送错误
        // 2 中是将 User 绑定到了变量 u，这样生命周期就不会立马结束了
        // ps：以上纯属个人猜测
        let uu = &User {
            //name: str.to_string(),
            name: str1,
            age: 0,
        };
        println!("{:#?}", uu);
    }
}

fn __true33__() {
    let str1 = String::from("li");
    let u = &User { name: str1, age: 0 };
    println!("{:#?}", u);
}

fn __true333__() {
    let u;
    let str1 = String::from("li");
    u = User { name: str1, age: 0 }; // 删掉 &，这样可能代表直接绑定
    println!("{:#?}", u);
}
