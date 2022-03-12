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
    - [Constructor](#Queue)
    - [Enqueue](#queueEnqueue)
    - [Dequeue](#queueDequeue)
    - [Clear](#queueClear)
    - [Clone](#queueClone)
    - [Contain](#Contain)
    - [Count](#Count)
    - [IsEmpty](#IsEmpty)
    - [Peek](#Peek)
    - [ToArray](#ToArray)
- [Stack](#Stack)
    - [Constructor](#Stack)
    - [Push](#stackPush)
    - [Pop](#stackPop)
    - [Clear](#stackClear)
    - [Clone](#stackClone)
    - [Contain](#Contain)
    - [Count](#Count)
    - [IsEmpty](#IsEmpty)
    - [Peek](#Peek)
    - [ToArray](#ToArray)
- [List](#List)

## Queue

### Constructor
```go
    /* Create an empty Queue */
    q := NewQueue()
    /* Create an Queue has objects */
    q := NewQueue("Hello", "World", "!")
```
### *queue* Enqueue
```go
    q := NewQueue()
    /* Use `ok` to check whether the Enqueue succeed */
    ok := q.Enqueue("Hello World")
    if !ok {
        panic()
    }

    /* You can Enqueue any type of object */
    q.Enqueue(20)           //support int object
    q.Enqueue(nil)          //support nil object
    q.Enqueue(&Person{})    //support custom stuct
```
### *queue* Dequeue
```go
    /* Use `ok` to check whether the Dequeue succeed */
    q := NewQueue("Hello World")
    obj, ok := q.Dequeue()
    if !ok {
        panic()
    }

    obj             // Use interface{} directly
    obj.(string)    // Transform interface{} to string
    obj.(int)       // Error because obj isn't an int
```
### *queue* Clear
```go
    q.Clear()
```
### *queue* Clone
```go
    q := NewQueue(10, 20)
    clone := q.Clone()

    num1 = q.Dequeue().(int)
    num2 = clone.Dequeue().(int)

    fmt.Println(num1 == num2)    //True
```
### *queue* Contain
```go
    q := NewQueue(10, 20)

    var answer bool
    answer = q.Contains(10)      //True
    answer = q.Contains(10, 30)  //True
    answer = q.Contains(30, 40)  //True
```
### *queue* Count
```go
    q := NewQueue(10, 20)
    count := q.Count()      //2
```
### *queue* IsEmpty
```go
    q := NewQueue(10, 20)
    q.IsEmpty()             //False
```
### *queue* Peek
```go
    q := NewQueue(10, 20)
    num1 := q.Peek()        //num1 == 10
    num2 := q.Dequeue()     //num1 == 10

```
### *queue* ToArray
```go
    q := NewQueue(10, 20)
    arr := q.ToArray()
    arr[0].(int)            //10
    arr[1].(int)            //20
```

## Stack
### Constructor
```go
    /* Create an empty Stack */
    s := NewStack()
    /* Create an Stack has objects */
    s := NewStack("Hello", "World", "!")
```
### *stack* Push
```go
    s := NewStack()
    /* Use `ok` to check whether the Push succeed */
    ok := s.Push("Hello World")
    if !ok {
        panic()
    }

    /* You can Push any type of object */
    s.Push(20)           //support int object
    s.Push(nil)          //support nil object
    s.Push(&Person{})    //support custom stuct
```
### *stack* Pop
```go
    /* Use `ok` to check whether the Pop succeed */
    s := NewStack("Hello World")
    obj, ok := q.Pop()
    if !ok {
        panic()
    }

    obj             // Use interface{} directly
    obj.(string)    // Transform interface{} to string
    obj.(int)       // Error because obj isn't an int
```
### *stack* Clear
```go
   s.Clear()
```
### *stack* Clone
```go
    s := NewStack(20, 10)
    clone := s.Clone()

    num1 = q.Pop().(int)
    num2 = clone.Pop().(int)

    fmt.Println(num1 == num2)    //True
```
### *stack* Contain
```go
    s := NewStack(20, 10)

    var answer bool
    answer = s.Contains(10)      //True
    answer = s.Contains(10, 30)  //True
    answer = s.Contains(30, 40)  //True
```
### *stack* Count
```go
    s := NewStack(20, 10)
    count := s.Count()      //2
```
### *stack* IsEmpty
```go
    s := NewStack(20, 10)
    s.IsEmpty()             //False
```
### *stack* Peek
```go
    s := NewStack(20, 10)
    num1 := s.Peek()        //num1 == 10
    num2 := s.Pop()         //num1 == 10

```
### *stack* ToArray
```go
    s := NewStack(20, 10)
    arr := s.ToArray()
    arr[0].(int)            //10
    arr[1].(int)            //20
```
## List
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


