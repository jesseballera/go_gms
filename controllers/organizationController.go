package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jesseballera/go_gms/database"
	"github.com/jesseballera/go_gms/models/iam"
	"gorm.io/gorm"
	"net/http"
)

type OrganizationRepository struct {
	Db *gorm.DB
}

func NewOrganizationRepository() *OrganizationRepository {
	db := database.InitDb()
	db.AutoMigrate(&iam.Organization{})
	return &OrganizationRepository{Db: db}
}

func (repository *OrganizationRepository) CreateOrganization(c *gin.Context) {
	var organization iam.Organization
	c.BindJSON(&organization)
	err := iam.CreateOrganization(repository.Db, &organization)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, organization)
}

func (repository *OrganizationRepository) GetOrganizations(c *gin.Context) {
	var organizations []iam.Organization
	err := iam.GetOrganizations(repository.Db, &organizations)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, organizations)
}
