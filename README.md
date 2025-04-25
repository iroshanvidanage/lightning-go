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

```go
type bot interface {
	getGreeting(string, int) (string, error)
	getBotVersion() float64
	respondToUser(user) string
}
```

- Interfaces are not generic types:
    - Go does not have _generic_ types.
- Interfaces are `implicit`: 
    - Automatically creates the link between interface and the othertypes from the defined functions.
- Interfaces are a contract to help us manage types.
- Interfaces are tough. Step #1 is understanding how to read them:
    - Understand how to read interfaces in the standard lib.


### Concrete Type vs Interface Type

| Concrete Type | Interface Type |
| ----- | ----- |
| Types you can create direct vaules of it. | Cannot create values directly out of these |
| `map`, `struct`, `int`, `string`, `englishBot` | `bot` |


## Concurrent Programming

- Allows multiple computations to execute in an overlapping time period, not necessarily simultaneously, by managing access to shared resources.
- It's a technique where two or more processes can start, run in an interleaved fashion through context switching, and complete within overlapping time periods.
- This is different from parallel programming, which involves tasks running truly simultaneously on multiple processors.


### Concurrency vs. Parallelism:

- Both involve multiple tasks, concurrency focuses on the overlapping execution of tasks, even on a single processor, while parallelism requires multiple processors to truly run tasks simultaneously. 


### Context Switching

- Concurrency on a single processor is achieved through context switching, where the CPU rapidly switches between different threads, creating the illusion of parallel execution. 

> [!NOTE]
>
> **Benefits**:
> - Concurrent programming can improve application performance by allowing tasks to overlap and make progress even when certain parts are blocked, for example, while waiting for user input.
>
> **Complexity**:
> - Concurrent programs can be complex due to the need to manage shared resources and avoid issues like race conditions and deadlocks.
>
> **Examples**:
> - Operating systems and database management systems are examples of concurrent systems.


## Go Routines

- Routines execute our Go code line by line. We can use go routines to launch functions into new routines so that the previous routines is continues while the second is starting.
- Placing `go` keyword before the code block you want to execute will launch a new routine for that code to be executed.
- Main routine will finish and exit even while the child routines are still in running.


## Channels

- Channels can be used to communicate / interact between the main and child go routines.
- Created as a variable and it's tight, meaning it's explict for one type (string/float/int).
- Sending data with channels.

| Syntax | Description |
| ----- | ----- |
| `channel <- 5` | Send the value 5 into this channel |
| `myNumber <- channel` | Wait for a value to be sent into the channel. When we get one, assign the value to *myNumber*. |
| `fmt.Println(<- channel)` | Wait for a value to be sent into the channel. When we get one, log it out immediately. |