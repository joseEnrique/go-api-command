package main

import "time"

type Command struct {
	Id       int       `json:"id"`
	Command  string    `json:"command"`
	PID      int       `json:"PID"`
	Executed time.Time `json:"executed"`
}

type Commands []Command
