package main

import (
	"github.com/juddrollins/twitter-dupe/cmd/config"
	"github.com/juddrollins/twitter-dupe/db"
)

//Not used at the moment

type CTX struct {
	cfig config.Config
	dao  *db.Dao
}
