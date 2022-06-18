package hello

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestTruth(t *testing.T) {
	if true != true {
		t.Error("everything I know is wrong")
	}
}

func TestBroken(t *testing.T) {
	if true != false {
		t.Error("expected", true, "got", false)
	}
	if 1+1 != 4 {
		t.Fatal("Can't proceed!", 1+1, 4)
	}

	// Go will not print "Test ends" because it stops with t.Fatal already
	fmt.Println("Test ends")
}

func TestError(t *testing.T) {
	dest := make([]byte, 0)
	if _, err := hex.Decode(dest, []byte{8}); err != nil {
		t.Error(err)
	}
}
