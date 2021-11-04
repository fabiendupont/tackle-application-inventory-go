package main

import (
	apiv1 "github.com/fabiendupont/tackle-hub/controllers/v1"
//	apiv2 "github.com/fabiendupont/tackle-hub/controllers/v2"
	"github.com/fabiendupont/tackle-hub/database"
	_ "github.com/fabiendupont/tackle-hub/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Tackle Hub - OpenAPI
// @version 3.0
// @description Tackle Hub provides an API to manipulate core Tackle objects
// @termsOfService http://swagger.io/terms/

// @contact.name Konveyor Tackle
// @contact.url https://konveyor.io
// @contact.email contact@konveyor.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	// Initialize database
	database.Setup()

	// Initialize Gin Gonic
	r := gin.Default()

	// Setup swagger
	swaggerUrl := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, swaggerUrl))

	// Register API v1 routes
	r.GET("/application-inventory/application", apiv1.GetApplications)
	r.GET("/application-inventory/application/:id", apiv1.GetApplication)
	r.POST("/application-inventory/application", apiv1.CreateApplication)
	r.PUT("/application-inventory/application/:id", apiv1.UpdateApplication)
	r.DELETE("/application-inventory/application/:id", apiv1.DeleteApplication)

	r.GET("/application-inventory/binary-repository", apiv1.GetBinaryRepositories)
	r.GET("/application-inventory/binary-repository/:id", apiv1.GetBinaryRepository)
	r.POST("/application-inventory/binary-repository", apiv1.CreateBinaryRepository)
	r.PUT("/application-inventory/binary-repository/:id", apiv1.UpdateBinaryRepository)
	r.DELETE("/application-inventory/binary-repository/:id", apiv1.DeleteBinaryRepository)

	r.GET("/application-inventory/review", apiv1.GetReviews)
	r.GET("/application-inventory/review/:id", apiv1.GetReview)
	r.POST("/application-inventory/review", apiv1.CreateReview)
	r.PUT("/application-inventory/review/:id", apiv1.UpdateReview)
	r.DELETE("/application-inventory/review/:id", apiv1.DeleteReview)

	r.GET("/application-inventory/source-repository", apiv1.GetSourceRepositories)
	r.GET("/application-inventory/source-repository/:id", apiv1.GetSourceRepository)
	r.POST("/application-inventory/source-repository", apiv1.CreateSourceRepository)
	r.PUT("/application-inventory/source-repository/:id", apiv1.UpdateSourceRepository)
	r.DELETE("/application-inventory/source-repository/:id", apiv1.DeleteSourceRepository)

	r.GET("/controls/business-service", apiv1.GetBusinessServices)
	r.GET("/controls/business-service/:id", apiv1.GetBusinessService)
	r.POST("/controls/business-service", apiv1.CreateBusinessService)
	r.PUT("/controls/business-service/:id", apiv1.UpdateBusinessService)
	r.DELETE("/controls/business-service/:id", apiv1.DeleteBusinessService)

	r.GET("/controls/job-function", apiv1.GetJobFunctions)
	r.GET("/controls/job-function/:id", apiv1.GetJobFunctions)
	r.POST("/controls/job-function", apiv1.CreateJobFunction)
	r.PUT("/controls/job-function/:id", apiv1.UpdateJobFunction)
	r.DELETE("/controls/job-function/:id", apiv1.DeleteJobFunction)

	r.GET("/controls/stakeholder", apiv1.GetStakeholders)
	r.GET("/controls/stakeholder/:id", apiv1.GetStakeholder)
	r.POST("/controls/stakeholder", apiv1.CreateStakeholder)
	r.PUT("/controls/stakeholder/:id", apiv1.UpdateStakeholder)
	r.DELETE("/controls/stakeholder/:id", apiv1.DeleteStakeholder)

	r.GET("/controls/stakeholder-group", apiv1.GetStakeholderGroups)
	r.GET("/controls/stakeholder-group/:id", apiv1.GetStakeholderGroup)
	r.POST("/controls/stakeholder-group", apiv1.CreateStakeholderGroup)
	r.PUT("/controls/stakeholder-group/:id", apiv1.UpdateStakeholderGroup)
	r.DELETE("/controls/stakeholder-group/:id", apiv1.DeleteStakeholderGroup)

	r.GET("/controls/tag-type", apiv1.GetTagTypes)
	r.GET("/controls/tag-type/:id", apiv1.GetTagType)
	r.POST("/controls/tag-type", apiv1.CreateTagType)
	r.PUT("/controls/tag-type/:id", apiv1.UpdateTagType)
	r.DELETE("/controls/tag-type/:id", apiv1.DeleteTagType)

	r.GET("/controls/tag", apiv1.GetTags)
	r.GET("/controls/tag/:id", apiv1.GetTag)
	r.POST("/controls/tag", apiv1.CreateTag)
	r.PUT("/controls/tag/:id", apiv1.UpdateTag)
	r.DELETE("/controls/tag/:id", apiv1.DeleteTag)

	// Register API v2 routes
	/*
	r.GET("/v2/applications", apiv2.GetApplications)
	r.GET("/v2/applications/:id", apiv2.GetApplication)
	r.POST("/v2/applications", apiv2.CreateApplication)
	r.PUT("/v2/applications/:id", apiv2.UpdateApplication)
	r.DELETE("/v2/applications/:id", apiv2.DeleteApplication)

	r.GET("/v2/binary_repositories", apiv2.GetBinaryRepositories)
	r.GET("/v2/binary_repositories/:id", apiv2.GetBinaryRepository)
	r.POST("/v2/binary_repositories", apiv2.CreateBinaryRepository)
	r.PUT("/v2/binary_repositories/:id", apiv2.UpdateBinaryRepository)
	r.DELETE("/v2/binary_repositories/:id", apiv2.DeleteBinaryRepository)

	r.GET("/v2/business_services", apiv2.GetBusinessServices)
	r.GET("/v2/business_services/:id", apiv2.GetBusinessService)
	r.POST("/v2/business_services", apiv2.CreateBusinessService)
	r.PUT("/v2/business_services/:id", apiv2.UpdateBusinessService)
	r.DELETE("/v2/business_services/:id", apiv2.DeleteBusinessService)

	r.GET("/v2/groups", apiv2.GetGroups)
	r.GET("/v2/groups/:id", apiv2.GetGroup)
	r.POST("/v2/groups", apiv2.CreateGroup)
	r.PUT("/v2/groups/:id", apiv2.UpdateGroup)
	r.DELETE("/v2/groups/:id", apiv2.DeleteGroup)

	r.GET("/v2/role_bindings", apiv2.GetRoleBindings)
	r.GET("/v2/role_bindings/:id", apiv2.GetRoleBinding)
	r.POST("/v2/role_bindings", apiv2.CreateRoleBinding)
	r.PUT("/v2/role_bindings/:id", apiv2.UpdateRoleBinding)
	r.DELETE("/v2/role_bindings/:id", apiv2.DeleteRoleBinding)

	r.GET("/v2/roles", apiv2.GetRoles)
	r.GET("/v2/roles/:id", apiv2.GetRole)
	r.POST("/v2/roles", apiv2.CreateRole)
	r.PUT("/v2/roles/:id", apiv2.UpdateRole)
	r.DELETE("/v2/roles/:id", apiv2.DeleteRole)

	r.GET("/v2/source_repositories", apiv2.GetSourceRepositories)
	r.GET("/v2/source_repositories/:id", apiv2.GetSourceRepository)
	r.POST("/v2/source_repositories", apiv2.CreateSourceRepository)
	r.PUT("/v2/source_repositories/:id", apiv2.UpdateSourceRepository)
	r.DELETE("/v2/source_repositories/:id", apiv2.DeleteSourceRepository)

	r.GET("/v2/tag_types", apiv2.GetTagTypes)
	r.GET("/v2/tag_types/:id", apiv2.GetTagType)
	r.POST("/v2/tag_types", apiv2.CreateTagType)
	r.PUT("/v2/tag_types/:id", apiv2.UpdateTagType)
	r.DELETE("/v2/tag_types/:id", apiv2.DeleteTagType)

	r.GET("/v2/tags", apiv2.GetTags)
	r.GET("/v2/tags/:id", apiv2.GetTag)
	r.POST("/v2/tags", apiv2.CreateTag)
	r.PUT("/v2/tags/:id", apiv2.UpdateTag)
	r.DELETE("/v2/tags/:id", apiv2.DeleteTag)

	r.GET("/v2/users", apiv2.GetUsers)
	r.GET("/v2/users/:id", apiv2.GetUser)
	r.POST("/v2/users", apiv2.CreateUser)
	r.PUT("/v2/users/:id", apiv2.UpdateUser)
	r.DELETE("/v2/users/:id", apiv2.DeleteUser)
	*/

	r.Run()
}
