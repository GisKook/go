package au

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

type au_jwt_claims_id struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

func AuJwtToken(id, secrect string, expire int64) (string, error) {
	claims := au_jwt_claims_id{
		id,
		jwt.StandardClaims{
			ExpiresAt: expire + time.Now().Unix(),
			Issuer:    "zhangkai",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secrect))
}

func AuJwtTokenValid(token_string string) (bool, string) {
	token, err := jwt.ParseWithClaims(token_string, &au_jwt_claims_id{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secrect), nil
	})

	if err != nil {
		log.Printf("<ERR> jwt parse error %s\n", err.Error())
		return false, ""
	}

	if claims, ok := token.Claims.(*au_jwt_claims_id); ok && token.Valid {
		return true, claims.ID
	}

	return false, ""
}
