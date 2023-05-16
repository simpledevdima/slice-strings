# Slice-strings
The package is a function to manage a slice of data containing strings.

## Installation
```
go get github.com/simpledevdima/slice-strings
```

## Functions
- Chuck - Chunk splits an array into parts;
- CountValues - counts the number of all slice values.


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

### CountValues
```go
CountValues (slice []string) map[string]int
```
// CountValues counts the number of all slice values.<br />
// Returns a map with indexes of slice values and values indicating the number of items found.<br />

### CountValues example
```go
s := []string{"apple", "body", "cloud", "apple", "apple", "cloud"}
cnt := slice.CountValues(s)
fmt.Println(cnt)
```

Result:
```
map[apple:3 body:1 cloud:2]
```