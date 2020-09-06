package web

import (
	"encoding/json"
	"net/http"
	"path"
	"regexp"

	"golang.org/x/crypto/bcrypt"

	"github.com/rs/zerolog/log"
)

type User struct {
	Id string
	Name string
	Email string
	Password password `json:",omitempty"`
}

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func (cfg *Config) UsersHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var u User
	switch r.Method {
	case `GET`:
		id := path.Base(r.URL.EscapedPath())
		err := cfg.db.GetUserInfo(ctx, id, &u)
		if err != nil {
			log.Info().Err(err).Msg(`could not find user`)
			http.Error(w, `could not find user`, http.StatusNotFound)
			return
		}
		err = json.NewEncoder(w).Encode(u)
		if err != nil {
			log.Error().Err(err).Msg(`could not find user`)
			http.Error(w, `could not find user`, http.StatusNotFound)
			return
		}
	case `POST`:
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			log.Info().Err(err).Msg(`failed to decode user create json`)
			http.Error(w, `name, email, password are required`, http.StatusBadRequest)
			return
		}
		if u.Name == `` || u.Email == `` || u.Password == `` {
			http.Error(w, `name, email, password are required`, http.StatusBadRequest)
			return
		}
		valid := emailRegex.Match([]byte(u.Email))
		if !valid {
			http.Error(w, `email is not a valid email`, http.StatusBadRequest)
			return
		}
		pw, err := cfg.HashPassword([]byte(u.Password))
		if err != nil {
			log.Error().Err(err).Msg(`failed to create user`)
			http.Error(w, `failed to create user.`, http.StatusBadRequest)
			return
		}
		id, err := cfg.db.CreateUser(ctx, u.Email, u.Name, string(pw))
		if err != nil {
			log.Error().Err(err).Msg(`failed to create user`)
			http.Error(w, `failed to create user.`, http.StatusBadRequest)
			return
		}
		u.Id = id
		u.Password = ``
		err = json.NewEncoder(w).Encode(u)
		if err != nil {
			log.Error().Err(err).Msg(`failed to create user`)
			http.Error(w, `failed to create user`, http.StatusBadRequest)
			return
		}
		log.Info().Str(`name`, u.Name).Str(`email`, u.Email).Msg(`user created`)
	default:
		http.Error(w, ``, http.StatusMethodNotAllowed)
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