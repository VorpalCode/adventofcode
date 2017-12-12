package main

import (
	"bufio"
	"log"
	"os"
)

func processStdin() []string {
	var output []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	log.Printf("lines: %v\n%v", len(output), output)

	return output
}

func debug(format string, a ...interface{}) {
	if false {
		log.Printf(format, a)
	}
}

type keypad struct {
	rows [3][3]int
	xPos int
	yPos int
}

func newKeypad() keypad {
	var kp keypad
	for i := range kp.rows {
		mod := i * 3
		kp.rows[i] = [3]int{1 + mod, 2 + mod, 3 + mod}
	}

	kp.xPos = 1
	kp.yPos = 1

	return kp
}

func (kp keypad) move(shift string) keypad {

	switch shift {
	case "U":
		if kp.yPos != 0 {
			kp.yPos--
		}
	case "D":
		if kp.yPos != 2 {
			kp.yPos++
		}
	case "L":
		if kp.xPos != 0 {
			kp.xPos--
		}
	case "R":
		if kp.xPos != 2 {
			kp.xPos++
		}
	}

	return kp
}

func main() {
	output := processStdin()
	kp := newKeypad()
	var buttons []int

	for _, line := range output {
		var move rune
		for _, move = range line {
			kp = kp.move(string(move))
			// log.Printf("%v", string(move))
		}
		log.Printf("kp = %v, current = %v", kp, kp.rows[kp.yPos][kp.xPos])
		buttons = append(buttons, kp.rows[kp.yPos][kp.xPos])
	}

	log.Printf("buttons: %v", buttons[:])
}
