package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	)

type color struct {
	name	string
	rgb	string
	red	byte
	green	byte
	blue	byte
	hue	byte
	sat	byte
	val	byte
}

func (c *color) populateFromRgb(rgb string) {
	rr, _ := strconv.ParseUint(rgb[0:2], 16, 8)
	gg, _ := strconv.ParseUint(rgb[2:4], 16, 8)
	bb, _ := strconv.ParseUint(rgb[4:6], 16, 8)
	c.rgb = "#" + rgb
	c.red = byte(rr)
	c.green = byte(gg)
	c.blue = byte(bb)
}

func main() {
	dat, err := ioutil.ReadFile("allcolor.csv")
	if err != nil {
		panic(err)
	}

	sdat := string(dat)
	clines := strings.Split(sdat, "\n")
	allcolors := make([]color, 0)
	for i := 0; i < len(clines); i++ {
		cparts := strings.Split(clines[i], ",")
		if len(cparts) < 8 {
			break
		}
		r, _ := strconv.ParseUint(cparts[4], 10, 8)
		g, _ := strconv.ParseUint(cparts[5], 10, 8)
		b, _ := strconv.ParseUint(cparts[6], 10, 8)
		h, _ := strconv.ParseUint(cparts[1], 10, 8)
		s, _ := strconv.ParseUint(cparts[2], 10, 8)
		v, _ := strconv.ParseUint(cparts[3], 10, 8)
		c := color {cparts[0], cparts[7],
			byte(r), byte(g), byte(b),
			byte(h), byte(s), byte(v)}
		allcolors = append(allcolors, c)
		fmt.Println(c)
	}
}

