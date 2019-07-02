package base

import (
	"fmt"
	"log"
	"runtime"
	"runtime/debug"
)

// Copyright 2016 - by Jim Lawless
// License: MIT / X11
// See: http://www.mailsend-online.com/license2016.php
//
// This code may not conform to popular Go coding idioms
func where_am_i(err error, depthList ...int) string {
	var depth int
	if depthList == nil {
		depth = 1
	} else {
		depth = depthList[0]
	}
	function, file, line, _ := runtime.Caller(depth)

	tips := ""
	if err != nil {
		tips = fmt.Sprintf("Error %s File: %s  Function: %s Line: %d", err.Error(), file, runtime.FuncForPC(function).Name(), line)
	} else {
		tips = fmt.Sprintf("Error %s File: %s  Function: %s Line: %d", "no error", file, runtime.FuncForPC(function).Name(), line)
	}

	return tips
}

func ErrorCheck(err error) {
	log.Println(where_am_i(err, 2))
}

func ErrorCheckPlus(err error, extra ...string) {
	log.Println(where_am_i(err, 2), extra)
}

func ErrorCheckFatal(err error, extra string) {
	if err != nil {
		ErrorCheckWithLevel(err, UPPER_LEVEL)
		log.Fatal(err.Error(), extra)
	}
}

const (
	UPPER_LEVEL int = 3
)

func ErrorCheckWithLevel(err error, level int) {
	log.Println(where_am_i(err, level))
}

func ErrorPrintStack() {
	debug.PrintStack()
}
