#include <string.h>
#include <memory>
#include <iostream>

void concatenate(const char *s1, const char *s2)
{
    size_t totalLen = strlen(s1) + strlen(s2);
    char *pc = new char[totalLen + 1]();
    std::unique_ptr<char[]> u(pc);
    size_t index = 0;
    while (*s1 != 0)
    {
        u[index++] = *s1++;
    }
    while (*s2 != 0)
    {
        u[index++] = *s2++;
    }

    std::cout << u.get() << std::endl;
}

int main()
{
    concatenate("hello,", " world!");
    return 0;
}
