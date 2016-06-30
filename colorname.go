package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	)

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

