package util

import (
	"errors"
	"log"
	"strings"

	"github.com/juddrollins/twitter-dupe/db"
)

func LoginUser(users []db.Entry, password string) (db.Entry, error) {
	for _, u := range users {
		var user_data = strings.Split(u.Data, "::")
		var user_password = user_data[1]

		// Check if password matches
		if CheckPasswordHash(password, user_password) {
			log.Println("Password matches")
			return u, nil
		}
	}
	log.Println("Password does not match")
	return db.Entry{}, errors.New("user password does not match")
}
