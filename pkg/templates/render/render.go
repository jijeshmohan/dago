package render

import (
	"github.com/jijeshmohan/dago/pkg/tasks"
	"github.com/jijeshmohan/dago/pkg/templates"
	"github.com/jijeshmohan/dago/pkg/xfilesystem"
	"github.com/jijeshmohan/dago/pkg/xlogger"
)

type Renderer interface {
	Render(template templates.Template, data map[string]interface{}, rwFS xfilesystem.RwFS) error
}

type FilesystemRenderer struct {
	logger       xlogger.Logger
	taskExecuter *tasks.TaskExecuter
}

func NewFilesystemRender(logger xlogger.Logger) Renderer {
	return FilesystemRenderer{
		logger:       logger,
		taskExecuter: tasks.NewTaskExecuter(logger),
	}
}

func (fs FilesystemRenderer) Render(template templates.Template, data map[string]interface{}, rwFS xfilesystem.RwFS) error {
	fs.logger.Info("generating template %s in %s", template.Name, rwFS.BasePath())
	if err := recursiveCopy(template.Filesystem(), rwFS, data); err != nil {
		if err := rwFS.CleanupNewFiles(); err != nil {
			fs.logger.Error("unable to clean up files %s", err)
		}
		return err
	}

	if err := fs.taskExecuter.ExecuteTasks(template.Tasks, data, rwFS.BasePath()); err != nil {
		if err := rwFS.CleanupNewFiles(); err != nil {
			fs.logger.Error("unable to clean up files %s", err)
		}
		return err
	}

	fs.logger.Info("template generated successfully")
	return nil
}
