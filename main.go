package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func findArduino() string {
	content, _ := ioutil.ReadDir("/dev")
	likelyDevices := [3]string{
		"ttyACM0",
		"ttyUSB.serial",
		"ttyUSB"}

	for _, f := range content {
		for i := 0; i < 3; i++ {
			if strings.Contains(f.Name(), likelyDevices[i]) {
				return "/dev/" + f.Name()
			}
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
