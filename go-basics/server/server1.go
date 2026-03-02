// Server1 — минимальный "эхо" сервер
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // каждый запрос вызывает handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler отражает путь URL запроса
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
