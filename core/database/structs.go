package database

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

const (
	ARTIDS_COL     = "artIds"
	ALERTS_COL     = "alerts" // unused... Will be used with @AlertItem
	EXPRESSION_COL = "expCnt"
)

type ArtItem struct {
	MGID        bson.ObjectId `json:"-" bson:"_id,omitempty"`
	ID          string        `json:"id"`
	Folder      string        `json:"folder"`
	FileName    string        `json:"filename"`
	OrgFileName string        `json:"orgfilename"`
	Expression  string        `json:"expression"`  // set if only one expression is used (can be empty)
	Expressions []string      `json:"expressions"` // set if multiable are used (will always be 0 if empty)
	Views       int           `json:"views"`
	Uploaded    time.Time     `json:"uploaded_on"`
	FileSize    int           `json:"file_size"`
	Width       int           `json:"width"`
	Height      int           `json:"height"`
}

type AlertItem struct {
	MGID  bson.ObjectId `json:"-" bson:"_id,omitempty"`
	Key   string        `json:"key"`
	Value string        `json:"message"`
	TTL   float64       `json:"ttl"` // Time To Live
}

type ExpressionItem struct {
	MGID       bson.ObjectId `json:"-" bson:"_id,omitempty"`
	Expression string        `json:"expression"`
	Usage      int           `json:"count"`
}
