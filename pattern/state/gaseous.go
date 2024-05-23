package main

import "fmt"

type GaseousState struct {
	name string
}

func NewGaseousState() *GaseousState {
	return &GaseousState{name: "gaseous"}
}

func (s *GaseousState) getName() string {
	return s.name
}

func (s *GaseousState) freeze(c *StateContext) {
	c.setState(NewLiquidState())
}

func (s *GaseousState) heat(c *StateContext) {
	fmt.Println("Nothing happens.")
}
