package main

import "fmt"

type Player struct {
	Name   string
	Health int
}

func (p *Player) Heal(health int) {
	p.Health += health
}

func main() {
	player := Player{"P1", 50}
	fmt.Println(player)

	player.Heal(10)
	fmt.Println(player)
}
