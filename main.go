package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

	log.Info("Running Example Machine 1")

	// Rotor 1 is the left most rotor and rotor 3 is the right most rotor
	rotor1 := GenerateRotorIII()
	rotor2 := GenerateRotorII()
	rotor3 := GenerateRotorI()

	var rotors RotorSet

	rotors.left = &rotor3
	rotors.middle = &rotor2
	rotors.right = &rotor1

	//Empty straight through plugboard
	straightThroughPlugBoard := Plugboard{map[string]string{}}

	em, err := CreateEnigmaMachine(rotors, "AAA", straightThroughPlugBoard, GenerateReflectorB(), GenerateMilitaryInputRotor())

	if err != nil {
		log.Fatal("There was an issue creating the machine: " + err.Error())
	}

	// Some testing for the rotor encoding logic.
	log.Println("Testing Rotor Encoding")
	fmt.Println("Right: " + em.rotors.right.EncodeRight("A"))
	fmt.Println("Left:  " + em.rotors.right.EncodeLeft("A"))

	log.Println("Testing Encryption")

	fmt.Println("ENCRYPTED: " + em.Encrypt("AAAAA")) // This should produce  BDZGO

	em.SetRotorPosition("left", 'A')
	em.SetRotorPosition("middle", 'A')
	em.SetRotorPosition("right", 'A')

	fmt.Println("DECRYPTED: " + em.Encrypt("BDZGO")) // This should produce AAAAA

	lorum := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure do"

	fmt.Println("LORUM: " + em.PrettyCrypt(lorum))

	// Encrypt from file test
	err = em.EncryptFromFile("shakespeare.txt", "ciphertext.txt")
	if err != nil {
		fmt.Println("There was an error: " + err.Error())
	}

}
