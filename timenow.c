#include <stdio.h>
#include <time.h>

int main()
{
time_t t = time(NULL);
struct tm* lt = localtime(&t);
char buffer[80];
strftime(buffer, 80, "%Y-%m-%d %H:%M:%S", lt);
printf("%s\n", buffer);
return 0;
}
