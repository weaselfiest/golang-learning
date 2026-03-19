package main

import "fmt"

// nonempty возвращает срез, содержащий только непустые строки.
// Базовый массив изменяется "на месте".
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0] // срез нулевой длины, но с тем же базовым массивом
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func main() {
	data := []string{"one", "", "two", "", "three"}
	fmt.Printf("%q\n", nonempty(data)) // ["one" "two" "three"]
}
