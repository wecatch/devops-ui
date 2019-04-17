package domain

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
	logger = log.Logger("apps.domain")
	appRouter := router.PathPrefix("/v1/domain").Subrouter()
	appRouter.Use(format.FormatResponseMiddleware)
	appRouter.HandleFunc("/list", getDomains)
	appRouter.HandleFunc("/create", createDomain).Methods("POST")
	appRouter.HandleFunc("/{id:[0-9]+}", updateDomain).Methods("PUT")
	appRouter.HandleFunc("/{id:[0-9]+}", deleteDomain).Methods("DELETE")
	return router
}

// getDomains domain list handle
func getDomains(w http.ResponseWriter, r *http.Request) {
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

	domains := helper.QueryDomain(page, limit, q.Get("name"), q.Get("url"))
	// var dataSource []resources.DomainForm
	// for _, v := range domains {
	// 	vv := v
	// 	dest := resources.DomainForm{}
	// 	utils.CopyStruct(&vv, &dest)
	// 	dataSource = append(dataSource, dest)
	// }
	jsonData, _ := json.Marshal(domains)
	w.Write(jsonData)
}

// createDomain handle
func createDomain(w http.ResponseWriter, r *http.Request) {
	dataSource := resources.DomainForm{}
	err := json.NewDecoder(r.Body).Decode(&dataSource)
	if err != nil {
		logger.Warn(err)
	}
	helper.CreateDomain(dataSource.Name, dataSource.Host, dataSource.Private, dataSource.IP)
}

// updateDomain handle
func updateDomain(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dataSource := resources.DomainForm{}
	err := json.NewDecoder(r.Body).Decode(&dataSource)
	if err != nil {
		logger.Warn(err)
	}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		logger.Warn("convert id to fails")
	}
	helper.UpdateDomain(id, dataSource.Name, dataSource.Host, dataSource.Private)
}

// deleteDomain handle
func deleteDomain(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		logger.Warn("convert id to fails")
	}
	helper.DeleteDomain(id)
}
