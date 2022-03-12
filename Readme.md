# Gollection

A package implement Queue, Stack and List like C# in go.



## Supported go versions
1.17.0 or higher



## Install
To install Gollection, use go get:
```shell
go get -u github.com/yanun0323/gollection
```



## Overview
- [Queue](#Queue) 
    - [Constructor](#Constructor-1)
    - [Enqueue](#Enqueue)
- [Stack](#Stack) 
- [List](#List) 
## Example
### Queue: 
* #### Constructor 
```go
    /* Queue is an implement of IQueue */
    /* Create an empty Queue */
    q := NewQueue()
```
* #### Enqueue
```go
    /* Use `ok` to check whether the Enqueue function succeed */
    ok := q.Enqueue("Hello World")
    if !ok {
        panic()
    }

    /* You can Enqueue any type of object */
    q.Enqueue(20)           //support int object
    q.Enqueue(nil)          //support nil object
    q.Enqueue(&Person{})    //support custom stuct
```
* #### Dequeue
```go
    /* Use `ok` to check whether the Dequeue succeed */
    obj, ok := q.Dequeue()
    if !ok {
        panic()
    }

    obj             // Use interface{} directly
    obj.(int)       // Transform interface{} to int type
    obj.(string)    // Error because obj isn't string
```
* #### Clear
```go
    q.Clear()
```
* #### Clone
```go
    clone := q.Clone()

    obj = q.Dequeue().(int)
    obj2 = clone.Dequeue().(int)

    fmt.Println(&obj == &obj2)   //true
```
* #### Contain
```go
    Contains()
```
* #### ContainAny
```go
    ContainsAny(...T) bool
```
* #### Count
```go
    Count() int
```
* #### IsEmpty
```go
    IsEmpty() bool
```
* #### Peek
```go
    Peek() T
```
* #### ToArray
```go
    ToArray() []T
```

### Stack: 
* #### Constructor 
```go
    /* Queue is an implement of IQueue */
    /* Create an empty Queue */
    q := NewQueue()
```
```go
NewStack()

Clear()
Clone() IStack
Contains(T) bool
ContainsAny(...T) bool
Count() int
IsEmpty() bool
Peek() T
Pop() T
Push(T)
ToArray() []T
```
### List: 
```go
NewList()

ADD(...T)
At(int) T
Clear()
Clone() IList
Contains(T) bool
ContainsAny(...T) bool
Count() int
Insert(int, ...T) bool
IsEmpty() bool
Remove(T) bool
RemoveAll(T) bool
RemoveAt(int) bool
Set(int, T) bool
ToArray() []T
```



License
---

© Yanun, 2022 ~ now

Released under the MIT License


