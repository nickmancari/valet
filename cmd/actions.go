package cmd

import (
	
	"fmt"
	"github.com/nickmancari/valet/database"
	"github.com/nickmancari/valet/runner"
)



func Create(f string) {

}

//Takes argument passed and hands it off to the database parser when the proper map is returned,
//the map is then handed to the runner to excute the alias.
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

//Takes argument passed from flag and returns a println of the JSON database.
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

//Takes argument passed from flag and generates a starting JSON file. 
func CreateDatabase(f bool) (interface{}, error) {
	if f == false {
		return "", nil
	} else {
		r := database.Generate()
		return r, nil
	}

	return "", nil
}
