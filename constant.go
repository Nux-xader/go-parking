package main

type InputLabels struct {
	label  string
	isChar bool
}
type Command struct {
	command string
	inputs  []InputLabels
}

var (
	Add = Command{
		command: "create_parking_lot",
		inputs: []InputLabels{
			{label: "Capacity"},
		},
	}

	Park = Command{
		command: "park",
		inputs: []InputLabels{
			{label: "Car Number", isChar: true},
		},
	}

	Leave = Command{
		command: "leave",
		inputs: []InputLabels{
			{label: "Car Number", isChar: true},
			{label: "Parking duration"},
		},
	}

	Status = Command{command: "status"}
)
