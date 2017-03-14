package main

import (
	"log"
	"os"
	"os/exec"
)

func execComand(cmdExec string, args *[]string, command *Command) {

	cmd := exec.Command(cmdExec, *args...)

	err := cmd.Start()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	command.PID = cmd.Process.Pid
	log.Print(command)
	cmd.Wait()

}
