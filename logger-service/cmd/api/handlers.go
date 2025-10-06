package main

import (
	"log"
	"net/http"

	"log-service/data"
)

type JSONPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	var payload JSONPayload

	err := app.readJSON(w, r, &payload)
	if err != nil {
		log.Println("Can't parse data", err)
		return
	}

	var entry data.LogEntry
	entry.Name = payload.Name
	entry.Data = payload.Data

	err = app.Models.LogEntry.Insert(entry)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	res := jsonResponse{
		Error:   false,
		Message: "Logged",
	}

	app.writeJSON(w, http.StatusAccepted, res)
}
