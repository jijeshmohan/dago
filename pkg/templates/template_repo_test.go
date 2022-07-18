package templates

import (
	"testing"
	"testing/fstest"

	"github.com/jijeshmohan/dago/pkg/xfilesystem"
	"github.com/stretchr/testify/require"
)

func TestNewRepository(t *testing.T) {
	t.Run("should load existing templates - ignore files", func(t *testing.T) {
		fsTest := xfilesystem.NewFileSystem("/tmp", fstest.MapFS{
			"sample.go":                    {},
			"template1/dago-template.yaml": {Data: []byte(`name: template1`)},
			"template2/dago-template.yaml": {Data: []byte(`name: template2`)},
		})

		repo, err := NewFSRepository(fsTest)
		require.NoError(t, err)

		templates := repo.GetAllTemplateNames()
		require.EqualValues(t, []string{
			"template1",
			"template2",
		}, templates)
	})
	t.Run("should ignore folders which are not templates", func(t *testing.T) {
		fsTest := xfilesystem.NewFileSystem("/tmp", fstest.MapFS{
			"sample.go":                    {},
			"template1/template.yaml":      {},
			"template2/dago-template.yaml": {Data: []byte(`name: template2`)},
		})

		repo, err := NewFSRepository(fsTest)
		require.NoError(t, err)

		templates := repo.GetAllTemplateNames()
		require.EqualValues(t, []string{
			"template2",
		}, templates)
	})
}

func Test_fileSystemRepository_GetTemplate(t *testing.T) {
	fsTest := xfilesystem.NewFileSystem("/tmp", fstest.MapFS{
		"template1/dago-template.yaml": &fstest.MapFile{
			Data: []byte(`name: sample`),
		},
	})

	repo, err := NewFSRepository(fsTest)
	require.NoError(t, err)

	template, err := repo.GetTemplate("sample")
	require.NoError(t, err)

	require.Equal(t, "sample", template.Name)

	_, err = repo.GetTemplate("template1")
	require.Error(t, err) // FIXME: Due to the mismatch in template name and folder
}
