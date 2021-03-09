package model

// User has user's personal information.
type User struct {
	ID      int
	Name    string
	Address string
}

// UserGroup has user group settings.
type UserGroup struct {
	ID      int
	Name    string
	Private bool
}
