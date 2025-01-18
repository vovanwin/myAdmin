package templates

import (
	"github.com/gin-contrib/multitemplate"
	"path/filepath"
)

type TemplateHtml struct {
	multitemplate.Renderer
}

// NewTemplateRenderer - инициализация шаблонов
func TemplateRenderer() *TemplateHtml {
	temp := loadTemplates("resources/templates")
	return &TemplateHtml{
		temp,
	}
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	adminLayouts, err := filepath.Glob(templatesDir + "/layout/base.gohtml")
	if err != nil {
		panic(err.Error())
	}

	admins, err := filepath.Glob(templatesDir + "/page/*/*.gohtml")
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our adminLayouts/ and admins/ directories
	for _, admin := range admins {
		layoutCopy := make([]string, len(adminLayouts))
		copy(layoutCopy, adminLayouts)
		layoutCopy = append(layoutCopy, admin)
		r.AddFromFiles(filepath.Base(admin), layoutCopy...)
	}
	return r
}
