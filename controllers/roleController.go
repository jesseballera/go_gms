package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jesseballera/go_gms/database"
	"github.com/jesseballera/go_gms/models/iam"
	"gorm.io/gorm"
	"net/http"
)

type RoleRepository struct {
	Db *gorm.DB
}

func NewRoleRepository() *RoleRepository {
	db := database.InitDb()
	db.AutoMigrate(&iam.Role{})
	return &RoleRepository{Db: db}
}

func (repository *RoleRepository) CreateRole(c *gin.Context) {
	var role iam.Role
	c.BindJSON(&role)
	err := iam.CreateRole(repository.Db, &role)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, role)
}

func (repository *RoleRepository) GetRoles(c *gin.Context) {
	var roles []iam.Role
	err := iam.GetAllRoles(repository.Db, &roles)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, roles)
}
