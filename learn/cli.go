package learn

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/bitfield/script"
)

func Run_Cli() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cwd)

	// data, err := script.File("README.md").String()

	re := regexp.MustCompile("go | Go")
	lines, err := script.File("README.md").MatchRegexp(re).CountLines()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("go is mentioned on %d lines\n", lines)

	data, err := script.File("README.md").MatchRegexp(re).First(3).Stdout()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)

}
