package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/noodles623/csp/handlers"
	"github.com/noodles623/csp/store"
)

type Args struct {
	conn string
	port string
}

func Run(args Args) error {
	router := mux.NewRouter().PathPrefix("/csp/").Subrouter()
	st := store.NewPostgresAssetStore(args.conn)
	hnd := handlers.NewAssetHandler(st)
	RegisterAllRoutes(router, hnd)

	log.Println("Starting server at port: ", args.port)
	return http.ListenAndServe(args.port, router)
}

func RegisterAllRoutes(router *mux.Router, hnd handlers.AssetHandler) {
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

	router.HandleFunc("/asset", hnd.Get).Methods(http.MethodGet)
	router.HandleFunc("/asset", hnd.Create).Methods(http.MethodPost)
	router.HandleFunc("/asset", hnd.Delete).Methods(http.MethodDelete)

	router.HandleFunc("/assets", hnd.List).Methods(http.MethodGet)
}
