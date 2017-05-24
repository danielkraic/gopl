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

# composite types

## arrays

```go
var a1 [3]int
var a2 [3]int = [3]int{1,2,3}
var a3 := [...]int{1,2,3}
```

```go
type Currency int
const (
        USD Currency = iota
        EUR
        GBP
)
symbol := [...]string{USD: "$", EUR: "e", GPG: "L"}
fmt.Printf("%s", symbol[EUR])
```

```go
a := [...]int{99: 0} //array of 100 items
```

```go
a1 := [2]int{1,2}
a2 := [..]int{1,2}
a3 := [..]int{1,2,3}
isSame1 := a1 == a2 // true
isSame2 := a1 == a3 // false
```

```go
func zero(arr *[32]byte) {
        for i := range arr {
                arr[i] = 0
        }
}
```

## slices

```go
a := []string{"a", "b"}
a = append(a, "c")
a = append(a, "d", "e")
```

```go
a := make([]string, 0, 2)
fmt.Printf("%v, len %d cap %d\n", a, len(a), cap(a))
```

```go
months := []string{"jan", "feb" /* ... */}
Q1 := months[1:4]
Q2 := months[4:7]
Q3 := months[7:10]
Q4 := months[10:]
```

```go
func nonempty(strings []string) []string {
	i := 0
	for _, val := range strings {
		if val != "" {
			strings[i] = val
			i++
		}
	}
	return strings[:i]
}
```

```go
func remove(strings []string, i int) []string {
	copy(strings[i:], strings[i+1:])
	return strings[:len(strings)-1]
}
```

## maps

```go
m1 := make(map[string]int)

m2 := map[string]int{
    "a": 1,
    "b": 2,
}

m2["c"] = 5
m2["d"] += 10
m2["e"]++

delete(m2, "a")

for k, c := range m2 {
        fmt.Printf("%s: %s\n", k, v)
}
```

```go
if val, ok := m2["somekey"]; ok {
    fmt.Printf("somekey val: %v\n", val)
} else {
    fmt.Printf("no value for somekey\n")
}
```

```go
keys := make([]string, 0, len(m2))
for k := range m2 {
    keys = append(keys, k)
}

sort.Strings(keys)
for _, k := range keys {
    fmt.Printf("%v: %v\n", k, m2[k])
}
```

## structs

```go
type User struct {
        id uint
        Name, Email string
}

var u1 User
u1.name = "N1"

var u1p *User := &u1
u1p.Name = "N2"

pname := &u1.Name
*name = "N3"

u4 := User{10, "N4"}
u5 := User{Name: "N5"}

u6 := &User{Name: "N6"}
```

```go
//string set
seen := make(map[string]struct{})
if _, ok := seen['N1']; !ok {
        seen[s] = struct{}{}
}
```

```go
type Point {x, y int}

func Scale(p *Point, factor int) {
        p.x *= factor
        p.y = p.y * factor         
}
```

```go
// comparing structs
type Point struct {x, y int}
p1 := Point{1, 2}
p2 := Point{2, 1}

isSame := p1 == p2 //false
 
// may be used as a key
 m1 := make(map[Point]int)
```

```go
// struct embedding and anonymous fields
type Point struct {
        X, Y int
}

type Circle {
        Point
        Radius int
}

type Wheel struct {
        Circle
        Spokes int
}

var w1 Wheel
w.X = 10
w.Radius = 11
w.Spokes = 12

w2 := Wheel{
        Circle: Circle{
                Point: Point{X: 10, Y: 11},
                Radius: 12,
        }
        Spokes: 13,
}
w3 := Wheel{Circle{Point{10,11}, 12}, 13}
```

## json

```go
type Movie struct {
	Title string
	Year int `json:"released"`
	Color bool `json:"color,omitempty"`
	Actors []string
}

data1, err1 := json.Marshal(&movies)
data2, err2 := json.MarshalIndent(&movies, "", " ")
var data3 []Movie
err3 := json.Unmarshal(data2, &data3)
```