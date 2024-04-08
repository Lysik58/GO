package main

import "fmt"

type User interface {
	getName(w int) string
	getAge(w int) int
	personIsDead() bool
}

type Person struct {
	Name    string
	Age     int
	Gender  string
	isAlive bool
}

type Programmer struct {
	Person
	Title   string
	isAdmin bool
	Level   string
}

func (p *Person) getName(w int) string {
	if w == 1 {
		fmt.Println(p.Name)
	}
	return p.Name
}

// Сделал параметр, чтобы было проще узнавать имя и сколько лет
func (p *Person) getAge(w int) int {
	if w == 1 {
		fmt.Println(p.Age)
	}
	return p.Age
}

func (p *Person) personIsDead() {
	if p.isAlive == true {
		fmt.Printf("%s жив", p.Name)
		fmt.Println()
	} else {
		fmt.Println("Потрачено")
	}
}

func (p *Person) personKill() {
	if p.isAlive == true {
		p.isAlive = false
		fmt.Println("Вам наследство!")
	} else {
		fmt.Println("Поздно")
	}
}

// Без звездочки, потому что этого требует задание. Я помню про правило разрабов
func (p Person) personIsWalking() {
	fmt.Printf("%s гуляет", p.Name)
	fmt.Println()
}

func main() {
	user0 := Person{
		Name:    "Jack",
		Age:     54,
		isAlive: false,
	}

	user1 := Person{
		Name:    "Denis",
		Age:     20,
		isAlive: true,
	}

	user0.getAge(1)
	user1.personIsDead()
	user1.personIsWalking()
	user1.personKill()
	user1.personIsDead()
	user1.personKill()
}
