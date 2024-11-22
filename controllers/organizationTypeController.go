package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jesseballera/go_gms/database"
	"github.com/jesseballera/go_gms/models/core"
	"github.com/jesseballera/go_gms/models/iam"
	"github.com/jesseballera/go_gms/utils"
	"gorm.io/gorm"
	"net/http"
)

type OrganizationTypeRepository struct {
	Db *gorm.DB
}

func NewOrganizationTypeRepository() *OrganizationTypeRepository {
	db := database.InitDb()
	db.AutoMigrate(&iam.OrganizationType{})
	return &OrganizationTypeRepository{Db: db}
}

func (repository *OrganizationTypeRepository) CreateOrganizationType(c *gin.Context) {
	var organizationType iam.OrganizationType
	c.BindJSON(&organizationType)
	err := iam.CreateOrganizationType(repository.Db, &organizationType)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, organizationType)
}

func (repository *OrganizationTypeRepository) GetOrganizationTypes(c *gin.Context) {
	var organizationTypes []iam.OrganizationType
	err := iam.GetOrganizationTypes(repository.Db, &organizationTypes)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, organizationTypes)
}

func (repository *OrganizationTypeRepository) GetOrganizationTypeById(c *gin.Context) {
	organizationTypeId := uuid.MustParse(c.Param("id"))
	var organizationType iam.OrganizationType
	err := iam.GetOrganizationTypeById(repository.Db, &organizationType, organizationTypeId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, organizationType)
}

// GetOrganizationTypeByName get organization type by name
func (repository *OrganizationTypeRepository) GetOrganizationTypeByName(c *gin.Context) {
	organizationTypeName := c.Param("name")
	var organizationType iam.OrganizationType
	err := iam.GetOrganizationTypeByName(repository.Db, &organizationType, organizationTypeName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, organizationType)
}

func (oraganizationTypeRepository *OrganizationTypeRepository) GetPaginatedOrganizationTypes(pagination utils.Pagination) (*utils.Pagination, error) {
	var categories []*iam.OrganizationType

	oraganizationTypeRepository.Db.Scopes(core.Paginate(categories, &pagination, oraganizationTypeRepository.Db)).Find(&categories)
	pagination.Rows = categories

	return &pagination, nil
}
