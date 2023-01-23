#include <stdio.h>
#include <strings.h>

char *max(const char *s1, const char *s2) {
    int ret = strcasecmp(s1, s2);
    if (ret < 0) {
        return s2;
    }
    return s1;
}

int main() {
    printf("%s\n", max("abc", "def"));
}
