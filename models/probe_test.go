package models

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
	"testing"
)

var status int
var session *mgo.Session
var collection *mgo.Collection
var expectedProbe Probe

func TestMain(m *testing.M) {
	before()
	status = m.Run()
	after()
	os.Exit(status)
}

func before() {
	expectedProbe = NewProbe("Is this test ok ?", []string{})
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

func TestSave(testing *testing.T) {
	probe := NewProbe("Can you save this ?", []string{})

	probe.Save()

	toto := findOneProbe()
	assert.Equal(testing, "Can you save this ?", toto.Question)
}

func TestUpdate(testing *testing.T) {
	probe := NewProbe("Can you save this ?", []string{})
	probe.Save()

	probes := AllProbes()
	probe = probes[0]
	probe.Question = "Can you save this again ?"
	probe.Save()

	assert.Equal(testing, 1, len(probes))
	toto := findOneProbe()
	assert.Equal(testing, "Can you save this again ?", toto.Question)
}

func testGetById(testing *testing.T) {
	probe := NewProbe("Can you save this ?", []string{})
	probe.Save()

	probe = GetById("000000000000000000000000")

	assert.Equal(testing, "Can you save this again ?", probe.Question)
}

func TestAllProbes(testing *testing.T) {
	probe := NewProbe("Can you save this ?", []string{})
	probe.Save()

	probes := AllProbes()

	assert.Equal(testing, 1, len(probes))
}

func findOneProbe() (probe *Probe) {
	probe = &Probe{}
	collection.Find(bson.M{}).One(probe)
	collection.RemoveAll(bson.M{})
	return
}
