package main

import (
	"encoding/json"
	"fmt"
	"github.com/sdiawara/probeit/models"
	"github.com/sdiawara/probeit/template"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

func StaticFilesHandler(writer http.ResponseWriter, request *http.Request) {
	staticFilesHandler := http.FileServer(http.Dir("static"))
	staticFilesHandler.ServeHTTP(writer, request)
}

func RespondProbe(writer http.ResponseWriter, request *http.Request) {
	session, c := getSessionAndProbeCollection()
	defer session.Close()

	decoder := json.NewDecoder(request.Body)
	var requestParam map[string]string

	err := decoder.Decode(&requestParam)
	if err != nil {
		panic(err)
	}

	update := bson.M{"$push": bson.M{"responses": requestParam["Response"]}}
	err = c.Update(bson.M{"_id": bson.ObjectIdHex(requestParam["probe_id"])}, update)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateProbe(writer http.ResponseWriter, request *http.Request) {
	p := make([]byte, request.ContentLength)
	request.Body.Read(p)

	var probe models.Probe
	err := json.Unmarshal(p, &probe)
	if err != nil {
		panic(err)
	}

	if probe.Question != "" {
		probe.Save()
		fmt.Fprint(writer, "ok")
	} else {
		writer.WriteHeader(http.StatusBadRequest)
	}
}

func getSessionAndProbeCollection() (*mgo.Session, *mgo.Collection) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	collection := session.DB("test").C("probe")
	return session, collection
}

func main() {
	http.HandleFunc("/", StaticFilesHandler)
	http.HandleFunc("/CreateProbe", CreateProbe)
	http.HandleFunc("/polls", template.PollTemplateHandler)
	port := "3000"
	fmt.Printf("Running on port %s...\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Erreur au démarrage du serveur : %s\n", err.Error())
	}
}
