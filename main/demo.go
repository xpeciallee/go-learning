package main

import "fmt"

func main() {
	man := Man{Name: "Eric"}
	man.eat()
	newMan, error := man.execute(29, "male")
	fmt.Println(error)
	fmt.Println(newMan.Name)
}

type Peaple interface {
	eat()
	run()
}

type Man struct {
	Peaple
	Name string
}

func (man *Man) eat() {
	fmt.Println(man.Name + " eat fast")
}

func (man *Man) run() {
	fmt.Print(man.Name + " run fast")
}

func (man *Man) execute(age int, sex string) (newMan *Man, error string) {
	error = fmt.Sprintf("%d %s", age, sex)
	man.Name = " new Name"
	return man, error
}
