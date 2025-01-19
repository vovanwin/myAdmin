package templates

import (
	"github.com/gin-contrib/multitemplate"
	"html/template"
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

	for _, admin := range admins {
		baseName := filepath.Base(admin)
		layoutCopy := make([]string, len(adminLayouts))
		copy(layoutCopy, adminLayouts)
		layoutCopy = append(layoutCopy, admin)

		// Регистрация шаблона с layout
		r.AddFromFiles(baseName, layoutCopy...)

		// Регистрация шаблона без layout
		tmpl := template.Must(template.New(baseName).ParseFiles(admin))
		r.Add("nolayout_"+baseName, tmpl)
	}

	return r
}
