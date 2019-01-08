package au

import "testing"
import "time"

func TestAuJwtToken(t *testing.T) {
	id := "190108141700"
	secrect := "world"
	expire := int64(15)
	token, _ := AuJwtToken(id, secrect, expire)
	t.Log(token)
	valid, new_token := AuJwtTokenValid(token, id, secrect, expire+15)
	if !valid {
		t.Errorf("should valid\n")
	}
	t.Log(valid)
	t.Log(new_token)
	time.Sleep(time.Duration(expire+1) * time.Second)
	valid, _ = AuJwtTokenValid(token, id, secrect, expire)
	t.Log(valid)
	if valid {
		t.Errorf("should invalid\n")
	}
	valid, _ = AuJwtTokenValid(new_token, id, secrect, expire)
	t.Log(valid)
	if !valid {
		t.Errorf("new token should valid\n")
	}
}
