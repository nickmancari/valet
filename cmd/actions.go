package cmd

import (

	"github.com/nickmancari/valet/database"
	"github.com/nickmancari/valet/runner"
)



func Create(f string) {

}

func RunCommand(f string) (interface{}, error) {
	if f == "" {
		return "", nil
	} else {
		c := database.MapFromDatabase(f)

		r := runner.ExecuteCommand(c)
		return r, nil
	}

	return "", nil

}
