package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

	log.Info("Running Example Machine 1")

	rotor1 := GenerateRotorI()
	rotor2 := GenerateRotorII()
	rotor3 := GenerateRotorIII()

	rotors := []*Rotor{&rotor3, &rotor2, &rotor1}

	em := EnigmaMachine{rotors, 0, GenerateReflectorB(), GenerateMilitaryInputRotor()}

	// Some testing for the rotor encoding logic.
	log.Println("Testing Rotor Encoding")
	fmt.Println("Right: " + em.rotors[0].EncodeRight("A"))
	fmt.Println("Left:  " + em.rotors[0].EncodeLeft("A"))

	log.Println("Testing Encryption")
	em.SetRotorPosition(0, 'A')
	em.SetRotorPosition(1, 'A')
	em.SetRotorPosition(2, 'A')

	fmt.Println("ENCRYPTED: " + em.Encrypt("AAAAA")) // This should produce  BDZGO

	em.SetRotorPosition(0, 'A')
	em.SetRotorPosition(1, 'A')
	em.SetRotorPosition(2, 'A')

	fmt.Println("DECRYPTED: " + em.Encrypt("BDZGO")) // This should produce AAAAA

}
