package template

import (
	"github.com/sdiawara/probeit/models"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http/httptest"
	"os"
	"testing"
)

var session *mgo.Session
var collection *mgo.Collection

func TestMain(m *testing.M) {
	before()
	status := m.Run()
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

func TestGetPollsTemplate(testing *testing.T) {
	templatePath = "../static/polls.html"
	probe := models.NewProbe("Is this test ok ?", []string{"Oui", "Non"})
	probe.Save()
	writer := httptest.NewRecorder()

	PollTemplateHandler(writer, nil)

	assert.Contains(testing, writer.Body.String(), "Is this test ok ?")
	assert.Contains(testing, writer.Body.String(), "Oui")
	assert.Contains(testing, writer.Body.String(), "Non")
}
