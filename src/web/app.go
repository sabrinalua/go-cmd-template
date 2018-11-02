package main

import (
	"fmt"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

type App struct {

}


func (a *App)Init() {
	fmt.Print("init\n")
}

func (a *App) Run(){
	r := mux.NewRouter()
	r.HandleFunc("/", loggingMiddleware(mainRouter))
	r.HandleFunc("/main", mainRouter)

	logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	  if err != nil {
	    panic(err)
	  }
	// routeLogger:= handlers.LoggingHandler(os.Stdout, r)
	routeLogger:= handlers.LoggingHandler(logFile, r)

	http.ListenAndServe(":1080", routeLogger)
}

func mainRouter(w http.ResponseWriter, r *http.Request) {
	fmt.Print("mainrouter")
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        next.ServeHTTP(w, r)
    })
}