package base

import (
	"fmt"
	"log"
	"runtime"
)

// Copyright 2016 - by Jim Lawless
// License: MIT / X11
// See: http://www.mailsend-online.com/license2016.php
//
// This code may not conform to popular Go coding idioms
func where_am_i(err error, depthList ...int) string {
	if err == nil {
		return ""
	}
	var depth int
	if depthList == nil {
		depth = 1
	} else {
		depth = depthList[0]
	}
	function, file, line, _ := runtime.Caller(depth)
	return fmt.Sprintf("Error %s File: %s  Function: %s Line: %d", err.Error(), file, runtime.FuncForPC(function).Name(), line)
}

func ErrorCheck(err error) {
	log.Println(where_am_i(err, 2))
}
