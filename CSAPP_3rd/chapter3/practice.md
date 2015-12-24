##3.1
%rax         Ox100 Register
Ox104        OxAB Absolute address
$0x108       Ox108 Immediate
(%rax)       OxFF Address Ox100
4(%rax)      OxAB Address Ox104
9(%rax,%rdx), Ox11 Address Ox10C
260(%rcx,%rdx) Ox13 Address Ox108
OxFC(, %rcx, 4) OxFF Address Ox100
(%rax, %rdx, 4) Ox11 Address OxlOC

##3.2
    movl %eax, (%rsp)
    movw (%rax),%dx
    movb $0xFF,%bl
    movb (%rsp,%rdx,4),%dl
    movq (%rdx), %rax
    movw %dx,(%rax)

##3.3
    movb $OxF, (%ebx)  %ebx,32-bit can not used as address reference
    movl %rax, (%rsp) movl is 32-bit,%rax iss 64-bit
    movw (%rax), 4(%rsp) can not both be memory
    movb %al, %sl no register named %sl
    movq %rax, $0x123 immediate can not be used as destination
    movl %eax,%rdx %rdx is 64-bit
    movb %si,8(%rbp) si is 16-bit

##3.4
    src_t *sp;
    dest_t *dp;
    *dp = (dest_t)*sp;

Assume the values of sp and dp are stored in register %rdi and %rsi, respectively.

    src_t dest_t
    long long movq (%rdi),%rax ; movq %rax, (%rsi)
    char int  movsbl (%rdi),%eax ; movl %eax, (%rsi)

    char -> int -> unsigned
    char unsigned movsbl (%rdi),%eax ; movl %eax, (%rsi)

    unsigned char long movzbl (%rdi), %eax ; movq %rax, (%rsi)

    int -> char, read 4 bytes , then store low-order byte
    int char movl (%rdi), %eax ; movb %al, (%rsi)
    unsigned unsigned char movl (%rdi), %eax ; movb %al, (%rsi)
    char short movsbw (%rsi), %ax ; movw %ax, (%rsi)

##3.5
    long x = *xp;
    long y = *yp;
    long z = *zp;
    *yp = x;
    *zp = y;
    *xp = z;

##3.6
%rax holds x, %rcx holds y
6x
x+y
x+4y
x+8x+7
4y + 0xA
x+2y+9

##3.7
5x + 2y + 8z

##3.8
0x100 0x100
0x108 0xBE
0x118 0x21
0x110 0x14
register 0x0
register 0xFD

##3.9
salq $4, %rax
sarq %cl, %rax

##3.10
long t1= x|y;
long t2 = t1 >> 3;
long t3 = ~t2;
long t4 = z - t3;
return t4;

##3.11
A. x = 0
B. movq $0, %rdx
C. xorq requires only 3 bytes; movq requires 7 bytes; xorl %edx, %edx(2 bytes); movl $0, %edx(5 bytes)

##3.12
void uremdiv(ulong x, ulong y, ulong *qp, ulong *rp)
movq %rdx, %r8
movq %rdi, %rax
mov1 $0, %edx
divq %rsi
movq %rax, (%r8)
movq %rdx, (%rcx)

##3.13
A. int <=
B. short >=
C. unsigned char <=
D. long, unsiged long, pointer !=

##3.14
A. long >=
B. short, unsigned short ==
C. unsigned char, >
D. int , <=

## 3.15
A. 4003fe
B. 400425
C. 400543 400545
D. 400560

## 3.16
    if(!p)
      goto direct_ret;
    if (*p >= a)
      goto direct_ret;
    *p = a;
    direct_ret:
      return ;

a && b, needs to calculate both a and b.

##3.17
    long gotodiff_se(long x, long y)
    {
      long result;
      if (x < y)
         goto x_l_y;
      ge_cnt++;
      result = x - y;
      return result;
    x_l_y:
      lt_cnt++;
      result = y - x;
      return result;
    }

Can you think of any reasons for choosing one rule over the other?
But the origin rule works better for the common case where there is no else statement.

##3.18
    long val = x+y+z;
    if (x<-3){
      if (y<z)
        val = x * y;
      else
        val = y * z;
    } else if(x>2)
      val = x * z;
    return val;

##3.19
too easy

##3.20
operator is "/", two's-complement division by power of 2 rounding up.

##3.21
Question : to be done

##3.22
test each n until n! overflows.
n! <= 0x7fffffff
n=1,2,3,...13

##3.23
x=%rax
y=%rcx
n=%rdx

x+=y;
(*p)++;
x=x+y+1

##3.24
    long result = 1;
    while(a < b){
      result = result * (a+b);
      a = a + 1;
    }
    return result;

##3.25
long result = b;
while (b > 0){
result = result * a;
b = b -a;
}

##3.26
jump to middle

long val = 0;
while (x != 0){
  val = val ^ x;
  x = x >> 1;
}
return val & 0x1;
x = [x(w-1), x(w-2), ..., x(0)]
x = x(0) ^ x(1) ^ x(2) ... ^ x(w-1)
返回0表明x的二进制位中1的个数为偶数;1为奇数

##3.27
    long fact_for(long n)
    {
      long i;
      long result = 1;
      for (i = 2; i <= n; i++){
        result *= i;
      }
      return result;
    }

    long fact_for_gd_goto(long n)
    {
      long i = 2;
      long result = 1;
      if (i > n)
        goto done;
    loop:
      result *= i;
      i++;
      if (i<=n)
        goto loop;
    done:
        return result;
    }

    movl $1, %eax
    movl $2, %edx
    cmpq %rax, %rdx
    jg .L1
    .L2:
      imulq %rdx, %rax
      addq $1, %rdx
      cmpq %rax, %rdx
      jle .L2
    .L1:
      rep;ret

##3.28
long fun_b(unsigned long x){
  long val = 0;
  long i;
for(i=64; i != 0; i--){
val = (val << 1) | (x & 0x1) 
x = x >> 1;
}
return val;
}

B. guarded-do transformation, i = 64, i !=0, so the intial test is not required.
C. The code reverse the bits in x, creating a mirror image.

##3.29
    long sum = 0;
    long i;
    for (i = 0; i < 10; i++){
      if (i & 1)
        continue;
      sum += i;
    }

    long sum = 0;
    long i = 0;
    while(i<10){
      if (i&1)
        continue;
      sum += i;
      i++;
    }

A. infinite loop
B. 

      long sum = 0;
      long i = 0;
      while(i<10){
        if (i&1)
          goto update_expr;
        sum += i;
    update_expr:
        i++;
      }

##3.30
.quad starts from 0,8,16...
(x+1)=0,1,2,3,4,5,6,7,8
L2 is the default, so x=-1,0,1,2,4,5,7

    L5,L7
    case 0/case 7
    case 2/case 4
##3.31
a=0,1,2,3,4,5,6,7
.L2 is the default case, so
a=0,2,4,5,7

.L5 repeats, so case 2/case 7
.L7 changes c, so c = b ^ 15;

    void switcher(long a, long b, long c, long *dest)
    {
    	long val;
    	switch (a)
    	{
    	case 5:
    		c = b ^ 15;
    	case 0:
    		val = c + 112;
    		break;
    	case 2:
    	case 7:
    		val = (b + c) << 2;
    		break;
    	case 4:
    		val = a;
    	default:
    		val = b;
    	}
    }

## 3.32
    0x400548 lea  10 -  -  0x7fffffffe818 0x400565 x+1
    0x40054c sub  10 11 -  0x7fffffffe818 0x400565 x-1
    0x400550 callq 9 11 -  0x7fffffffe818 0x400565 callq
    
    0x400540 mov   9  11 -  0x7fffffffe810 0x400555 mov
    0x400543 imul  9  11 99 0x7fffffffe810 0x400555 imul
    0x400547 retq  9  11 99 0x7fffffffe810 0x400555 retq
    
    0x400555 repz  9  11 99 0x7fffffffe818 0x400565 return 
    0x400565 mov   9  11 99 0x7fffffffe820 -        resume

## 3.33
    *u += a;
    *v += b;
    return sizeof(a) + sizeof(b);


    movslq %edi, %rdi
    addq   %rdi, (%rdx)
    addb   %sil, (%rcx)
    movl   $6, %eax
    ret

sizeof(a) + sizeof(b) = 6;//short,int
    size_t procprob(int a, short b, long *u, char *v);
    size_t procprob(int b, short a, long *v, char *u);

## 3.34
a function P, generate local values, named a0-a8
x, x+1,...x+5 are stored in the callee-saved
x+6, x+7  are stored on the stack

C. After storing six local variables, the program has used up the supply of callee-saved registers. It stores the remaining two local values on the stack.

## 3.35
A. store x in the callee-saved register

B.
    long rfun(unsigned long x){
      if(0 == x)
        return 0;
      unsigned long nx = x >> 2;
      long rv = rfun(nx);
      return x + rv;
    }

## 3.36
    S 2 14 x(s) x(s)+2i
    T 8 24 x(t) x(t)+8i
    U 8 48 x(u) x(u)+8i
    V 4 32 x(v) x(v)+4i
    W 8 32 x(w) x(w)+8i


##3.37
short S[N];
%rdx, x(s)
%rcx, index

    S+1     short*  x(s)+2           leaq $2(%rdx), %rax
    S[3]    short   M[x(s)+2*3]      movw $6(%rdx), %ax
    &S[i]   short*  x(s)+2i          leaq (%rdx, %rcx,2),%rax
    S[4*i+1]short   M[x(s)+2*(4i+1)] movw 2(%rdx, %rcx,8),%eax
    S+i-5   short*  x(s)+2*(i-5)     movw -10(%rdx, %rcx,2),%rax

##3.38
long P[M][N]; P[i][j] <--> x(P) + (i * N + j)

7i+j, N=7
5j+i, M=5

## 3.39
&A[i][0] = x(A) + sizeof(int) * (i*N+0) = x(A) + 64i
&B[0][k] = x(B) + 4(0 + k) = x(B)+4k
&B[N][k] = x(B) + 4(N*N+k) = x(B)+1024+4k

## 3.40
void fix_set_diag_opt(fix_matrix A, int val)
{
int *start = &A[0][0];
long offset = 0;
long end = N*(N+1);
do
{
  *(A+offset) = val;
  offset += 4*(N+1);
}while(offset != end)
}

## 3.41
A.
p : 0
s.x : 8
x.y : 12
next : 16

B.
24

C.
sp->s.x = sp->s.y;
sp->p = &(sp->s.x);
sp->next = sp;

##3.42
    long fun(struct ELE *ptr)
    {
      long ret = 0;
      while(ptr){
        ret += ptr->v;
        ptr = ptr->p;
      }
    }

Linked List.

##3.43
    up->t1.v short movw 8(%rdi), %ax; movl %eax, (%rsi)

    mine is "leaq 10(%rdi), %rax; movq %rax, (%rsi)"  
    &up->t1.w char* addq $10, %rdi; movq %rdi, (%rsi)

    mine is "int* leaq (%rdi), %rax; movq %rax, (%rsi)"
    up->t2.a int *  movq %rdi, (%rsi)

up->t2.a[up->t1.u] int 
movq (%rdi), %rax
movl (%rdi, %rax, 4), %eax
movl %eax, (%rsi)  

 
*up->t2.p
precedence : 
++ --
type() type{}
()
[]
. ->
*
&

char
movq 8(%rdi), %rax------mine is leaq 8(%rdi), %rax, totally wrong
movb (%rax), %al
movb %al, (%rsi)

## 3.44
16, 4
16, 8
10, 2
40, 8
40, 8, just use the rule that any primitive object of K bytes must have an address that is a multiple of K.

##3.45
56 bytes

double c, long g, char * a
float e, int h
short b, char d, char f

char d;
char f;
short b;
int h;
float e;
double c;
long g;
char *a;
40 bytes total

##3.46
E. should 
    malloc(strlen(buf)+1);
    check whether the return value of malloc is NULL

##3.47
2^13
2^13/2^7=2^6=64

##3.48
local variable v is closer to the top of the stack than buf. so corrupted stack will not affect v.

##3.49
Finally, I understand the question.
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/3/solution349.png "solution349")

##3.50
val1 = d
val2 = i
val3 = l
val4 = f

##3.51
dest_t cvt(src_t x)
{
  dest_t y = (dest_t)x;
  return y;
}

x is either in %xmm0 or in %rdi,%edi
one or more instructions are to be used to perform the type conversion and to copy the value to the %rax, %xmm0

T(x)   T(y) instructions
long   double vcvtsi2sdq %rdi, %xmm0
double int    vcvtsd2si %xmm0, %eax
double float  vmovddup %xmm0, %xmm0; vcvtpd2psx %xmm0, %xmm0
long   float  vcvtsi2ssq %rdi, %xmm0, %xmm0
float  long   vcvtss2siq %xmm0, %rax

##3.52
a in %xmm0, b in %rdi, c in %xmm1, d in %esi
a in %edi, b in %rsi, c in %rdx, d in %rcx
a in %rdi, b in %xmm0, c in %esi, d in %xmm1
a in %xmm0, b in %rdi, c in %xmm1, d in %xmm2

##3.53
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/3/problem353.png "problem353")
%xmm0 is float
%rsi is long
%edi is int
%xmm1 is double

%rsi -> %xmm2, long->float
%xmm0 = %xmm0 + %xmm2, float + float
%edi -> %xmm2, int -> float
%xmm0 = %xmm2 / %xmm0, float / float
%xmm0 -> double, float->double
%xmm0 = %xmm0 - %xmm1, double - double

p/(q+r) is float
s is double
line2-line5,  q and r(long, float), p is int

p=int,q=long, r=float, s=double
p=int,q=float, r=long, s=double

##3.54
line2:int x->float, (float)x
line3:float * float->%xmm1, y=x*y
line4-5: float %xmm1->double%xmm2
line6: long -> double, (double)z
line7:%xmm0/%xmm1, 
return x*y-w/z;

##3.55
M=1+f=1
E=(2^10+4-(2^10-1))=5
M*2^E=1*2^5=32

##3.56
fabs(x)
0*(x)
-(x)

## 3.57
double funct3(int *ap, double b, long c, float *dp)
{
%xmm1=(float)*dp
%xmm2=(double)*ap
b-(double)*ap<=0
%xmm0 = (float)c
%xmm1 = (float)c * *dp
return 
.L8
%xmm1 = %xmm1 + %xmm1
%xmm0 = (float)c
%xmm0 = (float)c + *dp * 2;
return double(c+*dp*2);
}

if (b - *ap <=0)
{
return c+*dp*2;
}
else
{
return *dp * c;
}