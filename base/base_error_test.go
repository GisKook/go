package base

import (
	"errors"
	"log"
	"testing"
)

func TestErrorCheck(t *testing.T) {
	err := errors.New("I am an error")
	ErrorCheck(err)
	ErrorCheckPlus(err, "hello", "world")
	log.Println("hello", "world")
	//t.Log(where)
}
