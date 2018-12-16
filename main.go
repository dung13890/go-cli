package main

import (
	"fmt"
)

var version = "develop"

const logo = `
  __  __         ___
 |  \/  |_ _    |   \
 | |\/| | '_|   | |) |
 |_|  |_|_| (_) |___/
 %23s
`

func main() {
	fmt.Println(fmt.Sprintf(logo, version))
}
