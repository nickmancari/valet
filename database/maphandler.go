package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/nickmancari/valet/pkg/color"
)

//Parses the JSON file and compares the given alias to all the keys. It loops over a map containing a string as
//a key with another map as a value. If it gets a match it pulls the value and returns it as a map type.
func MapFromDatabase(f string) (map[string]string, error) {
	content, err := ioutil.ReadFile("/usr/local/bin/valet/cmd_db")
	if err != nil {
		fmt.Println(err)
	}
	x := make(map[string]map[string]string)

	er := json.Unmarshal(content, &x)
	if er != nil {
		fmt.Println(er)
	}

	for k, v := range x {
		if strings.Contains(k, f) {
			return v, nil
		} else {
			e := make(map[string]string)
			return e, errors.New("Alias Not Found")
		}
	}

	y := make(map[string]string)
	return y, nil

}

//Returns a structured println of the JSON file database.
func AllCommands() (interface{}, error) {
	content, err := ioutil.ReadFile("/var/lib/valet/cmd_db")
	if err != nil {
		fmt.Println(err)
	}
	x := make(map[string]map[string]string)

	er := json.Unmarshal(content, &x)
	if er != nil {
		fmt.Println(err)
	}

	fmt.Printf("\n%-10s %-10s %-35s\n", "Alias", "Command", "Parameters")

	for k, v := range x {
		for kk, vv := range v {
			fmt.Printf(color.Green+"%-10s"+color.Reset+" "+color.Cyan+"%-10s"+color.Reset+" "+color.Yellow+"%-35s"+color.Reset+"\n---------------------------------------------------------\n", k, kk, vv)
		}
	}
	return "", nil
}
