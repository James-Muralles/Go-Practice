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
	csvFilename := flag.String("csv,", "problems.csv", "a csv in the format of 'question, answer'")
	timeLimit := flag.Int("limit", 2, "the time limit for the quiz in seconds" ) 
	flag.Parse()

	file, err := os.Open(*csvFilename)

	if err != nil {
		exit(fmt.Sprintf("Failed to open CSV file %s\n", *csvFilename))

	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV")
	}
	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	// <-timer.C
 
	correct := 0

	problemloop:
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		answerCh := make(chan string)
		go func(){
			var answer string
		fmt.Scanf("%s\n",  &answer)
		answerCh <- answer
		}()

		select{
		case <-timer.C:
			fmt.Println()
			break problemloop
		case answer := <- answerCh:	
		if answer == p.a {
			correct++

			}
			
		}

	}
	fmt.Printf("You scored %d out of %d.\n That is %.2f percent! ", correct, len(problems), float32(correct) / float32(len(problems))*100)

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

	func exit(msg string){
			fmt.Println(msg)
			os.Exit(1)

	}


