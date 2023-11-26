package main

import "fmt"

func main() {
	var drills = [4]string{"Attack", "Defense", "Health", "Possession"}
	fmt.Println(drills)

	var selectedDrills = drills[1:2]
	fmt.Println(len(selectedDrills))
	fmt.Println(cap(selectedDrills))

	fmt.Println(selectedDrills)
	selectedDrills = drills[:2]
	fmt.Println(selectedDrills)
	selectedDrills = drills[2:]
	fmt.Println(selectedDrills)
	selectedDrills = drills[:]
	fmt.Println(selectedDrills)

	fmt.Println(len(selectedDrills))
	fmt.Println(cap(selectedDrills))

	selectedDrills = append(selectedDrills, "Physic")
	fmt.Println(selectedDrills)
	fmt.Println(len(selectedDrills))
	fmt.Println(cap(selectedDrills))

	length := 2
	capacity := 3
	strategies := make([]string, length, capacity)
	strategies[0] = "Counter Attack"
	strategies[1] = "Offset Trap"
	fmt.Println(strategies)
	fmt.Println(len(strategies))
	fmt.Println(cap(strategies))

	strategies = append(strategies, "Conner Kick")
	fmt.Println(strategies)
	fmt.Println(len(strategies))
	fmt.Println(cap(strategies))

	copyStrategy := make([]string, len(strategies), cap(strategies))
	copy(copyStrategy, strategies)
	fmt.Println(copyStrategy)
	fmt.Println(len(copyStrategy))
	fmt.Println(cap(copyStrategy))

	activeDrills := []string{"Attack", "Defense", "Possession"}
	fmt.Println(activeDrills)
}
