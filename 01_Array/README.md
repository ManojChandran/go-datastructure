# ARRAY
An array is a collection of homogeneous (same type) data items stored in contiguous memory locations.

Declare an integer array of five elements.
```go
  var array [5] int
```

Declaring an array initializing specific elements
```go
  names := [3]string{"Manoj", "Archana", "Aami"}
```

Declaring an array and Go automatically calculating size, "len()" gives size of the array.
```go
  grades := [...]int{97,85,93}
  fmt.Println(grades, "len :", len(grades))
```

Go array's are considered as values, not like other language. When you create an Array its getting pointing to that. Consider the below example for reference.
```go
	a := [...]int{1, 2, 3}
	b := a  // value gets actually copied
	c := &a // c is pointing
	b[1] = 5
	c[1] = 0
	fmt.Println(a) // prints 1,2,3
	fmt.Println(b) //prints 1,5,3
	fmt.Println(c) //prints 1,0,3
```

Since go consider array as value, we can compare array.If an array's element type is comparable then the array type is comparable too and Array size should be same.
```go
	array1 := [3]int{1, 2, 3}
	array2 := [3]int{1, 2, 3}
	fmt.Println("array1 == array2:", array1 == array2)
```
Multi dimensional array
```go
  var identityMatrix [3][3]int = [3][3]int{ [3]int{1,0,0}, [3]int{0,1,0}, [3]int{0,0,1} }
```

# OUTPUT
Output for go program for reference.
```go
[0 0 0 0 0]
[Manoj Archana Amaya]
[97 85 93] len : 3
[1 0 3]
[1 5 3]
&[1 0 3]
array1 == array2: true
[[1 0 0] [0 1 0] [0 0 1]]
```

# TRAVERSING AN ARRAY

There are four ways of iterating through the ARRAY elements. 

#### Example 1
```go
	intArray := [5]int{10, 20, 30, 40, 50}

	for i := 0; i < len(intArray); i++ {
		fmt.Println(intArray[i])
    }
```
#### Example 2
```go
    intArray := [5]int{10, 20, 30, 40, 50}

	for index, element := range intArray {
		fmt.Println(index, "=>", element)

	}
```
#### Example 3
```go
    intArray := [5]int{10, 20, 30, 40, 50}
	
	for _, value := range intArray {
		fmt.Println(value)
    }
```

#### Example 3
```go
    intArray := [5]int{10, 20, 30, 40, 50}
		
	j := 0
	for range intArray {
		fmt.Println(intArray[j])
		j++
    }
```