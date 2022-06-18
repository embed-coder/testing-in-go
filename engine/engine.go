package engine

type Engine struct {
	speed int
}

func NewEngine() *Engine {
	return &Engine{}
}
func (e *Engine) Speed() int {
	return e.speed
}
func (e *Engine) Accel() {
	e.speed += 10
}
func (e *Engine) Decel() {
	e.speed -= 10
}
