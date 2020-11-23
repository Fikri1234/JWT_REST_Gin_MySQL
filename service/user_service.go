package service

import (
	"JWT_REST_Gin_MySQL/model"
	"JWT_REST_Gin_MySQL/repository"
	"JWT_REST_Gin_MySQL/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// RoutesUser ...
func RoutesUser(rg *gin.RouterGroup) {
	user := rg.Group("/user")

	user.GET("/:id", util.TokenAuthMiddleware(), getUserByID)
	user.GET("/", util.TokenAuthMiddleware(), getUsers)
	user.POST("/", util.TokenAuthMiddleware(), createUser)
	user.PUT("/", util.TokenAuthMiddleware(), updateUser)
	user.DELETE("/:id", util.TokenAuthMiddleware(), deleteUserByID)
}

func getUserByID(c *gin.Context) {
	var user model.MUser
	paramID := c.Param("id")
	varID, err := strconv.ParseInt(paramID, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user, err = repository.GetUserByID(varID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if (model.MUser{}) == user {
		c.JSON(http.StatusNotFound, user)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func getUsers(c *gin.Context) {

	var users []model.MUser
	users, err := repository.GetUserAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)

}

func createUser(c *gin.Context) {

	var user model.MUser

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid json"})
		return
	}

	user, err := repository.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func updateUser(c *gin.Context) {

	var user model.MUser

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid json"})
		return
	}

	usr, err := repository.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, usr)
}

func deleteUserByID(c *gin.Context) {

	var user model.MUser

	paramID := c.Param("id")
	varID, err := strconv.ParseInt(paramID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = repository.DeleteUserByID(varID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, user)
}
