package main

// Prototype 2

import (
	"errors"
	"strings"
	"unicode"
)

// RIGHTROTOR - The right most rotor in the machine is always at array position 0
// This is a bit counter inutative when looking at the array.
const RIGHTROTOR = 0

//Plugboard repesents the plugboard wiring
type Plugboard struct {
	// or int? or Rune?
	wiring map[string]string
}

// EnigmaMachine represents an Enigma machine using the abstracted types for each of the
// components
type EnigmaMachine struct {
	rotors RotorSet

	plugBoard Plugboard

	reflector Reflector

	// This is fixed in a physical machine, but can be changed here to emulate a commercial engima
	// OR military enigma. They had different entry wheels
	inputRotor InputRotor
}

// CreateEnigmaMachine - Create an enigma machine wired up as specified
func CreateEnigmaMachine(rotors RotorSet, rotorStartPosition string, plugBoard Plugboard, reflector Reflector, inputRotor InputRotor) (EnigmaMachine, error) {
	var em EnigmaMachine

	em.rotors = rotors
	em.reflector = reflector
	em.inputRotor = inputRotor

	// This might look a bit silly. If we used numbered rotors in an array then it would be a simple loop
	// but we use a struct with named rotors to make things closer to a real enigma machine so we
	// have to do some funky translation here to make this work.

	for rotNum, r := range rotorStartPosition {

		switch rotNum {
		case 0:
			em.SetRotorPosition("LEFT", r)
		case 1:
			em.SetRotorPosition("MIDDLE", r)
		case 2:
			em.SetRotorPosition("RIGHT", r)
		case 3:
			em.SetRotorPosition("FORTH", r)
		}
	}

	// Check for an impossible plugboard configuration. Each letter can only be wired once. A configuration like A<->B and B<->C is is impossible
	// So check that any values dont appear as keys in our plugboard.wiring map

	for _, val := range plugBoard.wiring {
		for innerKey := range plugBoard.wiring {
			if val == innerKey {
				return em, errors.New("Impossible plugboard setting. Check that no letter is plugged twice or plugged to itself")
			}
		}
	}

	em.plugBoard = plugBoard

	return em, nil
}

// Encrypt some text
func (machine *EnigmaMachine) Encrypt(plaintext string) string {

	//TODO: Return text in blocks of 5 letters

	var rotors []*Rotor

	// Construct the rotors array from the machines rotors
	if machine.rotors.forth != nil {
		rotors = []*Rotor{machine.rotors.right, machine.rotors.middle, machine.rotors.left, machine.rotors.forth}

	} else {
		rotors = []*Rotor{machine.rotors.right, machine.rotors.middle, machine.rotors.left}
	}

	// Make everything uppercase
	plaintext = strings.ToUpper(plaintext)
	cipherText := ""
	cIdx := -1

	// Reverse the plugboard map so we can go from a value to a key as well as key to value. Do this here so that we dont
	// need to iterate through the map on each iteration of the loop below

	var reversePlugBoard map[string]string
	reversePlugBoard = make(map[string]string)

	for key, val := range machine.plugBoard.wiring {
		reversePlugBoard[val] = key
	}

	for _, r := range plaintext {

		// Ignore anything that isnt a letter of the alphabet
		if r >= 65 && r <= 90 {
			// The rotors rotate BEFORE the encipherment is done. So rotate the rotors first
			machine.RotateRotors()

			inputLetter := string(r)

			// We need to translate the letter through the plugboard here

			if val, plugged := machine.plugBoard.wiring[inputLetter]; plugged {
				inputLetter = val
			} else if val, plugged := reversePlugBoard[inputLetter]; plugged {
				inputLetter = val
			}

			// Find the index of the letter in the entry wheel. The commercial and military enigma's had different entry wheels
			// For this implementation we can also make any wiring on the entry wheel.
			// This gives us the terminal on the entry wheel so we know where the signal enters the first (right most) rotor.

			inputIndex := (sliceIndex(len(machine.inputRotor.wiring), func(i int) bool { return machine.inputRotor.wiring[i] == inputLetter }))

			// Send the signal through the rotors

			var outputLetter string

			for _, rotor := range rotors {

				// Find the ciphertext letter in the wiring array. Use the alphabet ring offset and the current position of the rotort find the letter output by
				// the rotors wiring
				cIdx = (inputIndex - rotor.ringOffset + rotor.CurrentIndicator) % 26
				outputLetter = rotor.wiring[cIdx]

				// The input for the next rotor is the output index the letter based on the rotor wiring.
				// Adding 26 fixes for cases where the offset and indicator generate negative numbers
				inputIndex = (toAlphaNum(outputLetter) + rotor.ringOffset - rotor.CurrentIndicator + 26) % 26

			}

			// Pass the letter through the reflector
			inputIndex = machine.reflector.wiring[inputIndex]

			// Go back through the rotors from left to right. Use the inverse of the wiring to decode
			// This is repetative code, but I like it because it breaks the encryption stages into the same
			// stages as the physical machine. Its nice to see each step distinctly.

			for i := range rotors {
				rotor := rotors[len(rotors)-1-i]

				// I think I can use my encodeLeft logic here instead. But this was easier to trace
				wiringInverse := generateInverseWiring(rotor.wiring)

				cIdx := (inputIndex + rotor.ringOffset + rotor.CurrentIndicator) % 26

				outputLetter = wiringInverse[cIdx]
				inputIndex = (toAlphaNum(outputLetter) + rotor.ringOffset - rotor.CurrentIndicator + 26) % 26

			}

			outputLetter = machine.inputRotor.wiring[inputIndex]

			// Back through the plugboard

			if val, plugged := reversePlugBoard[outputLetter]; plugged {
				outputLetter = val
			} else if val, plugged := machine.plugBoard.wiring[outputLetter]; plugged {
				outputLetter = val
			}

			// Out the input rotor for the Final Encipherment. We do this because we can change the input rotor in our model.
			// There were different input rotor configurations between some variations of the machines. Particularly the input wiring
			// for of the commercial and military machines. These were a fixed part of the machine and not interchangeable.
			cipherText = cipherText + outputLetter
		}

	}

	return cipherText
}

//SetRotorPosition set a rotor to a position. Used in initial machine setup.
func (machine *EnigmaMachine) SetRotorPosition(rotorPos string, startPos rune) {
	p := unicode.ToUpper(startPos)
	pos := int(p) - 65

	// There might be a clever Go way of doing this...
	rp := strings.ToUpper(rotorPos)

	switch rp {
	case "LEFT":
		machine.rotors.left.CurrentIndicator = pos
	case "MIDDLE":
		machine.rotors.middle.CurrentIndicator = pos
	case "RIGHT":
		machine.rotors.right.CurrentIndicator = pos
	case "FORTH":
		if machine.rotors.forth != nil {
			machine.rotors.forth.CurrentIndicator = pos
		}
	}
}

// RotateRotors rotates the rotors in accordance with the setup
func (machine *EnigmaMachine) RotateRotors() error {

	// Build the rotors array from the rotor set
	var rotors []*Rotor

	// Construct the rotors array from the machines rotors
	if machine.rotors.forth != nil {
		rotors = []*Rotor{machine.rotors.right, machine.rotors.middle, machine.rotors.left, machine.rotors.forth}

	} else {
		rotors = []*Rotor{machine.rotors.right, machine.rotors.middle, machine.rotors.left}
	}

	// TODO: Add test for double turnover

	// A machine must have at least 3 rotors to be valid. Check for that here. Probably dont need to
	// do this for this version
	if len(rotors) < 3 {
		return errors.New("Not enough rotors installed in the machine")
	}

	// The right most rotor always rotates
	rotors[RIGHTROTOR].WillRotate = true

	for rotorNum, rotor := range rotors {

		CurrentIndicatorChar := string(toChar(rotor.CurrentIndicator))
		// If a rotor is at its turnover point and it will rotate, then trigger a rotate of the rotor to the left
		if (CurrentIndicatorChar == rotor.TurnOverPoint) && rotor.WillRotate {
			// Dont attempt to rotate anything if the current rotor is the left most
			if rotorNum+1 < len(rotors) {
				rotors[rotorNum+1].WillRotate = true
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
