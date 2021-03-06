package main

import (
	"flag"

	"github.com/nickmancari/valet/cmd"
)


var commandFlag = flag.String("cmd", "", "Run Command with Assigned Alias")
var listFlag = flag.Bool("all", false, "Lists All Current Aliases in the Database")
var createFlag = flag.Bool("init", false, "Create Database File")

//Handle flag flow to respected functions.
//Takes argument passed in from the commandline and hands it off to the designated function.
func main() {

	flag.Parse()
	cmd.RunCommand(*commandFlag)
	cmd.ListDatabase(*listFlag)
	cmd.CreateDatabase(*createFlag)

}
