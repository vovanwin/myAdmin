package devices

import (
	"github.com/gin-gonic/gin"
)

// Implementation содержит логику контроллера для работы с ролями.
type Implementation struct {
	app *gin.Engine
}

// RegisterRoutes регистрирует маршруты для работы с ролями.
func RegisterRoutes(app *gin.Engine) {
	controller := &Implementation{
		app: app,
	}

	// Регистрация маршрута для отображения экрана ролей.
	//app.GET("/devices", controller.DevicesScreen)
	app.GET("/dashboard", controller.DashboardScreen)
	app.GET("/profile", controller.ProfileScreen)
	app.GET("/users", controller.UsersScreen)
	app.GET("/users/:id/edit", controller.EditUserForm)
	app.GET("/devices", controller.DevicesScreen)
	app.GET("/devices/:id", controller.DeviceDetailScreen)
	app.GET("/devices/:id/edit", controller.EditDeviceForm)
}

func (i *Implementation) DashboardScreen(c *gin.Context) {
	c.HTML(200, "dashboard.gohtml", gin.H{
		"Title":   "Дашборд",
		"Content": "Добро пожаловать в админку!",
	})
}
func (i *Implementation) ProfileScreen(c *gin.Context) {
	// Пример данных пользователя (в реальном приложении данные будут браться из базы данных)
	user := map[string]string{
		"Name":  "Иван Иванов",
		"Email": "ivan@example.com",
	}

	c.HTML(200, "profile.gohtml", gin.H{
		"Title": "Профиль пользователя",
		"User":  user,
	})
}

func (i *Implementation) UpdateProfile(c *gin.Context) {
	//// Логика обновления профиля
	//name := c.PostForm("name")
	//email := c.PostForm("email")

	// Здесь можно добавить логику сохранения данных в базу данных

	c.Redirect(302, "/profile")
}

func (i *Implementation) UsersScreen(c *gin.Context) {
	// Пример данных пользователей (в реальном приложении данные будут браться из базы данных)
	users := []map[string]string{
		{"ID": "1", "Name": "Иван Иванов", "Email": "ivan@example.com"},
		{"ID": "2", "Name": "Петр Петров", "Email": "petr@example.com"},
	}

	c.HTML(200, "list_users.gohtml", gin.H{
		"Title": "Список пользователей",
		"Users": users,
	})
}

func (i *Implementation) EditUserForm(c *gin.Context) {
	userID := c.Param("id")

	// Пример данных пользователя (в реальном приложении данные будут браться из базы данных)
	user := map[string]string{
		"ID":    userID,
		"Name":  "Иван Иванов",
		"Email": "ivan@example.com",
	}

	c.HTML(200, "edit_user.gohtml", gin.H{
		"Title": "Редактирование пользователя",
		"User":  user,
	})
}

func (i *Implementation) UpdateUser(c *gin.Context) {

	// Логика обновления пользователя в базе данных

	c.Redirect(302, "/users")
}

func (i *Implementation) DevicesScreen(c *gin.Context) {
	// Пример данных устройств (в реальном приложении данные будут браться из базы данных)
	devices := []map[string]string{
		{"ID": "1", "Name": "Устройство 1", "Status": "Активно"},
		{"ID": "2", "Name": "Устройство 2", "Status": "Неактивно"},
	}

	c.HTML(200, "list_devices.gohtml", gin.H{
		"Title":   "Список устройств",
		"Devices": devices,
	})
}

func (i *Implementation) EditDeviceForm(c *gin.Context) {
	deviceID := c.Param("id")

	// Пример данных устройства (в реальном приложении данные будут браться из базы данных)
	device := map[string]string{
		"ID":     deviceID,
		"Name":   "Устройство 1",
		"Status": "Активно",
	}

	c.HTML(200, "edit_device.gohtml", gin.H{
		"Title":  "Редактирование устройства",
		"Device": device,
	})
}

func (i *Implementation) UpdateDevice(c *gin.Context) {
	//deviceID := c.Param("id")
	//name := c.PostForm("name")
	//status := c.PostForm("status")

	// Логика обновления устройства в базе данных

	c.Redirect(302, "/devices")
}

func (i *Implementation) DeviceDetailScreen(c *gin.Context) {
	deviceID := c.Param("id")

	// Пример данных устройства (в реальном приложении данные будут браться из базы данных)
	device := map[string]string{
		"ID":     deviceID,
		"Name":   "Устройство 1",
		"Status": "Активно",
	}

	c.HTML(200, "device_detail.gohtml", gin.H{
		"Title":  "Детали устройства",
		"Device": device,
	})
}
