package main

import (
	"testing"
)

// Create test machine with 3 test rotors
func createTestMachine() EnigmaMachine {
	r1 := TestRotor()
	r2 := TestRotor()
	r3 := TestRotor()

	var rotors RotorSet

	rotors.left = &r1
	rotors.middle = &r2
	rotors.right = &r3

	straightThroughPlugBoard := Plugboard{map[string]string{"A": "A"}}
	em := CreateEnigmaMachine(rotors, "AAA", straightThroughPlugBoard, GenerateReflectorA(), GenerateMilitaryInputRotor())
	return em
}

// Create test machine with 3 test rotors
func createMilitaryMachine() EnigmaMachine {
	r1 := GenerateRotorI()
	r2 := GenerateRotorII()
	r3 := GenerateRotorIII()

	var rotors RotorSet

	rotors.left = &r1
	rotors.middle = &r2
	rotors.right = &r3

	straightThroughPlugBoard := Plugboard{map[string]string{"A": "A"}}

	em := CreateEnigmaMachine(rotors, "AAA", straightThroughPlugBoard, GenerateReflectorB(), GenerateMilitaryInputRotor())
	return em
}

func TestEncryption(t *testing.T) {
	em := createMilitaryMachine()

	em.SetRotorPosition("left", 'A')
	em.SetRotorPosition("middle", 'A')
	em.SetRotorPosition("right", 'A')

	enc := em.Encrypt("AAAAA") // BDZGO

	if enc != "BDZGO" {
		t.Errorf("Encryption Failed. Expected BDZGO, got %s ", enc)
	}
}

func TestDecryption(t *testing.T) {
	em := createMilitaryMachine()

	em.SetRotorPosition("left", 'A')
	em.SetRotorPosition("middle", 'A')
	em.SetRotorPosition("right", 'A')

	enc := em.Encrypt("BDZGO") // AAAAA

	if enc != "AAAAA" {
		t.Errorf("Encryption Failed. Expected AAAAA, got %s ", enc)
	}
}

// This tests that any key will encrypt and decrypt properly
func TestEncryptDecrypt(t *testing.T) {
	em := createMilitaryMachine()

	em.SetRotorPosition("left", 'A')
	em.SetRotorPosition("middle", 'A')
	em.SetRotorPosition("right", 'A')

	enc := em.Encrypt("AAAA") // AAAAA

	em.SetRotorPosition("left", 'A')
	em.SetRotorPosition("middle", 'A')
	em.SetRotorPosition("right", 'A')

	enc = em.Encrypt(enc)

	if enc != "AAAA" {
		t.Errorf("Encryption/Decryption Failed. Expected AAAAA, got %s ", enc)
	}

}

func TestPlugBoard(t *testing.T) {
	em := createMilitaryMachine()

	// Create a plugboard mapping
	em.plugBoard = Plugboard{map[string]string{"A": "E", "B": "D", "C": "F"}}

	enc := em.Encrypt("AAAAA")

	if enc != "CLFWR" {
		t.Errorf("Encryption with plugboard mapping Failed. Expected CLFWR, got %s ", enc)
	}

}

func TestSetRotorPosition(t *testing.T) {
	em := createTestMachine()
	em.SetRotorPosition("middle", 'K')

	if em.rotors.middle.CurrentIndicator != 10 {
		t.Errorf("Start position of rotor 1 not set correctly. Expected 11, got %d", em.rotors.middle.CurrentIndicator)
	}
}

func TestRotateRotors(t *testing.T) {

	em := createTestMachine()

	// rotate once
	em.RotateRotors()

	if string(toChar(em.rotors.right.CurrentIndicator)) != "B" {
		t.Errorf("Rotate failed. Current indicator is wrong. Expected B, got %s", string(toChar(em.rotors.right.CurrentIndicator)))
	}

	// Rotate round to the last position
	for i := 2; i < 26; i++ {
		em.RotateRotors()
	}

	if string(toChar(em.rotors.right.CurrentIndicator)) != "Z" {
		t.Errorf("Rotate failed. Current indicator is wrong. Expected Z, got %s", string(toChar(em.rotors.right.CurrentIndicator)))
	}

	// Rotate one past the last position. This tests the Modulo arithmetic and array bounds checks and turnover of the 2nd rotor
	em.RotateRotors()
	if string(toChar(em.rotors.right.CurrentIndicator)) != "A" {
		t.Errorf("Rotate failed. Current indicator is wrong. Expected A, got %s", string(toChar(em.rotors.right.CurrentIndicator)))
	}

	if string(toChar(em.rotors.middle.CurrentIndicator)) != "B" {
		t.Errorf("Second Rotor Turnover failed. Current indicator is wrong. Expected B, got %s", string(toChar(em.rotors.middle.CurrentIndicator)))
	}

	// Cause the second rotor to rotate the way around
	for i := 1; i < 26; i++ {
		for j := 1; j < 26; j++ {
			em.RotateRotors()
		}
	}

	if string(toChar(em.rotors.middle.CurrentIndicator)) != "Z" {
		t.Errorf("Second Rotor Turnover failed. Current indicator is wrong. Expected Z, got %s", string(toChar(em.rotors.middle.CurrentIndicator)))
	}

	// Check the third rotor hasnt rotated
	if string(toChar(em.rotors.left.CurrentIndicator)) != "A" {
		t.Errorf("The third rotor moved when it shouldnt have. Current indicator is wrong. Expected A, got %s", string(toChar(em.rotors.left.CurrentIndicator)))
	}

	//Move the second rotor again and check the third rotor changes
	for i := 1; i < 26; i++ {
		em.RotateRotors()
	}
	em.RotateRotors()

	// Check the third rotor has rotated
	if string(toChar(em.rotors.left.CurrentIndicator)) != "B" {
		t.Errorf("The third rotor should have rotated but it didnt. Current input terminal is wrong. Expected B, got %s", string(toChar(em.rotors.left.CurrentIndicator)))
	}

	// Check the third rotor has rotated again
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			em.RotateRotors()
		}
	}

	if string(toChar(em.rotors.left.CurrentIndicator)) != "C" {
		t.Errorf("The third rotor should have rotated but it didnt. Current indicator is wrong. Expected C, got %s", string(toChar(em.rotors.left.CurrentIndicator)))
	}

	// Final state Check
	if string(toChar(em.rotors.right.CurrentIndicator)) != "A" && string(toChar(em.rotors.middle.CurrentIndicator)) != "A" && string(toChar(em.rotors.left.CurrentIndicator)) != "C" {
		t.Errorf("Inconsistent machine state. Rotors show: %s,%s,%s Expected: A,A,C", string(toChar(em.rotors.right.CurrentIndicator)), string(toChar(em.rotors.middle.CurrentIndicator)), string(toChar(em.rotors.left.CurrentIndicator)))
	}

}
