package criteria

import (
	"github.com/almighty/almighty-core/resource"
	"testing"
)

func TestGetParent(t *testing.T) {
	resource.Require(t, resource.UnitTest)
	l := Field("a")
	r := Literal(5)
	expr := Equals(l, r)
	if l.Parent() != expr {
		t.Errorf("parent should be %v, but is %v", expr, l.Parent())
	}
}
