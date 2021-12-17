package main

import (
	"fmt"
	"time"
)

func main() {
	for timer := 10; timer >= 0; timer-- {
		if timer%2 == 0 {
			//fmt.Println("This is time up Booms!!!")
			//break
			continue
		}
		fmt.Println(timer)
		time.Sleep(1 * time.Second)
	}
}
