# lightning-go

- Practising go ;)


## Main commands

- Following are a table of commonly used go commands.

| CLI Command | Description |
| ----- | ----- |
| `go build` | Compiles a bunch of go source code files and creates a executable file |
| `go run` | Compiles and executes one or two files |
| `go fmt` | Formats all the code in the each file in the current directory |
| `go install` | Compiles and installs a package |
| `go get` | Downloads the raw source code of someone else's package |
| `go test` | Runs any tests associated with the current project |


## Packages

- `main` is reserved keyword in GO.
- `package main` this is used when we want to create an executable package.
- `package rest` building this won't give an executable package. Reusable/Dependenc packages.
- The main package should contain a function called `main` as well.


## Basic go data types

| Type | Example |
| ----- | ----- |
| `bool` | `true` , `false` |
| `string` | "Hi" , "How is it going?" |
| `int` | 0 , -1000 , 99999 |
| `float64` | 10.00001 , 0.00009 , -100.03 |


## Functions

- Files in the same package can freely call functions in other files
- The command should contain all the files
    - `go run main.go deck.go`
    - `go run *.go`


## Data Structures

- Arrays:
    - Fixed length list of data.
    - More basic primitive type.
- Slices:
    - A list that can grow or shrink.
    - A type which have more features.
- Both types must be declared with a data type; string, int etc..


## Looping over a data set

- `for` loop explained;
    - `index` == index of this element in the array
    - `card` == Current card we are iterating over
    - `range deckCards` == Take the slice of `deckCards` and loop over it

```go
for index, card := range deckCards {
	fmt.Println(index, card)
}
```

- Basically there are 4 paterns in golang using `for` loop [construct](https://yourbasic.org/golang/for-loop/).

```go
// For-each range loop
for index, card := range deckCards {
	fmt.Println(index, card)
}

// Three component loop
for index := 0; index < 10; index++ {
    fmt.Println(index)
}

// Infinite loop
for {
    fmt.Println("hello")
}

// While loop
for n < 5 {
    fmt.Println("hello")
}
```

- How to exit a loop

```go
for index := range 10 {
    if index%2 != 0 {
        continue
    }
    fmt.Println(index)
}
```

> [!IMPORTANT]
> In golang if we declare a variable we must use it any point.
> If we doesn't need it but it should be declared we can replace it with a *_* (underscore).
> In loops where the range is used the first variable is always used for index, hence we can replace it with an underscore to ignore it.


## Receiver Function

- `func (d deck) print()` : means that any variable has type deck gets access to print method

```go
func (d deck) print() {
	for i, card := range d { // the d here is similar to this/self
		fmt.Println(i, card)
	}
}
```


## Slices

- Slices are zero-indexed.
- Slices can be used for slicing up.

```go

fruits := []string{"apple", "banana", "grape", "orange"}

fruits[startIndexInclusive : upToNotInclusive]

fruits[0:2] == "apple", "banana"

```


## 'error' Object

- This error object will be populated if some error to happen while the execution, otherwise it will have a `nil` value.
- `nil` object defines a value of nothing, a `null` value.


## Go Testing

- For testing need to create a `go.mod` file
- `go mod init <project_name>`


## Struct

- A collection of different properties.
- Similar to a dictionary.


## Pass by Value

- When a func gets a data structure, it makes a copy of that value and then execute the function onto that copy.
- _Pass by Value_ languages does this.


## Reference vs Value Types

- Value Types:
    - Passes the direct value of the data structure to a function.
    - `int`, `float`, `string`, `bool`, `structs`
- Reference Types:
    - Passes a reference to the actual data value stored memory address.
    - `slices`, `maps`, `channels`, `pointers`, `functions`

```go
package main
 
import "fmt"
 
func main() {
 name := "bill"
 
 namePointer := &name
 
 fmt.Println(&namePointer)
 printPointer(namePointer)
}
 
func printPointer(namePointer *string) {
 fmt.Println(&namePointer)
}
```

- Here, _the log statements will print different addresses because **everything** in go is pass by value_.


## Maps

```go
var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
b := map[KeyType]ValueType{key1:value1, key2:value2,...}
c := make(map[int]string)
```

### Maps vs Structs

| Map | Struct |
| ----- | ----- |
| All keys must be same type, All values must be same type | Values can be different type |
| Keys are indexed - can iterate over | Keys don't support indexing |
| Reference type | Value type |
| Use tp represent a collection of related properties | Use to represent a "collection" with a lot of different properties |
| No need to know all the keys at compile time | Need to know all the fields at compile time |


## Interfaces

- Interfaces are declared as a new data structure with function as fields
- Any function that is defined to use the parameter defined as the type of the interface type, that function can use/call any member functions of that interface type.

