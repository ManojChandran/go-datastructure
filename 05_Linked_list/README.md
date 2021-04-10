# LINKED LIST
A linked list is a linear collection of data elements whose order is not given by their physical placement in memory. Instead, each element points to the next. It is a data structure consisting of a collection of nodes which together represent a sequence. In its most basic form, each node contains: data, and a reference (in other words, a link) to the next node in the sequence. 

# Declare 

Linked list is a list of nodes, each node stores a data element and an address pointer to next node. First thing is to define the node and list. (We are talking about single linked list)
node 
```go
type Node struct {
    data int              // Data, it can be of any type
    next *Node            // Reference to next node
}
```

now lets define the list, list will have Head and tail to hold the startng of the list and node pointer.

```go
type llist struct {
	head *Node     // head of the node 
	tail *Node     // current node 
}
```

# Traverse the list

Example 1:
```go
// List all values
func (l llist) liatAll() {
	for l.tail = l.head; l.tail != nil; l.tail = l.tail.next {
		fmt.Println(l.tail.data)
	}
}
```go