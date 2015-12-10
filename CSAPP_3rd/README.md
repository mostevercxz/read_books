# csapp 3e changes
1. Introduction. Minor revisions.  Move the discussion of Amdahl's Law to here, since it applies across many aspects of computer systems.
2. Data.  Do some tuning to improve the presentation, without diminishing the core content.  Present fixed word size data types.
3. Machine code.  **A complete rewrite, using x86-64 as the machine language, rather than IA32.  Also update examples based on more a recent version of GCC (4.8.1)**.  Thankfully, GCC has introduced a new opimization level, specified with the command-line option `-Og' that provides a fairly direct mapping between the C and assembly code.  We will provide a web aside describing IA32.
4. Architecture.  Shift from Y86 to y86-64.  This includes having 15 registers (omitting %r15 simplifies instruction encoding.), and all data and addresses being 64 bits.  Also update all of the code examples to following the x86-64 ABI conventions.
5. Optimization.  All examples will be updated (they're mostly x86-64 already).
6. Memory Hierarchy.  Updated to reflect more recent technology.
7. Linking.  Rewritten for x86-64.  We've also expanded the discussion of using the GOT and PLT to create position-independent code, and added a new section on the very cool technique of library interpositioning.
8. Exceptional Control Flow.  More rigorous treatment of signal handlers, including async-signal-safe functions, specific guidelines for writing signal handlers, and using sigsuspendto wait for handlers.
VM.  Minor revisions.
9. System-Level I/O.  Added a new section on files and the file hierarchy.
10. Network programming.  Protocol-independent and thread-safe sockets programming using the modern getaddrinfo and getnameinfo functions, replacing the obsolete and non-reentrant gethostbyname and gethostbyaddr functions.
11. Concurrent programming.  Enhanced coverage of performance aspects of parallel multicore programs.


# csapp errta
1. 3rd edition errta: [3rd errta](http://csapp.cs.cmu.edu/3e/errata.html)
2. 2rd edition errta: [2en errta](http://csapp.cs.cmu.edu/2e/errata.html)