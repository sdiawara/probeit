package main

import (
	"encoding/json"
	"fmt"
	"github.com/sdiawara/probeit/Godeps/_workspace/src/gopkg.in/mgo.v2"
	"github.com/sdiawara/probeit/models"
	"log"
	"net/http"
	//"github.com/gorilla/mux"
)

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	staticFilesHandler := http.FileServer(http.Dir("static"))
	staticFilesHandler.ServeHTTP(writer, request)
}

func CreateProbe(writer http.ResponseWriter, request *http.Request) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("test").C("probe")

	decoder := json.NewDecoder(request.Body)
	var probe models.Probe
	err = decoder.Decode(&probe)
	if err != nil {
		panic(err)
	}

	err = c.Insert(probe)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", HelloHandler)
	fmt.Printf("Running on port 3000...\n")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Printf("Erreur au démarrage du serveur : %s\n", err.Error())
	}
}
