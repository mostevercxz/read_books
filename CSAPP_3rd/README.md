# Read CS:APP
## Chapter 1 : A Tour of computer systems
### 1.1 Infomation is bits + context
All infomation in a system-including disk files, programs stored in memory, user data stored in memory, and data transferred across a network--is represented as **a bunch of bits**. The only thins that distinguishes different data objects is *the context in which we view them*.
### 1.2 Programs are translted by other programs into different forms
On a Unix system, the translaion from source file to object file is performed by a *compiler driver* :

hello.c (source program text) -> Preprocessor(cpp)

hello.i (modified source program text) -> Compiler(cc1)

hello.s (Assembly program text) -> Assembler(as)

hello.o + print.o (Relocatable object programs binary) -> Linker(ld)

hello (Executable object program binary)

```
linux> gcc -o hello hello.c
```
  
* preprocessing phase : The preprocessor modifies the original C program according to directives that begin with the '#' character. ('#include <stdio.h>' tells the preprocessor to read the contents of stdio.h and insert it directly into the program text)
 
* Compilation phase : The compiler(cc1) trranslates the text file hello.i into an **assembly-language program** hello.s 

* Assembly phase : The assembler(as) translates hello.s into machine-language instructions, packages them in a form known as a **relocatable object program**

* Linking phase : The linker (ld) merges printf.o and hello.o and the result is the *hello* file.

### 1.3 It pays to understand how compilation systems work
Why programmers need to understand how compilation systems work ?

* Optimizing program performance

* Understanding link-time errors

* Avoiding security holes. (buffer overflow vulnerablities)

### 1.4 Processors read and interpret instructions stored in memory


    linux> ./hello
    hello, world
    linux>

#### 1.4.1 hardware organization of a system
Figure hardware organization: 

1. Buses : a collection of electrical conduits that carry bytes of infomation back and forth between the components. Buses are typically designed to transfer fixed-size chunks of bytes known as words(4 bytes-32bit, or 8 bytes-64 bit)

2. IO Devices : are the systems's connection to the external world(mouse, keyboard, screen, disk...) Each IO device is connected to the IO bus by either a controller or an adapter. 
 * Controllers are chip sets in the device itself or on the system's main printed circuit board(motherboard)
 * Adapter is a card that plugs into a slot on the motherboard(4G memory laptop) 
3. Main Memory : is a temporary storage device that holds both a program and the data it manipulates while the processor is executing the program.
 * Physically, main memory consists of a collection of dunamic random access memory(DRAM)
 * Logically, memory is organized as a linear array of bytes.
4. Processor : CPU is the engine that interprets instructions stored in main memory. At its core is a word-size storage device(register) called the program counter(PC). Simple ops:
 * Load (memory -> register)
 * Store (register -> memory)
 * Operate (registers -> ALU)
 * Jump (overwrite PC)

#### 1.4.2 Running hello
keyboard './hello' -> IO bridge -> Bus Interface -> Register -> IO -> memory

(Press Enter) Disk -> Disk Controller -> IO bridge -> memory

memory -> bus interface -> register -> bus interface -> IO bridge -> Graphics adapter -> Screen Display

Figures:

### 1.5 Caches Matter
The processor-memory gap(processor read 100 times faster than memory)

Figure Cache memories:


The idea behind caching is that a sys can get the effect of both a very large memory and a very fast one by exploiting locality, the tendency for programs to access data and code in localized regions.

**Application programmers who are aware of cache memories can exploit them to improve the performance of their programs by an order of magnitude.**

### 1.6 storage devices from a hierarchy
Figure L1 cache, L2 cache, L3 cache(SRAM)

### 1.7 os manages the hardware
We can think of os as a layer of software interposed between the application proram and the hardware.

The os has two primary purposes:

1. protect the hardware from misuse by runaway apps.
2. procide apps with simple and uniform mechanisms for manipulating complicated and low-level hardware devices.

#### 1.7.1 Processes
the notion of process : the most important and successful ideas in CS.

A process is the OS's abstraction for a running program.

OS keeps track of context(PC, register, memory) that the process needs in order to run.
When the OS decides to transfer control from the current process to new process, it performs a context switch by
 
* saving the context of the current process(shell)
* restoring the context og the new process(hello)
* passing control to the new process(shell)

#### 1.7.2 threads
A process can actually consists of multiple execution units called threads, each running in the context of the process and sharing the same code and global data.

#### 1.7.3 virtual memory
An abstraction that provides each process with the illusion that it has exclusive use of the main memory.(virtual address space, VAS)

VAS consists of a number of well-defined areas :

1. Program code and data (code begins at the same fixed address for all processes, followed by data locations that correspond to global C vars. initialized directly from the contents of an executable object file)
2. Heap, code and data areas are followed immediately by the run-time heap.
3. Shared libs
4. stack, at the top of the user's VAS
5. kernel virtual memory

#### 1.7.4 files
A file is a sequence of bytes. Every IO device(disk,keyborads, displays, networks) is modeled as a file.

### 1.8 system communicate with other systems using network
### 1.9 import themes
A system is a collection of intertwined 缠绕的,错综复杂的 hardware and systems software that must cooperate in order to achieve the ultimate goal of running application programs.

#### 1.9.1 Amdahl's law 阿姆达尔定律
The main idea is that when we speed up one part of a system, the effect on the overall system performance depends on both how significant this part was and how much it sped up.

Consider a system : execute some apps requires time T(old). Some part of the sys requires a fraction α, which means αT(old). If we improve its performance by a factor of k, The overall execution time would be :

    T(new) = (1-α)T(old) + (αT(old)) / k
           = T(old) [(1-α) + α / k]
    T(old) / T(new) = 1 / [1 + (α/k - α)]

    α/k - α < 0, α > 0, 1/k < 1, sys performance improves

k -> ∞, T(old) / T(new) = 1 / (1-α) 

#### 1.9.2 Concurrency and parallelism
Two demands : we want computers to do more, run faster.

*concurrency* : general concept of a sys with multiple, simultaneous activities

*parallelism* : the use of concurrency to make a sys run faster.

1. **Thread-level concurrency**
 * *uniprocessor sys* : most actual computing was done by a single processor, even if that processor had to switch among multiple tasks
 * *multiprocessor sys* : a sys consisting of multiple processors all under the control of a single OS kernel, becoming commonplace with the advent of *multi-core processors* and *hyperthreading*
 * multi-core processors,Intel core i7 organization : 
 * hyperthreading(simultaneous multi-threading) : a technology that allows a single CPU to execute multiple flows of control. Whereas a wnventional processor requires around 20:000 clock cycles to shift between different threads, a hyperthreaded processor decides which of its threads to execute on a cycle by cycle basis.
 * The use of multiprocessing can improve sys performance in two ways : 
     1. reduces the need to simulate concurrency when performing multiple tasks
     2. run a single app faster only if that app is expressed in terms of multiple threads that can effectively execute in parallel.
2. **Instruction-Level Parallelism**
  * *instruction-level parallelism* : a property that processors can execute multiple instructions at one time.
3. **Single-Instruction, Multiple-Data Parallelism（SIMD)**
  * SIMD is a model that allows a single instruction to cause multiple operations to be performed in parallel.  

#### 1.9.3 The importance of abstraction in Computer sys
1. The processor side : the instruction set architecture provides an abstraction of the actual processor hardware
2. OS side : 
  * files as an abstraction of IO devices; 
  * virtual memory as an abstraction of program memory;
  * processes as an abstraction of a running program
3. virtual machine : an abstraction of the entire computer

### 1.10 summary 


