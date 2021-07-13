package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type Commands struct {
	Info Data
}

type Data struct {
	Stuff string
}

func MapFromDatabase(f string) map[string]string {
	content, err := ioutil.ReadFile("cmd_db")
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
