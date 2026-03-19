package main

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// Есть место для роста — расширяем срез
		z = x[:zlen]
	} else {
		// Места нет — выделяем новый массив,
		// увеличиваем в два раза для линейной амортизированной сложности
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}
