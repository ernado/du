package du

import (
	"testing"
)

func TestGet(t *testing.T) {
	f, err := Get(".")
	if err != nil {
		t.Fatal(err)
	}
	if f.Total == 0 {
		t.Error("Total should not be zero.")
	}
}
