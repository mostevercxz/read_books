# Chapter 1

Question lists:

* Did Simula have co-routine mechanisms ?(The answer seems to be yes after 
    googling...) Did c++ have?
* An inline funcs must have a unique definition in a program, why??

## 1.1 Simula language

The contrast I perceived between **the rigidity of Pascal** and **the flexibility of
Simula** was essential for the dev of C++. Simula's class concept was seen as the
key difference, and ever sicne I have seen classes as the proper primary focus
of program design.

The cost of Simula arose from several language features and their interactions:

* run-time type checking
* guaranteed intialization of variables
* concurrency support
* garbage collection of user-allocated objs and procedure activation records

What did BS learned by designing and implementing the simulator? A suitable tool
for projects such as writing a significant simulator, an OS, and system programming tasks:

* Have simula's support for program organization -- classes, class hierarchies
    , some form of support for concurrency, and strong(static) checking of a type system based on classes.
* Produce programs that run as fast as BCPL programs and share BCPL's
    ability to easily combine separately compiled uints into a program.
* Allow for highly portable implementations. A tool must have mutiple
    source implementations, there should be no complicated run-time support
    system to port, and there should be only very limited integration between
    the tool and its host operating system.

The most important aspect of theses criteria is that they are only loose
connected with specific programming language features. Instead, they specify
constraints on a solution.

* the C++ model of protection is based on the notion of granting and 
  transferring access rights; 
* the distinction between intialization and assignment has its root in thoughts
  about transferring capabilities;
* C++'s notion of **const** is derived from hardware read/write access
  protection mechanisms;
* C++'s exception handling mechanism was influenced by work on fault-tolerant
  systems done by Brian Randell's group in Newcastle during the seventies.

## 1.2 C as systems programming

The general idea of a systems programming language thus determined the growth of
C++ to at least the same extent as the specific language-techinal details of C.

## 1.3 general background

The structure of a system reflects the structure of the organization or
individual that created it. BS think the overall structure of C++ was shaped as
much by his general 'world view' as it was shaped by the detailed computer
science concepts used to form its individual parts.

Study pure and applied mathematics(master's degree); long-term hobby (25 years)
is history and later studying philosophy. Feel most at home with the empiricists
rather than with the idealists.

Many c++ design decisions have their roots in my dislike for forcing people to
do things in some particular way. Often, when I was tempted to outlaw a feature
I personally disliked, I refrained from doing so because I did not think I had
the right to force my views on others.

I design C++ to solve a problem, not to prove a point, and it grew to serve its
users. I do not try to enforce a single style of design through a narrowly
defined programming language. People's ways of thinking and working are so
diverse that an attempt to force a single style would do more harm than good.


# Chapter 2, C with classes

## 2.1 The birth of C with Classes

C with classes provided general mechanism for organizing programs, rather than
support for specific application areas. The choice between providing support for
specialized applications or general abstraction mechanisms has come up
repeatedly. Each time the choice is to **improve the abstraction mechanisms**.

BS was concerned that improved program structure was not achieved **at the
expense of run-time overhead** compared to C. The explicit aim was to match 
C in terms of run-time, code compactness, and data
compactness. (3% systematic decrease in run-time efficiency as to C,removed...)
His another concern was to avoid restrictions on the domain where C with Classes
could be used. **C with Classes could be used for whatever C could be used for**.
This implied that in addition to matching C in efficiency, C with Classes could
not provide benefits at the expense of removing "dangerous" or "ugly" features
of C. There is no one right way of writing every program, and a language
designer has no business trying to force programmers to use a particular
style.The language designer does have an obligation to encourage and support a
variety of styles and practices that have proven effective and to provide
language features and tools to help programmers avoid the well-known traps and
pitfalls.


## 2.2 Feature overview

The 1980 implementation features can be summarized:

* Classes
* derived classes, not virtual functions yet
* public/private access control
* constructors and destructors
* call and return functions(later removed)
* friend classes
* type checking and conversion of function arguments
* inline functions
* default arguments
* overloading of the assignment operator

## 2.3 Classes

A class specifies:

* the type of class members
* the set of operations(functions)
* the access users have to these members

Key design decisions:

* C++ lets the programmer specify types from which variables(objects) can be
  created. A class is a type. Why didn't call *class* type? Because BS dislike
  inventing new terminology and found Simula's quite adequate in most cases
* The representation of objects of the user-defined type is part of the class
  declaration. It means that true local vars of a user-defined type can be
  implemented without heap store or gc. It also means that a function must be
  recompiled if the representation of an object it uses is changed.
* Compile-time access control is used to restrict access to the representation.
  public/private access control
* The full type(inclding both the return type and the argument types) of a
  function is specified for function members. Compile-time checking is based on
  this type specification.
* Function definitions are typically specified "elsewhere" to make a class more
  like an interface specification than a lexical mechanism for organizing source
  code. This implies separate compilcation for class member functions and their
  users is easy and the C linker is sufficient to support C++
* The funciton *new()* is a constructor, a function with a special meaning to
  the compiler. Such funcs provided guarantees about classes, *new()* guarantes
  that the constructor is guaranteed to be called to initialize every obj of its
  class before the first use of the obj.
* Like C, objs can be allocated in three ways : stack, static storage, heap.
  Unlike C, C++ provides specific operators *new* and *delete* for heap
  allocation/deallocation


## 2.4 Run-time efficiency

BS want global and local variables of class types compared to Simula, which is a
major source of inefficiency of Simula.

User-defined and built-in types should behave the same relative to the language
tools.

The general reason for the introduction of inline funcs was that the cost of
crossing a protection barrier would cause people to refrain from using classes
to hide representation. C++ was always considered as something to be used now or
next month rather than simply a research project to deliver something in a
couple of years. It is not sufficient to provide a feature, it had to be
provided in an afforded form. *afforded* was seen as meaning "affordable on
hardware common	among developers", not "affordable to researchers with high-end
equipment" or "affordable in a couple of years".

### 2.4.1 Inlining

How to provide inlining? Have the programmer select which functions the compiler
should try to inline.

1. BS has poor experience with languages that left the job of inlining to
   compilers. Inlining a function for which you don't know the source appears
   feasible given afvanced linker and optimizer technology, but such technology
   wasn't available at the time.
1. Techniques that require global analysis tend not to scale well to very large
   programs.

Member funcs could be inlined and the only way to request a fucntion to be
inlined was to place its body within the class declaration. The fact that this
made class declarations messier was observed at the time and seen as a good
thing to discourage overuse of inline funcs. You can also inline non-member
functions in C++.

An *inline* directive is only a **hint** that the compiler can and often does
ignore. This is a logical necessity because:

* one can write recursive inline
  functions that cannot at compile time be proven not to cause infinite recursions
  and trying to inline one of those funcs would lead to infinite compilations. 
* A practical advantage because it allows the compiler writer to handle
  "pathological" cases by simple not inlining.

An inline function must have a unique definition in a program.

## 2.5 The linkage model

* Separate compilation should be possible with traditional C style linkers
* Linkage should be type safe
* Linkage should not require any form of database(although one could be used to
  improve a given implementation)
* Linkage to program fragments written in other languages such as C should be
  easy and efficient.
