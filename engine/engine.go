package engine

import (
	"github.com/coreos/coreinit/machine"
	"github.com/coreos/coreinit/registry"
)

type Engine struct {
	scheduler *Scheduler
	listener  *EventListener
	machine   *machine.Machine
}

func New(reg *registry.Registry, mach *machine.Machine) *Engine {
	scheduler := NewScheduler(reg, mach)
	listener := NewEventListener(reg)
	engine := Engine{scheduler, listener, mach}
	return &engine
}

func (engine *Engine) Run() {
	go engine.scheduler.Schedule()
	engine.listener.Listen()
}
