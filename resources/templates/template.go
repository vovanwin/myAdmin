package templates

import (
	"html/template"
	"log"
)

type AdminTemplate struct {
	*template.Template
}

func ProvideTemplate() AdminTemplate {
	tmpl, err := template.ParseFiles(
		"resources/templates/layout/base.gohtml",

		"resources/templates/components/pagination.gohtml",
		"resources/templates/components/table.gohtml",
		"resources/templates/components/header.gohtml",
		"resources/templates/components/breadcrumb.gohtml",

		"resources/templates/layout/navbar.gohtml",
		"resources/templates/layout/footer.gohtml",

		"resources/templates/page/users/userList.gohtml",
	)
	if err != nil {
		log.Fatal(err)
	}

	return AdminTemplate{
		tmpl,
	}
}
