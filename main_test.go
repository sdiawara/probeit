package main

import (
	"encoding/json"
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
var expectedProbe models.Probe

func TestMain(m *testing.M) {
	before()
	status = m.Run()
	after()
	os.Exit(status)
}

func before() {
	expectedProbe = models.NewProbe("Is this test ok ?", []string{})
	var err error
	session, err = mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	collection = session.DB("test").C("probe")
	collection.RemoveAll(bson.M{})
}

func after() {
	if collection != nil {
		collection.RemoveAll(bson.M{})
	}
	if session != nil {
		session.Close()
	}
}

func TestStaticFilesHandler(testing *testing.T) {
	writer, request := createTestResponseAndRequest("")

	StaticFilesHandler(writer, request)

	assert.Equal(testing, true, strings.Contains(writer.Body.String(), "<img alt=\"logo\" src=\"/images/logo.svg\" id=\"logo\" width=\"150px\" />"))
	assert.Equal(testing, true, strings.Contains(writer.Body.String(), "<h1 class=\"cover-heading\">Nous les sondons pour vous.</h1>"))
}

func TestCreateProbe(testing *testing.T) {
	writer, request := createTestResponseAndRequest(`{"Question":"Do you like golang ?"}`)

	CreateProbe(writer, request)

	probe := findOneProbeAndRemoveIt()
	assert.Equal(testing, "Do you like golang ?", probe.Question)
	assert.Equal(testing, http.StatusOK, writer.Code)
}

func TestCanNotCreateInvalidProbe(testing *testing.T) {
	writer, request := createTestResponseAndRequest(`{"Question":""}`)

	CreateProbe(writer, request)

	assert.Equal(testing, http.StatusBadRequest, writer.Code)
}

func TestListProbe(testing *testing.T) {
	collection.RemoveAll(bson.M{})
	collection.Insert(expectedProbe)
	writer, request := createTestResponseAndRequest("")

	ListProbe(writer, request)

	probes := decode(writer)
	assert.Equal(testing, 1, len(probes))
	assert.Equal(testing, expectedProbe.Question, probes[0].Question)
	assert.Equal(testing, expectedProbe.Responses, probes[0].Responses)
}

func TestRespondProbe(testing *testing.T) {
	collection.RemoveAll(bson.M{})
	probe := models.NewProbe("Aimez-vous golang ?", []string{})
	probe.Id = bson.NewObjectId()
	collection.Insert(probe)
	request := createRequest(models.ProbeResponse{probe.Id, "Oui"})

	RespondProbe(nil, request)

	actualProbe := findOneProbeAndRemoveIt()
	assert.Equal(testing, 1, len(actualProbe.Responses))
	assert.Equal(testing, "Oui", actualProbe.Responses[0])
}

func decode(writer *httptest.ResponseRecorder) (probes []models.Probe) {
	decoder := json.NewDecoder(writer.Body)
	decoder.Decode(&probes)
	return
}

func createRequest(probeResponse models.ProbeResponse) (request *http.Request) {
	probeResponseJson, _ := json.Marshal(probeResponse)
	request, _ = http.NewRequest("", "/", strings.NewReader(string(probeResponseJson)))
	return
}

func createTestResponseAndRequest(data string) (writer *httptest.ResponseRecorder, request *http.Request) {
	writer = httptest.NewRecorder()
	request, _ = http.NewRequest("", "/", strings.NewReader(data))
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
