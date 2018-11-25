package main

import (
	"fmt"
)

var version = "develop"

const logo = `
    ______      ____
   / ____/___  / __ )____  __  __
  / / __/ __ \/ __  / __ \/ / / /
 / /_/ / /_/ / /_/ / /_/ / /_/ /
 \____/\____/_____/\____/\__, /
%23s /____/
`

func main() {
	fmt.Println(fmt.Sprintf(logo, version))
}
