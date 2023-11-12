package controllers

import (
	"strconv"
	"wareHouse/dao"
	"wareHouse/service"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	SetupRoutes(r *gin.Engine)
}

type controller struct {
	service service.Service
}

func NewController(service service.Service) Controller {
	return &controller{service}
}

func (c *controller) SetupRoutes(r *gin.Engine) {
	r.POST("/users", c.registerUser)
	r.GET("/users", c.getAllUsers)
	r.GET("/users/:id", c.getUserByID)
}

func (c *controller) registerUser(ctx *gin.Context) {
	var user dao.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := c.service.RegisterUser(&user); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, user)
}

func (c *controller) getAllUsers(ctx *gin.Context) {
	users, err := c.service.GetAllUsers()
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, users)
}

func (c *controller) getUserByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	user, err := c.service.GetUserByID(uint(id))
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, user)
}
