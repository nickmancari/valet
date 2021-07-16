package cmd

import (
	
	"fmt"
	"github.com/nickmancari/valet/database"
	"github.com/nickmancari/valet/runner"
)



func Create(f string) {

}

func RunCommand(f string) (interface{}, error) {
	if f == "" {
		return "", nil
	} else if f == "ls" {
		r, err := database.AllCommands()
		if err != nil {
			fmt.Println(err)
		}
		return r, nil
	} else {
		c := database.MapFromDatabase(f)

		r := runner.ExecuteCommand(c)
		return r, nil
	}

	return "", nil

}

func ListDatabase(f bool) (interface{}, error) {
	if f == false {
		return "", nil
	} else {
		r, err := database.AllCommands()
		if err != nil {
			fmt.Println(err)
		}
		return r, nil
	}

	return "", nil
}

func CreateDatabase(f bool) (interface{}, error) {
	if f == false {
		return "", nil
	} else {
		r := database.Generate()
		return r, nil
	}

	return "", nil
}
