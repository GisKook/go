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
func where_am_i(depthList ...int) string {
	var depth int
	if depthList == nil {
		depth = 1
	} else {
		depth = depthList[0]
	}
	function, file, line, _ := runtime.Caller(depth)
	return fmt.Sprintf("File: %s  Function: %s Line: %d", file, runtime.FuncForPC(function).Name(), line)
}

func ErrorCheck(err error) {
	if err != nil {
		log.Println(where_am_i())
	}
}
