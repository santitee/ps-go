package main

import (
	"fmt"
	"strings"
)

func main() {
	author := "santi teeragul"
	course := "getting started with kubernetes"

	fmt.Println(converter(author, course))
}

func converter(author, course string) (a, c string) {
	author = strings.ToUpper(author)
	course = strings.Title(course)

	return author, course
}
