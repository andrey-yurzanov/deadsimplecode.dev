package binarysearch

// Реализация бинарного поиска, поиск всех вхождений.
// array - отсортированный массив
// value - значение для поиска в массиве array
func BinarySearch(array []int, value int) []int {
	index := 0
	begin := 0
	end := len(array) - 1

	result := []int{}
	for begin <= end {
		// Вычисление середины массива
		index = (begin + end) / 2
		// Если значения равны, значит удалось найти элемент
		if array[index] == value {
			result = append(result, index)

			// Двигаемся влево до тех пор пока значения совпадают и есть элементы
			leftIndex := index - 1
			for leftIndex >= begin && array[leftIndex] == value {
				result = append(result, leftIndex)
				leftIndex--
			}

			// Двигаемся вправо до тех пор пока значения совпадают и есть элементы
			rightIndex := index + 1
			for rightIndex <= end && array[rightIndex] == value {
				result = append(result, rightIndex)
				rightIndex++
			}
			break
			// Если выбранный элемент больше, то переходим в левую часть
		} else if array[index] > value {
			end = index - 1
			// Если выбранный элемент меньше, то переходим в правую часть
		} else if array[index] < value {
			begin = index + 1
		}
	}
	// Если значение не найдено в массиве
	return result
}
