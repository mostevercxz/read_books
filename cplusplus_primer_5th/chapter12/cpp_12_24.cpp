#include <iostream>
#include <string>

int main()
{
	std::string s;
	std::cin >> s;
	char *pca = new char[s.length() + 1]();
	for (unsigned int i = 0; i < s.length(); ++i)
	{
		pca[i] = s[i];
	}
	std::cout << "You input:" << std::endl << pca << std::endl;
	return 0;
}
