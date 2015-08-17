package models

import (
	//"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Probe struct {
	Id                bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Question          string
	PossibleResponses []string
	Responses         []string
}

type ProbeResponse struct {
	ProbeId  bson.ObjectId `json:"probe_id"`
	Response string
}

func NewProbe(question string, possibleResponses []string) Probe {
	return Probe{"", question, possibleResponses, []string{}}
}

func (probe Probe) Save() {
	session := getSession()
	defer session.Close()

	collection := getCollection(session)

	if probe.Id == "" {
		save(collection, probe)
	} else {
		update(collection, probe)
	}
}

func save(collection *mgo.Collection, probe Probe) {
	err := collection.Insert(probe)
	if err != nil {
		panic(err)
	}
}

func update(collection *mgo.Collection, probe Probe) {
	err := collection.UpdateId(probe.Id, probe)
	if err != nil {
		panic(err)
	}
}

func GetById(id string) Probe {
	session := getSession()
	defer session.Close()

	collection := getCollection(session)

	var probe Probe
	collection.FindId(bson.ObjectIdHex(id)).One(&probe)

	return probe
}

func AllProbes() []Probe {
	session := getSession()
	defer session.Close()

	collection := getCollection(session)

	var probes []Probe
	collection.Find(bson.M{}).All(&probes)

	return probes
}

func getSession() (session *mgo.Session) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	return
}

func getCollection(session *mgo.Session) *mgo.Collection {
	return session.DB("test").C("probe")
}
