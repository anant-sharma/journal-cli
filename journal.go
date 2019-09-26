package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"time"

	"github.com/AlecAivazis/survey"
)

// JournalEntry struct
type JournalEntry struct {
	Date int64  `json:"date"`
	User string `json:"user"`
	Text string `json:"text"`
}

var journals []JournalEntry

func initNewEntry(user User) {

	qs := []*survey.Question{
		{
			Name:     "text",
			Prompt:   &survey.Input{Message: "Please Record Your Entry:"},
			Validate: survey.Required,
		},
	}

	j := JournalEntry{
		User: user.Email,
	}

	err := survey.Ask(qs, &j)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	j.Date = int64(time.Now().Unix())
	journals = append(journals, j)

	data, _ := json.Marshal(journals)
	_ = ioutil.WriteFile("./data/journals.json", data, 0644)
}

func showUserEntries(user User) {
	entries := make([]JournalEntry, 0)

	for _, j := range journals {
		if j.User == user.Email {
			entries = append(entries, j)
		}
	}

	sort.SliceStable(entries, func(i, j int) bool {
		return entries[i].Date > entries[j].Date
	})

	if len(entries) > 50 {
		entries = entries[:50]
	}

	for _, e := range entries {
		fmt.Println("["+time.Unix(e.Date, 0).Format("_2-Jan-2006 15:04:05")+"] -", e.Text)
	}
}

func loadJournals() {
	jsonFile, err := ioutil.ReadFile("./data/journals.json")
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(jsonFile, &journals)
}
