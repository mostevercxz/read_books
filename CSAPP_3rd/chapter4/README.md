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


## 4.2 Logic Design and Hardware control language(HCL)
In current technology, logic value 1 is represented by a high voltage of around 1.0 volt, while logic value 0 is represented by a low yoltage of around 0.0 volts. 

Three major components are required to implement a digital system: 

1. **combinational logic** to compute functions on the bits,
2. **memory elements** to store bits, and 
3. **clock signals** to regulate the updating of the memory elements.

### 4.2.1 Logic Gates
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/4/logic_gates.png "logic_gates")

We use these instead of the bit-level C operators &, |, and ~ , because logic gates operate on single-bit quantities, not entire words.

Logic gates are always active. If some input to a gate changes, then within some small amount of time, the output will change accordingly.

### 4.2.2 Combinational Circuits and HCL boolean exrpressions
By assembling a number of logic gates into a network, we can construct computational blocks known as **combinational circuits**, which have several restrictions:

1. Every logic gate input must be connected to exactly one of the following : 1) one of the system inputs(known as a primary input), 2) the output connection of some memory element, 3)the output of some logic gate
2. The output of two or more logic gates cannot be connected together
3. The network must be acyclic(非循环的)

Examples:
---
Bit equal:

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/4/bit_equal.png "bit_equal")

    bool eq = (a && b) || (!a && !b)

'=' associates a signal name with an expression, it is simply a way to give a name to an expression.
 
---
multiplexor(MUX, or multiplexer)
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/4/bit_mux.png "bit_mux")

    bool out = (s && a) || (!s && b) 

---

Differences between HCL and C expressions:

1. a combinational circuit has the property that the outputs continually respond to changes in the inputs. A C expression is only evaluated when it is encountered during the execution of a program
2. Logical expressions in C allow arguments to be arbitrary integers, while logic gates operate over the bit value 0 and 1.
3. C expression have the property that they might only be partially evaluated. Combinational logic does not have any partial evaluation rules. The gates simply respond to changing inputs.


### 4.2.3 Word-Level Combinational Circuits and HCL Integer expressions

Combinational circuits that perform word-level computations are constructed using logic gates to compute the individual bits of the output word, based on the individual bits of the input words. 

In HCL, we will declare any word-level signal as an int, without specifying the word size. This is done for simplicity. We'll draw word-level circuits using **medium-thickness** lines to represent the set of wires carrying the individual bits of the word, and we'll show a single-bit signal as a dashed line.

    bool Eq = (A == B);

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/4/word_equal.png "word_equal")

    word out = [
      s : A;
      1 : B;
    ]

---
case expressions:

Multiplexing funcs are described in HCL using case expressions. A case expression has the form:

    [
      select1 : expr1;
      select2 : expr2;
      ...
      selectk : exprk;
    ]

Unlike the switch statement of C, we do not require the different selection expressions to be mutually exclusive. Logically, the selection expressions are evaluated in sequence, and the case for the first one yielding 1 is selected.

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/4/four_mux.png "four_mux")

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/4/alu.png "alu")

One important combinational circuit, known as an
arithmetidlogic unit (ALU), has three inputs: two data inputs labeled A and B and a control input. Note the order, X - Y, the order is chosen in anticipation of the ordering of arguments in subq.(subq Y,X)

### 4.2.4 Set membership
Compare one signal against a number of possible matching signals.

General form of a set membership test is 
iexpr in {iexpr1, iexpr2, ..., iexprk}

where the value being tested (*iexpr*) and the candidate matches( *iexpr1-iexprk*) are all integer expressions.

    bool s1 = code ==2 || code ==3
    bool s1 = code in {2, 3}

### 4.2.5 Memory and Clock
Combinational circuits, by their very nature, do not store any information. They simply react to the signals at their inputs, generating outputs equal to some function of the inputs. To create **sequential circuits**, we must introduce devices that store infomation represented as bits.

Our storage devices are all controlled by a single clock, a periodic signal that determines when new values are to be loaded into the devices. We consider two classes of memory devices:

1. **Clocked registers** store individual bits or words. The clock signal controls the loading of the register with the value at its input.
2. **Random access memories** store multiple words, using an address to select which word should be read or written. Examples : 1)the virtual memory system of a processor; 2) the register file, where register identifies server as the address.

Thw word "register" means two slightly different things when speaking of **hardware language** Versus **machine language** :

* hardware, a register is directly connected to the rest of the circuit by its input and output wires.
* machine language, registers represent a small collection of addressable words in the CPU, where the addresses consist of register IDs. These words are generally stored in the register file.

Values only propagate from a register input to its output once every clock cycle at the rising clock edge. Our Y86-64 processors will use clocked registers to hold the program counter (PC), the condition codes (CC), and the program status (Stat).

---
typical register file:
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/4/typical_reg.png "typical_reg")

Such a multiported random access memory allows multiple read and write operations to take place simultaneously. In the register file diagranuued, the circuit can read the values of two program registers and update the state of a third. Each port has an address input indicating which program register should be selected, and a data output or input giving a value for that program register. The addresses are register identifiers.

Thw two read ports have **address inputs srcA and srcB**, **data outputs valA and valB**. The write port has **address input dstW** and **data input valW**. The **register file is not a combinational circuit**, since it has internal storage. In our implementation, however, data can be read from the register file as if it were a block of combinational logic having addresses as inputs and the data as outputs. When **srcA** is set to some register ID, after some delay, the value stored in the corresponding program register will appear on **valA**.

The writing of words to the register file is controlled by the clock signal in a manner similar to the loading of values into a clocked register. Every time the clock rises, the value on input valW is written to the program register indicated by the register ID on input dstW, when dstW=0xF, not register is written.

What happens if the circuit attempts to read and write the same register simultaneously?  there will be a transition on the read port's data output from the old value to the new.

---
random access memory

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/4/ram.png "ram")
This memory has a single address input, a data input for writing, and a data output for reading.

1. Set write control signal to 0, after some delay, the value stored at that address will appear on **data out**. Error signal will be set to 1 if address is out of range, 0 otherwise.
2. Set write signal to 1, set **address** to the desired address, **data in** to the desired value. When we then operate the clock, the specified location in the memory will be updated, as long as the address is valid.

---
Additional read-only memory for reading instructions.

## 4.3 Sequential Y86-64 implementations.
Our purpose in developing SEQ is to provide a first step toward our ultimate goal of implementing an efficient pipelined processor.
### 4.3.1 organizing processing into stages
