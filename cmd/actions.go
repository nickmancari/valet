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
	} else {
		c, err := database.MapFromDatabase(f)
		if err != nil {
			fmt.Println(err)
		}

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
