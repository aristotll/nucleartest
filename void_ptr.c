#include <stdio.h>

int main(int argc, char const *argv[])
{
    void * str = "abc\n";
    char *str1;
    str1 = str;
    str = str1;
    printf(str1);
    return 0;
}
