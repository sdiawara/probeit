package main

import (
	"encoding/json"
	"fmt"
	"github.com/sdiawara/probeit/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"html/template"
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

func ListProbe(request *http.Request) (probes []models.Probe) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("test").C("probe")

	err = c.Find(bson.M{}).All(&probes)
	if err != nil {
		log.Fatal(err)
	}
	return
}

type PollsPage struct {Polls []models.Probe}

func PollTemplateHandler(writer http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.New("poll").ParseFiles("static/polls.html"))
	
	pageParam := PollsPage{Polls: ListProbe(request)}
	if err := tmpl.Execute(writer, pageParam); err != nil {
		log.Fatalf("Erreur dans le template : %s", err.Error())
	}
}

func main() {
	http.HandleFunc("/", HelloHandler)
	http.HandleFunc("/polls", PollTemplateHandler)
	fmt.Printf("Running on port 3000...\n")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Printf("Erreur au démarrage du serveur : %s\n", err.Error())
	}
}
