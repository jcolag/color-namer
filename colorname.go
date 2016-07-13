package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
	)

type position struct {
    x float64
    y float64
    z float64
}

func (p *position) distance(q position) float64 {
	x := p.x - q.x
	y := p.y - q.y
	z := p.z - q.z
	return math.Sqrt(x*x + y*y + z*z)
}

type color struct {
	name	string
	rgb	string
	red	byte
	green	byte
	blue	byte
	hue	float64
	sat	float64
	val	float64
	distance float64
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

func (c *color) populateHsvFromRgb() {
    var min, max, delta float64
    
    min = math.Min(float64(c.red), math.Min(float64(c.green), float64(c.blue)))
    max = math.Max(float64(c.red), math.Max(float64(c.green), float64(c.blue)))
    c.val = max / 256
    delta = max - min
    if max == 0 {
        // All zero components
        c.sat = 0
        c.hue = 0
        return
    } else {
        c.sat = delta / max
    }
    
    if delta == 0 {
        // Hue is irrelevant
        c.hue = 0
    } else if max == float64(c.red) {
        // Hue is somewhere between yellow and magenta
        c.hue = (float64(c.green) - float64(c.blue)) / delta
    } else if max == float64(c.green) {
        // ...between cyan and yellow
        c.hue = 2 + (float64(c.blue) - float64(c.red)) / delta
    } else {
        // ...between magenta and cyan
        c.hue = 4 + (float64(c.red) - float64(c.green)) / delta
    }
    
    c.hue *= 60
    if c.hue < 0 {
        c.hue += 360
    }
    
    // Convert degrees to radians
    c.hue = c.hue * math.Pi / 180

func (c *color) populateDistance(d color) {
	cp := position{c.sat * math.Cos(c.hue), c.sat * math.Sin(c.hue), c.val}
	dp := position{d.sat * math.Cos(d.hue), d.sat * math.Sin(d.hue), d.val}
	delta := cp.distance(dp)
	c.distance = delta
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Cannot run without a color.")
		fmt.Printf("\t%s [RRGGBB]\n\n", os.Args[0])
		os.Exit(-1)
	}

	incolor := color{}
	incolor.populateFromRgb(args[0])
	incolor.name = "User Input"
	incolor.populateHsvFromRgb()

	dat, err := ioutil.ReadFile("allcolor.csv")
	if err != nil {
		panic(err)
	}

	sdat := string(dat)
	clines := strings.Split(sdat, "\n")
	allcolors := make([]color, 0)
	minDist := float64(2)
	minDistIdx := -1
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
			float64(h) * math.Pi / 180, // Convert degrees to radians
			float64(s) / 100,           // Normalize to unit circle
			float64(v) / 100,           // Normalize to unit height
			0}
		c.populateDistance(incolor)
		allcolors = append(allcolors, c)
		if c.distance < minDist {
			minDist = c.distance
			minDistIdx = i
		}
	}
}

