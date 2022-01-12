package main

import (
	"GinProject/model"
	"GinProject/routes"
)

// @title           Gin Blog API v1
// @version         1.0
// @description     Gin博客接口文档
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  942801422@qq.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:2333
// @BasePath  /api/v1
func main() {
	model.InitDB()
	routes.InitRouter()
}
