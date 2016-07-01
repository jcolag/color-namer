package main

import (
	"fmt"
	"io/ioutil"
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

