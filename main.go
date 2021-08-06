package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/piquette/finance-go/quote"
)

type quote_t struct {
	Symbol string
	Price  float64
}

func main() {
	http.HandleFunc("/csp/stock/", quoteHandler)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func parseParams(r *http.Request, prefix string, num int) ([]string, error) {
	url := strings.TrimPrefix(r.URL.Path, prefix)
	params := strings.Split(url, "/")
	if len(params) != num {
		return nil, fmt.Errorf("Bad format. Expecting exactly %d params", num)
	}
	for i := 0; i < num; i++ {
		if params[i] == "" {
			return nil, fmt.Errorf("Bad format. Expecting exactly %d params", num)
		}
	}
	return params, nil
}

func quoteHandler(w http.ResponseWriter, r *http.Request) {
	params, err := parseParams(r, "/csp/stock/", 1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	ticker := params[0]
	q, err := quote.Get(ticker)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	rq := quote_t{q.Symbol, q.RegularMarketPrice}
	b, err := json.Marshal(rq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
