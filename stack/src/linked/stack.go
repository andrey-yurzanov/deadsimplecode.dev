package stack

// Узел связного списка для хранения элементов стека.
// value - хранимое значение
// next  - указатель на следующий узел списка
type StackNode[T any] struct {
	value T
	next  *StackNode[T]
}

// Стек реализованный на базе связного списка.
// head - указатель на вершину стека
type Stack[T any] struct {
	head *StackNode[T]
}

// Добавление элемента в стек.
// element - новый элемент стека, который будет добавлен на вершину стека
func (stack *Stack[T]) Push(element T) {
	// Новый элемент указывается в качестве первого узла списка, старое
	// значение указывается в качестве следующего узла списка
	node := &StackNode[T]{value: element, next: stack.head}
	stack.head = node
}

// Извлечение элемента из стека, возвращает элемент или выбрасывает
// ошибку, если стек пуст.
func (stack *Stack[T]) Pop() T {
	// Если указатель содержит узел связного списка, то возвращаем
	// значение и указываем следующий узел в качестве первого узла
	if stack.head != nil {
		element := stack.head.value
		stack.head = stack.head.next
		return element
	}
	// Выбрасываем ошибку, если стек пуст
	panic("stack is empty")
}
