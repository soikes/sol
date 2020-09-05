package web

import (
	"context"
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/rs/zerolog/log"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func (cfg *Config) UsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case `POST`:
		var u struct { Name, Email string }
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			log.Info().Err(err).Msg(`failed to decode user create json`)
			http.Error(w, `name and email are required`, 400)
			return
		}
		if u.Name == `` || u.Email == `` {
			http.Error(w, `name and email are required`, 400)
			return
		}
		err = cfg.CreateUser(r.Context(), u.Email, u.Name)
		if err != nil {
			log.Error().Err(err).Msg(`failed to create user`)
			http.Error(w, `failed to create user.`, 500)
			return
		}
		log.Info().Str(`name`, u.Name).Str(`email`, u.Email).Msg(`user created`)
	}
}

func (cfg *Config) CreateUser(ctx context.Context, name, email string) error {
	return cfg.db.CreateUser(ctx, name, email)
}