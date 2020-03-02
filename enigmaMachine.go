package main

// Prototype 1

import (
	"errors"
	"strings"
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

	reflector [26]int

	// This is fixed in a physical machine, but can be changed here to emulate a commercial engima
	// OR military enigma. They had different entry wheels
	inputRotor [26]string
}

//Encrypt some text. This function will strip anything that isnt a letter
func (machine *EnigmaMachine) Encrypt(plaintext string) string {

	// Make everything uppercase
	plaintext = strings.ToUpper(plaintext)
	cipherText := ""
	cIdx := -1

	for _, r := range plaintext {

		// Ignore anything that isnt a letter of the alphabet
		if r >= 65 && r <= 90 {
			// The rotors rotate BEFORE the encipherment is done. So rotate the rotors first
			machine.RotateRotors()

			inputLetter := string(r)

			//TODO: Translate the inputLetter through the plugboard here

			// Find the index of the letter in the entry wheel. The commerical and military enigma's had different entry wheels
			// For this implementation we can also make any wiring on the entry wheel.
			// This gives us the terminal on the entry wheel so we know where the signal enters the first (right most) rotor.

			inputIndex := (sliceIndex(len(machine.inputRotor), func(i int) bool { return machine.inputRotor[i] == inputLetter }))

			// Send the signal through the rotors

			var outputLetter string

			for _, rotor := range machine.rotors {

				// Find the ciphertext letter in the wiring array. Use the alphabet ring offset and the current position of the rotort find the letter output by
				// the rotors wiring
				cIdx = (inputIndex - rotor.ringOffset + rotor.CurrentIndicator) % 26
				outputLetter = rotor.wiring[cIdx]

				// The input for the next rotor is the output index the letter based on the rotor wiring.
				// Adding 26 fixes for cases where the offset and indicator generate negative numbers
				inputIndex = (toAlphaNum(outputLetter) + rotor.ringOffset - rotor.CurrentIndicator + 26) % 26

			}

			// Pass the letter through the reflector
			inputIndex = machine.reflector[inputIndex]

			// Go back through the rotors from left to right. Use the inverse of the wiring to decode
			// This is repetative code, but I like it because it breaks the encryption stages into the same
			// stages as the physical machine. Its nice to see each step distinctly.

			for i := range machine.rotors {
				rotor := machine.rotors[len(machine.rotors)-1-i]

				// I think I can use my encodeLeft logic here instead. But this was easier to trace
				wiringInverse := generateInverseWiring(rotor.wiring)

				cIdx := (inputIndex + rotor.ringOffset + rotor.CurrentIndicator) % 26

				outputLetter = wiringInverse[cIdx]
				inputIndex = (toAlphaNum(outputLetter) + rotor.ringOffset - rotor.CurrentIndicator + 26) % 26

			}

			// Out the input rotor for the Final Encipherment. We do this because we can change the input rotor in our model.
			// There were different input rotor configurations between some variations of the machines. Particularly the input wiring
			// for of the commerical and military machines. These were a fixed part of the machine and not interchangeable.
			cipherText = cipherText + machine.inputRotor[inputIndex]
		}

	}

	return cipherText
}

//SetRotorPosition set a rotor to a position. Used in initial machine setup.
func (machine *EnigmaMachine) SetRotorPosition(rotorNumber int, startPos rune) {
	p := unicode.ToUpper(startPos)
	pos := int(p) - 65

	machine.rotors[rotorNumber].CurrentIndicator = pos
}

// RotateRotors rotates the rotors in accordance with the setup
func (machine *EnigmaMachine) RotateRotors() error {

	// TODO: Add test for double turnover

	// A machine must have at least 3 rotors to be valid. Check for that here
	if len(machine.rotors) < 3 {
		return errors.New("Not enough rotors installed in the machine")
	}

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

			newIndictor := (rotor.CurrentIndicator + 1) % 26
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

//toChar takes and int from 0-25 and returns the letter at that position in the alphabet
func toChar(i int) rune {
	return rune('A' + i)
}

func toAlphaNum(s string) int {
	outVal := -1

	for _, r := range s {
		outVal = int(r)

		if r >= 65 && r <= 90 {
			outVal = int(r) - 65
		}
	}

	return outVal
}
