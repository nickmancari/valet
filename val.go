package main

import (
	"flag"

	"github.com/nickmancari/valet/cmd"
)


var commandFlag = flag.String("cmd", "", "Run Command with Assinged Alias")

func main() {

	flag.Parse()
	cmd.RunCommand(*commandFlag)
}
