package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var results []string

func Index(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "Welcome!\n")
}

func Index2(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "prueb")
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(response, "Error reading request body",
			http.StatusInternalServerError)
	}
	results = append(results, string(body))

	log.Print(results)
	log.Print(request.Body)
	log.Print(err)
}

func CommandExecuted(w http.ResponseWriter, r *http.Request) {
	var command Command
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &command); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	command.Executed = time.Now()

	cmdName := "node"
	cmdArgs := []string{"/home/quique/lab/labcasa/api-post-example/index.js"}
	go execComand(cmdName, &cmdArgs, &command)
	time.Sleep(time.Millisecond)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(command); err != nil {
		panic(err)
	}

}
