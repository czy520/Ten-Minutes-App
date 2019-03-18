package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lotteryjs/ten-minutes-api/model"
)

// The UserDatabase interface for encapsulating database access.
type UserDatabase interface {
	CreateUser(user *model.User) error
	GetUsers() []*model.User
}

// The UserAPI provides handlers for managing users.
type UserAPI struct {
	DB UserDatabase
}

// GetUsers returns all the users
func (a *UserAPI) GetUsers(ctx *gin.Context) {
	ctx.JSON(200, a.DB.GetUsers())
}
