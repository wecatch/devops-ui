package apps

import (
	"github.com/gorilla/mux"
	"github.com/wecatch/devops-ui/apps/app"
	"github.com/wecatch/devops-ui/apps/computer"
	"github.com/wecatch/devops-ui/apps/consul"
	"github.com/wecatch/devops-ui/apps/domain"
	"github.com/wecatch/devops-ui/apps/gitlab"
	"github.com/wecatch/devops-ui/apps/home"
	"github.com/wecatch/devops-ui/apps/service"
	"github.com/wecatch/devops-ui/apps/cloud"
)

// InitRouter init all app route
func InitRouter() *mux.Router {
	router := mux.NewRouter()
	router = domain.InitRoutes(router)
	router = computer.InitRoutes(router)
	router = service.InitRoutes(router)
	router = consul.InitRoutes(router)
	router = app.InitRoutes(router)
	router = gitlab.InitRoutes(router)
	router = cloud.InitRoutes(router)
	// mux 使用 Use 来注册 middleware 这里 route 的注册顺序是该 router 的前面不能有使用未使用 middleware 的 router
	// 如果 home 在 gitlab 前面，则 middleware 不起作用
	router = home.InitRoutes(router)
	return router
}
