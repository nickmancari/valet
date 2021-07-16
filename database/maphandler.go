package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/nickmancari/valet/pkg/color"
)

func MapFromDatabase(f string) map[string]string {
	content, err := ioutil.ReadFile("/var/lib/valet/cmd_db")
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
			return v
		}
	}

	y := make(map[string]string)
	return y

}

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
