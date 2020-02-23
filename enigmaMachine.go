package main

// Prototype 1

import (
	"errors"
	"fmt"
)

// RIGHTROTOR - The right most rotor in the machine is always at array position 0
const RIGHTROTOR = 0

// EnigmaMachine contains all the parts of the machine
type EnigmaMachine struct {
	// 0 = Right most rotor
	rotors []*Rotor
	plugBoard int
	reflector  int
}

//SetRotorPosition set a rotor to a position. Used in initial machine setup. 
func (machine *EnigmaMachine) SetRotorPosition(rotorNumber int, startPos string) {
	machine.rotors[rotorNumber].CurrentInputTerminal = startPos
}

// RotateRotors rotates the rotors in accordance with the setup
func (machine *EnigmaMachine) RotateRotors() error {


	// A machine must have at least 3 rotors to be valid. Check for that here
	if len(machine.rotors) < 3 {
		return errors.New("Not enough rotors installed in the machine")
	}

	// DEBUG: Print the rotor state
	for _,rotor := range machine.rotors {
		fmt.Printf("%s,",rotor.CurrentInputTerminal)
	}
	fmt.Println()

	// The right most rotor always rotates
	machine.rotors[RIGHTROTOR].WillRotate = true	

	for rotorNum,rotor := range machine.rotors {

		if rotor.CurrentInputTerminal == rotor.TurnoverPoint {
			if rotoNum + 1 < len(machine.rotors) {
				machine.rotor[rotorNum + 1].WillRotate	= true
			}
		}

		if rotor.WillRotate {
			rotor.CurrentInputTerminal = getNextInputTerminal(*rotor,rotor.CurrentInputTerminal)
		}

	}

	return nil
}

func getNextInputTerminal(r Rotor, currentInputTerminal string ) string {
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

func sliceIndex(limit int, predicate func(i int) bool) int {
    for i := 0; i < limit; i++ {
        if predicate(i) {
            return i
        }
    }
    return -1
}

/*
//Rotate the rotor
func (rotor *Rotor) Rotate() {
	rotor.CurrentInputTerminal = (rotor.CurrentInputTerminal % 26) + 1

	if rotor.CurrentInputTerminal == 0 {
		rotor.CurrentInputTerminal = 1
	}
}

*/

