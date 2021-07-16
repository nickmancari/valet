package main

import (
	"flag"

	"github.com/nickmancari/valet/cmd"
)


var commandFlag = flag.String("cmd", "", "Run Command with Assinged Alias")
var listFlag = flag.Bool("all", false, "Lists All Current Aliases in the Database")
var createFlag = flag.Bool("generate", false, "Create Database File")
func main() {

	flag.Parse()
	cmd.RunCommand(*commandFlag)
	cmd.ListDatabase(*listFlag)
	cmd.CreateDatabase(*createFlag)

}
