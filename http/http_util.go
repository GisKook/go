package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/giskook/go/base"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
)

func GetClientIP(req *http.Request) string {
	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		base.ErrorCheck(err)
		return ""
	}

	return string(net.ParseIP(ip))
}

type Response struct {
	Code string      `json:"code"`
	Desc string      `json:"desc"`
	Data interface{} `json:"data,omitempty"`
}

func EncodeResponse(w http.ResponseWriter, code string, data interface{}, errmsg string) {
	gr := &Response{
		Code: code,
		Desc: errmsg,
		Data: data,
	}
	MarshalJson(w, gr)
}

func RecordReq(r *http.Request) {
	v, e := httputil.DumpRequest(r, true)
	if e != nil {
		return
	}
	log.Println(string(v))
}

// MarshalJson 把对象以json格式放到response中
func MarshalJson(w http.ResponseWriter, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	fmt.Fprint(w, string(data))
	log.Println(string(data))
	return nil
}

// UnMarshalJson 从request中取出对象
func UnMarshalJson(req *http.Request, v interface{}) error {
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	json.Unmarshal([]byte(bytes.NewBuffer(result).String()), v)
	return nil
}
