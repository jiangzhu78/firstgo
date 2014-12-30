package httpserver

import (
	"fmt"
	"log"
	"net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP (
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, "Congratulations! Here is the first GO!")
}

func Start(port int) {
	fmt.Println(port)
	var h Hello
	err := http.ListenAndServe("localhost:" + fmt.Sprintf("%d", port), h)
	if err != nil {
		log.Fatal(err)
	}
}
