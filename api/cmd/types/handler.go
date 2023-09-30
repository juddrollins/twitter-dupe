package types

import (
	"github.com/juddrollins/twitter-dupe/cmd/config"
	"github.com/juddrollins/twitter-dupe/db"
	"gopkg.in/go-playground/validator.v9"
)

type Handler struct {
	Validator *validator.Validate
	Cfig      config.Config
	Dao       *db.Dao
}
