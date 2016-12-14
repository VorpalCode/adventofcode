package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func processStdin() []string {
	var output []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	// log.Printf("lines: %v\n%v", len(output), output)

	return output
}

func debug(format string, a ...interface{}) {
	if false {
		log.Printf(format, a)
	}
}

type triangle [3]int

func newTrgString(str string) triangle {
	var sideInts triangle
	for i, lng := range strings.Fields(str) {
		sideInts[i], _ = strconv.Atoi(lng)
	}

	return sideInts
}

func (t triangle) slice() []int {
	return t[0:len(t)]
}

func (t triangle) valid() bool {
	ints := t.slice()
	sort.Ints(ints)
	log.Printf("%v - %T", ints, ints)
	return (ints[0] + ints[1]) > ints[2]
}

func main() {
	m := make(map[bool]int)
	for _, str := range processStdin() {
		m[newTrgString(str).valid()]++
	}
	log.Printf("results = %v", m)
}
