use std::fs::File;
use std::io;
use std::io::Read;

fn read_username_from_file() -> Result<String, io::Error> {
    let mut f = File::open("hello.txt")?;
    let mut s = String::new();
    f.read_to_string(&mut s)?;
    Ok(s)
}

pub fn __main__() {
    match read_username_from_file() {
        Ok(s) => println!("file content: {}", s),
        Err(e) => println!("error: {}", e),
    }
}
