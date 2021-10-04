package main

import (
	"github.com/fabiendupont/tackle-application-inventory-go/controllers"
	"github.com/fabiendupont/tackle-application-inventory-go/database"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.Setup()

	r.GET("/applications", controllers.GetApplications)
	r.GET("/applications/:id", controllers.GetApplication)
	r.POST("/applications", controllers.CreateApplication)
	r.PUT("/applications/:id", controllers.UpdateApplication)
	r.DELETE("/applications/:id", controllers.DeleteApplication)

	r.GET("/business_services", controllers.GetBusinessServices)
	r.GET("/business_services/:id", controllers.GetBusinessService)
	r.POST("/business_services", controllers.CreateBusinessService)
	r.PUT("/business_services/:id", controllers.UpdateBusinessService)
	r.DELETE("/business_services/:id", controllers.DeleteBusinessService)

	r.GET("/groups", controllers.GetGroups)
	r.GET("/groups/:id", controllers.GetGroup)
	r.POST("/groups", controllers.CreateGroup)
	r.PUT("/groups/:id", controllers.UpdateGroup)
	r.DELETE("/groups/:id", controllers.DeleteGroup)

	r.GET("/role_bindings", controllers.GetRoleBindings)
	r.GET("/role_bindings/:id", controllers.GetRoleBinding)
	r.POST("/role_bindings", controllers.CreateRoleBinding)
	r.PUT("/role_bindings/:id", controllers.UpdateRoleBinding)
	r.DELETE("/role_bindings/:id", controllers.DeleteRoleBinding)

	r.GET("/roles", controllers.GetRoles)
	r.GET("/roles/:id", controllers.GetRole)
	r.POST("/roles", controllers.CreateRole)
	r.PUT("/roles/:id", controllers.UpdateRole)
	r.DELETE("/roles/:id", controllers.DeleteRole)

	r.GET("/tag_types", controllers.GetTagTypes)
	r.GET("/tag_types/:id", controllers.GetTagType)
	r.POST("/tag_types", controllers.CreateTagType)
	r.PUT("/tag_types/:id", controllers.UpdateTagType)
	r.DELETE("/tag_types/:id", controllers.DeleteTagType)

	r.GET("/tags", controllers.GetTags)
	r.GET("/tags/:id", controllers.GetTag)
	r.POST("/tags", controllers.CreateTag)
	r.PUT("/tags/:id", controllers.UpdateTag)
	r.DELETE("/tags/:id", controllers.DeleteTag)

	r.GET("/users", controllers.GetUsers)
	r.GET("/users/:id", controllers.GetUser)
	r.POST("/users", controllers.CreateUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	r.Run()
}
