package common

import (
	"bufio"
	"log"
	"os"
)

//ProcessStdin converts stdin to slice of strings
func ProcessStdin() []string {
	var output []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	// log.Printf("lines: %v\n%v", len(output), output)

	return output
}

// Debug Prints a debug output
func Debug(format string, a ...interface{}) {
	if false {
		log.Printf(format, a)
	}
}
