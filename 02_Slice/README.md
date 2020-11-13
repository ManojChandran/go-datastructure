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

# Length and Capacity