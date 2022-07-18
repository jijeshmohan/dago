package generator

import (
	"context"
	"errors"
	"os"

	"github.com/jijeshmohan/dago/pkg/config"
	"github.com/jijeshmohan/dago/pkg/templates"
	"github.com/jijeshmohan/dago/pkg/templates/render"
	"github.com/jijeshmohan/dago/pkg/variables"
	"github.com/jijeshmohan/dago/pkg/xfilesystem"
	"github.com/jijeshmohan/dago/pkg/xlogger"
)

type Generator interface {
	Generate(ctx context.Context, templateName string, folderPath string) error
}

func NewGenerator(conf config.Config, logger xlogger.Logger) (Generator, error) {
	repo, err := templates.NewFSRepository(xfilesystem.NewFileSystem(conf.TemplatesPath, os.DirFS(conf.TemplatesPath)))
	if err != nil {
		return nil, err
	}

	return &defaultGenerator{
		conf:            conf,
		templateRepo:    repo,
		variablesGetter: variables.NewConsoleGetter(),
		templateRender:  render.NewFilesystemRender(logger),
		logger:          logger,
	}, nil
}

type defaultGenerator struct {
	conf            config.Config
	templateRepo    templates.Repository
	variablesGetter variables.VariablesGetter
	templateRender  render.Renderer
	logger          xlogger.Logger
}

func (g *defaultGenerator) Generate(ctx context.Context, templateName, folderPath string) error {
	template, err := g.templateRepo.GetTemplate(templateName)
	if err != nil {
		return err
	}

	data, err := g.variablesGetter.GetValues(template.Variables)
	if err != nil {
		return err
	}

	if err := g.templateRender.Render(template, data, xfilesystem.NewRWFileSystem(folderPath, os.DirFS(folderPath))); err != nil {
		g.logger.Error(err.Error())
		return errors.New("failed to render template")
	}

	return nil
}
