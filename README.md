# program structure

## names

### constants

```
true false iota nil
```

### types

```
int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr
float32 float64 complex64 complex128
bool byte rune string error
```

### functions

```
make len cap new append copy close delete
complex real imag
panic recover
```

## variables

```go
var s string
var i, j int
var a, b, c = true, 10, "str"
var a string = 10

i := 0
i, j := 0, 1
```

### pointers

```go
x := 1
p := &x
fmt.Println(*p)
*p = 2

p2 := new(int)
```

## types

```go
type Celsius float64
c := Celsius(11.3)
```

## program

```go
package somename

import (
    "fmt"
    "io"
)

const (
    //
)

func init() {
    //
}

func main() {
    //
}
```

# basic data types

* operators precedence (decreasing)

```
*  /  %  <<  >>  &  &^
*  -  |  ^
== != <  <=  >  >=
&&
||
```

## complex

```go
x := comples(1, 2)
r := real(x)
m := imag(x)
```

## strings

```go
s1 := "hello"
len(s1)

s2 := `a
b
c`
```

```go
strings.Contains -> bytes.Contains
strings.Count -> bytes.Count
strings.Fields -> bytes.Fields
strings.HasPrefix -> bytes.HasPrefix
strings.Index -> bytes.Index
strings.Join -> bytes.Join
```

## conversions

```go
strconv.Iota(10)
strconv.FormatInt(10, 2 /* base */)
fmt.Sprintf("%b", 63)

strconv.AtoI("123")
strconv.ParseInt("123", 10, 64)
fmt.Scanf()
```

## iota

```go
const (
    Mon = iota
    Tue
    Wed
    //...
)
```

```go
type Flags uint
const (
    Flag1 Flags = 1 << iota
    Flag2
    //...
)
```

```go
const (
    _ = 1 << (10 * iota)
    KiB
    MiB
    //...
)
```