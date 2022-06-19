package mytime

import (
	"testing"
	"time"
)

func expectTime(t *testing.T, mt *MyTime, s time.Time) {
	actual := mt.GetTime()
	if actual != s {
		t.Fatal("expected", s, "got", actual)
	}
}

func TestTime(t *testing.T) {
	currentTime := time.Now()
	mt := NewMyTime()
	mt.SetTime()
	expectTime(t, &mt, currentTime)
}
