package main

import (
	"encoding/json"
	"fmt"
	"github.com/sdiawara/probeit/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	staticFilesHandler := http.FileServer(http.Dir("static"))
	staticFilesHandler.ServeHTTP(writer, request)
}

func RespondProbe(writer http.ResponseWriter, request *http.Request) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("test").C("probe")

	decoder := json.NewDecoder(request.Body)
	var requestParam map[string]string

	err = decoder.Decode(&requestParam)
	if err != nil {
		panic(err)
	}

	update := bson.M{"$push": bson.M{"responses": requestParam["Responses"]}}
	err = c.Update(bson.M{"_id": bson.ObjectIdHex(requestParam["probe_id"])}, update)
	if err != nil {
		log.Fatal(err)
	}
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

func ListProbe(writer http.ResponseWriter, request *http.Request) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("test").C("probe")

	encoder := json.NewEncoder(writer)
	var probes []models.Probe

	err = c.Find(bson.M{}).All(&probes)
	if err != nil {
		log.Fatal(err)
	}

	err = encoder.Encode(&probes)
	if err != nil {
		panic(err)
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
