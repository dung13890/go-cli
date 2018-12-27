package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	cli "github.com/dung13890/go-cli"
)

func main() {
	cmd := cli.New("truck")
	cmd.Version = "1.0.0"
	cmd.Usage = "[command] [-option]"
	cmd.Run()
	//getName()
	// switchFlags()
	// showHelp()
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
	fmt.Println("version")
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
	case "help", "usage":
		showHelp()
	default:
		log.Fatalf("\nInvalid Command %s", os.Args[1])
	}
	fmt.Println(subArgs)
	fmt.Println("word:", *wordPtr)
}
