#chapter 2 practices
## 2.1 , 2.2, 2.3, 2.4 too easy

## 2.5
A. 21, 87
B. 21 43, 87 65
C. 21 43 65, 87 65 43

## 2.6
A. python -c "print(\"{0:020b}\".format(int(\"0x00359141\", 16)))"
python -c "print(\"{0:020b}\".format(int(\"0x4a564504\", 16)))"
       1101011001000101000001
1001010010101100100010100000100
How to find the max number of matching bits?which algorithm?

## 2.7
61 62 63 64 65 66

## 2.9
Black <-> White
Green <-> Magenta
Blue <-> Yellow
Cyan <-> Red

## 2.10
    void swap(int *x, int *y){
      *y = *x ^ *y;
      *x = *x ^ *y;
      *y = *x ^ *y;      
    }

Step　     *x                    *y
initially  a                     b
step1      a                     a ^ b
step2      a^(a^b) = 0^b = b     a ^ b
step3      b                     b^(a^b)=0^a=a

## 2.11
    first <= last -->
    first < last

## 2.12
x & 0xFF
~(x & 0xFFFFFF00) | (x & 0xFF)
(x & 0xFFFFFF00) | 0xFF

## 2.13(p56)
bis(x, m) = x | m
bic(x, m) = x & ~m
bool_or : result = bis(x, y)
bool_xor : x ⊕ y = (x & ~y) | (~x & y) = bic(x,y) | bic(y, x) = bis(bic(x,y), bic(u,x))

## 2.14 skipped
## 2.15
!(x ^ y)
## 2.16, 2.17 too easy
## 2.18 see it easy, -0x58 = -88
## 2.21
unsigned 0
int 1
unsigned 0 
int 1
unsigned 1

## 2.23
    int fun1(unsigned word){
        return (int)( (word<<24)>>24 );
    }
    int fun2(unsigned word){
        return ((int)(word<<24)) >> 24;
    }

    0x00000076 0x76=118 0x76=118
    0x87654321 0x21=33 0x21=33
    0x000000C9 0xC9=201 0xc9-2^8=-55
    0xEDCBA987 0X87=135 0X87-2^8=-121

## 2.24(p82)
unsigned (4->3) 
0 0
2 2
9 9%8=1
11 11%8=3
15 15%8=7

two's-complement(4->3)
0 0
2 2
-7 U2T((-7 + 2^4)%2^3)=U2T(1)=1
-5 U2T((-5+2^4)%2^3)=U2T(3)=3
-1 U2T(15%8)=U2T(3)(7)=7-8=-1

## 2.25
unsigned length
(unsigned)0-1=UMax, memory out of range
for (i=0; i < length; ++i)

##2.26
s length < t length
(unsigned)2-(unsigned)3 > 0

    int strlonger(char *s, char *t){
        return strlen(s) - strlen(t) > 0;
    }

##2.27
    int uadd_ok(unsigned x, unsigned y)
    {
      if (x+y<x){return 0;}
      return 1;
    }

## 2.28(p89)
0 0 0
5 11 B
8 8 8
13 3 3
15 1 1

##2.29 easy
##2.30
    int tadd_ok(int x, int y)
    {
      if(x<0 && y < 0 && (x+y) >=0)
      {return 0;}
      else if (x>0 && y >0 && (x+y) <= 0)
      {return 0;}
      else
      {return 1;}
    }

##2.31
proof(negative overflow):

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/2/practise213.png "detection")

So, overflow will still return 1

##2.32
x=0, y=-2^(w-1)

##2.33
0 0 0 0
5 5 -5 1011=B
8 -8 -8 8
D=1101 -3 3 3
F=1111 -1 1 1

##2.34,easy
##2.35
devise 设计想出发明 a mathematical justification of your approach, consider w-bit number x(x!=0),y,p,q, p is the result of performing two's-complement multiplication on x and y, q is the result of dividing p by x

1. show that x*y(the integer product of x and y) can be written in the form xy=p+t * 2^w, where t!=0 if and only if the computation of p overflows.
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/2/practise235.png "practise")
2. show that p can be written in the form p = x*q+r, where |r| < |x|
this is by the definition of integer division.
3. show that q=y if and only if r=t=0
x*y-t*2^w=x*q+r
suppose q=y, r = -t*2^w, |r| < 2^(w-1),t=0
suppose r=t=0, x*y=x*q, y=q

##2.36
    int tmul_ok(int x, int y)
    {
      int64_t result = (int64_t)x*y;
      return result == (int)result;
    }

##2.37(p101)
int and size_t are 32-bits.
A.no, the uint64_t asize passed to malloc() may still overflow.
B.if (asize == (int)asize){//go on}
else{return "memory not enough"}

##2.38
k=0,1,2,3; b=0 or a
2,4,8,3,5,9

##2.39
form a : (x<<n) + (x<<(n-1)) + .. + (x <<m)
form b : (x<<(n+1)) - (x << m)
if n is the most significant bit 
form b : (x<<n) - (x << m) + (x<<n)

##2.40
(x << 2) + (x << 1)
(x << 5) - x
(x << 1) - (x << 3)
(X << 6) - (X << 3) - X
##2.41(p103)
the rule is to choose from 
A when n=m;(form a:1 instruction;form b:2 insrtuction)
either form when n=m+1(2 instrction both)
B when n > m + 1(form b:2 instructions; form a more than 2 instruction)
 
##2.42
x > 0, bias=0;
x < 0, bias=(1 << 4)-1=15
int div16(int x)//return x/16
{
  int bias = (x >> 31) & 0xF;
  return (x + bias) >> 4;
}

##2.43
M=31, N=8

##2.44(p108)
-2^(w-1) <= x < 2^(w-1)

A. (x>0) || (x-1<0), false, x = -2^(w-1) =TMin
B. (x&7) !=7 || (x<<29<0), true
C. (x*x) >= 0, false, x = 65535=(2^16-1) **Question,how to find this number**
D. x < 0 || -x <= 0, true
E. x > 0 || -x >= 0, fasle, x = -2^(w-1)
F. x + y == uy + ux, true
G: x*~y + uy * ux == -x,  true, ~y=-y-1, x*-y-x-uy*ux = -x

##2.45 too easy
##2.46 binary approximation
Did not understand yet

##2.47(p117)
e E     2^E f    M    2^EM V  Decimal
0 1-1=0 1   0    0    0    0   0
0 0     1   1/4  1/4  1/4  1/4 1/4
0 0     1   1/2  1/2  1/2  1/2 1/2
0 0     1   3/4  3/4  3/4  3/4 3/4
1 e-1=0 1   0    1+f=1 1   1   1
1 0     1   1/4  5/4   5/4 5/4 1.25
1 0     1   1/2  3/2   3/2 3/2 1.5
1 0     1   3/4  7/4   7/4 7/4 1.75
2 e-1=1 2   0    1     2   2   2
2 1     2   1/4  5/4   5/2 5/2 2.5
2 1     2   1/2  3/2   3   3   3
2 1     2   3/4  7/4   7/2 7/2 3.5
infinity
NaN
NaN
NaN

##2.48 same as 12345
##2.49(p120)
A.2^(n+1)+1
B.2^24+1



