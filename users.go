package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/AlecAivazis/survey"
)

// User Struct
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var users []User

func initLogin() (User, error) {
	qs := []*survey.Question{
		{
			Name:     "email",
			Prompt:   &survey.Input{Message: "Please Enter Your Email:"},
			Validate: survey.Required,
		},
		{
			Name:     "password",
			Prompt:   &survey.Password{Message: "Please Enter Your Password:"},
			Validate: survey.Required,
		},
	}

	user := User{}

	err := survey.Ask(qs, &user)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	e := user.login()
	return user, e
}

func initSignup() (User, error) {
	qs := []*survey.Question{
		{
			Name:     "email",
			Prompt:   &survey.Input{Message: "Please Enter Your Email:"},
			Validate: survey.Required,
		},
		{
			Name:     "password",
			Prompt:   &survey.Password{Message: "Please Enter Your Password:"},
			Validate: survey.Required,
		},
	}

	user := User{}

	err := survey.Ask(qs, &user)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	e := user.find()
	if e == nil {
		users = append(users, user)
	}

	data, _ := json.Marshal(users)
	_ = ioutil.WriteFile("./data/users.json", data, 0644)

	return user, e
}

func (user *User) login() error {
	for _, u := range users {
		if u.Email == user.Email && u.Password == user.Password {
			return nil
		}
	}

	return errors.New("Invalid Credentials")
}

func (user *User) find() error {
	for _, u := range users {
		if u.Email == user.Email {
			return errors.New("User Already Exists")
		}
	}

	return nil
}

func loadusers() {
	jsonFile, err := ioutil.ReadFile("./data/users.json")
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(jsonFile, &users)
}
