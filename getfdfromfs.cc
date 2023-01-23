#include <iostream>
#include <fstream>
#include <string>
#include <unistd.h>

int main()
{
  std::string str("Hello World");

  std::ofstream fs("path",
    std::ofstream::binary | std::ofstream::out | std::ofstream::in);

  if (!fs.is_open())
    fs.open("path",
      std::ofstream::binary | std::ofstream::out);

  auto helper = [](std::filebuf& fb) -> int {
    class Helper : public std::filebuf {
    public:
      int handle() { return _M_file.fd(); }
    };

    return static_cast<Helper&>(fb).handle();
  };

  int fd = helper(*fs.rdbuf());

  fs.seekp(0, fs.beg);
  fs.write(str.data(), str.length());
  fsync(fd);
  fs.close();

  return 0;
}
