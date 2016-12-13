package main

import (
	"bufio"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func processStdin() string {
	var output string
	r := bufio.NewReader(os.Stdin)
	buf := make([]byte, 0, 4*1024)

	for {
		n, err := r.Read(buf[:cap(buf)])
		buf = buf[:n]
		if n == 0 {
			if err == nil {
				continue
			}
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		if err != nil && err != io.EOF {
			log.Fatal(err)
		}

		output = string(buf)
	}
	debug("Buf = %v", output)

	return output
}

func debug(format string, a ...interface{}) {
	if false {
		log.Printf(format, a)
	}
}

func sliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

type direction string

func (d direction) turn(way string) direction {
	compass := [4]direction{"N", "E", "S", "W"}
	current := sliceIndex(len(compass), func(i int) bool { return compass[i] == d })
	var newInt int

	if way == "L" {
		newInt = (current - 1) % len(compass)
		if newInt < 0 {
			newInt = len(compass) + newInt
		}
	} else if way == "R" {
		newInt = (current + 1) % len(compass)
	}

	d = compass[newInt]
	debug("newWay = %v", d)
	return d
}

type movement struct {
	turn     string
	distance int
}

type position struct {
	direction
	xPos int
	yPos int
}

func (p position) step(m movement) position {
	p.direction = p.turn(m.turn)

	switch p.direction {
	case direction("N"):
		p.yPos += m.distance
	case direction("S"):
		p.yPos -= m.distance
	case direction("E"):
		p.xPos += m.distance
	case direction("W"):
		p.xPos -= m.distance
	}

	return p
}

func (p position) distance() float64 {
	return math.Abs(float64(p.xPos)) + math.Abs(float64(p.yPos))
}

func main() {
	moveStrs := strings.Split(strings.TrimSpace(processStdin()), ", ")
	curPos := position{direction("N"), 0, 0}

	for _, moveStr := range moveStrs {
		dist, _ := strconv.Atoi(string([]rune(moveStr)[1:]))
		move := movement{turn: string([]rune(moveStr)[0]), distance: dist}
		curPos = curPos.step(move)

		debug("move = %v", move)
		debug("cP = %v", curPos)
		debug("----------")
	}

	log.Printf("x = %v / y = %v", curPos.xPos, curPos.yPos)
	log.Printf("total distance = %v", curPos.distance())
}
