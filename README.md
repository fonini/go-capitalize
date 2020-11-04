# go-capitalize #

go-capitalize is a Go library for capitalizing proper names.

## Usage ##

```go
import "github.com/fonini/go-capitalize/capitalize"
```

The capitalize package exports a single `Capitalize()` function. A simple example:

```go
name, _ := capitalize.Capitalize("jonnas fonini")
fmt.Print(name) // will output: "Jonnas Fonini"

name, _ = capitalize.Capitalize("ruA DA praça xv")
fmt.Print(name) // will output: "Rua da Praça XV"
```