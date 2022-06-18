package train

import Engine "hello/engine"

type Train struct {
	engine *Engine.Engine
}

func NewTrain() *Train {
	return &Train{
		engine: Engine.NewEngine(),
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
