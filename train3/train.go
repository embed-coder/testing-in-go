package train3

import (
	"time"
)

type Speeder interface {
	Speed() int
}

type Engine interface {
	Speeder
	Accel()
	Decel()
	Stop()
	Run()
}

type accelCommand struct {
	done chan bool
}

func newAccelCommand() *accelCommand {
	return &accelCommand{done: make(chan bool)}
}

type engine struct {
	speed    int
	stop     chan bool
	done     chan bool
	accelcmd chan *accelCommand
}

func newEngine() *engine {
	return &engine{
		accelcmd: make(chan *accelCommand),
		stop:     make(chan bool),
		done:     make(chan bool),
	}
}

func NewEngine() Engine {
	return newEngine()
}

func (e *engine) Run() {
	for {
		select {
		case c := <-e.accelcmd:
			e.speed += 10
			time.Sleep(100 * time.Millisecond)
			close(c.done)
		case <-e.stop:
			close(e.done)
			return
		}
	}
}

func (e *engine) Stop() {
	e.stop <- true
	<-e.done
}

func (e *engine) Speed() int {
	return e.speed
}

func (e *engine) Accel() {
	cmd := newAccelCommand()
	e.accelcmd <- cmd
	<-cmd.done
}

func (e *engine) Decel() {
	e.speed -= 10
}
