package routes

import (
	v1 "GinProject/api/v1"
	_ "GinProject/docs"
	"GinProject/middleware"
	"GinProject/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	//r := gin.Default()
	r := gin.New()
	r.Use(gin.Recovery())
	//r.Use(middleware.ZapLogger())
	r.Use(middleware.Logger())
	r.Use(middleware.Cors())

	router := r.Group("api/v1")
	{
		//用户模块的路由
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.POST("login", v1.Login)
		//分类模块的路由
		router.GET("category", v1.GetCategory)
		//文章模块的路由
		router.GET("article", v1.GetArticle)
		router.GET("article/info/:id", v1.GetArticleInfo)
		router.GET("article/catelist/:cid", v1.GetCateArt)
	}

	auth := r.Group("api/v1")
	//需要鉴权的接口
	auth.Use(middleware.JwtToken())
	{
		//router.GET("hello", func(ctx *gin.Context) {
		//	ctx.JSON(http.StatusOK, gin.H{
		//		"MSG": "OK",
		//	})
		//})
		//用户模块的路由
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		//分类模块的路由
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCategory)
		auth.DELETE("category/:id", v1.DeleteCategory)
		//文章模块的路由
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArticle)
		auth.DELETE("article/:id", v1.DeleteArticle)
		//上传文件
		auth.POST("upload", v1.Upload)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := r.Run(utils.HttpPort)
	if err != nil {
		fmt.Println("启动项目出错，错误提示为：", err)
		return
	}
}
