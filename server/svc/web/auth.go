package web

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/rs/zerolog/log"
)

func (cfg *Config) GrantToken(user string) (string, error) {
	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user,
		"nbf": time.Now().Unix(),
		"exp": time.Now().Add(1 * time.Hour).Unix(),
	})
	tknStr, err := tkn.SignedString([]byte(cfg.TokenSecret))
	if err != nil {
		return ``, fmt.Errorf(`failed to sign token string: %w`, err)
	}
	return tknStr, nil
}

func (cfg *Config) VerifyToken(tknStr string) error {
	_, err := jwt.Parse(tknStr, func(tkn *jwt.Token) (interface{}, error) {
		if _, ok := tkn.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf(`unexpected signing method: %v`, tkn.Header[`alg`])
		}
		return cfg.TokenSecret, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors & jwt.ValidationErrorExpired != 0 {
				return fmt.Errorf(`token expired`)
			}
		}
		return err
	}
	return nil
}