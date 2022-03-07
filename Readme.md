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



## Example
### Queue: 
```go
NewQueue()

Clear() bool
Clone() IQueue
Contains(T) bool
ContainsAny(...T) bool
Count() int
Dequeue() T.bool
Enqueue(T) bool
IsEmpty() bool
Peek() T, bool
ToArray() []T, bool
```

### Stack: 
```go
NewStack()

Clear() bool
Clone() IStack
Contains(T) bool
ContainsAny(...T) bool
Count() int
IsEmpty() bool
Peek() T, bool
Pop() T, bool
Push(T) bool
ToArray() []T, bool
```
### List: 
```go
NewSList()

ADD(...T)
At(int) T, bool
Clear() bool
Clone() IList
Contains(T) bool
ContainsAny(...T) bool
Count() int
Insert(int, ...T) bool
IsEmpty() bool
Remove(T) bool
RemoveAll(T) bool
RemoveAt(int) T, bool
Set(int, T) bool
ToArray() []T, bool
```



### License

© Yanun, 2022 ~ now

Released under the MIT License

