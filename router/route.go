package router

import (
	"github.com/gin-gonic/gin"
	css "myAdmin/resources"
)

func NewRouter(app *gin.Engine) error {
	app.GET("/public/css/styles.css", func(c *gin.Context) {
		// Отправляем содержимое CSS, встроенное в бинарник
		c.Data(200, "text/css", []byte(css.CssContent))
	})
	return nil
}
