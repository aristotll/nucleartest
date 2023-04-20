struct CleanupFile(String, i64);

fn main() {
    let file_path = String::from("file.txt");
    let file_size = 1024;
    let cleanup_file = CleanupFile(file_path, file_size);

    println!("path: {}, size: {}", cleanup_file.0, cleanup_file.1);

    let CleanupFile(path, size) = cleanup_file;
    println!("path: {}, size: {}", path, size);

}
