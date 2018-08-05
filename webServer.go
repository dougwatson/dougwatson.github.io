package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

var port = flag.String("port", "80", "port number")

func handler(w http.ResponseWriter, r *http.Request) {
	nanoTime := time.Now().UnixNano()      //in nanoseconds
	milliTime := nanoTime/1000000
	siteId := r.URL.Query().Get("siteId")
	mediaType := r.URL.Query().Get("mediaType")
	value := r.URL.Query().Get("value")
	eventTime := r.URL.Query().Get("eventTime")
	dbTime := r.URL.Query().Get("dbTime")
	fmt.Fprint(w, "%d,%v,%v,%v,\n",milliTime,dbTime,eventTime,siteId,mediaType,value)
}
func main() {
	flag.Parse()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+*port, nil)
}
