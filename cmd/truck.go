package main

import (
	"os/exec"

	cli "github.com/dung13890/go-cli"
)

func main() {
	ex := exec.Command("ls", "-al")
	stdout, err := ex.Output()

	if err != nil {
		println(err.Error())
		return
	}

	print(string(stdout))
	cmd := cli.New("truck")
	cmd.Version = "1.0.0"
	cmd.Usage = "[command] [-option]"
	cmd.Run()
}
