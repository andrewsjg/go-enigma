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
	straightThroughPlugBoard := Plugboard{map[string]string{"A": "A"}}

	em := CreateEnigmaMachine(rotors, "AAA", straightThroughPlugBoard, GenerateReflectorB(), GenerateMilitaryInputRotor())

	//em := EnigmaMachine{rotors, []string{"A", "A", "A"}, straightThroughPlugBoard, GenerateReflectorB(), GenerateMilitaryInputRotor()}

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

}
