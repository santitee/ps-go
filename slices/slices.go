package main

import "fmt"

func main() {

	courses := []string{"Docker & Kubernetes: The Big Picture",
		"Getting Started with Docker",
		"Getting Started with Kubernetes"}

	//courses[0] = "Docker & Kubernetes: The Big Picture"
	//courses[1] = "Getting Started with Docker"
	//courses[2] = "Getting Started with Kubernetes"

	fmt.Printf("Length of slice is %d and capacity %d\n",
		len(courses), cap(courses))

	for idx, i := range courses {
		fmt.Println(idx, i)
	}
}
