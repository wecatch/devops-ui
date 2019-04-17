package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	helper "github.com/wecatch/devops-ui/helpers/devops"
	"github.com/wecatch/devops-ui/resources"
	"github.com/wecatch/devops-ui/utils/format"
	"github.com/wecatch/devops-ui/utils/log"
)

var logger = log.MakeLogger()

// InitRoutes init domain app route
func InitRoutes(router *mux.Router) *mux.Router {
	logger = log.Logger("apps.service")
	appRouter := router.PathPrefix("/v1/service").Subrouter()
	appRouter.Use(format.FormatResponseMiddleware)
	appRouter.HandleFunc("/list", getServices).Methods("GET")
	appRouter.HandleFunc("/create", createService).Methods("POST")
	appRouter.HandleFunc("/{id:[0-9]+}", updateService).Methods("PUT")
	appRouter.HandleFunc("/{id:[0-9]+}", deleteService).Methods("DELETE")
	return router
}

// getServices service list handle
func getServices(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	var page, limit = 1, 10
	var err error
	if q.Get("page") != "" {
		page, err = strconv.Atoi(q.Get("page"))
		if err != nil {
			logger.Warn("page parse error")
		}
	}

	if q.Get("limit") != "" {
		limit, err = strconv.Atoi(q.Get("limit"))
		if err != nil {
			logger.Warn("limit parse error")
		}
	}

	ret := helper.QueryService(page, limit, q.Get("name"), q.Get("url"))
	jsonData, _ := json.Marshal(ret)
	w.Write(jsonData)
}

// createService handle
func createService(w http.ResponseWriter, r *http.Request) {
	dataSource := resources.ServiceForm{}
	err := json.NewDecoder(r.Body).Decode(&dataSource)
	if err != nil {
		logger.Warn(err)
	}
	helper.CreateService(dataSource)
}

// updateService handle
func updateService(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dataSource := resources.ServiceForm{}
	err := json.NewDecoder(r.Body).Decode(&dataSource)
	if err != nil {
		logger.Warn(err)
	}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		logger.Warn("convert id to fails")
	}
	helper.UpdateService(id, dataSource)
}

// deleteService handle
func deleteService(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		logger.Warn("convert id to fails")
	}
	helper.DeleteService(id)
}
