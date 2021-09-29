package strings

func Deduplicate(arr *[]string) {
	/*
	   Classic duplicate removal algorithm on sorted data
	   Data varies arr-place
	*/

	if len(*arr) <= 1 {
		return
	}

	j := 0

	for i := 1; i < len(*arr); i++ {
		if (*arr)[j] != (*arr)[i] {
			j++
			(*arr)[j] = (*arr)[i]
		}
	}
	*arr = (*arr)[:j+1]
	return
}

func Deduplicate2(arr []string) (out int) {
	/*
	   Дедуплицирует отсортированный массив на месте,
	   перемещая уникальные элементы в начало списка
	   Возвращает длину массива после дедупликации
	*/

	if len(arr) <= 1 {
		out = len(arr)
		return
	}

	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[out] {
			out += 1
			arr[out] = arr[i]
		}
	}
	out += 1
	return
}

func Deduplicate3(arr []string) (out []string) {
	/*
	   Возвращает новый массив без дубликтов из переданного отсортированного
	   (из группы одинаковых элементов берется последний)
	*/
	prev := arr[0]

	for _, curr := range arr {
		if prev != curr {
			out = append(out, prev)
			prev = curr
		}
	}

	out = append(out, prev)
	return
}

func Deduplicate4(arr []string) (out []string) {
	/*
	   Возвращает новый массив без дубликтов из переданного отсортированного
	   (из группы одинаковых элементов берется последний)
	*/
	prev := arr[0]

	for i := 1; i < len(arr); i++ {
		curr := arr[i]
		if prev != curr {
			out = append(out, prev)
			prev = curr
		}
	}

	out = append(out, prev)
	return
}

func Duplicates(arr []string) (out []string) {
	/*
	   Возвращает  массив повторяющихся элементов из отсортированного массива
	   Возвращает только один дубликат из группы
	*/

	prev := arr[0]
	cnt := 1

	for i := 1; i < len(arr); i++ {
		curr := arr[i]
		if prev != curr {
			prev = curr
			cnt = 1
		} else {
			if cnt == 1 {
				out = append(out, prev)
			}
			cnt += 1
		}
	}
	return
}

func Unique(arr []string) (out []string) {
	/*
	   Возвращает  массив уникальных (не имеющих дубликатов) элементов
	   из отсортированного массива
	*/
	prev := arr[0]
	cnt := 1

	for i := 1; i < len(arr); i++ {
		curr := arr[i]
		if prev != curr {
			if cnt == 1 {
				out = append(out, prev)
			}
			prev = curr
			cnt = 1
		} else {
			cnt += 1
		}
	}
	if cnt == 1 {
		out = append(out, prev)
	}

	return
}

// Варанты тех же функций с использованием доп. памяти
// (равную числу уникальных элементов) в виде словаря
// и основанные на предварительном подсчете числа вхождений каждого элемента
// Не требуют сортировки

func Counter(arr []string) (out map[string]int) {
	/*The dictionary for counting occurrences */
	out = make(map[string]int)

	for i := 0; i < len(arr); i++ {
		out[arr[i]] += 1
	}

	return
}

func DeduplicateCounter(arr []string) (out []string) {
	/*
	Возвращает новый массив без дубликтов
	Не требует сортировки
	*/

	counter := Counter(arr)
	for k, _ := range counter {
		out = append(out, k)
	}

	return
}

func DuplicatesCounter(arr []string) (out []string) {
	/* Возвращает новый массив из дубликтов (по одному из группы) */

	counter := Counter(arr)
	for k, v := range counter {
		if v > 1 {
			out = append(out, k)
		}
	}

	return
}

func UniqueCounter(arr []string) (out []string) {
	/* Возвращает новый массив из уникальных
	   (не имеющих повторов в исходном массиве) элементов
	*/

	counter := Counter(arr)
	for k, v := range counter {
		if v == 1 {
			out = append(out, k)
		}
	}

	return
}
