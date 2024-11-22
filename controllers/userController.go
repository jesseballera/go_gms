package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jesseballera/go_gms/database"
	"github.com/jesseballera/go_gms/models/iam"
	"gorm.io/gorm"
	"net/http"
	_ "strconv"
)

type UserRepository struct {
	Db *gorm.DB
}

func NewUserRepository() *UserRepository {
	db := database.InitDb()
	db.AutoMigrate(&iam.User{})
	return &UserRepository{Db: db}
}

// CreateUser create user
func (repository *UserRepository) CreateUser(c *gin.Context) {
	var user iam.User
	c.BindJSON(&user)
	err := iam.CreateUser(repository.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

func New() *UserRepository {
	db := database.InitDb()
	db.AutoMigrate(&iam.User{})
	return &UserRepository{Db: db}
}

// GetUsers get users
func (repository *UserRepository) GetUsers(c *gin.Context) {
	var user []iam.User
	err := iam.GetAllUsers(repository.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}
