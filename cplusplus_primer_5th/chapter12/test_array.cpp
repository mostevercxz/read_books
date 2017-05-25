#include <iostream>


int total_count = 0;
class A
{
    public:
    A() : m_count(++total_count) {
        std::cout << "constructor,Count is " << m_count << std::endl;
    }
    A(const A&r) : m_count(r.m_count){
        std::cout << "copy constructor" << m_count << std::endl;
    }
    ~A() {
        std::cout << "My Count is " << m_count << std::endl;
    }
    int m_count;
};

void f(A* p)
{
    delete[] p;
}

int main()
{
    /*
    A a1, a2;//two constructor
    A* ai = new A[2]{a1, a2};//two copy constructor
    */
    A *ai = new A[2]{A(), A()};//just two constructor
    //A *pA = new A(A(A(A(A{}))));//just one constructor
    f(ai);//two destructor
    //delete pA;//one destructor
    // two local variable destructor
    return 0;
}
