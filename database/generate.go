package database

import (
	"fmt"
	"os"
)


func Generate() interface{} {
	fmt.Println("\nWould you like to generate a database file? [Y/n]"+"\n")
	var generate string
	fmt.Scan(&generate)

	if generate == "Y" {
		err := os.Mkdir("/var/lib/valet", 0775)
		if err != nil {
			fmt.Println(err)
		}
		f, err := os.Create("/var/lib/valet/cmd_db")
		if err != nil {
			fmt.Println(err)
		}
		fresh := []byte(`
{
	"edit": {
		"vim": "/var/lib/valet/cmd_db"
	}
}`)

		f.Write(fresh)
		
		r, err := fmt.Println("Database file created at /var/lib/valet")
		if err != nil {
			fmt.Println(err)
		}
		return r
	} else {
		fmt.Println("\nDatabase file creation cancelled."+"\n")
	}
	return ""
}
