package domain

import "github.com/yoskeoka/go-example/mock/domain/model"

// User represents CRUD operation to User model.
type User interface {
	Read(userID int) (*model.User, error)
	Create(user *model.User) (*model.User, error)
	Update(user *model.User) error
	Delete(userID int) error
}

// UserGroup represents CRUD operation to UserGroup model.
type UserGroup interface {
	Read(grpID int) (*model.UserGroup, error)
	Create(grp *model.UserGroup) (*model.UserGroup, error)
	Update(grp *model.UserGroup) error
	Delete(grpID int) error
	ListUser(grpID int) ([]*model.User, error)
	AddUser(grpID int, userID int) error
	DeleteUser(grpID int, userID int) error
}
