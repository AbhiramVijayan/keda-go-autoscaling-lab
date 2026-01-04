package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/abhiramvijayan/paroduct-api/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{}

}

func (p *Products) ServeHTTP(rw http.ResponseWriter, h *http.Request) {
	lp := data.GetProducts()
	json_p, err := json.Marshal(lp)
	if err != nil {
		http.Error(rw, "Unable to marshal json ", http.StatusInternalServerError)
	}
	rw.Write(json_p)
}
