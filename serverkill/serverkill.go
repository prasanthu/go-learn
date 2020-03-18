package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func serverKill(pidFile string) error {
	data, err := ioutil.ReadFile(pidFile)
	if err != nil {
		return errors.Wrap(err, "PID file not found")
	}
	err = os.Remove(pidFile)
	if err != nil {
		fmt.Printf("Cannot remove file %s\n", pidFile)
	}
	pidString := strings.TrimSpace(string(data))
	pid, err := strconv.Atoi(pidString)
	if err != nil {
		return errors.Wrap(err, "Invalid file format")
	}
	fmt.Printf("Killing server with pid %d\n", pid)
	return nil
}

func main() {
	if err := serverKill("server.pid"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
