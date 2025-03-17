package queue

// Узел списка для хранения значения и указателя на следующий узел.
// value - значение в узле
// next  - указатель на следующий узел
type QueueNode[T any] struct {
	value T
	next  *QueueNode[T]
}

// Очередь на основе связного списка.
// head - указатель на первый узел списка
// tail - указатель на последний узел списка
type Queue[T any] struct {
	head *QueueNode[T]
	tail *QueueNode[T]
}

// Добавление нового элемента в конец очереди.
// element - добавляемый элемент
func (queue *Queue[T]) Insert(element T) {
	node := &QueueNode[T]{value: element}
	if queue.tail != nil {
		queue.tail.next = node
	} else {
		queue.head = node
	}
	queue.tail = node
}

// Удаление элемента из очереди, возвращает значение первого элемента или
// выбрасывает ошибку, если очередь пуста.
func (queue *Queue[T]) Remove() T {
	if queue.head != nil {
		value := queue.head.value
		queue.head = queue.head.next

		if queue.head == nil {
			queue.tail = nil
		}
		return value
	}
	panic("queue is empty")
}
