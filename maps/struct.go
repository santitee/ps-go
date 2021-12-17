package main

import (
	"fmt"
	"time"
)

type courseMeta struct {
	author       string  `json:"author,omitempty"`
	level        string  `json:"level,omitempty"`
	rating       float64 `json:"rating,omitempty"`
	courseDetail `json:"course_detail,omitempty"`
}

type courseDetail struct {
	courseID   string
	courseName string
	duration   time.Duration
}

func main() {

	//var gettingStartedWithK8s courseMeta
	//gettingStartedWithK8s := new(courseMeta)
	//fmt.Println(gettingStartedWithK8s)

	gettingStartedWithK8s := courseMeta{
		author: "Santi Teeragul",
		level:  "Basic",
		rating: 5,
	}

	gettingStartedWithK8sDetail := courseDetail{
		courseID:   "00001",
		courseName: "Getting Started with Kubernetes",
		duration:   30,
	}

	fmt.Println(gettingStartedWithK8s, gettingStartedWithK8sDetail)
	fmt.Println("Author of Getting Started with Kubernetes is", gettingStartedWithK8s.author)
	gettingStartedWithK8s.rating = 2.5
	fmt.Println("Course rating is currently", gettingStartedWithK8s.rating)
	//fmt.Println(gettingStartedWithK8sDetail)
}
