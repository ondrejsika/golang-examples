package main

import (
	"fmt"
	"log"

	"github.com/AlecAivazis/survey/v2"
)

func main() {
	answers := struct {
		Name     string
		Password string
		Pet      string
		PetName  string
		Age      int
	}{}

	err := survey.Ask([]*survey.Question{
		{
			Name:     "name",
			Prompt:   &survey.Input{Message: "What is your name?"},
			Validate: survey.Required,
		},
		{
			Name:     "password",
			Prompt:   &survey.Password{Message: "Please enter your password:"},
			Validate: survey.Required,
		},
		{
			Name: "pet",
			Prompt: &survey.Select{
				Message: "What's your pet?",
				Options: []string{"Dog", "Cat", "Iguana"},
			},
		},
		{
			Name: "petname",
			Prompt: &survey.Input{
				Message: "What is your pet's name?",
			},
		},
		{
			Name:   "age",
			Prompt: &survey.Input{Message: "How old is your pet?"},
		},
	}, &answers)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf(
		"Hello %s! You have a %s named %s that is %d years old.\n",
		answers.Name, answers.Pet, answers.PetName, answers.Age,
	)
	fmt.Printf("Your password is: \"%s\".\n", answers.Password)
}
