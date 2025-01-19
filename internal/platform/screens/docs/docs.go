package docs

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	templates "myAdmin/resources/templates/layout"
	docsHTML "myAdmin/resources/templates/page/docs"
)

// Implementation содержит обработчики для работы с пользователями.
type Implementation struct {
	app *gin.Engine
}

// UserScreen регистрирует маршруты и их обработчики.
func DocsScreen(app *gin.Engine) {
	controller := &Implementation{
		app: app,
	}

	// Регистрация маршрутов
	app.GET("/docs", controller.Swagger)
}

func (i *Implementation) getBaseURL() string {
	return "/swagger/file"
}

func (i *Implementation) Swagger(c *gin.Context) {
	render(c, 200, "Документация", docsHTML.Docs())

}

// render - универсальная функция рендеринга с поддержкой HTMX
func render(c *gin.Context, code int, title string, content templ.Component) {
	if c.GetHeader("HX-Request") != "" {
		// Если запрос HTMX, рендерим только контент
		content.Render(c.Request.Context(), c.Writer)
	} else {
		// Если обычный запрос, рендерим с базовым шаблоном
		templates.Base(title, content).Render(c.Request.Context(), c.Writer)
	}
}
