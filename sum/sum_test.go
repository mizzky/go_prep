package sum

import (
	"testing"
)

func TestSum(t *testing.T) {
	actual := Sum(1, 2)
	expect := 3
	if actual != expect {
		t.Errorf("got %v\n want %v", actual, expect)
	}
}
