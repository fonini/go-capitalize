# go-capitalize #

[![GoDoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/fonini/go-capitalize/capitalize)
[![Test Status](https://github.com/fonini/go-capitalize/workflows/tests/badge.svg)](https://github.com/fonini/go-capitalize/actions?query=workflow%3Atests)
[![codecov](https://codecov.io/gh/fonini/go-capitalize/branch/master/graph/badge.svg?token=FB25JPH4ED)](https://codecov.io/gh/fonini/go-capitalize)

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