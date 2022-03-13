# Gollection

Gollection implements Queue, Stack and List like C# in go.



## Supported go versions
1.17.0 or higher



## Install
To install Gollection, use go get:
```shell
go get -u github.com/yanun0323/gollection
```



## Overview
- [Queue](#Queue)
    - [Constructor](#Queue)
    - [Enqueue](#queueEnqueue)
    - [Dequeue](#queueDequeue)
    - [Clear](#queueClear)
    - [Clone](#queueClone)
    - [Contain](#queueContain)
    - [Count](#queueCount)
    - [IsEmpty](#queueIsEmpty)
    - [Peek](#queuePeek)
    - [ToArray](#queueToArray)
- [Stack](#Stack)
    - [Constructor](#Stack)
    - [Push](#stackPush)
    - [Pop](#stackPop)
    - [Clear](#stackClear)
    - [Clone](#stackClone)
    - [Contain](#stackContain)
    - [Count](#stackCount)
    - [IsEmpty](#stackIsEmpty)
    - [Peek](#stackPeek)
    - [ToArray](#stackToArray)
- [List](#List)
    - [Constructor](#List)
    - [Add](#listAdd)
    - [At](#listAt)
    - [Clear](#listClear)
    - [Clone](#listClone)
    - [Contains](#listContains)
    - [Count](#listCount)
    - [Insert](#listInsert)
    - [IsEmpty](#listIsEmpty)
    - [Remove](#listRemove)
    - [RemoveAll](#listRemoveAll)
    - [RemoveAt](#listRemoveAt)
    - [Set](#listSet)
    - [ToArray](#listToArray)


## Queue
Represents a first-in, first-out collection of objects.
Max stored element quantity is 2,147,483,647

### Constructor
Initializes a new instance of the Queue class that is empty and has the default initial capacity.
`O(1)`
```go
    /* Create an empty Queue */
    q := NewQueue()
    /* Create an Queue has objects */
    q := NewQueue("Hello", "World", "!")
```
### *queue*.Enqueue
Adds an object to the end of the Queue.
`O(1)`
```go
    q := NewQueue()
    q.Enqueue("Hello World")

    /* You can Enqueue any type of object */
    q.Enqueue(20)           // support int object
    q.Enqueue(nil)          // support nil object
    q.Enqueue(&Person{})    // support custom structure
```
### *queue*.Dequeue
Removes and returns the object at the beginning of the Queue.
Panic when the Queue is empty.
`O(1)`
```go
    q := NewQueue("Hello World")
    obj := q.Dequeue()

    obj             // Use interface{} directly
    obj.(string)    // Transform interface{} to string
    obj.(int)       // Error because obj isn't an int

    q.Dequeue()     // Panic!!! because the queue is empty
```
### *queue*.Clear
Removes all objects from the Queue.
`O(1)`
```go
    q.Clear()
```
### *queue*.Clone
Clone the Queue without clone the objects inside the Queue.
O(1)
```go
    q := NewQueue(10, 20)
    clone := q.Clone()
```
### *queue*.Contain
Determines whether any element is in the Queue.
`O(n)`
```go
    q := NewQueue(10, 20)

    answer := q.Contains(10)        // True
    answer := q.Contains(10, 30)    // True
    answer := q.Contains(30, 40)    // False
```
### *queue*.Count
Gets the number of elements contained in the Queue.
`O(1)`
```go
    q := NewQueue(10, 20)
    count := q.Count()              // 2
```
### *queue*.IsEmpty
Return true when the Queue is empty.
`O(1)`
```go
    q := NewQueue(10, 20)
    empty := q.IsEmpty()            // False
```
### *queue*.Peek
Returns the object at the beginning of the Queue without removing it.
Panic when the Queue is empty.
`O(1)`
```go
    q := NewQueue(10, 20)
    num := q.Peek()                 // 10

```
### *queue*.ToArray
Copies the Queue to a new slice.
Return nil when the Queue is empty.
`O(n)`
```go
    q := NewQueue(10, 20)
    arr := q.ToArray()              // {10 ,20}

    empty := NewQueue()
    arr = empty.ToArray()           // Nil 
```

## Stack
Represents a variable size last-in-first-out (LIFO) collection of instances of the same specified type.
Max stored object quantity is 2,147,483,647.

### Constructor
Initializes a new instance of the Stack class that is empty and has the default initial capacity.
`O(1)`
```go
    /* Create an empty Stack */
    s := NewStack()
    /* Create an Stack has objects */
    s := NewStack("Hello", "World", "!")
```
### *stack*.Push
Inserts an object at the top of the Stack.
`O(1)`
```go
    s := NewStack()
    s.Push("Hello World")

    /* You can Push any type of object */
    s.Push(20)           // support int object
    s.Push(nil)          // support nil object
    s.Push(&Person{})    // support custom structure
```
### *stack*.Pop
Removes and returns the object at the top of the Stack.
Panic when the Stack is empty.
`O(1)`
```go
    s := NewStack("Hello World")
    obj:= s.Pop()

    obj             // Use interface{} directly
    obj.(string)    // Transform interface{} to string
    obj.(int)       // Error because obj isn't an int in this case

    s := NewStack()
    obj:= s.Pop()   // Panic!!! because the queue is empty
```
### *stack*.Clear
Removes all objects from the Stack.
`O(1)`
```go
   s.Clear()
```
### *stack*.Clone
Clone the Stack without clone the objects inside the Stack.
`O(1)`
```go
    s := NewStack(20, 10)
    clone := s.Clone()
```
### *stack*.Contain
Determines whether any element is in the Stack.
`O(n)`
```go
    s := NewStack(20, 10)

    answer := s.Contains(10)        // True
    answer := s.Contains(10, 30)    // True
    answer := s.Contains(30, 40)    // False
```
### *stack*.Count
Gets the number of elements contained in the Stack.
`O(1)`
```go
    s := NewStack(20, 10)
    count := s.Count()              // 2
```
### *stack*.IsEmpty
Return true when the Stack is empty.
`O(1)`
```go
    s := NewStack(20, 10)
    empty := s.IsEmpty()            // False
```
### *stack*.Peek
Returns the object at the top of the Stack without removing it.
`O(1)`
```go
    s := NewStack(20, 10)
    num := s.Peek()                 // 10

```
### *stack*.ToArray
Removes and returns the object at the top of the Stack.
Panic when the Stack is empty.
`O(1)`
```go
    s := NewStack(20, 10)
    arr := s.ToArray()              // {10 ,20}

    empty := NewStack()
    arr := empty.ToArray()          // Nil
```
## List
Represents a strongly typed list of objects that can be accessed by index. Provides methods to search, sort, and manipulate lists.
Max stored element quantity is 2,147,483,647

### Constructor
Initializes a new instance of the List class that is empty and has the default initial capacity.
`O(1)`
```go
    /* Create an empty List */
    l := NewList()
    /* Create an List has objects */
    l := NewList("Hello", "World", "!")
```
### *list*.Add
Adds an object to the end of the List.
`O(1)`
```go
    l := NewList()

    l.Add(10)
    l.Add(20, 30, 40)
```
### *list*.At
Gets the element at the specified index.
Panic when the index of the List is empty.
`O(1)`
```go
    l := NewList(10, 20)

    obj := l.At(0)                  // 10
    obj := l.At(-1)                 // Panic!!! because the index is out of bounds
    obj := l.At(5)                  // Panic!!! because the index is out of bounds
```
### *list*.Clear
Removes all elements from the List.
`O(1)`
```go
    l.Clear()
```
### *list*.Clone
Clone the List without clone the objects inside the List.
`O(n)`
```go
    l := NewList(10, 20)
    clone := l.Clone()
```
### *list*.Contains
Determines whether any element is in the List.
`O(n)`
```go
    l := NewList(10, 20)

    answer := l.Contains(10)         // True
    answer := l.Contains(10, 30)     // True
    answer := l.Contains(30, 40)     // False
```
### *list*.Count
Gets the number of elements contained in the List.
`O(1)`
```go
    l := NewList(10, 20)
    count := l.Count()              // 2
```
### *list*.Insert
Inserts any element into the List at the specified index.
Panic when the index of the List is empty.
`O(n)`
```go
    l := NewList("E", "F")

    l.Insert(0, "D")
    l.Insert(0, "A", "B", "C")
    l.Insert(-1, "A")   // Panic!!! because the index is out of
    l.Insert(10, "A")   // Panic!!! because the index is out of

```
### *list*.IsEmpty
Return true when the List is empty.
`O(1)`
```go
    l := NewList(10, 20)
    empty := l.IsEmpty()            // False
```
### *list*.Remove
Removes the first object at the List.
`O(n)`
```go
    l := NewList(10, 20)
    l.Remove(10)
```
### *list*.RemoveAll
Removes all the elements that match the conditions defined by the specified predicate.
`O(n)`
```go
    l := NewList(10, 20)
    l.RemoveAll(10)
```
### *list*.RemoveAt
Removes the element at the specified index of the List.
Panic when the index of the List is empty.
`O(n)`
```go
    l := NewList(10, 20)
    obj := l.RemoveAt(0)            // 10
```
### *list*.Set
Sets the element at the specified index.
Panic when the index of List is empty.
`O(1)`
```go
    l := NewList("Hello", " ", "World")
    l.Set(0, "Hi")

    l.Set(-1, "Hi")     // Panic!!! because the index is out of bounds
    l.Set(5, "Hi")      // Panic!!! because the index is out of bounds
```
### *list*.ToArray
Copies the elements of the List to a new slice.
Return nil when the List is empty.
`O(n)`
```go
    l := NewList(10, 20)
    arr := l.ToArray()              // {10 ,20}

    empty := NewQueue()
    arr = empty.ToArray()           // Nil 
```



License
---

© Yanun, 2022 ~ now

Released under the MIT License


