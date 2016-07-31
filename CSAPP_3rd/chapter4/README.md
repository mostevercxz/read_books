**Questions**

1. Web aside ASM:EASM, how to write programs that combine C code with handwritten assembly code?
2. The two commands *ret*, *popq* both get the value of R[%rsp] two times, valA <-- R[%rsp],valB <-- R[%rsp], why ?
3. rrmovq, valE <-- valA + 0, why "+0"?
4. Page 504, psum1 CPE=9cycles/element, why ? How to calculate ?
4. Sequential implementation of Y86-64 instruction "cmovXX rA, rB" is not right.
5. What is the sequential implementation of Y86-64 instruction "halt"? (Cause the processor status to be set to HLT, causing it to halt operation.)

One important concept is that the actual way a modern processor operates can be quite different from the model of computation implied by the ISA. 

This idea of using clever tricks to improve performance while maintaining the functionality of a simpler and more abstract model is well known in computer science. Examples include the use of caching in Web browsers and information retrieval data structures such as balanced binary trees and hash tables.

Why should you learn about processor design?

1. It is intellectually interesting and important.
2. Understanding how the processor works aids in understanding how the overall computer system works.
3. Although few people design processors, many design hardware systems that contain processors.
4. You just might work on a processor design.


## 4.1 Y86-64 instruction set architecture
Defining an ISA includes:
 
1. **defining the different components of its state**, 
2. **the set of instructions and their encodings**, 
3. **a set of programming conventions**, 
4. and **the handling of exceptional events**.

### 4.1.1 Programmer(assemnly code or compiler-writer)-visiable state

15 program registers
%rsp is used as a stack pointer(push, pop, call, ret)

3 single-bit condition codes : ZF, SF and OF

memory is conceptually a large array of bytes.

a status code Stat indicates the overall state of programm execution.

### 4.1.2 Y86-64 instructions
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/4/y8664is.png "y8664 instruction set")

* irmovq, rrmovq, mrmovq, rmmovq indicates the form of the source and destination. The source is immediate(i), register(r), memory(m), designated by the first character in the instruction name. The destination is register(r) or memory(m). The memory references for the two memory movement instructions have a simple base and displacement format. We do not support the second index register or any scaling of a register's value.
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
In current technology, logic value 1 is represented by a high voltage of around 1.0 volt, while logic value 0 is represented by a low voltage of around 0.0 volts. 

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

In HCL, we will declare any word-level signal as an int, without specifying the word size. This is done for simplicity. 

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/4/word_level_logic.png "word_level")
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/4/word_level_multiplex.png "word_level_multiplex")

As shown in the picture above, We'll draw word-level circuits using **medium-thickness** lines to represent the set of wires carrying the individual bits of the word, and we'll show a single-bit signal as a dashed line.

    bool Eq = (A == B);

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/4/word_equal.png "word_equal")

    word out = [
      s : A;
      1 : B;
    ]

We'll use many forms of multiplexors in our processor designs as they allow us to select a word from a number of sources depending on some control condition.

---
case expressions:

Multiplexing funcs are described in HCL using case expressions. (each case i consists of a Boolean expression select(i), indicating when this case should be selected) A case expression has the form:

    [
      select1 : expr1;
      select2 : expr2;
      ...
      selectk : exprk;
    ]

Unlike the switch statement of C, we do not require the different selection expressions to be mutually exclusive. Logically, the selection expressions are evaluated in sequence, and **the case for the first one yielding 1 is selected**.

Consider 4-way multiplexor 

    word out4 = [
      !A && !B : D0;#00
      !A       : D1;#01
      !B       : D2;#10
      1        : D3;#11
    ]

The selection expressions can sometimes be simplified, since only the first matching case is selected. !A && B can be simplified as !A

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/4/four_mux.png "four_mux")

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/4/alu.png "alu")

One important combinational circuit, known as an
arithmetidlogic unit (ALU), has three inputs: two data inputs labeled A and B and a control input. Note the order, X - Y, the A input is subtracted from the B input, the order is chosen in anticipation of the ordering of arguments in subq.(subq Y,X)

### 4.2.4 Set membership
We'll find many examples where we want to compare one signal against a number of possible matching signals, such as to test whether the code for some instruction being processed matches some category of instruction  codes.

General form of a set membership test is 
iexpr in {iexpr1, iexpr2, ..., iexprk}

where the value being tested (*iexpr*) and the candidate matches( *iexpr1-iexprk*) are all integer expressions.


Suppose we want to generate the signals s1 and s0 for the 4-way multiplexor as follows:

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/4/set_membership.png "set_membership")

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

The registers serve as barriers between the combinational logic in different parts of the circuit. Values only propagate from a register input to its output once every clock cycle at the rising clock edge. Our Y86-64 processors will use clocked registers to hold the program counter (PC), the condition codes (CC), and the program status (Stat).

---
**typical register file:**
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/4/typical_reg.png "typical_reg")

Two read ports: A and B. One write port: W. Such a multiported random access memory allows multiple read and write operations to take place simultaneously. In the register file diagrammed, the circuit can read the values of two program registers and update the state of a third. Each port has an address input indicating which program register should be selected, and a data output or input giving a value for that program register. The addresses are register identifiers.

Thw two read ports have **address inputs srcA and srcB**, **data outputs valA and valB**. The write port has **address input dstW** and **data input valW**. The **register file is not a combinational circuit**, since it has internal storage. In our implementation, however, data can be read from the register file as if it were a block of combinational logic having addresses as inputs and the data as outputs. When **srcA** is set to some register ID, after some delay, the value stored in the corresponding program register will appear on **valA**.

The writing of words to the register file is controlled by the clock signal in a manner similar to the loading of values into a clocked register. Every time the clock rises, the value on input valW is written to the program register indicated by the register ID on input dstW, when dstW=0xF, not register is written.

What happens if the circuit attempts to read and write the same register simultaneously?  t**here will be a transition on the read port's data output from the old value to the new.**

---
**random access memory**

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/4/ram.png "ram")
This memory has a single address input, a data input for writing, and a data output for reading.

1. Set write control signal to 0, after some delay, the value stored at that address will appear on **data out**. Error signal will be set to 1 if address is out of range, 0 otherwise.
2. Set write signal to 1, set **address** to the desired address, **data in** to the desired value. When we then operate the clock, the specified location in the memory will be updated, as long as the address is valid.

The signal is generated by combinational logic, since the required bounds checking is purely a function of the address input and does not involve saving any state.

---
Additional read-only memory for reading instructions.

## 4.3 Sequential Y86-64 implementations.
Our purpose in developing SEQ(for "sequential",a processor) is to provide a first step toward our ultimate goal of implementing an efficient pipelined processor.
### 4.3.1 organizing processing into stages
The following is an informal description of the stages
and the operations performed within them :

1. **Fetch**. Reads the bytes of an instruction from memory, using the program counter (PC) as the memory address. From the instruction it extracts the two 4-bit portions of the instruction specifier byte, referred to as icode (the instruction code) and ifun (the instruction function). It computes **valP** to be the address of the instruction following the current one in sequential order.
2. **Decode**. Reads up to two operands from the register file, giving values **valA** and/or **valB**.(rA, rB, some instructions pushq, reads %rsp)
3. **Execute**. ALU(arithmetic/logic unit) either performs the operation specified by the instruction(ifun), computes the effective address of a memory reference, or increments(decrements) the stack pointer. The condition codes are possibly set.(cmovXX, evaluate the condition codes and move condition, enable the updating of the destination register only if the condition holds.) (For jump instruction, it determines whether or not the branch should be taken.)
4. **Memory**. Write data to memory, may read data from memory.
5. **Write back**. writes up to two results to the register file.
6. **PC update**. PC is set to the address of the next instruction.

The processor loops infinitely, will stop when any exception occurs. One way to minimize the complexity is to have the different instructions share as much of the hardware as possible.

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/4/instruction_sequence.png "instruction_sequence")

----
Instruction types OPq(addq, subq, andq, xorq), rrmovq, irmovq

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/4/computations.png "computations")

The control logic blocks labeled "icode" and "ifun" then compute the instruction and function codes as equaling either the values read from memory or, in the event that the instruction address is not valid, the values corresponding to a nop instruction.

nop, halt computations

### 4.3.2 SEQ hardware structure
Our method of drawing processors with the flow going from bottom to top is unconventional. We will explain the reason for our convention when we start designing pipelined processors. Here is the question, Why ?

Some stages worth notification:
Write back :
The register file has two write ports. Port E is used to write values computed by the ALU, while Port M is used to write values read from the data memory.

PC Update:
The new value of PC is set to be either valP, valC(jXX Dest) or valM(ret).

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/4/abstract_SEQ.png "abstract SEQ view")

Below is a more detailed view of the hardware required to implement SEQ:
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/4/more_detailed_SEQ.png "more_detailed_SEQ")

We use the following drawing conventions:

1. Clocked registers are shown as white rectangles, the PC register
2. Hardware units are shown as light blue boxes. These include the memories, the ALU, and so forth.(Treat these units as black boxes)
3. Control logic blocks are drawn as gray rounded rectangles.
4. Wire names are indicated in white circles.These are simply labels on the wires, not any kind of hardware element.
5. Word-wide data connections are shown as medium lines.(Each of these lines actually represents a bundle of 64 wires, connected in parallel, for transferring a word from one part of the hardware to another.)
6. Byte and narrower data connections are shown as thin lines.(four or eight wires)
7. Single-bit connections are shown as dotted lines. 

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/4/computation_steps.png "computation_steps")
The second column includes 4 register ID signals:
srcA, the source of valA;
srcB, the source of valB;
dstE, the register to which valE gets written;
dstM, the register to which valM gets written.

### 4.3.3 SEQ Timing
Our implementation of SEQ consists of combinational logic and two forms of memory devices: clocked registers (the program counter and condition code register) and random access memories (the register file, the instruction memory,and the data memory).

We also assume that reading from a random access memory operates much like combinational logic. This is a reasonable assumption for smaller memories (such as the register file), and we can mimic this effect for larger circuits using special clock circuits. (Question, why reasonable assumption?)

**We are left with just four hardware units that require an explicit control over their sequencing**:

1. the program counter,loaded with a new instruction address every clock cycle.
2. the condition code register,loaded only when an integer operation instruction is executed.
3. the data memory, written only when rmmovq, call, pushq is executed.
4. the register file

We have organized the computations in such a way that our design obeys the following principle: **No reading back**.(The processor never needs to read back the state updated by an instruction in order to complete the processing of this instruction, example: pushq generates valE, not use %rsp directly)

No instruction must both set and then read the condition codes. Even though the condition codes are not set until the clock rises to begin the next clock cycle, they will be updated before any instruction attempts to read them.

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/4/trace_two_cycles.png "trace_two_cycles")

### 4.3.4 SEQ Stage implementations
The constants we use are documented below. By convention, we use uppercase names for constant values.

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/4/constant_values_HCL.png "constant_values_HCL")

---
Fetch Stage
This unit read 10 bytes from memory at a time, using the PC as the address of the first byte (byte 0). Control logic blocks labeled "icode" and "ifun" then compute the instruction and function codes as equaling:

1. the values read from memory 「in the event that the instruction address is not valid (as indicated by the signal imem_error, generated when the instruction address
is out of bounds)」 
2. the values corresponding to a nop instruction. 

Based on the value of icode, we can compute three 1-bit signals(shown as dashed lines):

1. instr_valid. Does this byte correspond to a legal Y86-64 instruction?
2. need_regids. Does this instruction include a register specifier byte?
3. need_valC. Does this instruction include a constant word?

Note:

    The signals instr_valid , imem_error are used to generate the status code in the memory stage.

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/4/SEQ_fetch.png "SEQ fetch")

The HCL description for need_regids:

    bool need_regids = icode in { IRRMOVQ, IOPQ, IPUSHQ, IPOPQ, IIRMOVQ, IRMMOVQ, IMRMOVQ };

---
Decode and Write-back Stages
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/4/SEQ_decode.png "SEQ_decode")

RRSP is the register ID of %rsp

    word srcA = [
      icode in {IRRMOVQ, IRMMOVQ, IOPQ, IPUSHQ} : rA;
      icode in {IPOPQ, IRET} : RRSP;
      1 : RNONE;
    ]
    word srcB = [
      icode in {IOPQ, IRMMOVQ, IMRMOVQ} : rB;
      icode in {IPUSHQ, IPOPQ, ICALL, IRET} : RRSP;
      1 : RNONE;
    ]

Register ID dstE indicates which register for write port E, where the computed value valE is stored. If we ignore for the moment the conditional move instructions, we can give the following HCL description of dstE:

    word dstE = [
      icode in {IRRMOVQ} : rB;
      icode in {IIRMOVQ, IOPQ} : rB;
      icode in {IPUSHQ, IPOPQ, ICALL, IRET} : RRSP;
      1 : RNONE;
    ]    
    word dstM = [
      icode in {IMRMOVQ, IPOPQ} : rA;
      1 : RNONE;
    ]

We'll revisit how to implement conditional moves when we examine the execute stage.

---
Execute Stage
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/4/SEQ_execute.png "SEQ_execute")
The ALU unit performs the operation ADD, SUBTRACT, AND, or EXCLUSIVE-OR on inputs **aluA** and **aluB** based on the setting of the **alufun** signal. These data and control signals are generated by three control blocks, as diagrammed in Figure 4.29.

    word aluA = [
      icode in {IRRMOVQ, IOPQ} : valA;
      icode in {IIRMDVQ, IRMMOVQ, IMRMOVQ} : valC;
      icode in { ICALL, IPUSHQ } : -8;
      icode in { IRET, IPOPQ } : 8;
    ]

    word aluB = [
      icode in {IOPQ, IRMMOVQ, IMRMOVQ, IPUSHQ, IPOPQ, ICALL, IRET} : valB;
      icode in {IRRMOVQ, IIRMOVQ} : 0;
    ]
    bool set_cc = icode in {IOPQ};

The unit "cond" uses combination of the condition codes and the function code to determine whether a conditional branch or data transfer should take place. It generates the Cnd signal used both for the setting of dstE with conditional moves and in the next PC logic for conditional branches.

---
Memory Stage
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/4/SEQ_memory.png "SEQ_memory")

When a read operation is performed, the data memory generates the value valM.

Observe that the address for memory reads and writes is always valE or valA.

    word mem_addr = [
      icode in {IRMMDVQ, IPUSHQ, ICALL, IMRMDVQ} : valE;
      icode in {IPOPQ, IRET} : valA;
    ]

The data for memory writes are always either valA or valP.
word mem_data = [
]

---
PC Update Stage

    word new_pc = [
      icode == ICALL : valC;
      icode == IJXX && Cnd : valC;
      icode == IRET : valM;
      1 : valP;
    ]

---
Surveying SEQ
The only problem with SEQ is that it is too slow.

## 4.4 General Principles of Pipelining
### 4.4.1 Computational Pipelines
### 4.4.2 A detailed look at Pipeline operation
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/4/pipeline_operation.png "pipeline_operation")
### 4.4.3 Limitations of Pipelining
### 4.4.4 Pipelining a system with feedback


## 4.5 Pipelined Y86-64 implementations
### 4.5.1 SEQ + : Rearranging the Computation Stages
PC update stage comes at the beginning of the clock cycle, rather than at the end. With SEQ+, we create state registers to hold the signals computed during an instruction. Then, as a new clock cycle begins, the values propagate through the exact same logic to compute the PC for the now-current instruction.(pIcode to indicate that on any given cycle, they hold the control signals generated during the **previous cycle**.)

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/4/SEQ+hardware.png "SEQ+ hardware")

### 4.5.2 Inserting pipeline registers
we insert pipeline registers between the stages of SEQ+ and rearrange signals somewhat, yielding the PIPE– processor, where the “–” in the name signifies that this processor has somewhat less performance than our ultimate processor design. In Figure 4.41, these white boxes are real hardware components.

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/csapp/4/PIPE-hardware.png "PIPE- hardware")

Avoid Data hazards by Stalling(货摊,托辞,停止,拖延)