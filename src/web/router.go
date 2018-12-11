package main 

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	// "log"
	"os"
	// "time"
)

var (
	count int 
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

func sendChunk(w http.ResponseWriter, r *http.Request) {
	// path:= r.URL.Path 
	chunksize:=100
	flusher, _:= w.(http.Flusher)
	w.Header().Set("X-Content-Type-Options","nosniff")
	file, err := os.Open("s.log")
	if err!=nil{
		fmt.Printf("multipart_upload (failed to open file): %s\n", err)
	}
	defer file.Close()

	// filesize := config.objectsize 
	buffer:= make([]byte, chunksize)
	counter:=0
	for {
		bytesread, err:= file.Read(buffer)
		if err!=nil{
			break
		}
		counter++
		fmt.Println("bytes read ", bytesread)
		flusher.Flush()
		fmt.Println("counter ", counter)
		w.Write(buffer[:bytesread])
		// time.Sleep(500 *time.Millisecond)
	}


	// for i:=1; i< 4; i++{
	// 	// fmt.Fprint(w, "%s\n", i)
	// 	flusher.Flush()
	// 	w.Write([]byte(fmt.Sprintf("i %d\n",i)))
	// }
	// vars := mux.Vars(r)
	// w.Write([]byte(fmt.Sprintf("bucket{%s} location",vars["bucket"])))
}