package tasks

import (
	"fmt"
	"path/filepath"

	"github.com/jijeshmohan/dago/pkg/textrender"
	"github.com/jijeshmohan/dago/pkg/xlogger"
)

type TaskExecuter struct {
	logger xlogger.Logger
}

func NewTaskExecuter(logger xlogger.Logger) *TaskExecuter {
	return &TaskExecuter{
		logger: logger,
	}
}

func (t *TaskExecuter) ExecuteTasks(tasks []Task, data map[string]interface{}, path string) error {
	for _, task := range tasks {
		commandPath := path
		if task.Path != "" {
			renderPath, err := textrender.RenderString(task.Path, data)
			if err != nil {
				return fmt.Errorf("invalid path in command %s: %v", task.Command, err)
			}

			commandPath = filepath.Join(path, renderPath)
		}
		output, err := task.execute(data, commandPath, task.Arguments...)
		if err != nil {
			if task.IgnoreError {
				t.logger.Warn("ignoring error for command %s", err)
				continue
			}

			return err
		}

		t.logger.Info(output)
	}

	return nil
}
