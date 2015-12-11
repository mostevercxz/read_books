#Part I, Program Structure and Execution
1. void swap(int x, int y) {y = x ^ y; x = x ^ y; y = x ^ y;}; swap function is not suitable for exchange array members. 异或版本的swap不适合在转置奇数数组中使用
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

The behaviour is not guaranteed for C programs however, and so shift amouts should be kept less than the number of bits in the value being shifted.

### 2.2 Integer representations
#### 2.2.1 integral data types
The only machine-dependent range indicated is for size designator *long*. 64-bit 8byte, 32-bit 4 byte.

The range of negative numbers extends one further than the range of positive numbers.

The C standards define minimum ranges of values that each data type must be able to represent. (guaranteed ranges for C integral data types). int could be 2-byte, long 4-byte ([this is the standard link](http://www.open-std.org/jtc1/sc22/wg14/www/docs/n1256.pdf))

#### 2.2.2 unsigned encodings
1. **Principle : Definition of unsigned encoding,**
for vector x = [x(w-1),x(w-2),...x(0)]
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/2/formula_21.png "b2u")

2. **Range of values:**
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/2/unsigned_range.png "unsigned range")
The unsigned binary representation has the important property that **every number between 0 and 2^w-1 has a unique encoding as a w-bit value.**

3. **Principle:Uniqueness of unsigned encoding,**
Function B2U(w) is a bijection.


The mathematical term **bijection** refers to a function f that goes two ways: it maps a value x to a value y where y = f(x), but it can also operate in reverse, since for every y, there is a unique value x such that f(x) = y. This is given by the inverse function f-1, where, for our example, x = f-1(y).
B2U(w), U2B(w)

#### 2.2.3 two's-complement encodings(二补数)
two's-complement is defined by interpreting the most significant bit of the word to have negative weight.
**B2T(w)**

1. **Principle : Definition of two's-complement encoding**, for x = [x(w-1), x(w-2),...x(0)]
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/2/formula23.png "two's-complement")
![alt-text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/2/b2t.png "b2t examples")

2. **range of values, B2T(w)** is a mapping of bit patterns of length w to numbers between TMin(w) and TMax(w): **The least representable value is given by bit vector [10...0], TMin(w) = -2^(w-1); the max value is given by bit vector [01...1], TMax(w) = 求和从0->(w-2)2 ^ i = 2^(w-1) - 1.** 假设w=4, TMin(4)=-8, TMax(4)=7
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/2/b2t_map.png "b2t map")

**3. Priciple : uniqueness of two's-complement encoding**: function B2T(w) is a bijection. 

Define T2B(w) (two's-complement --> binary) to be the inverse of B2T(w).
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/2/important_nums.png "important numbers")

*UMax(w), TMax(w), TMin(w)*

We will drop the subscript w and refer to the values UMax, TMin and TMax when w can be inferred from context or is not central to the discussion.

**4. Observations:**

 1. |TMin| = |TMax| + 1
 2. UMax = 2TMax + 1 

fixed size integer types(int32_t,int64_t):

    #include <inttypes.h>
    printf("x=%" PRId32 ",y=%" PRIu64 "\n", x, y)

    #define PRId32 "ld"
    #define PRIu64 "llu"

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/2/guaranteed_ranges.png "guaranteed ranges")

Programmers who are concerned with maximizing portability across all possible machines should not assume any particular range of representable values, beyond the ranges indicated in Figure 2.11, nor should they assume any particular representation of signed numbers.

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/2/two_other_standard_representation.png "other representations")

Both of the representations have the curious property that there are two different encodings of number 0. +0 = [00...0],
 
 * sign-magnitude, -0 = [10...0]
 * one's-complement, -0 = [11...1]

Note the different position of apostrophes two's complement versus ones' complement. 省略符号,呼语

**Questions,why?**

1. Two's complement : for nonnegative x we compute a w-bit representation of -x as 2^w-x(a signle two);
2. One's complement : -x as [111...1] - x(multiple ones)

Two's complement Explain:

    -t = -2^(w-1) + 2^(w-1) - t , 
    most siginificant bit is 1, the total value of other bits is 2 ^ (w-1) - t,
    so the total unsigned value is:
    most siginificant bit value 2^(w-1) + other bits value 2 ^(w-1) -t = 2 ^ w - t
    -t = x^w - t


One's complement explain:

    -t = -(2 ^ (w-1) - 1) + 2 ^(w-1) -1 -t
    most siginificant bit is 1, total value of other bits is 2^(w-1) -1-t
    so the total unsigned value is:
    2 ^ (w-1) + 2 ^(w-1) -1-t = 2^w-1-t = [111...1] - t
    -t = [111...1] - t

#### 2.2.4 conversions between signed and unsigned

**1. From a mathematical perspective:**

 * we want to preserve any value that can be represented in both forms.
 * converting a negative value to unsigned might yield
zero
 * Converting an unsigned value that is too large to be represented in two'scomplement form might yield TMax

**2. For most implementations of C**, based on a bit-level perspective(conversions between unsigned and signed numbers with the same word size):

  * casting from *short* to *unsigned short* changed the numeric value, but not the bit representation.


    0 <= x <= UMax(w), U2B(w)(x) unique
    Tmin(w) <= x <= TMax(w), T2B(w)(x) unique

>signed and unsigned are **the same word size**
>
>T2U(w)(x) = B2U(w)(T2B(w)(x)), Tmin(w) <= x <= TMax(w)
>U2T(w)(x) = B2T(w)(U2T(w)(x)), 0 <= x <= UMax(w) 

**3. Principle : Conversion from two's complement --> unsigned**:

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/2/formula25.png  "guaranteed  ranges")

**4. derivation : Conversion from two's complement --> unsigned**:
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/2/proof_t2u.png  "proof")


Going in the other direction, we can get:

**5.Principle : Conversion from unsigned --> two's complement**:

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/2/proof_u2t.png  "proof u2t")

**6.derivation:Conversion from unsigned --> two's complement**:

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/2/proof_u2t_derivition.png  "proof_u2t_derivition.png")

**7.summary**:

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/2/summary224.png  "proof u2t")

#### 2.2.5 signed versus unsigned in C
**1.explicit, implicit, printf conversion**
Although the C standard does not specify precisely how this conversion should be made, most systems follow the rule that the underlying bit representation does not change.

*printf* does not make use of any type info.
printf("-1=%u", -1);
-1=4294967295

**2.operation conversion**
When an operation is performed where one operand is signed and the other is unsigned, C implicitly casts the signed argument to unsigned and performs the operations assuming the numbers are nonnegative.
(gdb)p -1u

#### 2.2.6 expanding the bit representation of a number
Convert from a smaller to a larger data type:

 1. unsigned number --> larger data type, add leading zeros to the representation(zero extension)
 2. two's-complement --> larger, sign extension, adding copies of the most significant bit to the representation.

**1.principle:expansion of an unsigned number by zero extension**:
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/2/zero_extension.png  "zero")

**2.principle:expansion of a two's-complement number by sign extension**:
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/2/sign_extension.png  "sign")

**3.derivation:expansion of a two's-complement by sign extension**:
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/2/proof_sign_extension.png "proof sign extension")

#### 2.2.7 Truncating numbers
**1.method**:
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/2/truncate_method.png "truncate method")
Drop the high-order w-k bits.

**2.principle:truncation of unsigned number**:
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/2/truncate_unsigned.png "unsigned")

**3.derivation:truncation of an unsigned number**:
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/2/truncate_unsigned_proof.png "truncate_unsigned_proof")

**4.principle:truncation of atwo's-complement number:**
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/2/truncate_two.png "truncate two")

**5.derivation:truncation of two's-complement number**:
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/2/truncate_two_proof.png "truncate two proof")

#### 2.2.8 advice on signed versus unsigned
**Bug code**

    void fun(float a[], unsigned length){
      int result = 0;
      for (int i = 0; i <= length-1; ++i)
      {
        result += a[i];
      }
    }

    void strlonger(char *s, char *t){
      return strlen(s) - strlen(t) > 0;
    }

### 2.3 Integer arithmetic
1. two positive numbers can yield a negative result
2. x < y **NOT EQUAL** x - y < 0

#### 2.3.0 Mathmetical knowledge
1. **Modular arithmetic**(同余) [wikipedia Modular_arithmetic](https://en.wikipedia.org/wiki/Modular_arithmetic) is a system of arithmetic for integers, where numbers "wrap around" upon reaching a certain value-the modulus. A familiar use of modular arithmetic is in the 12-hour clock, in which the day is divided into two 12-hour periods. If the time is 7:00 now, then 8 hours later it will be 3:00. Usual addition would suggest that the later time should be 7+8=15, but this is not the answer because clock time "wraps around" every 12 hours; in 12-hour time, there is no "15 o'clock".
2. **Abelian group**(阿贝尔群) [wikipedia Abelian group](https://en.wikipedia.org/wiki/Abelian_group). In abstract algebra, an abelian group,also called a commutative group. is a group in which the result of applying the group operation to two groups elements that does not depend the order in which they are written. Definition: An abelian group is a set, A, together with an operation • that combines any two elements a and b to form another element denoted a • b. The symbol • is a general placeholder for a concretely given operation. To qualify as an abelian group, the set and operation, (A, •), must satisfy five requirements known as the abelian group axioms:
 1. Closure : For all a, b in A, the result of the operation a • b is also in A.
 2. Associativity : For all a, b and c in A, the equation (a • b) • c = a • (b • c) holds.
 3. Identity element : There exists an element e in A, such that for all elements a in A, the equation e • a = a • e = a holds.
 4. Inverse element : For each a in A, there exists an element b in A such that a • b = b • a = e, where e is the identity element.
 5. Commutativity : For all a, b in A, a • b = b • a.
3. space

#### 2.3.1 unsigned addition
0 <= x,y < 2^w, w-bit
0 <= x+y <= 2^(w+1)-2,need (w+1)-bit

**1. Define the operation +uw for arguments** x,y,0 <= x,y < 2^w, as the result of truncating the integer sum *x+y* to be w-bits long and then viewing the result as an unsigned number. 

Security vulnerablity(maxlen=-1):

    void *memcpy(void *dest, void *src, size_t n);
    int copy_from_kernel(void *user_dest, int maxlen)
    {
      memcpy(user_dest, buf, maxlen);
    }

**2.principle:unsigned addition**:
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/2/unsigned_addition.png "unsigned addition")

**3.derivation:unsigned addition**:

x+y<2^w, leading bit in the (w+1)-bit representation is 0; 2^w<=(x+y)<2^(w+1),leading bit in the (w+1)-bit representation is 1, discarding leading bit is equalent to subtracting 2^w from the sum.

**4.definition of *overflow***:

An arithmetic operation is said to *overflow* when the full integer result cannot fit within the word size limits of the data type.

executing C programs, overflows are not signed as errors.

**5.principle:detecting overflow og unsigned addition**:
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/2/addition_overflow.png "addition overflow")

**6.derivation:detecting overflow of unsigned addition**:

    if s is overflow, s >= 2^w, s = x+y-2^w, y<2^w
    s = x + (y-2^w) < x

**7.principle:unsigned negation**:
**8.derivation:unsigned negation**:

Modular addition forms abelian group. It is commutative, associative, identify-element 0, every element has an additive inverse. This additive inverse operation can be characterized as follows:

For any number 0<=x<=2^w, its unsigned negation is:
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/2/unsigned_negation.png "unsigned negation")

(x+2^w-x) mode 2^w = 2^w mod 2^w = 0, hence 2^w-x is the inverse of x.

#### 2.3.2 two's-complement addition

#### 2.3.3 two's-complement negation
#### 2.3.4 unsigned multiplication
#### 2.3.5 two's-complement multiplication
#### 2.3.6 multiplying by constants
#### 2.3.7 dividing by powers of 2
#### 2.3.8 final thoughts on integer arithmetic

### 2.4 floating point
#### 2.4.1 fractional binary numbers
#### 2.4.2 IEEE floating-point representation
#### 2.4.3 example numbers
#### 2.4.4 rounding
#### 2.4.5 floating-point operations
#### 2.4.6 floating point in C

### 2.5 Summary
