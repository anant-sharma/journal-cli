package main

import (
	"fmt"
	"os"
)

func start() {
	homeOption := showHomePrompt()

	switch homeOption {
	case "Login":
		{
			user, err := initLogin()
			if err != nil {
				fmt.Println(err)
				start()
			}

			onUserLogin(user)
		}
	case "Signup":
		{
			user, err := initSignup()
			if err != nil {
				fmt.Println(err)
				start()
			}

			onUserLogin(user)
		}
	case "Exit":
		{
			os.Exit(0)
		}

	}
}

func onUserLogin(user User) {
	opt := showLoggedInUserPrompt()

	switch opt {
	case "View Entries":
		{
			showUserEntries(user)
			onUserLogin(user)
		}
	case "Add New Entry":
		{
			initNewEntry(user)
			onUserLogin(user)
		}
	case "Logout":
		{
			start()
		}
	}

}

func main() {
	ensureDir()
	ensureFile("./data/users.json")
	ensureFile("./data/journals.json")
	loadusers()
	loadJournals()
	start()
}
