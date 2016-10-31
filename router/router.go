package router

import (
	"github.com/go-martini/martini"

	"code.oa.com/sean/api-doc/controllers"
)

func InitRouter(m *martini.ClassicMartini) {
	var DocsController controllers.DocsController
	var QueryController controllers.QueryController

	m.Post("/", controllers.Index)

	m.Get("/doc", DocsController.Index)

	//慢查询列表
	m.Get("/query/slow", QueryController.SlowList)

	//js跨域请求
	m.Post("/(beta|gamma)_[a-z]+/[a-z_A-Z0-9/]+", DocsController.RequestRemote)

	m.Get("/doc/detail", DocsController.Detail)

	m.Group("/test", func(r martini.Router) {
		r.Get("/", func() string {
			return "get"
		})

		r.Post("/", func() string {
			return "create"
		})
	})
}
