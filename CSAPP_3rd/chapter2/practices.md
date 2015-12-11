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
