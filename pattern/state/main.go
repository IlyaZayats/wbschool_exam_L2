package main

import "fmt"

type StateContext struct {
	state State
}

func NewStateContext() *StateContext {
	return &StateContext{state: NewSolidState()}
}

func (c *StateContext) freeze() {
	fmt.Println("Freezing " + c.state.getName() + " substance...")
	c.state.freeze(c)
}

func (c *StateContext) heat() {
	fmt.Println("Heating " + c.state.getName() + " substance...")
	c.state.heat(c)
}

func (c *StateContext) setState(state State) {
	fmt.Println("Changing state to " + state.getName() + "...")
	c.state = state
}

func (c *StateContext) getState() State {
	return c.state
}

type State interface {
	getName() string
	freeze(c *StateContext)
	heat(c *StateContext)
}

func main() {
	sc := NewStateContext()
	sc.heat()
	sc.heat()
	sc.heat()
	sc.freeze()
	sc.freeze()
	sc.freeze()
}
