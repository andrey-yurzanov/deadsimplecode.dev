package queue

// Кол-во элементов, которое будет выделено при достижении границы массива.
const ALLOCATE_SIZE = 10

// Очередь на основе массива.
// head     - индекс указывающий на первый элемент в очереди
// tail     - индекс указывающий на последний элемент очереди
// elements - массив элементов очереди
type Queue[T any] struct {
	head     int
	tail     int
	elements []T
}

// Добавление нового элемента в конец очереди.
// element - добавляемый элемент
func (queue *Queue[T]) Insert(element T) {
	// При первом добавлении элемента необходимо инициализировать
	// индексы начальными значениями
	size := len(queue.elements)
	if size == 0 {
		queue.head = 0
		queue.tail = -1
	}

	// Если исчерпали свободные ячейки для добавления, то необходимо расширить
	// массив на ALLOCATE_SIZE, значения при этом должны быть скопированны
	// в новый массив
	index := queue.tail + 1
	if (index >= size && queue.head == 0) || index == queue.head {
		allocated := make([]T, size+ALLOCATE_SIZE)

		headIndex := 0
		if size > 0 {
			// Копирование элементов от head и до конца массива
			for headIndex+queue.head < size {
				allocated[headIndex] = queue.elements[headIndex+queue.head]
				headIndex++
			}

			// Копирование элементов от начала массива и до tail, только в
			// том случае, если начало очереди не совпадает с началом массива
			if queue.tail < queue.head {
				tailIndex := 0
				for tailIndex <= queue.tail {
					allocated[headIndex] = queue.elements[tailIndex]
					headIndex++
					tailIndex++
				}
			}
		}
		queue.head = 0
		index = headIndex
		queue.elements = allocated
		// Если достигли конца массива, но в самом начале есть пустые ячейки
		// для добавления элементов, тогда необходимо перейти в начало массива
	} else if index >= size && queue.head > 0 {
		index = 0
	}

	queue.tail = index
	queue.elements[index] = element
}

// Удаление элемента из очереди, возвращает значение первого элемента или
// выбрасывает ошибку, если очередь пуста.
func (queue *Queue[T]) Remove() T {
	// Если очередь содержит элементы, то обработка продолжается,
	// иначе выбрасываем ошибку
	size := len(queue.elements)
	if size > 0 && queue.tail >= 0 {
		value := queue.elements[queue.head]

		// Если достигли конца массива, но в самом начале есть элементы,
		// тогда необходимо перейти в начало массива
		index := queue.head + 1
		if index >= size && queue.head > queue.tail {
			index = 0
			// Если извлекли последний элемент, то возвращаем индексы
			// к начальным значениям
		} else if queue.head == queue.tail {
			index = 0
			queue.tail = -1
		}
		queue.head = index

		return value
	}
	panic("queue is empty")
}
