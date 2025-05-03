package main

import "fmt"

type ParkingLot struct {
	capacity  int
	slots     []string
	carToSlot map[string]int
}

func NewParkingLot(capacity int) *ParkingLot {
	return &ParkingLot{
		capacity:  capacity,
		slots:     make([]string, capacity),
		carToSlot: make(map[string]int),
	}
}

func (pl *ParkingLot) Park(carNumber string) {
	for i := 0; i < pl.capacity; i++ {
		if pl.slots[i] == "" {
			pl.slots[i] = carNumber
			pl.carToSlot[carNumber] = i + 1
			fmt.Printf("Allocated slot number: %d\n", i+1)
			return
		}
	}
	fmt.Println("Sorry, parking lot is full")
}

func (pl *ParkingLot) Leave(carNumber string, hours int) {
	slotNum, ok := pl.carToSlot[carNumber]
	if !ok {
		fmt.Printf("Registration number %s not found\n", carNumber)
		return
	}
	charge := 10
	if hours > 2 {
		charge += 10 * (hours - 2)
	}
	fmt.Printf("Registration number %s with Slot Number %d is free with Charge $%d\n", carNumber, slotNum, charge)
	pl.slots[slotNum-1] = ""
	delete(pl.carToSlot, carNumber)
}

func (pl *ParkingLot) Status() {
	fmt.Println("Slot No.\t Registration No.")
	for i := 0; i < pl.capacity; i++ {
		if pl.slots[i] != "" {
			fmt.Printf("%d\t\t %s\n", i+1, pl.slots[i])
		}
	}
}
