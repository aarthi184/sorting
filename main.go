package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/aarthi184/sorting/bubble"
	"github.com/aarthi184/sorting/insertion"
)

const (
	optInsertion = "INSERTION"
	optBubble    = "BUBBLE"
)

func main() {
	var (
		sort string
		in   []int
		err  error
	)

	if len(os.Args) < 2 {
		sort = "INSERTION"
		log.Println("No specification. Doing Insertion Sort")
	} else {
		sort = strings.ToUpper(os.Args[1])
	}

	in = []int{5, 8, 1, 56, 87, 13, 4, 12, 4, 6, 5}
	in = []int{5, 6,7,8,19,20,21,23,27,30, 40}
	log.Printf("Input: %v\n", in)

	start := time.Now()
	switch sort {
	case optInsertion:
		err = insertion.Sort(in)
	case optBubble:
		err = bubble.Sort(in)
	default:
		log.Println("Unknown sort", sort)
	}
	duration := time.Since(start)
	if err != nil {
		log.Printf("Error: %s\n", err)
		return
	}
	log.Printf("%sSort:%v\n", sort, in)
	log.Println("Time taken:", duration)
}
