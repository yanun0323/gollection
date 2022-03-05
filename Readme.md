# Gollection

A package implement Queue, Stack and List like C# in go.



## Requirements
Go 1.11.0 or higher



## Install
```
$go get github.com/yanun0323/gollection
```



## Overview
### Queue: 
     NewQueue()

     Count

     Clear()
     Clone() IQueue
     Contains(T) bool
     ContainsAny(...T) bool
     Dequeue() T
     Enqueue(T)
     IsEmpty() bool
     Peek() T
     ToArray() []T

### Stack: 
    NewStack()

    Count

	Clear()
	Clone() IStack
	Contains(T) bool
	ContainsAny(...T) bool
	IsEmpty() bool
	Peek() T
	Pop() T
	Push(T)
	ToArray() []T

### List: 
     NewSList()

     Count
     
	 ADD(...T)
	 At(int) T
	 Clear()
	 Clone() IList
	 Contains(T) bool
	 ContainsAny(...T) bool
	 Insert(int, ...T) bool
	 IsEmpty() bool
	 Remove(T) bool
	 RemoveAll(T) bool
	 RemoveAt(int) bool
	 Set(int, T) bool
	 ToArray() []T



License
---

© Yanun, 2022 ~ now

Released under the MIT License

