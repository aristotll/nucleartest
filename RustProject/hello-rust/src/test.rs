#[cfg(test)]
mod test {
    use crate::demo;
    use crate::issuenote;

    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4);
    }

    #[test]
    fn option() {
        demo::option::__main__();
    }

    #[test]
    fn option1() {
        let v = demo::option::search_from_db("10001");
        match v {
            Some(v) => println!("{:#?}", v),
            None => println!("not found"),
        }
    }

    #[test]
    fn move_test() {
        demo::move_test::__main__();
    }

    #[test]
    fn result() {
        demo::result::__main__();
    }

    #[test]
    fn ptr() {
        demo::ptr::__main__();
    }

    #[test]
    fn method() {
        demo::method::__main__();
    }

    #[test]
    fn lifetime() {
        demo::lifetime::__main__();
    }

    #[test]
    fn issue() {
        issuenote::issue_note_temporary_var::__main__();
    }

    #[test]
    // test --package hello-rust --bin hello-rust test::test::net -- --exact --nocapture
    // --nocapture 打印输出
    fn net() {
        demo::net::start_tcp_server();
    }
}
