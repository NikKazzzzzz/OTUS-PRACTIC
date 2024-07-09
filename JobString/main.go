package main

import "fmt"

type Item struct {
	prev  *Item
	next  *Item
	value interface{}
}

type List struct {
	head *Item
	tail *Item
	len  int
}

func (l *List) Len() int {
	return l.len
}

func (l *List) First() *Item {
	return l.head
}

func (l *List) Last() *Item {
	return l.tail
}

func (l *List) PushFront(v interface{}) {
	newItem := &Item{
		value: v,
	}

	if l.len == 0 {
		l.head = newItem
		l.tail = newItem
	} else {
		newItem.next = l.head
		l.head.prev = newItem
		l.head = newItem
	}

	l.len++
}

func (l *List) PushBack(v interface{}) {
	newItem := &Item{
		value: v,
	}

	if l.len == 0 {
		l.head = newItem
		l.tail = newItem
	} else {
		newItem.prev = l.tail
		l.tail.next = newItem
		l.tail = newItem
	}

	l.len++
}

func (l *List) Remove(i *Item) {
	if i == nil {
		return
	}

	if i.prev != nil {
		i.prev.next = i.next
	} else {
		l.head = i.next
	}

	if i.next != nil {
		i.next.prev = i.prev
	} else {
		l.tail = i.prev
	}

	l.len--
}

func (i *Item) Value() interface{} {
	return i.value
}

func (i *Item) Next() *Item {
	return i.next
}

func (i *Item) Prev() *Item {
	return i.prev
}

func main() {
	list := &List{}

	list.PushFront(3)
	list.PushFront(2)
	list.PushFront(1)

	list.PushBack(4)
	list.PushBack(5)
	list.PushBack(6)

	fmt.Println("Список в прямом порядке:")
	for item := list.First(); item != nil; item = item.Next() {
		fmt.Println(item.Value())
	}

	fmt.Println("Список в обратном порядке:")
	for item := list.Last(); item != nil; item = item.Prev() {
		fmt.Println(item.Value())
	}

	item := list.First().Next()
	fmt.Println("Удален элемент со значением:", item.Value())
	list.Remove(item)

	fmt.Println("Список после удаления элемента:")
	for item := list.First(); item != nil; item = item.Next() {
		fmt.Println(item.Value())
	}
}
