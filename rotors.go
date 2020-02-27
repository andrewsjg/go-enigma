package main

// This file contains the basic rotor struct and sets of predefined rotors that represent historical enigma machines
// https://en.wikipedia.org/wiki/Enigma_rotor_details

// Rotor type represents an enigma machine rotor
type Rotor struct {
	// The index of the wiring slice represents 1 terminal on the rotor, the value at each index is the output letter associated with that terminal.
	// For this verison a string is used to make it easier to visualise and read, however it would probably be simpler in the implementation to use integers
	wiring [26]string

	ringOffset    int    // How far offset from position 1 on the rotor is the alphabet ring. This is used to map the indicator to wiring for cipher operations
	TurnOverPoint string // TODO: This will need updating to support multiple turnover points

	CurrentIndicator int  // The letter (or number) currently shown in the indicator window. This is how the rotor is rotated relative to the alphabet ring
	WillRotate       bool // Will the rotor rotate on the next keypress or stay static. Important for the turnover mechanism
}

// TestRotor creates a rotor with straight through wiring, a turnover point at Z, no offset on the alphabet ring and default indicator showing A when inserted
func TestRotor() Rotor {

	wiring := [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	r := Rotor{wiring, 0, "Z", 0, false}

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
