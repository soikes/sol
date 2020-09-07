package web

import (
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"
)

func (cfg *Config) RequiresJWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearer := r.Header.Get(`Authorization`)
		if bearer == `` {
			http.Error(w, `authentication required`, http.StatusUnauthorized)
			return
		}
		parts := strings.Split(bearer, ` `)
		if len(parts) != 2 {
			http.Error(w, `authentication required`, http.StatusUnauthorized)
			return
		}
		tkn := parts[1]
		err := cfg.VerifyToken(tkn)
		if err != nil {
			log.Info().Err(err).Str(`path`, r.URL.EscapedPath()).Msg(`unauthorized access attempt`)
			http.Error(w, `authentication required`, http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}