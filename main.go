package main

import (
	"fmt"
	"flag"
	"bufio"
	"os"
	"log"
)

var version = "v1.0.0"

const logo = `
  __  __         ___
 |  \/  |_ _    |   \
 | |\/| | '_|   | |) |
 |_|  |_|_| (_) |___/
 %s
`

func main() {
	getName()
	switchFlags()
	//showHelp()
}

func showHelp() {
	fmt.Print(`
USAGE: CLI Template [OPTIONS]

OPTIONS:
 -h, --help         	Display usage
 -v, --version      	Display version

Commands:
 app			App arguments
	`)
}

func showVersion() {
	fmt.Println(fmt.Sprintf(logo, version))
}

func getName() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is your name? ")
	text, _ := reader.ReadString('\n')
	fmt.Printf("Hello: %s, May I help you?\n", text)
}

func switchFlags() {

	wordPtr := flag.String("word", "foo", "a string")
	flag.Parse()

	if len(flag.Args()) < 1 {
		return
	}

	subArgs := flag.Args()[1:]
	switch flag.Args()[0] {
		case "version":
			showVersion()
		case "help":
			showHelp()
		default:
			log.Fatalf("\nInvalid Command %s", os.Args[1])
	}
	fmt.Println(subArgs)
	fmt.Println("word:", *wordPtr)
}
