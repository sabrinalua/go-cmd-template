package main 

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"

)

func registerRouter(router *mux.Router) {
	apiRouter:= router.PathPrefix("/").Subrouter()
	var routers []*mux.Router
	routers = append(routers, apiRouter.PathPrefix("/{bucket}").Subrouter())
	for _, bucket := range routers {
		bucket.Methods("GET").Path("/{object:.+}").HandlerFunc(router1)
		bucket.Methods("GET").HandlerFunc(router2).Queries("location","")
	}
	apiRouter.Methods("GET").HandlerFunc(router3)
}

func router1(w http.ResponseWriter, r *http.Request) {
	// path:= r.URL.Path 
	vars := mux.Vars(r)
	fmt.Print(vars["object"])
	w.Write([]byte(fmt.Sprintf("bucket{%s}",vars["bucket"])))
}

func router2(w http.ResponseWriter, r *http.Request) {
	// path:= r.URL.Path 
	vars := mux.Vars(r)
	w.Write([]byte(fmt.Sprintf("bucket{%s} location",vars["bucket"])))
}

func router3(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("list buckets"))
}