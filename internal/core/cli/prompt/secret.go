package prompt

import (
	"errors"

	"github.com/manifoldco/promptui"
)

func validateSecret(input string) error {
	if len(input) < 3 {
		return errors.New("Please make sure to enter more than 3 characters for the secret")
	}
	return nil
}

func PromptForSecret(question string) (string, error) {
	prompt := promptui.Prompt{
		Label:    question,
		Validate: validateSecret,
	}

	result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return result, nil
}
