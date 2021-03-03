package service_test

import (
	"reflect"
	"testing"

	"github.com/yoskeoka/go-example/mock/domain/model"
	"github.com/yoskeoka/go-example/mock/service"
)

func TestUser_Create_1(t *testing.T) {

	userRepo := &userRepoMock{
		FakeCreate: func(user *model.User) (*model.User, error) {
			created := &model.User{ID: 7, Name: user.Name, Address: user.Address}
			return created, nil
		},
	}
	userGrpRepo := &userGrpRepoMock{
		FakeCreate: func(grp *model.UserGroup) (*model.UserGroup, error) {
			created := &model.UserGroup{ID: 9, Name: grp.Name, Private: grp.Private}
			return created, nil
		},
		FakeAddUser: func(grpID int, userID int) error { return nil },
	}

	userSvc := service.NewUser(userRepo, userGrpRepo)

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
			userRepo := &userRepoMock{}
			userRepo.FakeCreate = tt.userFakes.Create
			userGrpRepo := &userGrpRepoMock{}
			userGrpRepo.FakeCreate = tt.userGrpFakes.Create
			userGrpRepo.FakeAddUser = tt.userGrpFakes.AddUser
			u := service.NewUser(userRepo, userGrpRepo)
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
