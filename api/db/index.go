package db

import (
	"github.com/juddrollins/twitter-dupe/cmd/config"
)

type Entry struct {
	PK   string `json:"PK"`
	SK   string `json:"SK"`
	Data string `json:"Data"`
}

type EntryService interface {
	CreateRecord(m Entry) (Entry, error)
	GetRecord(PK string, SK string) (Entry, error)
}

type Dao struct {
	db *config.Database
}

func InitDb(db *config.Database) *Dao {
	return &Dao{
		db: db,
	}
}
