package main

import (
	"testing"
	"fmt"
)


func TestRotateRotors(t *testing.T) {

	r1 := TestRotor()
	r2 := TestRotor()
	r3 := TestRotor()

	em := EnigmaMachine{[]*Rotor{&r1,&r2,&r3},1,1}

	// rotate once
	em.RotateRotors()

	if em.rotors[0].CurrentInputTerminal != "B" {
		t.Errorf("Rotate failed. Current input terminal is wrong. Expected B, got %s",em.rotors[0].CurrentInputTerminal)
	}

	// Rotate round to the last position
	for i := 2;i < len(em.rotors[0].wiring);i++ {
		em.RotateRotors()
	}

	if em.rotors[0].CurrentInputTerminal != "Z" {
		t.Errorf("Rotate failed. Current input terminal is wrong. Expected Z, got %s",em.rotors[0].CurrentInputTerminal)
	}

	// Rotate one past the last position. This tests the Modulo arithmetic and array bounds checks and turnover of the 2nd rotor
	em.RotateRotors()
	if em.rotors[0].CurrentInputTerminal != "A" {
		t.Errorf("Rotate failed. Current input terminal is wrong. Expected A, got %s",em.rotors[0].CurrentInputTerminal)
	}

	if em.rotors[1].CurrentInputTerminal != "B" {
		t.Errorf("Second Rotor Turnover failed. Current input terminal is wrong. Expected B, got %s",em.rotors[1].CurrentInputTerminal)
	}

	// Cause the second rotor to rotate the way around
	for i := 1; i <len(em.rotors[0].wiring);i++ {
		for j := 1;j < len(em.rotors[0].wiring);j++ {
			em.RotateRotors()
		}
	}

	if em.rotors[1].CurrentInputTerminal != "Z" {
		t.Errorf("Second Rotor Turnover failed. Current input terminal is wrong. Expected Z, got %s",em.rotors[1].CurrentInputTerminal)
	}

	// Check the third rotor hasnt rotated 
	if em.rotors[2].CurrentInputTerminal != "A" {
		t.Errorf("The third rotor moved when it shouldnt have. Current input terminal is wrong. Expected A, got %s",em.rotors[2].CurrentInputTerminal)
	}

	//Move the second rotor again and check the third rotor changes
	for i := 1; i <len(em.rotors[0].wiring);i++ {
		em.RotateRotors()
	}
	em.RotateRotors()

	// Check the third rotor has rotated 
	if em.rotors[2].CurrentInputTerminal != "B" {
		t.Errorf("The third rotor should have rotated but it didnt. Current input terminal is wrong. Expected B, got %s",em.rotors[2].CurrentInputTerminal)
	}

	
	// Check the third rotor has rotated again
	for i := 0; i <len(em.rotors[0].wiring);i++ {
		for j := 0;j < len(em.rotors[0].wiring);j++ {
			em.RotateRotors()
		}
	}
	if em.rotors[2].CurrentInputTerminal != "C" {
		t.Errorf("The third rotor should have rotated but it didnt. Current input terminal is wrong. Expected C, got %s",em.rotors[2].CurrentInputTerminal)
	}

	// Final state Check
	if em.rotors[0].CurrentInputTerminal != "A" && em.rotors[1].CurrentInputTerminal != "A" && em.rotors[1].CurrentInputTerminal != "C" {
		t.Errorf("Inconsistent machine state. Rotors show: %s,%s,%s Expected: A,A,C",em.rotors[0].CurrentInputTerminal,em.rotors[1].CurrentInputTerminal,em.rotors[2].CurrentInputTerminal)
	}

}
