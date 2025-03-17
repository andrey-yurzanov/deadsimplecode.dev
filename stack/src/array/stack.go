package stack

// Кол-во элементов, которое будет выделено при достижении границы массива.
const ALLOCATE_SIZE = 10

// Стек реализованный на базе массива, содержит два поля:
// index    - индекс указывающий на свободную ячейку доступную для добавления нового элемента
// elements - элементы стека
type Stack[T any] struct {
	index    int
	elements []T
}

// Добавление элемента в стек.
// element - новый элемент стека, который будет добавлен на вершину стека
func (stack *Stack[T]) Push(element T) {
	// Если указатель добрался до границ массива, необходимо выделить место
	// для новых элементов размером ALLOCATE_SIZE
	if stack.index >= len(stack.elements)-1 {
		allocated := make([]T, len(stack.elements)+ALLOCATE_SIZE)
		copy(allocated, stack.elements)

		stack.elements = allocated
	}

	// Добавление нового элемента в массив и сдвиг указателя
	stack.elements[stack.index] = element
	stack.index++
}

// Извлечение элемента из стека, возвращает элемент или выбрасывает
// ошибку, если стек пуст.
func (stack *Stack[T]) Pop() T {
	// Если указатель больше нуля, извлекаем элемент из стека
	if stack.index > 0 {
		stack.index--
		element := stack.elements[stack.index]
		return element
	}
	// Выбрасываем ошибку, если стек пуст
	panic("stack is empty")
}
