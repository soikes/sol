package web

import (
	"encoding/json"
	"net/http" 

	"github.com/rs/zerolog/log"
)

func (cfg *Config) LoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case `POST`:
		ctx := r.Context()
		var u struct { Email, Password string }
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			log.Info().Err(err).Msg(`failed to decode user create json`)
			http.Error(w, `email, password are required`, 400)
			return
		}
		if u.Email == `` || u.Password == `` {
			http.Error(w, `email, password are required`, 400)
			return
		}
		hashed, err := cfg.db.GetUserPassword(ctx, u.Email)
		if err != nil {
			http.Error(w, `failed to login. check email and password.`, 400)
			return
		}
		valid := cfg.VerifyPassword([]byte(hashed), []byte(u.Password))
		if !valid {
			http.Error(w, `failed to login. check email and password.`, 400)
			return
		}
		tkn, err := cfg.GrantToken(u.Email)
		if err != nil {
			log.Error().Err(err).Msg(`login failed`)
			http.Error(w, `failed to login. check email and password.`, 400)
			return
		}
		rsp := struct { Token string }{ Token: tkn }
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(rsp)
		if err != nil {
			log.Error().Err(err).Msg(`login failed`)
			http.Error(w, `failed to login. check email and password.`, 400)
			return
		}
	}
}