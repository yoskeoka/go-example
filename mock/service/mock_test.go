package service_test

import (
	"github.com/yoskeoka/go-example/mock/domain"
	"github.com/yoskeoka/go-example/mock/domain/model"
	"github.com/yoskeoka/go-example/mock/registry"
)

type serviceRegistryMock struct {
	registry.ServiceRegistryInterface
	userRepoMock
	userGrpRepoMock
}

func (sr *serviceRegistryMock) User() domain.User {
	return &sr.userRepoMock
}

func (sr *serviceRegistryMock) UserGroup() domain.UserGroup {
	return &sr.userGrpRepoMock
}

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
