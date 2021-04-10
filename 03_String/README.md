# STRING
A string is an immutable sequence of bytes. Strings can be created by enclosing a set of characters inside double quotes " ".

# Delcaring String
```go
	var fname string
	fname = "MANOJ"
	lname := "CHANDRAN"
```

# Traversing String
Same as the ARRAY...
```go
    s:= "MANOJ"
    for i := 0; i < len(s); i++ {
        fmt.Println(string(s[i]))
    }
    for _, c := range s {
        fmt.Println(string(c))
    }	
```
# Output
```go
    s := "hello world"
    fmt.Println(s) // print hello world
    fmt.Println(s[:]) // print hello world
    fmt.Println(s[0:5]) // print hello
```