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