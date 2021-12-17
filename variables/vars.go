package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	name := "Santi Teeragul"
	course := "Getting Started with Kubernetes"
	module := "4"
	clip := 2
	//courseComplete := false

	fmt.Println("Name and course are set to", name, "and", course, ".")
	fmt.Println("Module and clip are set to", module, "and", clip, ".")
	fmt.Println("Name is of type", reflect.TypeOf(name))
	fmt.Println("Module is of type", reflect.TypeOf(module))
	//total := module + clip
	//fmt.Println("Module plus clip equal", total)

	iModule, err := strconv.Atoi(module)
	if err == nil {
		total := iModule + clip
		fmt.Println("Module plus clip equal", total)
	}

	fmt.Println("Memory address if *course* variable is", &course)

	var ptr *string = &course
	fmt.Println("Pointing course variable at address,", ptr, "which holds this value,", *ptr)

}
