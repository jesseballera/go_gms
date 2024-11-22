package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jesseballera/go_gms/controllers"
	"github.com/jesseballera/go_gms/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	router := setupRouter()
	_ = router.Run(":9080")

}

func setupRouter() *gin.Engine {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/users"
	//v1 := router.Group("/api/v1")
	//{
	//	users := v1.Group("/users")
	//	{
	//		users.GET("/", func(c *gin.Context) {
	//			c.JSON(200, gin.H{
	//				"message": "users",
	//			})
	//		})
	//	}
	//}

	organizationTypeRepository := controllers.NewOrganizationTypeRepository()
	router.POST("/organization-types", organizationTypeRepository.CreateOrganizationType)
	router.GET("/organization-types", organizationTypeRepository.GetOrganizationTypes)
	router.GET("/organization-types?name:name", organizationTypeRepository.GetOrganizationTypeByName)
	router.GET("/organization-types?id:id", organizationTypeRepository.GetOrganizationTypeById)

	//paginate := organizationTypeRepository.GetPaginatedOrganizationTypes()

	//router.GET("/organization-types", organizationTypeRepository.GetPaginatedOrganizationTypes(new utils.Pagination{}))

	organizationRepository := controllers.NewOrganizationRepository()
	router.POST("/organizations", organizationRepository.CreateOrganization)
	router.GET("/organizations", organizationRepository.GetOrganizations)

	roleRepository := controllers.NewRoleRepository()
	router.POST("/roles", roleRepository.CreateRole)
	router.GET("/roles", roleRepository.GetRoles)

	userRepository := controllers.NewUserRepository()
	router.POST("/users", userRepository.CreateUser)
	router.GET("/users", userRepository.GetUsers)

	url := ginSwagger.URL("localhost:8080/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, url))

	//router.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})

	return router
}
