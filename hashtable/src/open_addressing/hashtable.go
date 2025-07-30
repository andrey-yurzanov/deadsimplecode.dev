package hashtable

import (
	"fmt"
)

// Начальный размер хеш-таблицы
const INIT_SIZE = 100

// Уровень заполнености при котором произойдет выделение нового места
const LOAD_FACTOR = 0.5

// Значение на которое будет увеличен размер таблицы
const INCREASE_FACTOR = 2

// Константы для алгоритма хеширования djb2.
const DJB2_SHIFT = 5
const DJB2_INIT_HASH = 5381

// Узел для хранения данных хеш-таблицы.
type HashTableNode[K comparable, V any] struct {
	key   K
	value V
}

// Реализация хеш-таблицы открытой адресации с использованием двойного хеширования.
type HashTable[K comparable, V any] struct {
	size   uint64
	values []*HashTableNode[K, V]
}

// Добвление элемента в хеш-таблицу.
// key   - ключ
// value - значение
func (htable *HashTable[K, V]) Put(key K, value V) {
	// Инициализация при первом обращении
	if htable.values == nil {
		htable.values = make([]*HashTableNode[K, V], INIT_SIZE)
	} else {
		// Если таблица заполнена больше, чем на LOAD_FACTOR,
		// необходимо перенести данные в таблицу большего размера
		loaded := float64(htable.size) / float64(len(htable.values))
		if loaded >= LOAD_FACTOR {
			values := htable.values
			htable.values = make([]*HashTableNode[K, V], htable.size*INCREASE_FACTOR)
			htable.size = 0

			// Перехеширование всех элементов таблицы
			for _, node := range values {
				if node != nil {
					htable.Put(node.key, node.value)
				}
			}
		}
	}

	index := hash(key) % uint64(len(htable.values))
	// Вычисление сдвига относительно текущего индекса
	shift := DJB2_SHIFT - index%DJB2_SHIFT
	for htable.values[index] != nil && htable.values[index].key != key {
		index += shift
		index %= uint64(len(htable.values))
	}
	htable.values[index] = &HashTableNode[K, V]{key: key, value: value}
	htable.size++
}

// Получение значения по ключу.
// key - ключ для поиска значения в хеш-таблице
func (htable *HashTable[K, V]) Get(key K) V {
	var result V
	if htable.values != nil {
		index := hash(key) % uint64(len(htable.values))
		// Вычисление сдвига относительно текущего индекса
		shift := DJB2_SHIFT - index%DJB2_SHIFT
		for htable.values[index] != nil {
			if htable.values[index].key == key {
				result = htable.values[index].value
				break
			}

			index += shift
			index %= uint64(len(htable.values))
		}
	}
	return result
}

// Удаление значения по указанному ключу.
// key - ключ для поиска значения в хеш-таблице
func (htable *HashTable[K, V]) Delete(key K) {
	if htable.values != nil {
		index := hash(key) % uint64(len(htable.values))
		// Вычисление сдвига относительно текущего индекса
		shift := DJB2_SHIFT - index%DJB2_SHIFT
		for htable.values[index] != nil {
			if htable.values[index].key == key {
				htable.values[index] = nil
				break
			}

			index += shift
			index %= uint64(len(htable.values))
		}
		htable.size--
	}
}

// Получение количества элементов в хеш-таблице.
func (htable *HashTable[K, V]) GetSize() uint64 {
	return htable.size
}

// Простая реализация хеш-функции на базе алгоритма djb2.
func hash(value any) uint64 {
	text := fmt.Sprint(value)
	hash := uint64(DJB2_INIT_HASH)
	for _, symbol := range text {
		hash = ((hash << DJB2_SHIFT) + hash) + uint64(symbol)
	}
	return hash
}
