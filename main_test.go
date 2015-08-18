package main

import (
	"bytes"
	"github.com/sdiawara/probeit/models"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var status int
var session *mgo.Session
var collection *mgo.Collection

func TestMain(m *testing.M) {
	before()
	status = m.Run()
	after()
	os.Exit(status)
}

func before() {
	var err error
	session, err = mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	collection = session.DB("test").C("probe")

}

func after() {
	if collection == nil {		
		collection.RemoveAll(bson.M{})
	}
	if session != nil {
		session.Close()
	}
}

func TestHelloHandler(testing *testing.T) {
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

func TestListProbe(testing *testing.T) {
	collection.RemoveAll(bson.M{})
	expectedProbe := models.Probe{"", "Is this test ok ?", []string{}}
	collection.Insert(expectedProbe)

	request, _ := http.NewRequest("", "/", strings.NewReader(""))

	polls := ListProbe(request)

	assert.Equal(testing, 1, len(polls))
	assert.Equal(testing, expectedProbe.Question, polls[0].Question)
	assert.Equal(testing, expectedProbe.Responses, polls[0].Responses)
}

func TestRespondProbe(testing *testing.T) {
	CreateProbe(nil, createRequest(`{"Question":"Aimez-vous golang ?"}`))
	probe := findOneProbe()
	request := createRequest(`{"probe_id": "` + probe.Id.Hex() + `", "Responses": "Oui"}`)

	RespondProbe(nil, request)

	probe = findOneProbeAndRemoveIt()
	assert.Equal(testing, 1, len(probe.Responses))
	assert.Equal(testing, "Oui", probe.Responses[0])
}

func createRequest(param string) (request *http.Request) {
	probeJson := []byte(param)
	request, _ = http.NewRequest("", "/", bytes.NewBuffer(probeJson))
	return
}

func findOneProbe() (probe *models.Probe) {
	probe = &models.Probe{}
	collection.Find(bson.M{}).One(probe)
	return
}

func findOneProbeAndRemoveIt() (probe *models.Probe) {
	probe = findOneProbe()
	collection.Remove(bson.M{})
	return
}
