#[derive(Debug)]
struct Rectangle {
    width: u32,
    height: u32,
}

impl Rectangle {
    // &self 代表借用
    fn area_no_borrow(&self) -> u32 {
        self.height * self.width
    }

    // self 代表获取所有权，需要通过返回值归还
    // 注意这里的返回值定义顺序，需要 u32 在前，self 在后，如果 self 在前，那么
    // 在计算 self 的面积之前，self 自身已经归还了，不能再使用其 width 和 height 变量
    fn area_borrow(self) -> (u32, Self) {
        (self.width * self.height, self)
    }

    fn can_hold(self, oth: &Rectangle) -> bool {
        (self.width * self.height) > (oth.width * oth.height)
    }
}

pub fn __main__() {
    let r = Rectangle {
        width: 10,
        height: 20,
    };
    let (ret, r_) = r.area_borrow();
    // ------ `r` moved due to this method call
    // println!("{:#?}", r);
    //                   ^ value borrowed here after move
    println!("area: {}", ret);
    println!("{:#?}", r_);

    let r1 = &Rectangle {
        width: 5,
        height: 10,
    };
    // 这里注意使用 area_borrow 归还的 r_ 而不是 r，因为 r 的所有权已经转移了
    println!("{}", r_.can_hold(r1));
}
