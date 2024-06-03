package main

import "fmt"

type PC struct {
	Motherboard string
	CPU         string
	GPU         string
	AVRFps      float32
	Title       string
}

var spec = map[int]string{
	0: "Office",
	1: "Middle",
	2: "Gaming",
}

func (p *PC) changeTitlePc(title string) {
	p.Title = title
}

func main() {

	PC0 := PC{
		Motherboard: "Asus",
		CPU:         "Intel i5",
		GPU:         "RTX 4090",
		AVRFps:      90.3,
		Title:       spec[1],
	}
	PC0.changeTitlePc(spec[2])
	fmt.Println(PC0)
}

// var numbers [5]int

// var users []string

// func (p *PC) getSpec
