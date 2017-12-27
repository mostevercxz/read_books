# Packages and the Go Tool

## 10.1 Introduction

The purpose of any package system: Making the design and maintenance of large programs practical, by grouping releated features together into units that can be easily understood and changed, independent of the other packages of the system.

Each package defines a distinct name space that encloses its identifiers.

Go's compilation is notably faster than most other compiled languages:

1. all imports must be explicitly listed at the beginning of each source file(the compiler does not need to read and process an entire file to determine its dependencies)
1. the dependencies of a package form a directed acyclic graph.(packages can be compiled separately and perhaps in parallel)
1. the object file for a compiled Go package records export info not just for the package itself, but for its denpendencies too.(the compiler only need to read one object file for each import)

## 10.2 Import Paths

Each package is identified by a unique string called its *import path*. Import paths are the strings that appear in import declarations.

The Go language does not define the meaning of these strings or how to determine a package's import path, but leaves these issues to the tools.

## 10.3 The package declaration

A **package** declaration is required at the start of every Go source file : to determine the default identifier for that package when it is imported by another package.

Conventionally the package name is the last segment of the import path. **math/rand** and **crypto/rand**

Three exception to the "last segment" convention:

1. a package that defining a command(an exectuable Go program) always has the name **main**, regardless of the package's import path. (This is a signal to **go build** that it must invoke the linker to make an exectuable file)
1. Some files in the directory may have the suffix **_test** on their package name if the file ends with **_test.go**.(Section 11.2.4)
1. some tools for dependency management appends version number suffixs to the package import paths, such as **gopkg.in/yaml.v2**. The package name excludes the suffix, so in this case the package name is **yaml**.

## 10.4 Import Declarations

Two import forms:

1. specify the import path of a single package, one line per package
1. specify multiple packages in a parenthesized list as the code below:

```go
import (
    "fmt"
    "os"
)
```

If we need to import two packages with the same name:

```go
import (
    "crypto/rand"
    mrand "math/rand"
)
```

The code above for **mrand** is called **renaming import**. Each import establishes a dependency from the current package to the imported package. **go build** will report an error if these dependencies form a cycle.

## 10.5 Blank Imports

On occasion we must import a package merely for the side effects of doing so: evaluation of the initialzer expressions of its package-level variables and execution of its **init** function. To suppress the "unused import" error, we would use the blank identifier "_". Blank identifier can never be referenced.

This is called a *blank import*. It is most often used to implement a compile-time mechanism whereby the main program can enable optional features by blank-importing additional packages.

## 10.6 Packages and Naming

Be descriptive and unambiguous where possible.

## 10.7 The Go Tool

### 10.7.1  workspace organization

```bash
export GOPATH=$HOME/gobook
export http_proxy="ip:port"
export https_proxy="ip:port"
go get gopl.io/...
```

### 10.7.2 Downloading packages

**go get** can download a single package or an entire subtree or repository using the ... notion. **go get** also download all the dependencies of the initial packages.

Once **go get** has downloaded all the packages, it builds them and installs the libraries and commands.

**go get** creates true clients of the remote repository.  Packages can use a custom domain name in their import path while being hosted by **github.com**.

## 10.7.3 Building packages

* If the package is a library, **go build** merely checks that the package is free of compile errors
* If the package is named **main**,**go build** invokes the linker to create an exectuable in the current directory. The name of the exe is taken from the last segment of the package's import path.

How to specify a package:

1. by import path `gopl.io/ch1/helloworld`
1. by relative path, `./src/gopl.io/ch1/helloworld`
1. if no argument is provided, the current directory is used.
1. specify a list of file names, if the package name is **main**,the exe name comes from the basename of the first .go file.

**go install** is very similar to **go build**, except that it saves the compiled code for each package and command instead of throwing it away.

Build for specific target:

1. If a file name includes an OS or processor architecture name like net_linux.go or asm_amd64.s, the go tool will compile the file only when building for that target
1. *build tags* give more fine-grained control.

```go
// +build linux darwin
package mypkg
// +build ignore
go doc go/build
```

### 10.7.4 Documenting pkgs

### 10.7.5 Internal pkgs

A way to define identifiers that are visiable to a small set of trusted pkgs.(big->small, share utility funcs inside part of a project's files, experiment with a new pkg)

**go build** treats a pkg specially if its import path contains a path segment named **internal**, such pkgs are called *internal packages*

An internal pkg may be imported only by another pkg that is inside the tree rooted at the parent of the internal directory.

### 10.7.6 querying pkgs

**go list** reports info about available pkgs.


```go
go list fmt
go list ...
go list gopl.io/ch3/...
go list ...xml...
go list -json hash
```