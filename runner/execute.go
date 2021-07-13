package runner

import (
	"os"
	"os/exec"
	"syscall"
)

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
