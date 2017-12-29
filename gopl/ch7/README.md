# Interfaces

## 7.1 Interfaces as contracts

An interface type is an abstract type. When you have a value of an interface, you know nothing about what it is, you know only what it can do, what behaviors are provided by its methods.

This freedom to substitute one type for another that satisfies the same interface is called **substitutablity**, and is a hallmark of **OOP**.

```go
type ByteCounter int
func (bc *ByteCounter) Write(p []byte) (int, error) {
    *c += ByteCounter(len(p))
    return len(p), nil
}

var c ByteCounter
c.Write([]byte("hello"))
c = 0
fmt.Fprintf(&c, "hello, %s", "world~")
fmt.Println(c)
```

## 7.2 Interface types

An interface type specifies a set of methods that a concrete type must possess to be considered an instance of that interface.

We can declare new interface types as combinations of existing ones, which is called *embedding* an interface. 

**The order in which the methods appear does not matter. All that matters is the set of methods.**

## 7.3 Interface satisfaction

A type *satisfy* an interface if it possesses all the methods the interface requires. We can say the type **"is a"** particular interface type. 

>a *bytes.Buffer is an io.Writer; an *os.File is an io.ReadWriter

The assignablity rule for interfaces is very simple : an expression(may be an interface itself) may be assigned to an interface only if **its type satisfies the interface**.

a value of type T does not possess all the methods that a *T pointer does. Why is it legal to call a *T method on an argument of type T so long as the argument is a *variable*? Merely a syntactic sugar.

**interface{}**, which is called the *empty interface type*, is not dispensable. Because the empty interface type places no demands on the types that satisfy it, we can assign any value to the empty interface.

How to assert the relationship between a interface type and a concrete type?
>Eg. var _ io.Writer = (*bytes.Buffer)(nil)

Types that satisfy interfaces:

* pointer types
* slice types with methods(**geometry.Path**)
* map types with methods(**url.Values**)
* function type with methods(**http.HandlerFunc**)
* basic types, time.Duration

## 7.4 Parsing Flags with flag.Value
