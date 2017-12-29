# Functions

## 5.1 function declarations

```go
func name(para-list) (result-list) {
    body
}
```

The type of a func is sometimes called its *signature*. Two funcs have the same type or signature if they have the same sequence of para types and the same sequence of result types.

Go has no concept of default para values, nor any way to specify arguments by name.

Arguments are passed by **value**. If the argument contains some kind of reference(pointer,slice,map,function,channel).

## 5.2 recursion

## 5.3 multiple return values

## 5.4 errors

## 5.5 function values

## 5.6 anonymous functions

Named funcs can be declared only at the package level. But we can use a *function literal* to denote a func value within any expression. 

A function literal is written like a function, but without a name following the **func** keyword. It is an expression and its value is called an **anonymous function**.

## 5.7 variadic functions

## 5.8 deferred function calls

## 5.9 panic

## 5.10 recover