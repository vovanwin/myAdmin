package dto

type Device struct {
	ID   string
	Name string
	Type string
}

type DeviceDashboard struct {
	ID     string
	Name   string
	Status string
}

type User struct {
	ID    string
	Name  string
	Email string
}

type Profile struct {
	ID    string
	Name  string
	Email string
	Role  string
}
