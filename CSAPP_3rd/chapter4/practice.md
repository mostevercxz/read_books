## 4.1
30F3 0F00 0000 0000 0000
2031
4013 FDFF FFFF FFFF FFFF
6031
70   0c01 0000 0000 0000

##4.2
A.

    irmovq $-4, %rbx
    rmmovq %rsi, 0x800(%rbx)

B.

    pushq %rsi
    call proc
    halt
    proc:
      irmovq 0xa, %rbx
      ret

C.

    mrmovq 0x7(%rsp), %rbp
    nop
    invalid bytes f0
    popq %rcx

D.

    loop:
      subq %rcx, %rbx
      je loop
      halt

E.


    xorq %rsi, %rdx
invalid byte, a0f0

##4.3
iaddq V,rB

    sum:
      xorq %rax, %rax
      andq %rsi, %rsi
      jmp test
    loop:
      mrmovq (%rdi), %r8
      addq %r8, %rax
      iaddq $8, %rdi
      iaddq $-1, %rsi
    test:
      jne loop
      ret

## 4.4

long rsum(long *start, long count)
{
  if (count <= 0)
  {return 0;}

  return *start + rsum(start+1, count-1);
}

x86-64 code(gcc -Og -S rsum.c)

    rsum:
      movl $0, %eax
      testq %rsi, %rsi
      jle .L6
      pushq %rbx
      movq (%rdi), %rbx
      addq $8, %rdi
      subq $1, %rsi
      call rsum
      addq %rbx, %rax
      popq %rbx
    .L6:
      rep;ret

Y86-64 code:

.pos 0
rsum:
  xorq %rax, %rax(irmovq $0, %rax)
  andq %rsi, %rsi
  jle end_now
  pushq %rbx
  mrmovq (%rdi), %rbx
  irmovq $8, %r8
  addq %r8, %rdi
  irmovq $1, %r9
  subq %r9, %rsi
  call rsum
  addq %rbx, %rax
  popq %rbx
end_now:
  ret

## 4.5
absSum that computes the sum of absolute values of an array.

loop:
  mrmovq (%rdi), %r10
  xorq %r11, %r11
  subq %r10, %r11
  jle x_greater_zero
  rrmovq %r11, %r10
x_greater_zero:
  addq %r10, %rax
  addq %r8, %rdi
  subq %r9, %rsi

------ my answer(ignore the fact OPq set CC) ------

    loop:
      mrmovq (%rdi), %r10
      andq %r10, %r10
      jge great_zero
      xorq %r11, %r11
      subq %r11, %r10
    great_zero:
      addq %r10, %rax
      addq %r8, %rdi
      subq %r9, %rsi

--------my answer end -------

## 4.6
absSum using conditional move instruction

loop:
  mrmovq (%rdi), %r10
  xorq %r11, %r11
  subq %r10, %r11
  cmovg %r11, %r10
  addq %r10, %rax
  addq %r8, %rdi
  subq %r9, %rsi

## 4.7
the original value

## 4.8
the value read from memory

mrmovq (%rsp), %rsp

## 4.9
bool xor = (a && !b) || (!a && b)

complements of each other.

## 4.10
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/4/word_xor.png "word_xor")


## 4.11
word min3 = [
  A <= B && A <= C : A;
  B <= C : B;
  1 : C;
]

## 4.12
word median3 = [
  A <= B && B <= C : B;
  C <= B && B <= A : B;
  B <= A && A <= C : A;
  C <= A && A <= B : A;
  1 : C;
]

