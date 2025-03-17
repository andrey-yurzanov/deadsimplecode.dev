package linkedlist

// Узел списка.
// value - полезные данные, добавляемые в список
// next  - указатель на следующий узел списка
// prev  - указатель на предыдущий узел списка
type LinkedListNode[T any] struct {
	value T
	next  *LinkedListNode[T]
	prev  *LinkedListNode[T]
}

// Связный список.
// size - длина списка
// head - первый узел списка
// tail - последний узел списка
type LinkedList[T any] struct {
	size uint
	head *LinkedListNode[T]
	tail *LinkedListNode[T]
}

// Добавление элемента в начало списка.
// element - новый элемент для добавления
func (list *LinkedList[T]) InsertFirst(element T) {
	node := &LinkedListNode[T]{value: element, next: list.head}
	if list.head == nil {
		list.tail = node
	} else {
		list.head.prev = node
	}
	list.head = node
	list.size++
}

// Добавление элемента в конец списка.
// element - новый элемент для добавления
func (list *LinkedList[T]) InsertLast(element T) {
	node := &LinkedListNode[T]{value: element, prev: list.tail}
	if list.tail == nil {
		list.head = node
	} else {
		list.tail.next = node
	}
	list.tail = node
	list.size++
}

// Удаление первого элемента списка.
func (list *LinkedList[T]) RemoveFirst() {
	if list.head != nil {
		next := list.head.next
		if next != nil {
			next.prev = nil
		} else {
			list.tail = nil
		}

		list.head = next
		list.size--
	}
}

// Удаление последнего элемента списка.
func (list *LinkedList[T]) RemoveLast() {
	if list.tail != nil {
		prev := list.tail.prev
		if prev != nil {
			prev.next = nil
		} else {
			list.head = nil
		}
		list.tail = prev
		list.size--
	}
}

// Получение длины списка.
func (list *LinkedList[T]) Len() uint {
	return list.size
}

// Получение первого элемента списка.
func (list *LinkedList[T]) GetFirst() T {
	if list.head != nil {
		return list.head.value
	}
	panic("list is empty")
}

// Получение последнего элемента списка.
func (list *LinkedList[T]) GetLast() T {
	if list.tail != nil {
		return list.tail.value
	}
	panic("list is empty")
}
