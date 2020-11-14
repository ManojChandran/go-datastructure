# LINKED LIST
A linked list is a linear collection of data elements whose order is not given by their physical placement in memory. Instead, each element points to the next. It is a data structure consisting of a collection of nodes which together represent a sequence. In its most basic form, each node contains: data, and a reference (in other words, a link) to the next node in the sequence. 

node 
```go
type Node struct {
    data int              // Data 
    next *Node            // Reference to next node
}
```