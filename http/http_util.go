package http

import (
	"github.com/giskook/go/base"
	"net"
	"net/http"
)

func GetClientIP(req *http.Request) string {
	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		base.ErrorCheck(err)
		return ""
	}

	return string(net.ParseIP(ip))
}
