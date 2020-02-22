package main

import "testing"


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


}
