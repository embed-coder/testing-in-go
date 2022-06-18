package train2

type Speeder interface {
	Speed() int
}

type Engine interface {
	Speeder
	Accel()
	Decel()
}

type engine struct {
	speed int
}

func NewEngine() Engine {
	return &engine{}
}

func (e *engine) Speed() int {
	return e.speed
}

func (e *engine) Accel() {
	e.speed += 10
}

func (e *engine) Decel() {
	e.speed -= 10
}

type Train struct {
	engine Engine
}

func NewTrain() *Train {
	return &Train{
		engine: NewEngine(),
	}
}

func (t *Train) Go() {
	t.engine.Accel()
	t.engine.Accel()
}

func (t *Train) Stop() {
	t.engine.Decel()
	t.engine.Decel()
}

func (t *Train) Speed() int {
	return t.engine.Speed()
}
