package slice

func RemoveFast(s []string, i int) []string {
	/* If order is not important */

	// 1. Копировать последний элемент в индекс i.
	s[i] = s[len(s)-1]

	// 2. Удалить последний элемент (записать нулевое значение).
	s[len(s)-1] = ""

	// 3. Усечь срез.
	s = s[:len(s)-1]
	return s
}

func RemoveWithSaveOrder(s []string, i int) []string {
	/* very slow */

	// 1. Выполнить сдвиг s[i+1:] влево на один индекс.
	copy(s[i:], s[i+1:])

	// 2. Удалить последний элемент (записать нулевое значение).
	s[len(s)-1] = ""

	// 3. Усечь срез.
	s = s[:len(s)-1]

	return s
}

func RemoveWithSaveOrder2(s []string, i int) []string {
	/*very slow */
	return append(s[:i], s[i+1:]...)
}
