package tasks

import (
	"errors"
	"fmt"
	"os/exec"

	"github.com/jijeshmohan/dago/pkg/textrender"
)

var ErrCommand = errors.New("error running command")

type Task struct {
	Command     string   `yaml:"command"`
	Arguments   []string `yaml:"arguments"`
	Path        string   `yaml:"path"`
	IgnoreError bool     `yaml:"ignore-error"`
}

func (c *Task) execute(data map[string]interface{}, path string, args ...string) (string, error) {
	processedArgs := make([]string, len(args))
	for i, arg := range args {
		parg, err := textrender.RenderString(arg, data)
		if err != nil {
			return "", err
		}

		processedArgs[i] = parg
	}

	cmd := exec.Command(c.Command, processedArgs...)
	cmd.Dir = path
	stdout, err := cmd.Output()
	if err != nil {
		err = updateError(c.Command, err)
		return "", err
	}

	return string(stdout), nil
}

func updateError(command string, err error) error {
	exitError, ok := err.(*exec.ExitError)
	if ok {
		return fmt.Errorf("%w: %s\n%s", ErrCommand, command, exitError.Stderr)
	}
	return fmt.Errorf("%w: %s\n%s", ErrCommand, command, err)
}
