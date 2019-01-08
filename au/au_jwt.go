package au

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"os"
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

func au_jwt_token_gen(machine_no, secret string, d int64) (string, error) {
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

func AuJwtEncode(comment, machine_no, secret string, expire_unix int64) ([]byte, error) {
	jwt, err := au_jwt_token_gen(machine_no, secret, expire_unix)
	if err != nil {
		return nil, err
	}

	return json.Marshal(&au_jwt_lic{
		Comment:    comment,
		ExpireDate: time.Unix(expire_unix, 0).Format("2006-01-02 15:04:05"),
		Token:      base64.StdEncoding.EncodeToString([]byte(jwt)),
	})
}

func AuJwtValid(jwt_token, secret string) bool {
	token, err := jwt.ParseWithClaims(jwt_token, &au_jwt{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		log.Printf("<ERR> [lic error] jwt parse error %s\n", err.Error())
		return false
	}
	if _, ok := token.Claims.(*au_jwt); ok && token.Valid {
		return true
	} else {
		log.Println("<ERR> [lic error] jwt is not valid")
		return false
	}
}

func AuJwtValidFile(lic_path, secret string) bool {
	file, err := os.Open(lic_path)
	if err != nil {
		log.Printf("<ERR> [lic error] open %s failed.\n", lic_path)
		return false
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	au_lic := &au_jwt_lic{}
	err = decoder.Decode(au_lic)
	if err != nil {
		log.Printf("<ERR> [lic error] parse %s failed. \n", lic_path)
		return false
	}
	jwt_token, err := base64.StdEncoding.DecodeString(au_lic.Token)
	if err != nil {
		log.Println("<ERR> [lic error] jwt token decode base64 failed.")
		return false
	}

	return AuJwtValid(string(jwt_token), secret)

}
