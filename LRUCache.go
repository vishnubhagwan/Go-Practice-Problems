package main

import (
	"container/list"
	"fmt"
)

type LRU struct {
	cap   int
	cache map[int]*list.Element
	list  *list.List
}

type entry struct {
	key, val int
}

func NewLRU(cap int) *LRU {
	return &LRU{cap, make(map[int]*list.Element), list.New()}
}

func (l *LRU) Get(key int) int {
	if e, ok := l.cache[key]; ok {
		l.list.MoveToFront(e)
		return e.Value.(entry).val
	}
	return -1
}

func (l *LRU) Put(key, val int) {
	if v, ok := l.cache[key]; ok {
		l.list.MoveToFront(v)
		v.Value = entry{key: key, val: val}
		return
	}
	l.list.PushFront(entry{key: key, val: val})
	l.cache[key] = l.list.Front()
	if l.list.Len() > l.cap {
		v := l.list.Back()
		delete(l.cache, v.Value.(entry).key)
		l.list.Remove(l.list.Back())
	}
}

func main() {
	l := NewLRU(2)
	l.Put(2, 1)
	l.Put(1, 1)
	l.Put(2, 3)
	l.Put(4, 1)
	fmt.Println(l.Get(1))
	fmt.Println(l.Get(2))
}
