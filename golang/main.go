package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func ReadInput(filename string) []string {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	cases := strings.Split(string(contents), "\n")
	return cases[1 : len(cases)-1]
}

func WriteOutput(filename string, answers []string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for i, v := range answers {
		fmt.Println(v)
		f.WriteString(fmt.Sprintf("Case #%v: %v\n", i+1, v))
	}
}

func main() {
	inputFilename := flag.String("i", "input.txt", "Input file containing all the cases")
	outputFilename := flag.String("o", "output.txt", "Output file containing the results")

	flag.Parse()

	results := Solve(ReadInput(*inputFilename))

	WriteOutput(*outputFilename, results)
}
