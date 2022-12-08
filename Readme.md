# Gollection

Gollection implements Queue, Stack and List like C# in go.



## Supported go versions
1.18.0 or higher



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
    - [Len](#queueLen)
    - [Peek](#queuePeek)
    - [ToSlice](#queueToSlice)
- [Stack](#Stack)
    - [Constructor](#Stack)
    - [Push](#stackPush)
    - [Pop](#stackPop)
    - [Len](#stackLen)
    - [Peek](#stackPeek)
    - [ToSlice](#stackToSlice)


## Queue
Represents a first-in, first-out collection of objects.<br/>
Max stored element quantity is 2,147,483,647

### Constructor
Initializes a new instance of the Queue class that is empty and has the default initial capacity.
```go
    /* Create an empty Queue */
    q := NewQueue()
```
### *queue*.Enqueue
Adds an object to the end of the Queue.
```go
    q := NewQueue()
    q.Enqueue("Hello World")

    /* You can Enqueue any type of object */
    q.Enqueue(20)           // support int object
    q.Enqueue(nil)          // support nil object
    q.Enqueue(&Person{})    // support custom structure
```
### *queue*.Dequeue
Removes and returns the object at the beginning of the Queue.<br/>
Returns `nil` when the Queue is empty.
```go
    q := NewQueue()
    q.Enqueue("Hello World")
    obj := q.Dequeue()

    obj             // Use interface{} directly
    obj.(string)    // Transform interface{} to string
    obj.(int)       // Error because obj isn't an int

    q.Dequeue()     // nil
```
### *queue*.Len
Gets the number of elements in the Queue.
```go
    q := NewQueue()
    q.Enqueue(10, 20)
    length := q.Len()              // 2
```
### *queue*.Peek
Returns the object at the beginning of the Queue without removing it.<br/>
Returns `nil` when the Queue is empty.
```go
    q := NewQueue()
    q.Enqueue(10, 20)
    num := q.Peek()                 // 10

```
### *queue*.ToSlice
Copies the Queue to a new slice.<br/>
Returns empty slice when the Queue is empty.
```go
    q := NewQueue()
    q.Enqueue(10, 20)
    arr := q.ToArray()              // {10 ,20}

    empty := NewQueue()
    arr = empty.ToSlice()           // {} 
```

## Stack
Represents a variable size last-in-first-out (LIFO) collection of instances of the same specified type.<br/>
Max stored object quantity is 2,147,483,647.

### Constructor
Initializes a new instance of the Stack class that is empty and has the default initial capacity.
```go
    /* Create an empty Stack */
    s := NewStack()
```
### *stack*.Push
Inserts an object at the top of the Stack.
```go
    s := NewStack()
    s.Push("Hello World")

    /* You can Push any type of object */
    s.Push(20)           // support int object
    s.Push(nil)          // support nil object
    s.Push(&Person{})    // support custom structure
```
### *stack*.Pop
Removes and returns the object at the top of the Stack.<br/>
Returns `nil` when the Stack is empty.
```go
    s := NewStack()
    s.Push("Hello World")
    obj:= s.Pop()

    obj             // Use interface{} directly
    obj.(string)    // Transform interface{} to string
    obj.(int)       // Error because obj isn't an int in this case

    s := NewStack()
    obj:= s.Pop()   // nil
```
### *stack*.Len
Gets the number of elements in the Stack.
```go
    s := NewStack()
    s.Push(20, 10)
    length := s.Len()              // 2
```
### *stack*.Peek
Returns the object at the top of the Stack without removing it.<br/>
Returns `nil` when the Stack is empty.
```go
    s := NewStack()
    s.Push(20, 10)
    num := s.Peek()                 // 10

```
### *stack*.ToSlice
Copies the Stack to a new slice.<br/>
Returns empty slice when the Stack is empty.
```go
    s := NewStack()
    s.Push(20, 10)
    arr := s.ToSlice()              // {10 ,20}

    empty := NewStack()
    arr := empty.ToSlice()          // {}
```

License
---

© Yanun, 2022 ~ now

Released under the MIT License


