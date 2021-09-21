package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	{
		// User模块的路由接口
		routeUser(router)
		// 分类模块的路由接口
		routeCategory(router)
		// 文章模块的路由接口
		routeArticle(router)
	}

	r.Run(utils.HttpPort)
}

// 用户路由
func routeUser(router *gin.RouterGroup) {
	router.POST("user/add", v1.AddUser)
	router.GET("users", v1.GetUsers)
	router.PUT("user/:id", v1.EditUser)
	router.DELETE("user/:id", v1.DeleteUser)
}

// 分类路由
func routeCategory(router *gin.RouterGroup) {
	router.POST("category/add", v1.AddCategory)
	router.GET("category", v1.GetCategory)
	router.PUT("category/:id", v1.EditCategory)
	router.DELETE("category/:id", v1.DeleteCategory)
}

// 文章路由
func routeArticle(router *gin.RouterGroup) {
	router.POST("article/add", v1.AddArticle)
	router.GET("article", v1.GetArticle)
	router.GET("article/catelist", v1.GetCategoryArticle)
	router.GET("article/info/:id", v1.GetArticleInfo)
	router.PUT("article/:id", v1.EditArticle)
	router.DELETE("article/:id", v1.DeleteArticle)
}
