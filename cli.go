package cli

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type Cli struct {
	Name     string
	HelpName string
	Version  string
	Usage    string
	Commands []Command
	Options  []Option
}

func New(name string) *Cli {
	if name == "" {
		panic("can't construct an app without a name")
	}
	c := &Cli{Name: name}
	c.setUp()
	return c
}

func (c *Cli) setUp() {
	help := Option{
		Name:  "help",
		Short: "h",
		Usage: "Display this help message",
	}

	version := Option{
		Name:  "version",
		Short: "v",
		Usage: "Display this application version",
	}

	c.Options = append(c.Options, help)
	c.Options = append(c.Options, version)
}

func (c *Cli) Run() {
	flag.Parse()

	if len(flag.Args()) < 1 {
		c.printHelp()
		return
	}

	subArgs := flag.Args()[0:]

	if len(subArgs) == 0 {
		c.printHelp()
		return
	}

	switch flag.Args()[0] {
	case "version", "-v":
		fmt.Println(c.Version)
	case "help", "-h":
		c.printHelp()
		return
	default:
		log.Fatalf("\nInvalid Command %s", os.Args[1])
	}

	return
}
