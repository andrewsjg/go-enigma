package main

import (
	"testing"
)

// TestRotor creates a rotor with straight through wiring
func testRotor() Rotor {

	wiring := [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	r := Rotor{wiring, 0, "Z", 1, false}

	return r
}

// Create test machine with 3 test rotors
func createTestMachine() EnigmaMachine {
	r1 := testRotor()
	r2 := testRotor()
	r3 := testRotor()

	return EnigmaMachine{[]*Rotor{&r1, &r2, &r3}, 1, 1, GenerateMilitaryEntryWheel()}
}

func TestSetRotorPosition(t *testing.T) {
	em := createTestMachine()
	em.SetRotorPosition(1, 'K')

	if em.rotors[1].CurrentIndicator != 11 {
		t.Errorf("Start position of rotor 1 not set correctly. Expected 11, got %d", em.rotors[1].CurrentIndicator)
	}
}

func TestRotateRotors(t *testing.T) {

	em := createTestMachine()

	// rotate once
	em.RotateRotors()

	if string(toChar(em.rotors[0].CurrentIndicator)) != "B" {
		t.Errorf("Rotate failed. Current indicator is wrong. Expected B, got %s", string(toChar(em.rotors[0].CurrentIndicator)))
	}

	// Rotate round to the last position
	for i := 2; i < 26; i++ {
		em.RotateRotors()
	}

	if string(toChar(em.rotors[0].CurrentIndicator)) != "Z" {
		t.Errorf("Rotate failed. Current indicator is wrong. Expected Z, got %s", string(toChar(em.rotors[0].CurrentIndicator)))
	}

	// Rotate one past the last position. This tests the Modulo arithmetic and array bounds checks and turnover of the 2nd rotor
	em.RotateRotors()
	if string(toChar(em.rotors[0].CurrentIndicator)) != "A" {
		t.Errorf("Rotate failed. Current indicator is wrong. Expected A, got %s", string(toChar(em.rotors[0].CurrentIndicator)))
	}

	if string(toChar(em.rotors[1].CurrentIndicator)) != "B" {
		t.Errorf("Second Rotor Turnover failed. Current indicator is wrong. Expected B, got %s", string(toChar(em.rotors[0].CurrentIndicator)))
	}

	// Cause the second rotor to rotate the way around
	for i := 1; i < 26; i++ {
		for j := 1; j < 26; j++ {
			em.RotateRotors()
		}
	}

	if string(toChar(em.rotors[1].CurrentIndicator)) != "Z" {
		t.Errorf("Second Rotor Turnover failed. Current indicator is wrong. Expected Z, got %s", string(toChar(em.rotors[1].CurrentIndicator)))
	}

	// Check the third rotor hasnt rotated
	if string(toChar(em.rotors[2].CurrentIndicator)) != "A" {
		t.Errorf("The third rotor moved when it shouldnt have. Current indicator is wrong. Expected A, got %s", string(toChar(em.rotors[2].CurrentIndicator)))
	}

	//Move the second rotor again and check the third rotor changes
	for i := 1; i < 26; i++ {
		em.RotateRotors()
	}
	em.RotateRotors()

	// Check the third rotor has rotated
	if string(toChar(em.rotors[2].CurrentIndicator)) != "B" {
		t.Errorf("The third rotor should have rotated but it didnt. Current input terminal is wrong. Expected B, got %s", string(toChar(em.rotors[2].CurrentIndicator)))
	}

	// Check the third rotor has rotated again
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			em.RotateRotors()
		}
	}

	if string(toChar(em.rotors[2].CurrentIndicator)) != "C" {
		t.Errorf("The third rotor should have rotated but it didnt. Current indicator is wrong. Expected C, got %s", string(toChar(em.rotors[2].CurrentIndicator)))
	}

	// Final state Check
	if string(toChar(em.rotors[0].CurrentIndicator)) != "A" && string(toChar(em.rotors[1].CurrentIndicator)) != "A" && string(toChar(em.rotors[2].CurrentIndicator)) != "C" {
		t.Errorf("Inconsistent machine state. Rotors show: %s,%s,%s Expected: A,A,C", string(toChar(em.rotors[0].CurrentIndicator)), string(toChar(em.rotors[1].CurrentIndicator)), string(toChar(em.rotors[2].CurrentIndicator)))
	}

}
