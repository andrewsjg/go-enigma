package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

	log.Info("Running Example Machine 1")

	rotor1 := TestRotor()
	rotor2 := TestRotor()
	rotor3 := TestRotor()

	rotors := []*Rotor{&rotor1, &rotor2, &rotor3}

	em := EnigmaMachine{rotors, 1, 1, GenerateMilitaryEntryWheel()}

	em.Encrypt("AAA")
	//em.RotateRotors()

	//em.SetRotorPosition(1, 'K')
}
