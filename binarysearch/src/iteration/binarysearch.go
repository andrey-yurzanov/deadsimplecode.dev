package binarysearch

// Реализация бинарного поиска, итеративный вариант.
// array - отсортированный массив
// value - значение для поиска в массиве array
func BinarySearch(array []int, value int) int {
	index := 0
	begin := 0
	end := len(array) - 1

	for begin <= end {
		// Вычисление середины массива
		index = (begin + end) / 2
		// Если значения равны, значит удалось найти элемент
		if array[index] == value {
			return index
			// Если выбранный элемент больше, то переходим в левую часть
		} else if array[index] > value {
			end = index - 1
			// Если выбранный элемент меньше, то переходим в правую часть
		} else if array[index] < value {
			begin = index + 1
		}
	}
	// Если значение не найдено в массиве
	return -1
}
