package main

import (
	"log"
	"math"
	"strconv"
	"strings"

	c "github.com/VorpalCode/aoc2016/common"
)

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
	var newInt int
	compass := [4]direction{"N", "E", "S", "W"}
	current := sliceIndex(len(compass), func(i int) bool { return compass[i] == d })
	turns := map[string]int{"L": -1, "R": 1}

	newInt = int(math.Abs(float64(len(compass)+current+turns[way]))) % len(compass)

	d = compass[newInt]
	c.Debug("newWay = %v", d)
	return d
}

type movement struct {
	turn     string
	distance int
}

func newMovementFromStr(str string) movement {
	dist, _ := strconv.Atoi(string([]rune(str)[1:]))
	return movement{turn: string([]rune(str)[0]), distance: dist}
}

type coords struct {
	xPos int
	yPos int
}

type position struct {
	direction
	coords
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
	moveStrs := strings.Split(c.ProcessStdin()[0], ", ")
	curPos := position{direction("N"), coords{0, 0}}

	for _, moveStr := range moveStrs {
		move := newMovementFromStr(moveStr)
		curPos = curPos.step(move)

		c.Debug("move = %v", move)
		c.Debug("cP = %v", curPos)
		c.Debug("----------")
	}

	log.Printf("x = %v / y = %v", curPos.xPos, curPos.yPos)
	log.Printf("total distance = %v", curPos.distance())
}
