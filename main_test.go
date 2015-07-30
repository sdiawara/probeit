package main

import (
	"bytes"
	"github.com/sdiawara/probeit/models"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMain(testing *testing.T) {
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("", "/", strings.NewReader(""))

	HelloHandler(writer, request)

	assert.Equal(testing, true, strings.Contains(writer.Body.String(), "<img alt=\"logo\" src=\"/images/logo.svg\" id=\"logo\" width=\"150px\" />"))
	assert.Equal(testing, true, strings.Contains(writer.Body.String(), "<h1 class=\"cover-heading\">Nous les sondons pour vous.</h1>"))
}

func TestCreateProbe(testing *testing.T) {
	request := createRequest(`{"Question":"Aimez-vous golang ?"}`)

	CreateProbe(nil, request)

	probe := findOneProbeAndRemoveIt()
	assert.Equal(testing, "Aimez-vous golang ?", probe.Question)
}

func TestRespondProbe(testing *testing.T) {
	CreateProbe(nil, createRequest(`{"Question":"Aimez-vous golang ?"}`))
	probe := findOneProbe()
	request := createRequest(`{"probe_id": "` + probe.Id.Hex() + `", "Responses": "OK"}`)

	RespondProbe(nil, request)

    probe = findOneProbeAndRemoveIt()
	assert.Equal(testing, 1, len(probe.Responses))
	assert.Equal(testing, "OK", probe.Responses[0])
}

func createRequest(param string) (request *http.Request) {
	probeJson := []byte(param)
	request, _ = http.NewRequest("", "/", bytes.NewBuffer(probeJson))
	return
}

func getCollection() (collection *mgo.Collection) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	collection = session.DB("test").C("probe")
	return
}


func findById(id string) (probe *models.Probe) {
	probe = &models.Probe{}
	getCollection().Find(bson.M{"_id" : bson.ObjectIdHex(id)}).One(probe)
	return
}

func findOneProbe() (probe *models.Probe) {
	probe = &models.Probe{}
	getCollection().Find(bson.M{}).One(probe)
	return
}

func findOneProbeAndRemoveIt() (probe *models.Probe) {
	probe = findOneProbe()
	getCollection().Remove(bson.M{})
	return
}
