// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"gofucking/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)

	beego.Router("/", &controllers.IndexController{}, "get:Get;post:Post;*:Index")
	beego.Router("/index", &controllers.IndexController{}, "*:Index")
	beego.Router("/book", &controllers.IndexController{}, "get:Book")

	beego.Router("/info", &controllers.InfoController{}, "get:Get;post:Post")
	beego.Router("/userinfo/get", &controllers.UserInfoController{}, "get:Get")
	beego.Router("/userinfo/insert", &controllers.UserInfoController{}, "post:Insert")
	beego.Router("/userinfo/update", &controllers.UserInfoController{}, "post:Update")
	beego.Router("/userinfo/query", &controllers.UserInfoController{}, "get:Query")
}
