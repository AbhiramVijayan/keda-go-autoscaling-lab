package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHelloHandler(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	d, _err := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "Hello %s \n", d)
	h.l.Printf("Data %s", d)
	if _err != nil {
		http.Error(w, " r Error", http.StatusBadRequest)
		h.l.Printf("Error reading body: %v", _err)
		return
	}
}
