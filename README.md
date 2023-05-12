# Slice-strings
The package is a function to manage a slice of data containing strings.

## Installation
```
go get github.com/simpledevdima/slice-strings
```

## Functions
- Chuck - Chunk splits an array into parts;


### Chuck
```go
Chuck (slice []string, length int, flags ...bool) ([][]string, error)
```
Chunk splits an array into parts.<br />
Splits the array into multiple arrays of length elements.<br />
The last array retrieved may contain fewer values than length.<br />
The first argument to flags can be set to the boolean value true to indicate that a new native array should be created for the returned result.


### Chuck example
```go
s := []string{"apple", "body", "cloud", "diamond", "each"}
chunk, err := slice.Chunk(s, 2, true)
if err != nil {
    log.Fatalln(err)
}
fmt.Println(chunk)
```

Result:
```
[[apple body] [cloud diamond] [each]]
```