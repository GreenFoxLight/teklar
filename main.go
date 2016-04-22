package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func findArduino() string {
	content, _ := ioutil.ReadDir("/dev")

	for _, f := range content {
		if strings.Contains(f.Name(), "tty.usbserial") ||
			strings.Contains(f.Name(), "ttyUSB") {
			return "/dev/" + f.Name()
		}
	}
	return ""
}

func main() {
	dev := findArduino()

	if dev == "" {
		fmt.Println("No Arduino detected!")
	} else {
		fmt.Print("Arduino on: ")
		fmt.Println(dev)
	}
}
