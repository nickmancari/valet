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
//		if err != nil {
//			fmt.Println(err)
//		}

		r := runner.ExecuteCommand(c)
		return r, nil
	}

	return "", nil

}
