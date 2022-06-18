package train3

import (
	"testing"
)

func expectSpeed(t *testing.T, speeder Speeder, speed int) {
	actual := speeder.Speed()
	if actual != speed {
		t.Fatal("expected", speed, "got", actual)
	}
}

func TestEngine(t *testing.T) {
	e := NewEngine()
	go e.Run()
	defer e.Stop()
	expectSpeed(t, e, 0)
	e.Accel()
	expectSpeed(t, e, 10)
	e.Decel()
	expectSpeed(t, e, 0)
}

func TestEngineAccel(t *testing.T) {
	e := newEngine()

	done := make(chan bool)
	go func() {
		e.Accel()
		close(done)
	}()

	c, ok := <-e.accelcmd
	if c == nil || !ok {
		t.Fatal("expected acceleration command")
	}

	close(c.done)
	<-done
}
