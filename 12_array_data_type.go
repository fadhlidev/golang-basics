package main

import "fmt"

func main() {
	var drills [4]string
	drills[0] = "Attack"
	drills[1] = "Defense"
	drills[2] = "Health"
	drills[3] = "Possession"
	fmt.Println(drills)
	fmt.Println(len(drills))

	var activeDrills = [...]string{"Attack", "Defense", "Health"}
	fmt.Println(activeDrills)
	activeDrills[2] = "Possession"
	fmt.Println(activeDrills)
}
