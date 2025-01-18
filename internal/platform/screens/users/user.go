package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Implementation содержит обработчики для работы с пользователями.
type Implementation struct {
	app *gin.Engine
}

// UserScreen регистрирует маршруты и их обработчики.
func UserScreen(app *gin.Engine) {
	controller := &Implementation{
		app: app,
	}

	// Регистрация маршрутов
	app.GET("/", controller.HomePage)
	app.GET("/users", controller.Users)
	app.GET("/users/edit/:id", controller.UserEdit)
	app.DELETE("/users/:id", controller.UserDelete)
}

// HomePage обрабатывает запрос на главную страницу.
func (i *Implementation) HomePage(c *gin.Context) {

	c.HTML(200, "page/users/userList", gin.H{
		"title":   "Home Page",
		"content": "Welcome to the Home Page!",
	})
}

// UsersPage обрабатывает запрос на страницу пользователей.
func (i *Implementation) Users(c *gin.Context) {
	c.HTML(200, "userList.gohtml", gin.H{
		"Title": "Список пользователей",
		"Users": users,
	})

}

func (i *Implementation) UserEdit(c *gin.Context) {
	user := User{
		ID:    1,
		Name:  "Lol",
		Email: "lo.email.ru",
	}

	c.HTML(http.StatusOK, "update.gohtml", gin.H{
		"Title": "Edit User",
		"User":  user,
	})
}

func (i *Implementation) UserDelete(c *gin.Context) {
	user := User{
		ID:    1,
		Name:  "Lol",
		Email: "lo.email.ru",
	}

	c.HTML(http.StatusOK, "edit_user.html", gin.H{
		"Title": "Edit User",
		"User":  user,
	})
}

type User struct {
	ID    int
	Name  string
	Email string
}

type Device struct {
	ID   int
	Name string
	Type string
}

var users = []User{
	{ID: 1, Name: "John Doe1", Email: "john@example.com"},
	{ID: 2, Name: "Jane Doe222", Email: "jane@example.com"},
}

var devices = []Device{
	{ID: 1, Name: "Device 1", Type: "Type A"},
	{ID: 2, Name: "Device 2", Type: "Type B"},
}
