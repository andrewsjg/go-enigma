package main

// Prototype 1

import (
	"errors"
	"unicode"
)

// RIGHTROTOR - The right most rotor in the machine is always at array position 0
const RIGHTROTOR = 0

// EnigmaMachine contains all the parts of the machine
type EnigmaMachine struct {
	// 0 = Right most rotor
	rotors []*Rotor

	// Not yet implemented
	plugBoard int

	// Not yet implemented
	reflector int

	// This is fixed in a physical machine, but can be changed here to emulate a commercial engima
	// OR military enigma. They had different entry wheels
	entrywheel [26]string
}

/*
//Encrypt some text. This function will strip anything that isnt a letter
func (machine *EnigmaMachine) Encrypt(plaintext string) string {
	// The rotors rotate BEFORE the encipherment is done. So rotate the rotors first
	machine.RotateRotors()

	for _, r := range plaintext {
		c := string(r)

		//TODO: Translate C (current letter) through the plugboard here

		// Find where the current letter hits the entrywheel
		inputLocation := sliceIndex(len(machine.entrywheel), func(i int) bool { return machine.entrywheel[i] == c })

		// Run it through the rotors
		for _,rotor := range machine.rotors {

		}
	}

	return ""
}
*/

//SetRotorPosition set a rotor to a position. Used in initial machine setup.
func (machine *EnigmaMachine) SetRotorPosition(rotorNumber int, startPos rune) {
	p := unicode.ToUpper(startPos)
	pos := int(p) - 64

	machine.rotors[rotorNumber].CurrentIndicator = pos
}

// RotateRotors rotates the rotors in accordance with the setup
func (machine *EnigmaMachine) RotateRotors() error {

	// A machine must have at least 3 rotors to be valid. Check for that here
	if len(machine.rotors) < 3 {
		return errors.New("Not enough rotors installed in the machine")
	}

	/* DEBUG: Print the rotor state
	for _, rotor := range machine.rotors {
		c := toChar(rotor.CurrentIndicator)
		fmt.Printf("%s ", string(c))
	}
	fmt.Println() */

	// The right most rotor always rotates
	machine.rotors[RIGHTROTOR].WillRotate = true

	for rotorNum, rotor := range machine.rotors {

		CurrentIndicatorChar := string(toChar(rotor.CurrentIndicator))
		// If a rotor is at its turnover point and it will rotate, then trigger a rotate of the rotor to the left
		if (CurrentIndicatorChar == rotor.TurnOverPoint) && rotor.WillRotate {
			// Dont attempt to rotate anything if the current rotor is the left most
			if rotorNum+1 < len(machine.rotors) {
				machine.rotors[rotorNum+1].WillRotate = true
			}
		}

		if rotor.WillRotate {

			// This is silly. There must be a better way, but this can stay for now
			newIndictor := (rotor.CurrentIndicator + 1) % 27
			if newIndictor == 0 {
				newIndictor = 1
			}

			rotor.CurrentIndicator = newIndictor
			rotor.WillRotate = false
		}

	}

	return nil
}

// Helper function to return the next input terminal according to the wiring array
func getNextInputTerminal(r Rotor, currentInputTerminal string) string {
	currentIndex := sliceIndex(len(r.wiring), func(i int) bool { return r.wiring[i] == currentInputTerminal })

	// Should probably make SliceIndex return an error
	if currentIndex != -1 {

		newIndex := currentIndex + 1

		if newIndex == 26 {
			newIndex = 0
		}

		return r.wiring[newIndex]
	}

	// TODO: FIX THIS. Should return an error
	return "?"
}

// A nifty go like way to find the index of an element in a slice. Thanks Stackoverflow: https://stackoverflow.com/a/18203895
func sliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

//toChar takes and int from 1-26 and returns the letter at that position in the alphabet
func toChar(i int) rune {
	return rune('A' - 1 + i)
}
