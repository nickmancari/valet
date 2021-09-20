package runner

import (
	"os"
	"os/exec"
	"syscall"
)

//Takes a map, loops over it, and runs the key as the command and the value as the parameters for that command.
func ExecuteCommand(m map[string]string) interface{} {
	for k, v := range m {
		connect, err := exec.LookPath(k)
		if err != nil {
			panic(err)
		}

		syscall.Exec(connect, []string{k, v}, os.Environ())
	}
	return ""
}
