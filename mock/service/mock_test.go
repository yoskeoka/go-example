package service_test

import (
	"github.com/yoskeoka/go-example/mock/domain"
	"github.com/yoskeoka/go-example/mock/domain/model"
)

type userRepoMock struct {
	domain.User
	FakeCreate func(user *model.User) (*model.User, error)
}

func (m *userRepoMock) Create(user *model.User) (*model.User, error) {
	return m.FakeCreate(user)
}

type userGrpRepoMock struct {
	domain.UserGroup
	FakeCreate  func(grp *model.UserGroup) (*model.UserGroup, error)
	FakeAddUser func(grpID int, userID int) error
}

func (m *userGrpRepoMock) Create(grp *model.UserGroup) (*model.UserGroup, error) {
	return m.FakeCreate(grp)
}

func (m *userGrpRepoMock) AddUser(grpID int, userID int) error {
	return m.FakeAddUser(grpID, userID)
}
