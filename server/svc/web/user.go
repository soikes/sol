package web

import (
	"encoding/json"
	"net/http"
	"regexp"

	"golang.org/x/crypto/bcrypt" 

	"github.com/rs/zerolog/log"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func (cfg *Config) UsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case `POST`:
		var u struct { Name, Email, Password string }
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			log.Info().Err(err).Msg(`failed to decode user create json`)
			http.Error(w, `name, email, password are required`, 400)
			return
		}
		if u.Name == `` || u.Email == `` || u.Password == `` {
			http.Error(w, `name, email, password are required`, 400)
			return
		}
		valid := emailRegex.Match([]byte(u.Email))
		if !valid {
			http.Error(w, `email is not a valid email`, 400)
			return
		}
		pw, err := cfg.HashPassword([]byte(u.Password))
		if err != nil {
			log.Error().Err(err).Msg(`failed to create user`)
			http.Error(w, `failed to create user. Try again later.`, 500)
			return
		}
		err = cfg.db.CreateUser(r.Context(), u.Email, u.Name, string(pw))
		if err != nil {
			log.Error().Err(err).Msg(`failed to create user`)
			http.Error(w, `failed to create user. Try again later.`, 500)
			return
		}
		log.Info().Str(`name`, u.Name).Str(`email`, u.Email).Msg(`user created`)
	}
}

func (cfg *Config) HashPassword(pw []byte) ([]byte, error) {
	hashed, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	if err != nil {
		return []byte{}, err
	}
	return hashed, nil
}

func (cfg *Config) VerifyPassword(hashed, pw []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashed, pw)
	if err != nil {
		return false
	} else {
		return true
	}
}