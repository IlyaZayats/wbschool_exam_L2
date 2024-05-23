package main

import "fmt"

type Receiver struct {
	data string
}

func (r *Receiver) StartCommand(someStr string) {
	r.data = someStr
	fmt.Println(r.data)
}

func (r *Receiver) StopCommand() {
	fmt.Println(r.data + " stopped")
	r.data = ""
}

type CommandInterface interface {
	Execute(someStr string)
	UnExecute()
}

type CommandA struct {
	receiver *Receiver
}

func (c *CommandA) Execute(someStr string) {
	c.receiver.StartCommand(someStr)
}

func (c *CommandA) UnExecute() {
	c.receiver.StopCommand()
}

type CommandB struct {
	receiver *Receiver
}

func (c *CommandB) Execute(someStr string) {
	c.receiver.StartCommand(someStr)
}

func (c *CommandB) UnExecute() {
	c.receiver.StopCommand()
}

type Invoker struct {
	commandA       *CommandA
	commandB       *CommandB
	currentCommand CommandInterface
}

func (i *Invoker) InvokeCommandA() {
	i.currentCommand = i.commandA
	i.currentCommand.Execute("CommandA")
}

func (i *Invoker) InvokeCommandB() {
	i.currentCommand = i.commandB
	i.currentCommand.Execute("CommandB")
}

func (i *Invoker) Stop() {
	if i.currentCommand != nil {
		i.currentCommand.UnExecute()
		i.currentCommand = nil
	} else {
		fmt.Println("Current command is not defined!")
	}
}

func main() {
	receiver := &Receiver{}
	invoker := Invoker{commandA: &CommandA{receiver: receiver}, commandB: &CommandB{receiver: receiver}}
	invoker.InvokeCommandA()
	invoker.Stop()
	invoker.InvokeCommandB()
	invoker.Stop()
}
