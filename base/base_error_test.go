package base

import (
	"testing"
)

func TestWhereAmI(t *testing.T) {
	where := WhereAmI()
	t.Log(where)
	where = WhereAmI(2)
	t.Log(where)
}
