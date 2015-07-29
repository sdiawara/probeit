package main

import (
	"bytes"
	"github.com/sdiawara/probeit/Godeps/_workspace/src/github.com/stretchr/testify/assert"
	"github.com/sdiawara/probeit/Godeps/_workspace/src/gopkg.in/mgo.v2"
	"github.com/sdiawara/probeit/Godeps/_workspace/src/gopkg.in/mgo.v2/bson"
	"github.com/sdiawara/probeit/models"
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
	writer := httptest.NewRecorder()

	var probeJson = []byte(`{"Question":"Aimez-vous golang ?"}`)
	request, _ := http.NewRequest("", "/", bytes.NewBuffer(probeJson))

	CreateProbe(writer, request)

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	probe := models.Probe{}
	c := session.DB("test").C("probe")

	c.Find(bson.M{}).One(&probe)

	c.Remove(bson.M{})
	assert.Equal(testing, "Aimez-vous golang ?", probe.Question)
}
