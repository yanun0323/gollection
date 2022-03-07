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
- [Stack](#Stack) 
- [List](#List) 
### Queue: 
```go
NewQueue()

Clear()
Clone() IQueue
Contains(T) bool
ContainsAny(...T) bool
Count() int
Dequeue() T
Enqueue(T)
IsEmpty() bool
Peek() T
ToArray() []T
```

### Stack: 
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
NewSList()

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

