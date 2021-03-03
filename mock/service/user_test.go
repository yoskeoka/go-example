package service_test

import (
	"reflect"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/yoskeoka/go-example/mock/domain"
	"github.com/yoskeoka/go-example/mock/domain/model"
	"github.com/yoskeoka/go-example/mock/service"
)

func TestUser_Create_1(t *testing.T) {

	r := &serviceRegistryMock{}
	r.userRepoMock.FakeCreate = func(user *model.User) (*model.User, error) {
		created := &model.User{ID: 7, Name: user.Name, Address: user.Address}
		return created, nil
	}

	r.userGrpRepoMock.FakeCreate = func(grp *model.UserGroup) (*model.UserGroup, error) {
		created := &model.UserGroup{ID: 9, Name: grp.Name, Private: grp.Private}
		return created, nil
	}
	r.userGrpRepoMock.FakeAddUser = func(grpID int, userID int) error { return nil }

	userSvc := service.NewUser(r)

	userInput := model.User{Name: "John", Address: "Kyoto"}
	got, err := userSvc.Create(&userInput)
	if err != nil {
		t.Fatal(err)
	}

	if got.ID != 7 {
		t.Errorf("User.Create() should return model.User.ID = 7, but got = %d", got.ID)
	}

	if got.Name != userInput.Name {
		t.Errorf("User.Create() should return model.User.Name = %s, but got = %s", userInput.Name, got.Name)
	}

	// snip...
}

func TestUser_Create_TableDrivenTests(t *testing.T) {
	type userFakes struct {
		Create func(user *model.User) (*model.User, error)
	}
	type userGrpFakes struct {
		Create  func(grp *model.UserGroup) (*model.UserGroup, error)
		AddUser func(grpID int, userID int) error
	}
	type args struct {
		user *model.User
	}
	tests := []struct {
		name         string
		userFakes    userFakes
		userGrpFakes userGrpFakes
		args         args
		want         *model.User
		wantErr      bool
	}{
		{"create successfully",
			userFakes{
				Create: func(user *model.User) (*model.User, error) {
					created := &model.User{ID: 7, Name: user.Name, Address: user.Address}
					return created, nil
				},
			},
			userGrpFakes{
				Create: func(grp *model.UserGroup) (*model.UserGroup, error) {
					created := &model.UserGroup{ID: 9, Name: grp.Name, Private: grp.Private}
					return created, nil
				},
				AddUser: func(grpID int, userID int) error {
					return nil
				},
			},
			args{&model.User{Name: "John", Address: "Kyoto"}},
			&model.User{ID: 7, Name: "John", Address: "Kyoto"},
			false,
		},
		{"create error",
			userFakes{
				Create: func(user *model.User) (*model.User, error) {
					created := &model.User{ID: 7, Name: user.Name, Address: user.Address}
					return created, nil
				},
			},
			userGrpFakes{
				Create: func(grp *model.UserGroup) (*model.UserGroup, error) {
					created := &model.UserGroup{ID: 9, Name: grp.Name, Private: grp.Private}
					return created, nil
				},
				AddUser: func(grpID int, userID int) error {
					return nil
				},
			},
			args{&model.User{Name: "John", Address: "Kyoto"}},
			&model.User{ID: 7, Name: "John", Address: "Kyoto"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &serviceRegistryMock{}
			r.userRepoMock.FakeCreate = tt.userFakes.Create
			r.userGrpRepoMock.FakeCreate = tt.userGrpFakes.Create
			r.userGrpRepoMock.FakeAddUser = tt.userGrpFakes.AddUser
			u := service.NewUser(r)
			got, err := u.Create(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_Create_2_mockgen(t *testing.T) {

	ctrl := gomock.NewController(t)

	userRepo := NewMockUser(ctrl)
	userGrpRepo := NewMockUserGroup(ctrl)

	userRepo.EXPECT().
		Create(gomock.Any()).
		DoAndReturn(func(user *model.User) (*model.User, error) {
			return &model.User{ID: 7, Name: user.Name, Address: user.Address}, nil
		}).
		AnyTimes()

	userGrpRepo.EXPECT().
		Create(gomock.Any()).
		DoAndReturn(func(grp *model.UserGroup) (*model.UserGroup, error) {
			return grp, nil
		}).
		AnyTimes()

	userGrpRepo.EXPECT().
		AddUser(gomock.Any(), gomock.Any()).
		DoAndReturn(func(grpID int, userID int) error {
			return nil
		}).
		AnyTimes()

	r := NewMockServiceRegistryInterface(ctrl)
	r.EXPECT().
		User().
		DoAndReturn(func() domain.User { return userRepo }).
		AnyTimes()

	r.EXPECT().
		UserGroup().
		DoAndReturn(func() domain.UserGroup { return userGrpRepo }).
		AnyTimes()

	userSvc := service.NewUser(r)

	userInput := model.User{Name: "John", Address: "Kyoto"}
	got, err := userSvc.Create(&userInput)
	if err != nil {
		t.Fatal(err)
	}

	if got.ID != 7 {
		t.Errorf("User.Create() should return model.User.ID = 7, but got = %d", got.ID)
	}

	if got.Name != userInput.Name {
		t.Errorf("User.Create() should return model.User.Name = %s, but got = %s", userInput.Name, got.Name)
	}

	// snip...
}
