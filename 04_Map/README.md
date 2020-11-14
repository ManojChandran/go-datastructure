# MAP
Golang Maps is a collection of unordered pairs of key-value. It is widely used because it provides fast lookups and values that can retrieve, update or delete with the help of keys.

Declaring MAP
```go
  var map1 map[string]int
  map2 := make(map[string]int) // construct => map[key]value{}

```
In the maps, a key must be unique and always in the type which is comparable using == operator or the type which support != operator. So, most of the built-in type can be used as a key like an int, float64, rune, string, comparable array and structure, pointer, etc. The data types like slice and noncomparable arrays and structs or the custom data types which are not comparable don’t use as a map key.


Check a particular value with key.
```go
  pop, ok := m["B"] // check out the value in map
  if ok {
    fmt.Println("value of B =", pop)
  }

  if pop, ok := m["C"]; ok {
    fmt.Println("value of B =", pop)
  }
  
```
Add and delete value from the MAP
```go
  m["C"] = 3001 // single insert
  delete(m, "C") // delete with key value "B"
```

# Traversing MAP
Similar to ARRAY
```go
  for k,v := range m { // iterate map
    fmt.Println(k,"value = ",v)
  }
```