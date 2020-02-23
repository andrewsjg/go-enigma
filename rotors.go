package main

// This file contains the basic rotor struct and sets of predefined rotors that represent historical enigma machines
// https://en.wikipedia.org/wiki/Enigma_rotor_details

// Rotor type represents an enigma machine rotor
type Rotor struct {
	wiring        [26]string
	ringOffset    int
	TurnOverPoint string

	CurrentInputTerminal string // Change to current indicator?
	WillRotate           bool
}

// TestRotor creates a rotor with straight through wiring
func TestRotor() Rotor {

	wiring := [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	r := Rotor{wiring, 1, "Z", "A", false}

	return r
}

// GenerateCommercialEntryWheel builds a commercial entry wheel
func GenerateCommercialEntryWheel() [26]string {
	return [26]string{"Q", "W", "E", "R", "T", "Z", "U", "I", "O", "P", "A", "S", "D", "F", "G", "H", "J", "K", "L", "Y", "X", "C", "V", "B", "N", "M"}
}

// GenerateMilitaryEntryWheel builds a commercial entry wheel
func GenerateMilitaryEntryWheel() [26]string {
	return [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
}
