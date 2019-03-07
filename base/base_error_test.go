package base

import (
	"errors"
	"testing"
)

func TestErrorCheck(t *testing.T) {
	err := errors.New("I am an error")
	ErrorCheck(err)
	//t.Log(where)
}
