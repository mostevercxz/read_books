**Questions**
1. Web aside ASM:EASM, how to write programs that combine C code with handwritten assembly code?

One important concept is that the actual way a modern processor operates can be quite different from the model of computation implied by the ISA. 

This idea of using clever tricks to improve performance while maintaining the functionality of a simpler and more abstract model is well known in computer science. Examples include the use of caching in Web browsers and information retrieval data structures such as balanced binary trees and hash tables.

Why should you learn about processor design?

1. It is intellectually interesting and important.
2. Understanding how the processor works aids in understanding how the overall computer system works.
3. Although few people design processors, many design hardware systems that contain processors.
4. You just might work on a processor design.


## 4.1 Y86-64 instruction set architecture
Defining an ISA includes defining the different components of its state, the set of instructions and their encodings, a set of programming conventions, and the handling of exceptional events.

### 4.1.1 Programmer(assemnly code or compiler-writer)-visiable state

15 program registers
%rsp is used as a stack pointer(push, pop, call, ret)

3 single-bit condition codes : ZF, SF and OF

memory is conceptually a large array of bytes.

a status code Stat indicates the overall state of programm execution.

### 4.1.2 Y86-64 instructions
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/4/y8664is.png "y8664 instruction set")

* irmovq, rrmovq, mrmovq, rmmovq indicates the form of the source and destination. The source is immediate(i), register(r), memory(m), designated by the first character in the instruction name. The destination is register(r) or memory(m). The memory references for the rwo memory movement instructions have a simple base and displacement format. We do not support the second index register or ant scaling of a register's value.
* 4 integer operation instructions, addq, subq, andq, xorq, operating only on register data and setting three condition codes ZF,CF,OF.
* 7 jump instructions are jmp, jle, je, jne, jge, jg
* 6 conditional move instrictions are cmovle, cmovl, cmove, cmovge, cmovne, cmovg
* call , ret
* *halt* instruction stops instruction execution. (x86-64 hlt, but is not allowed to use)

### 4.1.3 Instruction encoding
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/4/functioncodes.png "functioncodes")

Each instruction require 1-10 bytes.

Register require 1 bytes.(15 registers) ranging from 0 to 0xE. 0xF is used in the instruction encodings and within our hardware designs when we need  to indicate that no register should be accessed.

First, there can be an additional register specifier byte, specifying either one or two registers.(rA and rB). instructions that have no register operands do not have a register specifier byte. Those that require just one register operand have the other register specifier set to 0xF. 

Branch and call destinations are given as absolute addresses rather than using the PC-relative addressing

### 4.1.4 Y86-64 exceptions
A status code Stat describing the overall state of the executing program. Possible values are

    1 AOK Normal operation
    2 HLT halt instruction encountered
    3 ADR Invalid address encountered
    4 INS Invalid instruction encountered

We limit the maximum address and any access to an address beyond this limit will trigger an ADR exception. We'll simply have the processor stop executing instructions when it encounters any of the exceptions listed.

### 4.1.5 Y86-64 Programs

    long sum(long *start, long count)
    {
      long sum = 0;
      while (count){
        sum+= *start;
        start++;
        count--;
      }
      return sum;
    }


x86-64 code

    sum:
      movl $0, %eax
      jmp .L2
    .L3:
      addq (%rdi), %rax
      addq $8, %rdi
      subq $1, %rsi
    .L2:
      testq	%rsi, %rsi
      jne .L3
      rep;ret

Y86-64 code

    sum:
      irmovq $8, %r8
      irmovq $1, %r9
      xorq %rax, %rax (2 bytes, irmovq $0, %rax needs 10 bytes)
      andq %rsi, %rsi
      jmp test
    loop:
      mrmovq (%rdi), %r10
      addq %r10, %rax
      addq %r8, %rdi
      subq %r9, %rsi
    test:
      jne loop
      ret

Some differences:

* the y86-64 code loads constants into registers, since it cannot use immediate data in arithmetic instructions
* y86-64 code requires two instructions to read a value from memory and add it to a register
* OPq(subq) instruction also sets the CC(consition codes), so we do not need(testq %rsi, %rsi)


.pos 0 indicates that the assembler should begin generating code starting at address 0. This is the starting address for all Y86-64 programs.


### 4.1.6 some y86-64 instruction details
Two unusual instruction combinations require special attention:

1. pushq both decrements the stack pointer by 8 and writes a register value to memory. (pushq %rsp; 1.push the original value of %rsp; 2. push the decremented value of %rsp). The original value.
2. popq %rsp, 1.the value read from memory; 2. the incremented stack pointer. The value.