package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

var inputFile = flag.String("inputFile", "", "Input file")

func main() {
	flag.Parse()
	if *inputFile == "" {
		myName, _ := os.Executable()
		fmt.Printf("Usage: %s -inputFile <source>\n", myName)
		os.Exit(1)
	}
	input, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		fmt.Printf("Couldn't open %s for read: %v\n", *inputFile, err)
		os.Exit(1)
	}
	linkID := regexp.MustCompile(`(?m)^\[(\d+)\]: (.*)$`)
	var re *regexp.Regexp

	document := string(input)

	var match [][]string
	match = linkID.FindAllStringSubmatch(document, -1)
	for i := range match {
		re, err = regexp.Compile(fmt.Sprintf("\\[(.*)\\]\\[%s\\]", match[i][1]))
		if err != nil {
			fmt.Printf("Couldn't compile the regexp for %s: %v\n", match[i][1], err)
			os.Exit(1)
		}
		document = re.ReplaceAllString(document, "[$1|"+match[i][2]+"]")
	}

	fmt.Println(document)

}
