package main

// This file contains the basic rotor struct and sets of predefined rotors that represent historical enigma machines
// https://en.wikipedia.org/wiki/Enigma_rotor_details

// Rotor type represents an enigma machine rotor
type Rotor struct {
	wiring        [26]string
	ringSetting   int
	TurnOverPoint string

	CurrentInputTerminal string
}

// TestRotor creates a rotor with straight through wiring
func TestRotor() Rotor {
	// 1=A,
	wiring := [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	r := Rotor{wiring, 1, "Z", "A"}

	return r
}
