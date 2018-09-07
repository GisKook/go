package au

import (
	"encoding/base64"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type au_jwt struct {
	MachineNO string `json:"machine_no"`
	jwt.StandardClaims
}

type au_jwt_lic struct {
	Comment    string `json:"comment“"`
	ExpireDate string `json:"expire_date"`
	Token      string `json:"token"`
}

func au_jwt_token(machine_no, secret string, d int64) (string, error) {
	claims := au_jwt{
		machine_no,
		jwt.StandardClaims{
			ExpiresAt: d,
			Issuer:    "tang_yun_chao",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func AuJwtEncode(comment, machine_no, secret string, d int) ([]byte, error) {
	expire_unix := time.Now().Add(time.Hour * time.Duration(24*d)).Unix()
	jwt, err := au_jwt_token(machine_no, secret, expire_unix)
	if err != nil {
		return nil, err
	}

	return json.Marshal(&au_jwt_lic{
		Comment:    comment,
		ExpireDate: time.Unix(expire_unix, 0).Format("2006-01-02 15:04:05"),
		Token:      base64.StdEncoding.EncodeToString([]byte(jwt)),
	})
}
