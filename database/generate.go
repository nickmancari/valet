package database

import (
	"fmt"
	"os"
)

//Creates a JSON file from which to start the database.
func Generate() interface{} {
	fmt.Println("\nWould you like to generate a database file? [Y/n]"+"\n")
	var generate string
	fmt.Scan(&generate)

	if generate == "Y" {
		err := os.Mkdir("/usr/local/bin/valet", 0775)
		if err != nil {
			fmt.Println(err)
		}
		f, err := os.Create("/usr/local/bin/valet/cmd_db")
		if err != nil {
			fmt.Println(err)
		}
		fresh := []byte(`
{
	"edit": {
		"vim": "/usr/local/bin/valet/cmd_db"
	}
}`)

		f.Write(fresh)
		
		r, err := fmt.Println("Database file created at /usr/local/bin/valet")
		if err != nil {
			fmt.Println(err)
		}
		return r
	} else {
		fmt.Println("\nDatabase file creation cancelled."+"\n")
	}
	return ""
}
