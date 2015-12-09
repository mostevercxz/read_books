#Part I, Program Structure and Execution
1. void swap(int x, int y) {y = x ^ y; x = x ^ y; y = x ^ y;}; swap function is not suitable for exchange array members. 异或版本的swap不适合在转置奇数数组中使用
2. 
## chapter 2,Representing and Manipulating info
We consider the three most important representations of numbers:

 * *Unsigned* encodings
 * *Two's-complement* encodings, signed integers
 * *Floating-point* encodings, real numbers
### 2.1 info storage
8 bits, or bytes : the smallest addressable unit of memory.

A machine-level program views memory as a very large array of bytes, referred to as *virtual memory*

Every byte of memory is identified by a unique number, known as its address, and the set of all possible addresses is known as the *virtual address space*.

The actual implementation uses a combination of dynamic random access memory (DRAM), flash memory, disk storage, special hardware, and operating system software to provide the program with what appears to be a monolithic 整体的,巨大的,完全统一的 byte array.


#### 2.1.1 Hexadecimal Notation

When a value x is a power of 2, x = 2 ^ n for some non-negative integer n:

**n = i + 4j**, 0 <= i <= 3, we can write x with a leading hex digit of 1(i=0), 2(i=1), 4(i=2), 8(i=3), followed by j hexadecimal 0s.

To convert a decimal number x to hexadecimal, we can repeatedly divide x by 16, giving a quotient q and a remainder r, such that x = q * 16 + r, We then use the hexadecimal digit representing r as the least significant digit·and generate the remaining digits by repeating the process on q.

Conversely, to convert a hexadecimal number to decimal, we can multiply each of the hexadecimal digits by the appropriate power of 16.


#### 2.1.2 data sizes
Every computer has a word size, indicating the nominal size of pointer data. Since a virtual address is encoded by such a word, the most important system parameter determined by the word size is **the maximum size of the virtual address space**. That is, for a machine with a w-bit word size, the virtual addresses can range from 0 to 2^w - 1, giving the program access to at most 2^w bytes.

> linux > gcc -m32 program.c

then the program will run correctly either on 32-bit or 64-bit.

    char, 1 byte
    short, 2 bytes
    int, 4
    long, 4 bytes(32-bit), 8 bytes(64-bit)
    float, 4
    double,8
    int 32_t, 4
    int64_t, 8
    char *, 4 bytes(32-bit), 8 bytes(64-bit)

To avoid the vagaries of relying on "typical" sizes and diffierent compiler settings,
ISO C99 introduced a class of data types where the data sizes are fixed
regardless of compiler and machine settings. Among these are data types *int32_ t*
and *int64_ t*, having exactly 4 and 8 bytes, respectively. 

**Using fixed-size integer
types is the best way for programmers to have close control over data representations.**

A pointer uses the full word size of the program.

#### 2.1.3 addressing and byte ordering
For program objs that span multiple bytes, we must establish two conventions:

 1. what the address of the obj will be?
 2. how we will order the bytes in memory?

In virtually all machines, a multi-byte obj is stored as a contiguous sequence of bytes. The address of the obj is given by the smallest address of the bytes used.

**Little endian VS big endian**:

>Assume a w-bit integer has a bit representation [x(w-1), x(w-2),...,x(1), x(0)], where x(w-1) is the most signigicant bit and x(0) is the least.
>Assume w is a multiple of 8, those bits can be grouped as bytes, with the most significant byte having bits [x(w-1), x(w-2),...,x(w-8)], the lease significant byte having bits [x(7),x(6),...x(0)]

Some machines choose to store the obj in memory ordered from least significant byte to most(the least significant byte comes first, **little endian**); from most significant byte to least(the most bytes comes first, **big endian**)

>most Intel-compatible : little-endian
>
>most IBM, Oracle : big-endian

When byte order becomes import:

 1. send binary data over network
 2. looking at the byte sequences representing integer data
 3. when programs are written that circumvent 包围陷害绕行 the normal type system

#### 2.1.4 representing strings

A string in C is encoded by an array of characters terminated by the null (having
value O) character.

#### 2.1.5 representing code
A fundamental concept of computer systems is that a program , from the perspective of the machine, is simply a sequence of bytes.

#### 2.1.6 introduction to boolean algebra
George Boole

    NOT logical : ¬， Boolean : ~
    AND          ^ &
    OR           ∨ |
    EXCLUSIVE-OR ⊕ ^ 

extend Boolean operations to operate on bit vectors, strings of zeros and ones of some fixed length *w*

Let a and b denote the bit vectors [a(w-1),a(w-2),..., a(0)] and [b(w-1),b(w-2),..., b(0)] respectively. 
We define a & b to also be a bit vector of length w, where the ith element equals a(i) & b(i), for 0 <= i < w. 
The operations |,~, and ^ are extended to bit vectors in a similar fashion.

**Web Aside data : bool**

Unique to integer rings:

    a + -a = 0  

Unique to boolean algebras:

    a | (b & c) = (a | b) & (a | c)
    complement 余角: a | ~a = 1, a & ~a = 0
    idempotency 幂等性 : a & a = a, a | a = a
    absorption : a | (a & b) = a, a & (a | b) = a
    DeMorgan's laws: ~(a&b) = ~a | ~b, ~(a|b) = ~a & ~b

One useful app of bit vectors is to represent finite sets. We can encode any subset A 包含于 {0, 1,...,w-1} with a bit vector [a(w-1),...,a(1),a(0)] where a(i)=1 if and only if i ∈ A, 
    
    bit vector a=[01101001] encodes A={0,3,5,6}

#### 2.1.7 bit-level operations in C
The best way to determine the effect of a bit-level expression is :
> expand the hexadecimal arguments to their binary representations.

#### 2.1.8 Logical ops in C
1. The logical operations treat
any nonzero argument as representing TRUE and argument 0 as representing FALSE
2. the logical operators do not evaluate their second argument if t,he result of the expression can be determined by evaluating the first argument. Thus, for example, the, expression a && 5/a will never cause a division by zero, and the exNession p && *p++ will ,never cause the dereferencing of a null pointer.

#### 2.1.9 shift ops in C
*Left shift*: That is, x is shifted k bits to the left, dropping off the k most significant
bits and filling the right end with k zeros

*right shift* two forms:

 1. Logical, fills the left end with k zeros, [0,...,0,x(w-1),x(w-2),...,x(k)]
 2. Arithmetic, fills the left end with k repetitions of the most siginificant bit [x(w-1),...,x(w-1),x(w-1),x(w-2),...,x(k)]


**almost all compiler/machine combinations use arithmetic right shifts for signed data. For unsigned data, right shifts must be logical.**

The c standards carefully avoid stating what should be done in such a case. On many machines, the shift instructions consider only the lower log2(w) bits of the shift amount when shifting a w bit value, and so 
the shift amount is computed as k mod w(let w=32, log2(w) = 5, k = 100, k & 00011111 = k mod 32 = k mod w).

The behaviour is not guaranteed for C programs however, and so shift amouts should be kept less than the word size.

### 2.2 Integer representations
func : B2U binary to unsigned 
#### 2.2.1 integral data types
The only machine-dependent range indicated is for size designator *long*. 64-bit 8byte, 32-bit 4 byte.

The range of negative numbers extends one further than the range of positive numbers.

The C standards define minimum ranges of values that each data type must be able to represent. (guaranteed ranges for C integral data types). int could be 2-byte, long 4-byte ([this is the standard link](http://www.open-std.org/jtc1/sc22/wg14/www/docs/n1256.pdf))

#### 2.2.2 unsigned encodings
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/b2u.png "b2u")

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/2/unsigned_range.png "unsigned range")

The unsigned binary representation has the important property that every number between 0 and 2^w-1 has a unique encoding as a w-bit value.

**Principle**:
Uniqueness of unsigned encoding, Function B2U(w) is a bijection.

The mathematical term bijection refers to a function f that goes two ways: it maps a value x to a value y where y = f(x), but it can also operate in reverse, since for every y, there is a unique value x such that f(x) = y. This is given by the inverse function f-1, where, for our example, x = f-1(y).

B2U(w), U2B(w)

#### 2.2.3 two's-complement encodings(二补数)
two's-complement is defined by interpreting the most significant bit of the word to have negative weight.B2T(w)

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/2/two-complement.png "two's-complement")

![alt-text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/2/b2t.png "b2t examples")

The least representable value is given by bit vector [10...0], TMin(w) = -2^(w-1); the max value is given by bit vector [01...1], TMax(w) = 求和从0->(w-2)2^i = 2^(w-1) - 1. 假设w=4, TMin(4)=-8, TMax(4)=7

B2T(w) is a mapping of bit patterns of length w to numbers between TMin(w) and TMax(w):
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/2/b2t_map.png "b2t map")

**Priciple**: Uniqueness of two's-complement encoding, function B2T(w) is a bijection. 

Define T2B(w) (two's-complement --> binary) to be the inverse of B2T(w).
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/2/important_nums.png "important numbers")

*UMax(w), TMax(w), TMin(w)*

We will drop the subscript w and refer to the values UMax, TMin and TMax when w can be inferred from context or is not central to the discussion.

|TMin| = |TMax| + 1