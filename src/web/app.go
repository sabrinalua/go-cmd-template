package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

type App struct {

}


func (a *App)Init() {
	fmt.Print("init\n")
}

func (a *App) Run(){
	r := mux.NewRouter().SkipClean(true)
	r.Methods("GET").Path("/chunky").HandlerFunc(sendChunk)

	// registerRouter(r)

	logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	  if err != nil {
	    panic(err)
	  }
	// routeLogger:= handlers.LoggingHandler(os.Stdout, r)
	routeLogger:= handlers.CombinedLoggingHandler(logFile, r)

	srv:= &http.Server{
		Addr: "0.0.0.0:1080",
		Handler: routeLogger,
		WriteTimeout: 15*time.Second,
		ReadTimeout: 15*time.Second,
	}
	srv.ListenAndServe()
}

func mainRouter(w http.ResponseWriter, r *http.Request) {
	fmt.Print("mainrouter")
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        next.ServeHTTP(w, r)
    })
}