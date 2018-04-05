/*
Part 2
Adapt your program from part 1 to add a timer. The default time limit should be 30 seconds, but should also be customizable via a flag.

Your quiz should stop as soon as the time limit has exceeded. That is, you shouldn’t wait for the user to answer one final questions but should ideally stop the quiz entirely even if you are currently waiting on an answer from the end user.

Users should be asked to press enter (or some other key) before the timer starts, and then the questions should be printed out to the screen one at a time until the user provides an answer. Regardless of whether the answer is correct or wrong the next question should be asked.

At the end of the quiz the program should still output the total number of questions correct and how many questions there were in total. Questions given invalid answers or unanswered are considered incorrect.
*/

package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("file", "problems.csv", "a csv file with Question and Answer")
	timelimit := flag.Int("timelimit", 10, "timeout for test")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	errorHandle(err, fmt.Sprintf("Could not open csv file: %s\n", *csvFilename))

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	errorHandle(err, "Failed ReadAll")

	problems := parseLines(lines)
	correct := 0
	var timer *time.Timer

	var ready string
	fmt.Printf("Ready to start? (y/n) ")
	fmt.Scanf("%s\n", &ready)
	if ready == "y" {
		// start timer
		timer = time.NewTimer(time.Duration(*timelimit) * time.Second)
		go func() {
			<-timer.C
			fmt.Println("Times UP!")
			os.Exit(1)

		}()
	} else {
		fmt.Println("If that's what you want!")
		os.Exit(1)
	}
	fmt.Println("You gave: "+ready+" and time value is ", *timelimit)

	for i, p := range problems {
		fmt.Printf("Problem: #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
		}
	}
	//stop := timer.Stop()

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

func errorHandle(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		os.Exit(1)
	}
}
