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
