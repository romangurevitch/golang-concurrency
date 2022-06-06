package prompter

import (
	"github.com/manifoldco/promptui"
)

//go:generate mockgen -source=prompter.go -destination=mocks/prompter.go Prompter
type Prompter interface {
	ReadStringFromPrompt(label, defVal string, password bool) (string, error)
}

func New() Prompter {
	return &prompter{}
}

type prompter struct {
}

func (p *prompter) ReadStringFromPrompt(label, defVal string, password bool) (string, error) {
	prompt := promptui.Prompt{
		Label:   label,
		Default: defVal,
	}
	if password {
		prompt.Mask = '*'
	}

	return prompt.Run()
}
