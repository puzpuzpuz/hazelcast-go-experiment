package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hazelcast/hazelcast-go-client"
)

const (
	APP_PORT = ":8080"
)

var client hazelcast.Client

func CreateEntryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	imap, err := client.GetMap(vars["map"])
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	valueData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	imap.Set(vars["key"], string(valueData))

	w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("updated"))
}

func main() {
	config := hazelcast.NewConfig()
	var err error
	client, err = hazelcast.NewClientWithConfig(config)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Println("Hazelcast client started")

	r := mux.NewRouter()
	r.HandleFunc("/api/{map}/{key}", CreateEntryHandler).Methods("POST")
	log.Println("Go service is running on port", APP_PORT)
	log.Fatal(http.ListenAndServe(APP_PORT, r))
}
