package mgolog

import (
	"errors"
	"github.com/kidstuff/WebUtil/dblog"
	"io"
	"labix.org/v2/mgo"
	"net/http"
)

var (
	ErrNotInit = errors.New("kidstuff/WebUtil/dblog/mgolog: package need to be init")
)

type MgoConfiguratorProvider struct {
	DBSess *mgo.Session
	DBName string
}

func (p *MgoConfiguratorProvider) OpenLogger(*http.Request) (io.Writer, error) {
	if DefaultProvider.DBSess == nil {
		return nil, ErrNotInit
	}
	db := p.DBSess.Clone().DB(p.DBName)

	return NewMgoLogManager(db), nil
}

var DefaultProvider = &MgoConfiguratorProvider{}

func Register(dbsess *mgo.Session, dbname string) {
	DefaultProvider.DBSess = dbsess.Clone()
	DefaultProvider.DBName = dbname
}

func init() {
	dblog.LoggerRegister("mgolog", DefaultProvider)
}
