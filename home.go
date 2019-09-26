package main

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey"
)

func showHomePrompt() string {
	qs := []*survey.Question{
		{
			Name: "Ans",
			Prompt: &survey.Select{
				Message: "Select Option",
				Options: []string{"Login", "Signup", "Exit"},
				Default: "Login",
			},
			Validate: survey.Required,
		},
	}

	a := struct {
		Ans string
	}{}

	err := survey.Ask(qs, &a)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return a.Ans
}

func showLoggedInUserPrompt() string {
	qs := []*survey.Question{
		{
			Name: "Ans",
			Prompt: &survey.Select{
				Message: "Select Option",
				Options: []string{"Add New Entry", "View Entries", "Logout"},
				Default: "Add New Entry",
			},
			Validate: survey.Required,
		},
	}

	a := struct {
		Ans string
	}{}

	err := survey.Ask(qs, &a)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return a.Ans
}
