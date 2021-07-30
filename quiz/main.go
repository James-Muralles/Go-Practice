package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFilename := flag.String("csv,", "problems.csv", "a csv in the format of 'question, answer'")
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

	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
		}

	}
	fmt.Printf("You scored %d out of %d.\n That is %d percent ", correct, len(problems), correct/len(problems)*5)

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

