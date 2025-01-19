package devices

import (
	"bytes"
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"math/rand"
	"myAdmin/internal/dto"
	templates "myAdmin/resources/templates/layout"
	dashboardHTML "myAdmin/resources/templates/page/dashboard"
	devicesHTML "myAdmin/resources/templates/page/devices"
	profileHTML "myAdmin/resources/templates/page/profile"
	usersHTML "myAdmin/resources/templates/page/users"
	"time"
)

// Implementation содержит логику контроллера для работы с устройствами.
type Implementation struct {
	app *gin.Engine
}

// RegisterRoutes регистрирует маршруты для работы с устройствами.
func RegisterRoutes(app *gin.Engine) {
	controller := &Implementation{
		app: app,
	}

	// Регистрация маршрутов
	app.GET("/dashboard", controller.DashboardScreen)
	app.GET("/users", controller.UsersScreen)
	app.GET("/devices", controller.DevicesScreen)
	app.GET("/profile", controller.ProfileScreen)
	app.GET("/devices-status", controller.DashboardDevicesScreen)
	app.GET("/devices-status-sse", controller.sseDevices)
}

// DashboardScreen отображает дашборд.
func (i *Implementation) DashboardScreen(c *gin.Context) {
	chart := createBarChart()
	render(c, 200, "Дашборд", dashboardHTML.Dashboard(chart))
}

// DashboardScreen отображает дашборд.
func (i *Implementation) DashboardDevicesScreen(c *gin.Context) {
	// Пример данных устройств
	devices := []dto.DeviceDashboard{
		{ID: "1", Name: "Устройство 1", Status: randomStatus()},
		{ID: "2", Name: "Устройство 2", Status: randomStatus()},
		{ID: "3", Name: "Устройство 3", Status: randomStatus()},
	}
	render(c, 200, "Дашборд", dashboardHTML.DevicesStatus(devices))
}
func (i *Implementation) sseDevices(c *gin.Context) {
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")

	// Отправляем обновления каждые 5 секунд
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-c.Writer.CloseNotify():
			// Клиент закрыл соединение
			return
		case <-ticker.C:
			// Отправляем текущее состояние устройств
			// Генерируем HTML с помощью шаблона

			devices := []dto.DeviceDashboard{
				{ID: "1", Name: "Устройство 1", Status: randomStatus()},
				{ID: "2", Name: "Устройство 2", Status: randomStatus()},
				{ID: "3", Name: "Устройство 3", Status: randomStatus()},
			}
			// Генерация HTML как строки
			var buf bytes.Buffer
			dashboardHTML.DevicesStatus(devices).Render(c, &buf)
			html := buf.String()

			c.SSEvent("message", html)
			c.Writer.Flush()
		}
	}
}

// UsersScreen отображает список пользователей.
func (i *Implementation) UsersScreen(c *gin.Context) {
	// Пример данных пользователей
	users := []dto.User{
		{ID: "1", Name: "Иван Иванов", Email: "ivan@example.com"},
		{ID: "2", Name: "Петр Петров", Email: "petr@example.com"},
	}
	render(c, 200, "Список пользователей", usersHTML.ListUsers(users))
}

// DevicesScreen отображает список устройств.
func (i *Implementation) DevicesScreen(c *gin.Context) {
	// Пример данных устройств
	devices := []dto.Device{
		{ID: "1", Name: "Устройство 1", Type: "Активно"},
		{ID: "2", Name: "Устройство 2", Type: "Неактивно"},
	}
	render(c, 200, "Список устройств", devicesHTML.ListDevices(devices))
}

// DevicesScreen отображает список устройств.
func (i *Implementation) ProfileScreen(c *gin.Context) {
	// Пример данных устройств
	admin := dto.Profile{
		ID:    "1",
		Name:  "Владимир",
		Email: "",
		Role:  "Администратор",
	}
	render(c, 200, "Список устройств", profileHTML.Profile(admin))
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

// randomStatus генерирует случайный статус для устройства.
func randomStatus() string {
	if rand.Intn(2) == 0 {
		return "open"
	}
	return "closed"
}
