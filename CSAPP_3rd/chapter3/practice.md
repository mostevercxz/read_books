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
