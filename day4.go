package main

import (
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"

	c "github.com/VorpalCode/aoc2016/common"
)

type room struct {
	code      string
	sector    int
	checksum  string
	charCount map[string]int
}

func newRoomFromStr(line string) room {
	var r room

	re := regexp.MustCompile(`^([\w-]+)-(\d+)\[(\w+)\]$`)
	matches := re.FindStringSubmatch(line)
	r.code = matches[1]
	r.sector, _ = strconv.Atoi(matches[2])
	r.checksum = matches[3]
	r.charCount = make(map[string]int)

	for _, c := range r.code {
		r.charCount[string(c)]++
	}

	return r
}

func (r room) top5() string {
	topStrings := map[int][]string{}
	counts := []int{}
	ordered := []string{}

	for k, v := range r.charCount {
		topStrings[v] = append(topStrings[v], k)
	}
	for k := range topStrings {
		counts = append(counts, k)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(counts)))
	c.Debug("%v / %v", topStrings, counts)

	for _, k := range counts {
		letters := topStrings[k]
		sort.Strings(letters)
		for _, s := range letters {
			if s == "-" {
				continue
			}
			ordered = append(ordered, s)
		}
	}
	rslt := strings.Join(ordered[0:5], ``)
	c.Debug("%v - %T", rslt, rslt)
	return rslt
}

func main() {
	m := make(map[bool]int, 2)
	for _, str := range c.ProcessStdin() {
		r := newRoomFromStr(str)
		m[r.top5() == r.checksum] += r.sector

		c.Debug("-----------")
	}
	log.Printf("%v", m)
}
