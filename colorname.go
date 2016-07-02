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
	for i := 0; i < len(clines); i++ {
		cparts := strings.Split(clines[i], ",")
		if len(cparts) < 8 {
			break
		}
		fmt.Println(clines[i])
	}
}

