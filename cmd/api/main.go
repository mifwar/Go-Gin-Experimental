package main

import (
	"github.com/gin-gonic/gin"
	admin "online-course.mifwar.com/internal/admin/injector"
	oauth "online-course.mifwar.com/internal/oauth/injector"
	productCategory "online-course.mifwar.com/internal/product_category/injector"
	profile "online-course.mifwar.com/internal/profile/injector"
	register "online-course.mifwar.com/internal/register/injector"
	mysql "online-course.mifwar.com/pkg/db/mysql"
)

func main() {

	db := mysql.DB()

	r := gin.Default()

	register.InitializedService(db).Route(&r.RouterGroup)
	oauth.InitializedService(db).Route(&r.RouterGroup)
	profile.InitializedService(db).Route(&r.RouterGroup)
	admin.InitializedService(db).Route(&r.RouterGroup)
	productCategory.InitializedService(db).Route(&r.RouterGroup)

	r.Run()
}
