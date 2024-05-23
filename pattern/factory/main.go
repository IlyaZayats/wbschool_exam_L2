package main

import "fmt"

type ProductInterface interface {
	doSomething()
}

type ConcreteProductA struct {
	Name string
}

func (p *ConcreteProductA) doSomething() {
	fmt.Println(p.Name)
}

type ConcreteProductB struct {
	Name string
}

func (p *ConcreteProductB) doSomething() {
	fmt.Println(p.Name)
}

type CreatorInterface interface {
	build() ProductInterface
}

type ConcreteCreatorA struct{}

func (c *ConcreteCreatorA) build() ProductInterface {
	return &ConcreteProductA{Name: "ConcreteProductA"}
}

type ConcreteCreatorB struct{}

func (c *ConcreteCreatorB) build() ProductInterface {
	return &ConcreteProductB{Name: "ConcreteProductB"}
}

func main() {
	creators := []CreatorInterface{&ConcreteCreatorA{}, &ConcreteCreatorB{}}
	for _, c := range creators {
		product := c.build()
		product.doSomething()
	}
}
