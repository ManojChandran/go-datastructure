# LINKED LIST
A linked list is a linear collection of data elements whose order is not given by their physical placement in memory. Instead, each element points to the next. It is a data structure consisting of a collection of nodes which together represent a sequence. In its most basic form, each node contains: data, and a reference (in other words, a link) to the next node in the sequence. 

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