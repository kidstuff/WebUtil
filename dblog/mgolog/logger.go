package mgolog

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

type logInfo struct {
	Id      bson.ObjectId `bson:"_id"`
	Message string        `bson:"Message"`
	On      time.Time
}

type MgoLogManager struct {
	LogColl *mgo.Collection
}

func (m *MgoLogManager) Write(p []byte) (int, error) {
	err := m.LogColl.Insert(&logInfo{
		Id:      bson.NewObjectId(),
		Message: string(p),
		On:      time.Now(),
	})

	return len(p), err
}

func NewMgoLogManager(db *mgo.Database) *MgoLogManager {
	return &MgoLogManager{db.C("mgolog")}
}
