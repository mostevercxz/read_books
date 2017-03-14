## Chapter 12, dynamic memory

Operations common to `shared_ptr` and `unique_ptr`:

    shared_ptr<T> sp;
    unique_ptr<T> up;
    p
    *p
    p->mem    synonym for (*p).mem
    p.get()
    swap(p, q)
    p.swap(q)
    
Operations specific to shared_ptr:

    make_shared<T>(args)
    shared_ptr<T> p(q)
    p = q        Decrements p's reference count and increments q's count.
    p.unique()   Return true if p.use_count()==1
    p.use_count()
    
We can deduce the type of the object we want to allocate from the initializer, while only single initializer inside parentheses is allowed beacause the compile uses the initializer's type to deduce the type to allocate.

    MyClass obj;
    auto p1 = new auto(obj);
    
It is legal to use new to allocate const objects.

    const int *pci = new const int();
    const string *cs = new const string;
    
### 12.1.3Using shared_ptrs with new
Other ways to define and change shared_ptrs:

    shared_ptr<int> p2(new int(42));
    shared_ptr<T> p(q);//q must point to memory allocated by new,can be convertible to T*
    shared_ptr<T> p(u);//assumes ownership from unique_ptr u, makes u null
    shared_ptr<T> p(q, d);//p will use the callable object d in place of delete to free q
    shared_ptr<T> p(p2, d);
    p.reset();
    p.reset(q);
    p.reset(q, d);
    
The smart pointer constructors that take pointers are **explicit**.

    shared_ptr<int> pi = new int();//error, must use direct initialization
    shared_ptr<int> pi2(new int());//ok

### 12.1.4 Smart Pointers and Exceptions

    void f()
    {
        int *pi = new int();
        // code that throws an exception that is not caught inside f
        delete pi;
    }

If an exception happens between `new` and `delete`, and is not caught inside f, then this memory can never be freed.

Smart Pointer pitfalls:
 1. Do not use the same built-in pointer value to initialize(or reset) more than one smart pointer(may be deleted by smart pointer)
 2. Do not `delete` the pointer returned from get()
 3. Do not use `get()` to initialize or reset another smart pointer;
 4. If you use a pointer returned by `get()`, remember that the pointer will become invalid when the last corresponding smart pointer goes away.
 5. Is you use a smart pointer to manage a resource rather than memory allocated by `new`, remember to pass a deleter. 
 
###  12.1.5 unique_ptr

    unique_ptr<T> u1;//Null unique_ptr that points to objects of type T.
    unique_ptr<T, D> u2;//u1 use delete to free a pointer, u2 use a callable object of type D to free its pointer 
    unique_ptr<T, D> u3(d);//u3 uses d to free its pointer, which must be an object of type D in place of `delete`
    u = nullptr;//delete the object to which u points; makes u null.
    u.release();//relinquishes control of the pointer u had held, returns the pointer u and makes u null.
    u.reset();//delete the object to which u points.
    u.reset(q);
    u.reset(nullptr);
    
Passing a deleter to unique_ptr

    void f(destination &dest)
    {
        connection c = connect(&dest);
        unique_ptr<connection, decltype(end_connection)*> p(&c, end_connection);
    }

### 12.1.6 weak_ptr

A weak_ptr is a smart pointer that does not control the lifetime of the object to which it points.  Binding a weak_ptr to a shared_ptr does not change the reference count of the shared_ptr.

    weak_ptr<T> w;
    weak_ptr<T> w(sp);//sp is a shared_ptr. T must be convertible to the type sp points.
    w = p;//p is shared_ptr or weak_ptr
    w.reset();
    w.use_count();//The number of shared_ptrs that share ownership with w.
    w.expired();
    w.lock();//if `expired()` is true, returns a null shared_ptr; otherwise returns a shared_ptr to the object w points.


When a `weak_ptr` is useful?

`weak_ptr` is used to solve the dangling pointer problem. By using raw pointers, it is impossible to know if the referenced data has been freed or not. Instead, use `shared_ptr` to manage the data, supply `weak_ptr` to users of the data, the users can check validity of the data by calling `expired()` or `lock()`.

(A good example may be cache. You want to keep frequently accessd objects in memory  using `shared_ptr`, and get rid of the non-frequently accessed objects. But if the object is in use and some other code holds a `shared_ptr` to it? If the cache gets rid of its pointer to the object, it can never be found again(freed). A `weak_ptr` allows you to locate an object if it's still valid, but does not keep it around if nothing else needs it.) 

    int main()
    {
        std::shared_ptr<int> spi(new int(10));
        std::weak_ptr<int> wp(spi);
        spi.reset(new int(5));//the origin pointer is freed, acquires a new pointer
        if (auto tmp = wp.lock())
        {
            std::cout << "pointer is valid,value=" << *tmp << std::endl;
        }
        else
        {
             std::cout << "Shit, the origin pointer is freed" << std::endl;
        }
    }
    
## 12.2 Dynamic Arrays

Two ways to allocate an array of objects at once:
1. The language defines a second kind of `new` expression that allocates and initializes an array of objects. 
2. The library includes a template class named `allocator` that lets us separate allocation from initialization. 

Using an allocator generally provides better performance and more flexible memory management, Why ??

Libraries that support the new standard tend to be dramatically faster than previous releases, Why ??

**Initialize an array of dynamically allocated objects**

    int *pia0 = new int[10];//10 uninitialized ints
    int *pia1 = new int[10]();//10 ints valued initialized to 0
    int *pia2 = new int[10](1);//error, must use braced list of element initializers
    int *pia = new int[10]{0, 1, 2, 3};//the first 4 is initialized, other are value initialized.
    string *psa = new string[5]{"hello", string(3, 'x')};

**Legal to dynamically allocate an empty array.**

    // new returns a valid,nonzero pointer, guaranteed to be distinct from any other pointer returned by new, acting as the off-the-end pointer.
    char *pca = new char[0];
    int ia[] = {0, 1, 2, 3};
    // `end` returns a pointer one past the last element in the array list.(off-the-end pointer)
    int *pbeg = begin(ia); *pend = end(ia);
    int sum = 0; 
    while (pbeg != pend)
    {
         sum += *pbeg;
    }
    
**Freeing dynamically arrays**

    delete [] pa;//Elements in an array are destroyed in reverse order.
    
Test example, Why freed in reverse order?
    
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

    uname@server:~$ g++ -g -std=c++11 test_array.cpp -o test_array
    uname@server:~$ ./test_array 
    constructor,Count is 1
    constructor,Count is 2
    My Count is 2
    My Count is 1
    
gdb test_array,断点 main 函数, run
disas 查看汇编代码(g++ -O0,没做任何优化),构造部分在main函数中:

       0x0000000000400a8c <+11>:  sub    $0x10,%rsp
    => 0x0000000000400a90 <+15>:  mov    $0x10,%edi//传参,请求分配0x10=16字节(前8个字节存储数组元素个数,后面8=4+4,两个A对象的大小)
       0x0000000000400a95 <+20>:  callq  0x400870 <_Znam@plt> //分配内存
       0x0000000000400a9a <+25>:  mov    %rax,%r12
       0x0000000000400a9d <+28>:  movq   $0x2,(%r12)//将个数2写进前8个字节
       0x0000000000400aa5 <+36>:  lea    0x8(%r12),%rbx
       0x0000000000400aaa <+41>:  mov    $0x1,%r13d
       0x0000000000400ab0 <+47>:  mov    %rbx,%rdi
       0x0000000000400ab3 <+50>:  callq  0x400b86 <A::A()>
       0x0000000000400ab8 <+55>:  lea    0x4(%rbx),%rax
       0x0000000000400abc <+59>:  sub    $0x1,%r13
       0x0000000000400ac0 <+63>:  mov    %rax,%rdi
       0x0000000000400ac3 <+66>:  callq  0x400b86 <A::A()>
       0x0000000000400ac8 <+71>:  lea    0x8(%r12),%rax
       0x0000000000400acd <+76>:  mov    %rax,-0x28(%rbp)
       0x0000000000400ad1 <+80>:  mov    -0x28(%rbp),%rax
       0x0000000000400ad5 <+84>:  mov    %rax,%rdi
       0x0000000000400ad8 <+87>:  callq  0x400a26 <f(A*)>//调用函数 f(A*)准备析构

析构部分在 f(A*) 中:

     0x0000000000400a3a <+20>:  mov    -0x18(%rbp),%rax//指针存在 -0x18(%rbp)中
     0x0000000000400a3e <+24>:  sub    $0x8,%rax
     0x0000000000400a42 <+28>:  mov    (%rax),%rax
     0x0000000000400a45 <+31>:  lea    0x0(,%rax,4),%rdx//rdx中存储数组所有元素所占字节数 2*4=8
     0x0000000000400a4d <+39>:  mov    -0x18(%rbp),%rax
     0x0000000000400a51 <+43>:  lea    (%rdx,%rax,1),%rbx//rbx代表 off-the-end pointer,用作终止条件
     0x0000000000400a55 <+47>:  cmp    -0x18(%rbp),%rbx
     0x0000000000400a59 <+51>:  je     0x400a69 <f(A*)+67>
     0x0000000000400a5b <+53>:  sub    $0x4,%rbx//rbx-4得到数组最后一个元素的地址
     0x0000000000400a5f <+57>:  mov    %rbx,%rdi
     0x0000000000400a62 <+60>:  callq  0x400be2 <A::~A()>
     0x0000000000400a67 <+65>:  jmp    0x400a55 <f(A*)+47>
     0x0000000000400a69 <+67>:  mov    -0x18(%rbp),%rax
     0x0000000000400a6d <+71>:  sub    $0x8,%rax//将数组地址前面的8个字节也释放掉
     0x0000000000400a71 <+75>:  mov    %rax,%rdi
     0x0000000000400a74 <+78>:  callq  0x4008d0 <_ZdaPv@plt>

由析构代码可以看出,汇编代码先算出off-the-end pointer 的位置,然后用 while 循环,每次减掉数组元素大小个字节，调用析构函数。
是逆序析构,直到所有元素都被析构完后,再将数组地址的前8个字节(`sub $0x8, %rax`)也归还给操作系统(_ZdaPv是系统释放内存函数).

那为什么不用 while 循环，从数组第1个元素开始加一直到结束地址呢（算出rbx后,调用 `add $0x4,%rax`）？
同样是只用2个寄存器，为啥要逆序析构呢？
个人认为原因是，g++做优化时候用逆序可以少做一步赋值操作(`mov    -0x18(%rbp),%rax`)
函数入口先将 rdi 寄存器里的值保存到某寄存器(r13)中，再用rdi取出元素个数，算出终止的地址位置后
，再使用 r13 和终止地址做比较，调用完析构函数后,最终将 r13 的地址再赋给 rdi,调用系统 free 函数.

mov %rdi, %r13
mov -0x8(%rdi), %rax
lea (%rdi, %rax, 4), %rbp
cmp %rbp, %r13
je .L2//如果相等就该结束了
sub $0x4, %rbp//每次减4个字节,依次传地址析构

.L2:
lea -0x8(%r13), %rdi//将数组首地址倒退8个字节赋值给rdi
callq _ZdaPv@plt//释放内存

(Makefile 中调用 g++ -O1,发现汇编代码确实和上面的相似)

    
Smart pointers and dynamically arrays
