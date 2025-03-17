package binarysearch

// Реализация бинарного поиска, рекурсивный вариант.
// array - отсортированный массив
// value - значение для поиска в массиве array
func BinarySearch(array []int, value int) int {
	end := len(array) - 1
	if end >= 0 {
		// Вычисление середины массива
		index := end / 2
		// Если значения равны, значит удалось найти элемент
		if array[index] == value {
			return index
			// Если выбранный элемент больше, то переходим в левую часть
		} else if array[index] > value {
			if result := BinarySearch(array[:index], value); result >= 0 {
				return result
			}
			// Если выбранный элемент меньше, то переходим в правую часть
		} else if array[index] < value {
			begin := index + 1
			if result := BinarySearch(array[begin:end+1], value); result >= 0 {
				return begin + result
			}
		}
	}
	// Если значение не найдено в массиве
	return -1
}
