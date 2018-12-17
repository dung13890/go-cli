package main

import (
	"fmt"
    "flag"
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

    wordPtr := flag.String("word", "foo", "a string")
    flag.Parse()

    fmt.Println("word:", *wordPtr)
}
