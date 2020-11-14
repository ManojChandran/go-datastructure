# SLICE

A Slice is just like an array, which is a container to hold elements of same data type. But, Slice can vary in size.

Declaring Slice
```go
    var strSlice []string
    intSlice := []int{1,2,3}
```

":" operator can be used to slice out elements from the Slice
```go
  row0 := []int{1,2,3,4,5,6,7,8,9,0}
  fmt.Println("row0 :",row0) // row0 : [1 2 3 4 5 42 7 8 9 0]
  row1 := row0[:] //slice of all element
  fmt.Println("row1 :",row1) // row1 : [1 2 3 4 5 42 7 8 9 0]
  row2 := row0[3:] //slice from 4th element to end
  fmt.Println("row2 :",row2) // row2 : [4 5 42 7 8 9 0]
  row3 := row0[:6] // slice first 6 element
  fmt.Println("row3 :",row3) // row3 : [1 2 3 4 5 42]
  row4 := row0[3:6] // slice the 4th,5th,6th elements
  fmt.Println("row4 :",row4) // row3 : [1 2 3 4 5 42]
  row0[5] = 42 // changes 5th element value in all rows
```

# Functions
copy function
```go
	src := []string{"Sublime", "VSCode", "IntelliJ", "Eclipse"}
  dest := make([]string, 2)
  // func copy(dst, src []T) int
  numElementsCopied := copy(dest, src)
```

Adding elements
```go
  slice1 := []string{"C", "C++", "Java"}
  // func append(s []T, x ...T) []T
  slice2 := append(slice1, "Python", "Ruby", "Go")
```  

# Length and Capacity
A slice has both a length and a capacity. The length of a slice is the number of elements it contains. The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
```go
package main

import (
  "fmt"
)

func main() {
  a := []int{}
  fmt.Println(a)
  fmt.Printf("length : %v\n ", len(a))
  fmt.Printf("capacity : %v\n ", cap(a))
  a = append(a,1) // slice capacity changes
  fmt.Println(a)
  fmt.Printf("length : %v\n ", len(a))
  fmt.Printf("capacity : %v\n ", cap(a))
  a = append(a,2,3,4,5) // slice capacity changes
  fmt.Println(a)
  fmt.Printf("length : %v\n ", len(a))
  fmt.Printf("capacity : %v\n ", cap(a))
  a = append(a,[]int{6,7,8,9}...) // slice capacity changes 
  fmt.Println(a)
  fmt.Printf("length : %v\n ", len(a))
  fmt.Printf("capacity : %v\n ", cap(a))
}
```

output 
```go
[]
length : 0
 capacity : 0
 [1]
length : 1
 capacity : 1
 [1 2 3 4 5]
length : 5
 capacity : 6
 [1 2 3 4 5 6 7 8 9]
length : 9
 capacity : 12
```

# Traversing Slice
```go
package main 
import "fmt"

func main() { 
    s1 := []int{23, 56, 89, 34} 
    for i, d := range s1{
     fmt.Println(i, "=",d) 
    }
} 
```