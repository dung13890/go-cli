package cli

import (
	"flag"
	"fmt"
)

type Cli struct {
	commands Commands
}

func (c *Cli) Commands *Commands {
	return &c.commands
}

func (c *Cli) Run() (err error) {

}
