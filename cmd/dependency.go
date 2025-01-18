package cmd

import (
	"github.com/gin-gonic/gin"
	"log"
	"myAdmin/resources/templates"
)

func NewServerAdmin(temp *templates.TemplateHtml) (*gin.Engine, error) {
	router := gin.Default()

	router.HTMLRender = temp.Renderer
	return router, nil
}

func ServerAdminRun(app *gin.Engine) {
	if err := app.Run(":3001"); err != nil {
		log.Fatal(err)
	}
}
