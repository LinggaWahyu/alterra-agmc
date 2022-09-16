package database

import (
	"alterra-agmc-day-4/config"
	"alterra-agmc-day-4/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	config.InitDB()
	users, err := GetUsers()

	assert.NoError(t, err)
	assert.NotEmpty(t, users)
}

func TestUserGetById(t *testing.T) {
	config.InitDB()
	user := models.User{
		Name:     "test",
		Email:    "test@test.com",
		Password: "test",
	}

	userCreated, err := CreateUser(user)
	assert.NoError(t, err)
	assert.NotEmpty(t, userCreated)

	getUserSuccess, err := GetUserById(int(userCreated.ID))
	assert.NoError(t, err)
	assert.NotEmpty(t, getUserSuccess)

	getUserError, err := GetUserById(-1)
	assert.Error(t, err)
	assert.Empty(t, getUserError)
}

func TestCreateUser(t *testing.T) {
	config.InitDB()
	user := models.User{
		Name:     "test",
		Email:    "test@test.com",
		Password: "test",
	}

	userCreatedSuccess, err := CreateUser(user)
	assert.NoError(t, err)
	assert.NotEmpty(t, userCreatedSuccess)
}

func TestUpdateUser(t *testing.T) {
	config.InitDB()
	user := models.User{
		Name:     "test",
		Email:    "test@test.com",
		Password: "test",
	}

	userCreated, err := CreateUser(user)
	assert.NoError(t, err)
	assert.NotEmpty(t, userCreated)

	userCreated.Name = "test update"
	userCreated.Email = "test_update@test.com"
	userCreated.Password = "test_update"

	err = UpdateUser(userCreated)
	assert.NoError(t, err)
}

func TestDeleteUser(t *testing.T) {
	config.InitDB()
	user := models.User{
		Name:     "test",
		Email:    "test@test.com",
		Password: "test",
	}

	userCreated, err := CreateUser(user)
	assert.NoError(t, err)
	assert.NotEmpty(t, userCreated)

	err = DeleteUser(int(userCreated.ID))
	assert.NoError(t, err)
}

func TestLoginUser(t *testing.T) {
	config.InitDB()
	user := models.User{
		Name:     "test",
		Email:    "test@test.com",
		Password: "test",
	}

	userCreated, err := CreateUser(user)
	assert.NoError(t, err)
	assert.NotEmpty(t, userCreated)

	loginUserSuccess, err := LoginUser(userCreated)
	assert.NoError(t, err)
	assert.NotEmpty(t, loginUserSuccess)

	err = DeleteUser(int(userCreated.ID))
	assert.NoError(t, err)

	loginUserError, err := LoginUser(userCreated)
	assert.Error(t, err)
	assert.Empty(t, loginUserError)
}
