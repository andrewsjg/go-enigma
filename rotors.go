package main

// This file contains the basic rotor struct and sets of predefined rotors that represent historical enigma machines
// https://en.wikipedia.org/wiki/Enigma_rotor_details

// TODO: Add a function for creating a complete set of rotors?

// Rotor type represents an enigma machine rotor
type Rotor struct {
	// The index of the wiring slice represents 1 terminal on the rotor, the value at each index is the output letter associated with that terminal.
	// For this verison a string is used to make it easier to visualise and read (for me!), however it would probably be simpler in the implementation to use integers
	wiring [26]string

	ringOffset    int    // How far offset from position 1 on the rotor is the alphabet ring. This is used to map the indicator to wiring for cipher operations
	TurnOverPoint string // TODO: This will need updating to support multiple turnover points

	CurrentIndicator int  // The letter (or number) currently shown in the indicator window. This is how the rotor is rotated relative to the alphabet ring
	WillRotate       bool // Will the rotor rotate on the next keypress or stay static. Important for the turnover mechanism
}

// RotorSet is an abstraction to make it easier to visualise were the rotors are physically
// located in the machine
type RotorSet struct {
	left   *Rotor
	middle *Rotor
	right  *Rotor
	forth  *Rotor // For machines that had four rotor
}

// Reflector represents the fixed reflector
type Reflector struct {
	wiring [26]int
}

//InputRotor Represents an Input Rotor
type InputRotor struct {
	wiring [26]string
}

// TestRotor creates a rotor with straight through wiring, a turnover point at Z, no offset on the alphabet ring and default indicator showing A when inserted
func TestRotor() Rotor {

	wiring := [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	r := Rotor{wiring, 0, "Z", 0, false}

	return r
}

// EncodeRight takes a letter and returns a letter transformed according to the wiring if the current enters from the right
func (rtr *Rotor) EncodeRight(letter string) string {
	return rtr.wiring[toAlphaNum(letter)]
}

// EncodeLeft akes a letter and returns a letter transformed according to the wiring if the current enters from the right
func (rtr *Rotor) EncodeLeft(letter string) string {
	idx := sliceIndex(len(rtr.wiring), func(i int) bool { return rtr.wiring[i] == letter })
	return string(toChar(idx))
}

func generateInverseWiring(wiring [26]string) [26]string {
	var inverse [26]string

	for i, letter := range wiring {
		invIndex := toAlphaNum(letter)
		inverse[invIndex] = string(toChar(i))
	}

	return inverse
}

//GenerateRotorI Creates a Rotor I configuration
func GenerateRotorI() Rotor {
	wiring := [26]string{"E", "K", "M", "F", "L", "G", "D", "Q", "V", "Z", "N", "T", "O", "W", "Y", "H", "X", "U", "S", "P", "A", "I", "B", "R", "C", "J"}
	r := Rotor{wiring, 0, "Q", 0, false}

	return r
}

//GenerateRotorII Creates a Rotor II configuration
func GenerateRotorII() Rotor {
	wiring := [26]string{"A", "J", "D", "K", "S", "I", "R", "U", "X", "B", "L", "H", "W", "T", "M", "C", "Q", "G", "Z", "N", "P", "Y", "F", "V", "O", "E"}
	r := Rotor{wiring, 0, "E", 0, false}

	return r
}

//GenerateRotorIII Creates a Rotor III configuration
func GenerateRotorIII() Rotor {
	wiring := [26]string{"B", "D", "F", "H", "J", "L", "C", "P", "R", "T", "X", "V", "Z", "N", "Y", "E", "I", "W", "G", "A", "K", "M", "U", "S", "Q", "O"}
	r := Rotor{wiring, 0, "V", 0, false}

	return r
}

//GenerateRotorIV Creates a Rotor IV configuration
func GenerateRotorIV() Rotor {
	wiring := [26]string{"E", "S", "O", "V", "P", "Z", "J", "A", "Y", "Q", "U", "I", "R", "H", "X", "L", "N", "F", "T", "G", "K", "D", "C", "M", "W", "B"}
	r := Rotor{wiring, 0, "J", 0, false}

	return r
}

//GenerateRotorV Creates a Rotor V configuration
func GenerateRotorV() Rotor {
	wiring := [26]string{"V", "Z", "B", "R", "G", "I", "T", "Y", "U", "P", "S", "D", "N", "H", "L", "X", "A", "W", "M", "J", "Q", "O", "F", "E", "C", "K"}
	r := Rotor{wiring, 0, "Z", 0, false}

	return r
}

// GenerateCommercialInputRotor builds a commercial entry wheel
func GenerateCommercialInputRotor() InputRotor {
	ir := InputRotor{[26]string{"Q", "W", "E", "R", "T", "Z", "U", "I", "O", "P", "A", "S", "D", "F", "G", "H", "J", "K", "L", "Y", "X", "C", "V", "B", "N", "M"}}
	return ir
}

// GenerateMilitaryInputRotor builds a commercial entry wheel
func GenerateMilitaryInputRotor() InputRotor {
	ir := InputRotor{[26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}}
	return ir
}

// GenerateReflectorA generates a type A reflector. The each array index represents a letter of the alphabet
// The value is the number of the letter that is reflected back
func GenerateReflectorA() Reflector {
	ref := Reflector{[26]int{4, 9, 12, 25, 0, 11, 24, 23, 21, 1, 22, 5, 2, 17, 16, 20, 14, 13, 19, 18, 15, 8, 10, 7, 6, 3}}
	return ref

}

// GenerateReflectorB generates a type B reflector. The each array index represents a letter of the alphabet
func GenerateReflectorB() Reflector {
	ref := Reflector{[26]int{24, 17, 20, 7, 16, 18, 11, 3, 15, 23, 13, 6, 14, 10, 12, 8, 4, 1, 5, 25, 2, 22, 21, 9, 0, 19}}
	return ref

}

// Quick and dirty utilityUsed to paste wiring configs from wikipedia and generate the strings for the rotor wirings generated above.
func wiringGenerator(input string) string {
	output := ""
	for _, r := range input {
		output = output + "\"" + string(r) + "\","
	}

	return output
}
