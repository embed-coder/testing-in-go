package train2

import "testing"

func expectSpeed(t *testing.T, speeder Speeder, speed int) {
	actual := speeder.Speed()
	if actual != speed {
		t.Fatal("expected", speed, "got", actual)
	}
}

func TestEngine(t *testing.T) {
	e := NewEngine()
	expectSpeed(t, e, 0)
	e.Accel()
	expectSpeed(t, e, 10)
	e.Decel()
	expectSpeed(t, e, 0)
}

type FakeEngine struct {
	AccelCalls int
	DecelCalls int
}

func NewFakeEngine() *FakeEngine {
	return &FakeEngine{}
}

func (e *FakeEngine) Accel() {
	e.AccelCalls += 1
}

func (e *FakeEngine) Decel() {
	e.DecelCalls += 1
}

func (e *FakeEngine) Speed() int {
	return 0
}

func trainWithFakeEngine() (*Train, *FakeEngine) {
	t := NewTrain()
	e := NewFakeEngine()
	t.engine = e
	return t, e
}

func TestTrainGo(t *testing.T) {
	tr, e := trainWithFakeEngine()
	tr.Go()
	if e.AccelCalls != 2 {
		t.Error("expected", 2, "got", e.AccelCalls)
	}
}

func TestTrainStop(t *testing.T) {
	tr, e := trainWithFakeEngine()
	tr.Stop()
	if e.DecelCalls != 2 {
		t.Error("expected", 2, "got", e.DecelCalls)
	}
}
