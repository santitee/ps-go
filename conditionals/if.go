package main

import "fmt"

func main() {
	dddLengthMins := 275 // Docker Deep Dive Course
	cawLengthMins := 30  // Containers on AWS Wavelength course

	if dddLengthMins > cawLengthMins {
		fmt.Println("Docker Deep Dive is longer than Containers on AWS Wavelength")
	} else if dddLengthMins == cawLengthMins {
		fmt.Println("Both courses are the same length")
	} else {
		fmt.Println("Containers on AWS Wavelength must be longer than Docker Deep Dive")
	}
}
