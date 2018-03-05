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

No notes need to be taken...

## 7.5 Interface Values

An interface value has two components : a concrete type(**dynamic type**) and a value of that type(**dynamic value**).

For statically typed language Go, types are a compile-time concept, a type is not a value. In our conceptual model, a set of values called *type descriptor* provide infomation(name,methods...) about each type. In an interface value, the type component is represented by the appropriate type descriptor.

>Why could not we know at compile time?

We can not know at compile time what the dynamic type of an interface value will be, so a call through an interface must use **dynamic dispath**.

The compile must generate code to obtain the address of the method named **Write** from the type descriptor, then make an indirect call to that address. The receiver argument for the call is a copy of the interface's dynamic value.

Two interfaces are equal if :

* both are nil
* of if their dynamic types are identical and their dynamic values are equal according to the usual behavior of == for that type.(if that type is not comparable, comparing the two interfaces will panic. **Only compare interface values if you are certain that they contain dynamic values of comparable types**)

A slice can contains itself:

```go
s := []interface{}("a", nil)
s[1] = s
```

### 7.5.1 Caveat : An interface contains a nil pointer is Non-nil

```go
var nilBuf *bytes.Buffer
if (false){
    nilBuf = new(bytes.Buffer)
}
f(nilBuf)
func f(out io.Writer) {
    if out != nil{
        out.Write([]byte("hello"))
    }
}
```

The code above will cause the program to panic. Whereas the dynamic value of out is nil, its dynamic type is *bytes.Buffer, a non-nil interface contains a nil pointer. `out != nil` is true. The solution is change nilBuf's type to io.Writer.

## 7.6 Sorting with sort.Interface

```go
package sort
type Interface interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}
type reverse struct {Interface}
func (r reverse) Less(i, j int) bool {return r.Interface.Less(j, i)}
func Reverse(data Interface) Interface {return reverse(data)}

sort.Sort(sort.Interface)
```

## 7.7 the http.Handler Interface

A ServeMux aggregates a collection of http.Handlers into a single http.Handler.

```go
mux := http.NewServeMux()
mux.Handle("/list", http.HandlerFunc(db.list))
mux.Handle("/price", http.HandlerFunc(db.price))

package http
type HandlerFunc func(w ResponseWriter, r *Request)
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
    f(w, r)
}
```

**HandleFunc** is thus an adapter that lets a function value satisfy an interface, where the function and the interface's sole method have the same signature.

```go
mux.HandleFunc("/list", db.list)
mux.HandleFunc("/price", db.price)
```

**net/http** provides a global **ServeMux** instance called DefaultServeMux and  function **http.Handle** and **http.HandleFunc**. To use DefaultServeMux, we should pass nil to ListenAndServe

```go
http.HandleFunc("/list", db.list)
http.HandleFunc("/price", db.price)
http.ListenAndServe("0.0.0.0:8080", nil)
```

## 7.8 the error Interface

```go
type error interface {
    Error() string
}

packages errors
func New(text string) error {return &errorString{text}}
type errorString struct {text string}
func (e *errorString) Error() string {return e.text}

errors.New("EOF") == errors.New("EOF")//false
```

The underlying type of errorString is a struct, not a string, to protext its representation from inadvertent (or premeditated) updates.

```go
var err error = syscall.Errno(2)
```

The code above creates an interface holding the Errno value 2(type=syscall.Errno,value=2).

## 7.9 Example : Expression Evaluator

## 7.10 Type assertions

A type assertion is an operation applied to an interface value. x.(T)