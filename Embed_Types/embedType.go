package Embed_Types

import "fmt"

type Person struct {
	Name string
}

type Android struct {
	Person
	Model string
}

func (p *Android) setName(Name string) {
	p.Name = Name
}

func (p *Android) setModel(Model string) {
	p.Model = Model
}

func (p *Android) Talk() {
	fmt.Println("Hi, my name is", p.Name)
	fmt.Println("Hi, my model is", p.Model)
}

func Embed_Test() {
	a := new(Android)
	a.setName("Ramu")
	a.setModel("Human")
	a.Talk()
}
