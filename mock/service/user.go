package service

import (
	"errors"
	"fmt"

	"github.com/yoskeoka/go-example/mock/domain"
	"github.com/yoskeoka/go-example/mock/domain/model"
)

// User manages user's personal information.
type User struct {
	userRepo domain.User
	grpRepo  domain.UserGroup
}

// NewUser initializes User service.
func NewUser(
	userRepo domain.User,
	grpRepo domain.UserGroup,
) *User {
	return &User{
		userRepo: userRepo,
		grpRepo:  grpRepo,
	}
}

// Create creates a new user alongside his/her default group.
func (u *User) Create(user *model.User) (*model.User, error) {
	if user.Name == "" {
		return nil, errors.New("user name required")
	}

	createdUser, err := u.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	createdGrp, err := u.grpRepo.Create(
		&model.UserGroup{
			Name:    fmt.Sprintf("%s's default group", user.Name),
			Private: true,
		},
	)

	err = u.grpRepo.AddUser(createdGrp.ID, createdUser.ID)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
