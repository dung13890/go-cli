package main

import (
	cli "github.com/dung13890/go-cli"
)

func main() {
	cmd := cli.New("truck")
	cmd.Version = "1.0.0"
	cmd.Usage = "[command] [-option]"
	cmd.Run()
}
