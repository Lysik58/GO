package main

import (
	"fmt"
)

func Tab(i int) {
	if i == 1 {
		fmt.Println()
	}
	if i == 2 {
		fmt.Println()
		fmt.Println()
	}
}

type User interface {
	getName(w int) string
	getAge(w int) int
	personIsDead() bool
}

type Progger interface {
	getTitle() string
	makeProgAdmin()
	progIsAdmin() bool
}

type Person struct {
	Name    string
	Age     uint
	Gender  string
	isAlive bool
}

type Programmer struct {
	Person
	Title   string
	isAdmin bool
	Level   uint
}

var progTitle = map[int]string{
	0: "Junior",
	1: "Middle",
	2: "Senior",
}

func (p *Person) getName(w int) string {
	if w == 1 {
		fmt.Println(p.Name)
	}
	return p.Name
}

// Сделал параметр, чтобы было проще узнавать имя и сколько лет
func (p *Person) getAge(w int) uint {
	if w == 1 {
		fmt.Println(p.Age)
	}
	return p.Age
}

func (p *Person) personIsDead() {
	if p.isAlive == true {
		fmt.Printf("%s жив \n", p.Name)
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
	fmt.Printf("%s гуляет \n", p.Name)
}

func (p *Programmer) progIsAdmin() {
	if p.isAdmin == true {
		fmt.Printf("%s является администратором \n", p.Name)
	} else if p.isAdmin == false { //сделал строгое elseif для практики0_0
		fmt.Printf("%s не является администратором \n", p.Name)
	}
}

func (p *Programmer) makeProgAdmin() {
	if p.isAdmin == false {
		p.isAdmin = true
		fmt.Println("Готово")
	} else {
		fmt.Printf("%s уже админ \n", p.Name)
	}
}

func (p *Programmer) getTitle() {
	fmt.Println(p.Title)
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
	prog0 := Programmer{
		Person: Person{
			Name:    "Alyx",
			Age:     33,
			Gender:  "Male",
			isAlive: true,
		},
		Title:   progTitle[0],
		isAdmin: false,
		Level:   12,
	}

	user0.getAge(1)
	user1.personIsDead()
	user1.personIsWalking()
	user1.personKill()
	user1.personIsDead()
	user1.personKill()
	Tab(1)
	prog0.progIsAdmin()
	prog0.makeProgAdmin()
	prog0.progIsAdmin()
	prog0.makeProgAdmin()
	prog0.getTitle()
	prog0.personIsDead()
	prog0.personKill()

}
