package cli

import (
	"fmt"
)

type Command struct {
	Name string
	Help string
	Usage string
}

type Commands struct {
	list []*Command
}
