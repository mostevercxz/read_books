#include <iostream>

int total_count = 0;
class A
{
    public:
    A() : m_count(++total_count) {
        std::cout << "constructor,Count is " << m_count << std::endl;
    }
    ~A() {
        std::cout << "My Count is " << m_count << std::endl;
    }
    private:
        int m_count;
};

void f(A* p)
{
    delete[] p;
}

int main()
{
    A* ai = new A[2]{A(), A()};
    f(ai);
    return 0;
}
