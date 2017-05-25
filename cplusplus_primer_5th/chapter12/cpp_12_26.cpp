#include <memory>
#include <iostream>
#include <string>

#define STRING_NUM 2

int main()
{
	std::allocator<std::string> a;
	auto const p = a.allocate(STRING_NUM);
	std::string s;
	std::string *ps = p;
	while (std::cin >> s && ps != p + STRING_NUM)
	{
		a.construct(ps);
		*ps++ = s;
	}

	ps = p;
	while (ps != p + STRING_NUM)
	{
		std::cout << "The string input is " << *ps++ << std::endl;
	}
	a.deallocate(p, STRING_NUM);
	return 0;
}
