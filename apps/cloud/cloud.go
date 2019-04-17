package cloud

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	helper "github.com/wecatch/devops-ui/helpers/devops"
	"github.com/wecatch/devops-ui/resources"
	"github.com/wecatch/devops-ui/utils/log"
	"github.com/wecatch/devops-ui/utils/ucloud"
	"github.com/wecatch/devops-ui/utils/format"
)

var logger = log.MakeLogger()

// InitRoutes init domain app route
func InitRoutes(router *mux.Router) *mux.Router {
	logger = log.Logger("apps.app")
	appRouter := router.PathPrefix("/v1/cloud").Subrouter()
	appRouter.Use(format.FormatResponseMiddleware)
	appRouter.HandleFunc("/region", getRegion).Methods("GET")
	appRouter.HandleFunc("/region", createRegion).Methods("POST")
	appRouter.HandleFunc("/sync", syncHost).Methods("POST")
	return router
}

// 获取可用区
func getRegion(w http.ResponseWriter, r *http.Request) {
	ret := helper.QueryRegion()
	jsonData, _ := json.Marshal(ret)
	w.Write(jsonData)
}

// 创建可用区
func createRegion(w http.ResponseWriter, r *http.Request) {
	dataSource := resources.CloudAvailableRegionForm{}
	err := json.NewDecoder(r.Body).Decode(&dataSource)
	if err != nil {
		logger.Warn(err)
	}
	helper.CreateRegion(dataSource)
}

// 同步主机
func syncHost(w http.ResponseWriter, r *http.Request) {
	ucloud.UcloundClient.SyncAllHosts(helper.QueryRegion())
}
