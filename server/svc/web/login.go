package web

import (
	"encoding/json"
	"net/http"
)

func (cfg *Config) LoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case `POST`:
		ctx := r.Context()
		var u User
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			cfg.Log.Info().Err(err).Msg(`failed to decode user create json`)
			http.Error(w, `failed to login. email and password are required`, http.StatusBadRequest)
			return
		}
		if u.Email == `` || u.Password == `` {
			http.Error(w, `failed to login. email and password are required`, http.StatusBadRequest)
			return
		}
		hashed, err := cfg.db.GetUserPassword(ctx, u.Email)
		if err != nil {
			http.Error(w, `failed to login. check email and password.`, http.StatusBadRequest)
			return
		}
		valid := cfg.VerifyPassword([]byte(hashed), []byte(u.Password))
		if !valid {
			http.Error(w, `failed to login. check email and password.`, http.StatusBadRequest)
			return
		}
		tkn, err := cfg.GrantToken(u.Id, u.Email)
		if err != nil {
			cfg.Log.Error().Err(err).Msg(`login failed`)
			http.Error(w, `failed to login. check email and password.`, http.StatusBadRequest)
			return
		}
		rsp := struct { Token string }{ Token: tkn }
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(rsp)
		if err != nil {
			cfg.Log.Error().Err(err).Msg(`login failed`)
			http.Error(w, `failed to login. check email and password.`, http.StatusBadRequest)
			return
		}
	}
}