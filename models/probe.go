package models

import (
    "gopkg.in/mgo.v2/bson"
)

type Probe struct {
	Id       bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Question string
	Responses []string
}
