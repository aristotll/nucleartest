#include <filesystem>
#include <iostream>

int main() {
  // 创建临时文件夹
  std::filesystem::path temp_dir = std::filesystem::temp_directory_path();
  std::cout << temp_dir << std::endl;
  temp_dir /= "my_temp_folder";
  std::cout << temp_dir << std::endl;
  std::filesystem::create_directory(temp_dir);

  // 程序结束时删除文件夹
  std::filesystem::remove_all(temp_dir);

  return 0;
}
