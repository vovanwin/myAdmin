package users

import (
	"log"
	"myAdmin/resources/templates"
	"myAdmin/router"
	"net/http"
)

type Implementation struct {
	template templates.AdminTemplate
}

func UserScreen(
	r router.AdminRouter,
	template templates.AdminTemplate,
) {
	controller := &Implementation{
		template,
	}

	r.HandleFunc("", controller.UserScreen).Name("home")
}

func (i *Implementation) UserScreen(w http.ResponseWriter, r *http.Request) {

	user := User{
		ID:     1,
		Name:   "Иван Иванов",
		Role:   "Администратор",
		Email:  "ivan@example.com",
		Avatar: "/path/to/avatar.jpg",
	}

	dataUsers := struct {
		Title string
		Users []UserList
		Year  int
	}{
		Title: "Управление пользователями",
		Users: []UserList{
			{Name: "Админ m1", Email: "admin@m1.com", LastEdited: "24 ноя 2024"},
			{Name: "Админ m1", Email: "admin@m1.com", LastEdited: "24 ноя 2024"},
			{Name: "Админ m1", Email: "admin@m1.com", LastEdited: "24 ноя 2024"},
			{Name: "Админ m1", Email: "admin@m1.com", LastEdited: "24 ноя 2024"},
			{Name: "Админ m1", Email: "admin@m1.com", LastEdited: "24 ноя 2024"},
			{Name: "Админ m1", Email: "admin@m1.com", LastEdited: "24 ноя 2024"},
			{Name: "Админ m1", Email: "admin@m1.com", LastEdited: "24 ноя 2024"},
			{Name: "Админ m1", Email: "admin@m1.com", LastEdited: "24 ноя 2024"},
			{Name: "Админ m1", Email: "admin@m1.com", LastEdited: "24 ноя 2024"},
			{Name: "Админ m1", Email: "admin@m1.com", LastEdited: "24 ноя 2024"},
			{Name: "Админ m1", Email: "admin@m1.com", LastEdited: "24 ноя 2024"},
			{Name: "Админ m1", Email: "admin@m1.com", LastEdited: "24 ноя 2024"},
			{Name: "Админ m1", Email: "admin@m1.com", LastEdited: "24 ноя 2024"},
			{Name: "Админ m1", Email: "admin@m1.com", LastEdited: "24 ноя 2024"},
			{Name: "Админ m1", Email: "admin@m1.com", LastEdited: "24 ноя 2024"},
			{Name: "Админ m1", Email: "admin@m1.com", LastEdited: "24 ноя 2024"},
			{Name: "Админ m1", Email: "admin@m1.com", LastEdited: "24 ноя 2024"},
			{Name: "Админ m1", Email: "admin@m1.com", LastEdited: "24 ноя 2024"},
			{Name: "Админ m1", Email: "admin@m1.com", LastEdited: "24 ноя 2024"},
			{Name: "Админ m1", Email: "admin@m1.com", LastEdited: "24 ноя 2024"},
			{Name: "Админ m1", Email: "admin@m1.com", LastEdited: "24 ноя 2024"},
			{Name: "Админ m1", Email: "admin@m1.com", LastEdited: "24 ноя 2024"},
			{Name: "Админ m1", Email: "admin@m1.com", LastEdited: "24 ноя 2024"},
			{Name: "Админ m1", Email: "admin@m1.com", LastEdited: "24 ноя 2024"},
			{Name: "Админ m1", Email: "admin@m1.com", LastEdited: "24 ноя 2024"},
			{Name: "Админ m1", Email: "admin@m1.com", LastEdited: "24 ноя 2024"},
			{Name: "Админ m1", Email: "admin@m1.com", LastEdited: "24 ноя 2024"},
			{Name: "Админ m1", Email: "admin@m1.com", LastEdited: "24 ноя 2024"},
			{Name: "Админ m1", Email: "admin@m1.com", LastEdited: "24 ноя 2024"},
			{Name: "Админ m1", Email: "admin@m1.com", LastEdited: "24 ноя 2024"},
			{Name: "Админ m1", Email: "admin@m1.com", LastEdited: "24 ноя 2024"},
			{Name: "Админ m1", Email: "admin@m1.com", LastEdited: "24 ноя 2024"},
			{Name: "Админ m1", Email: "admin@m1.com", LastEdited: "24 ноя 2024"},
			{Name: "Админ m1", Email: "admin@m1.com", LastEdited: "24 ноя 2024"},
			{Name: "Админ m1", Email: "admin@m1.com", LastEdited: "24 ноя 2024"},
			{Name: "Админ m1", Email: "admin@m1.com", LastEdited: "24 ноя 2024"},
			{Name: "Админ m1", Email: "admin@m1.com", LastEdited: "24 ноя 2024"},
			{Name: "Админ m1", Email: "admin@m1.com", LastEdited: "24 ноя 2024"},
			{Name: "Админ m1", Email: "admin@m1.com", LastEdited: "24 ноя 2024"},
		},
		Year: 2024,
	}

	data := map[string]interface{}{
		"Title":   "Панель управления",
		"Users":   dataUsers,
		"Sidebar": RenderSidebar(user),
	}

	err := i.template.ExecuteTemplate(w, "userList.gohtml", data)
	if err != nil {
		log.Fatal(err)
	}
}

// Пример функции для генерации данных
func RenderSidebar(user User) map[string]interface{} {
	return map[string]interface{}{
		"UserName":   user.Name,
		"UserRole":   user.Role,
		"UserAvatar": user.Avatar,
	}
}

type User struct {
	ID         int
	Name       string
	Role       string
	Email      string
	Avatar     string
	LastEdited string
}

type UserList struct {
	Name       string
	Email      string
	LastEdited string
}
