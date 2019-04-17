package computer

import (
	"encoding/json"
	// "log"
	"fmt"
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
	logger = log.Logger("apps.computer")
	appRouter := router.PathPrefix("/v1/computer").Subrouter()
	appRouter.Use(format.FormatResponseMiddleware)
	appRouter.HandleFunc("/list", getComputers).Methods("GET")
	appRouter.HandleFunc("/app", getAppComputers).Methods("GET")
	appRouter.HandleFunc("/app/group", getAppGroupComputers).Methods("GET")
	appRouter.HandleFunc("/create", createComputer).Methods("POST")
	appRouter.HandleFunc("/role", createRole).Methods("POST")
	appRouter.HandleFunc("/role", deleteRole).Methods("DELETE")
	appRouter.HandleFunc("/{id:[0-9]+}", updateComputer).Methods("PUT")
	appRouter.HandleFunc("/{id:[0-9]+}", deleteComputer).Methods("DELETE")
	appRouter.HandleFunc("/disk/create", createDisk).Methods("POST")
	appRouter.HandleFunc("/disk/{id:[0-9]+", updateDisk).Methods("PUT")
	appRouter.HandleFunc("/disk/{id:[0-9]+", deleteDisk).Methods("DELETE")
	return router
}

func getAppGroupComputers(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	var page, limit = 1, 1000
	var err error
	// var tag, name = q.Get("tag"), q.Get("name")
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

	appRet := helper.QueryApp(page, limit, "", "")
	var appIds []int
	for _, app := range appRet {
		appIds = append(appIds, app.ID)
	}
	fmt.Println(appIds)

	ret := helper.QueryAppComputer(appIds, 0, 10000, "")
	jsonData, _ := json.Marshal(ret)
	w.Write(jsonData)
}

// getAppComputers computer list handle
func getAppComputers(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	var page, limit, appID = 1, 50, 0
	var err error
	// var tag, name = q.Get("tag"), q.Get("name")
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

	if q.Get("app_id") != "" {
		appID, err = strconv.Atoi(q.Get("app_id"))
		if err != nil {
			logger.Warn("app_id parse error")
		}
	}

	ret := helper.QueryAppComputer(appID, page, limit, q.Get("app_name"))
	// var dataSource []resources.ComputerForm
	// for _, v := range ret {
	// 	vv := v
	// 	dest := resources.ComputerForm{}
	// 	utils.CopyStruct(&vv, &dest)
	// 	dataSource = append(dataSource, dest)
	// }
	jsonData, _ := json.Marshal(ret)
	w.Write(jsonData)
}

// getComputers computer list handle
func getComputers(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	var page, limit = 1, 50
	var err error
	var tag, name, searchValue = q.Get("tag"), q.Get("name"), q.Get("search_value")
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

	ret := helper.QueryComputer(tag, name, searchValue, page, limit)
	jsonData, _ := json.Marshal(ret)
	w.Write(jsonData)
}

// createComputer handle
func createComputer(w http.ResponseWriter, r *http.Request) {
	dataSource := resources.ComputerForm{}
	err := json.NewDecoder(r.Body).Decode(&dataSource)
	if err != nil {
		logger.Warn(err)
	}
	helper.CreateComputer(dataSource)
}

// updateComputer handle
func updateComputer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dataSource := resources.ComputerForm{}
	err := json.NewDecoder(r.Body).Decode(&dataSource)
	if err != nil {
		logger.Warn(err)
	}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		logger.Warn("convert id to fails")
	}
	helper.UpdateComputer(id, dataSource)
}

// deleteComputer handle
func deleteComputer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		logger.Warn("convert id to fails")
	}
	helper.DeleteComputer(id)
}

// createDisk handle
func createDisk(w http.ResponseWriter, r *http.Request) {
	dataSource := resources.DiskForm{}
	err := json.NewDecoder(r.Body).Decode(&dataSource)
	if err != nil {
		logger.Warn(err)
	}
	helper.CreateDisk(dataSource)
}

// updateDisk handle
func updateDisk(w http.ResponseWriter, r *http.Request) {
	dataSource := resources.DiskForm{}
	vars := mux.Vars(r)
	err := json.NewDecoder(r.Body).Decode(&dataSource)
	if err != nil {
		logger.Warn(err)
	}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		logger.Warn(err)
	}
	helper.UpdateDisk(id, dataSource)
}

// deleteDisk handle
func deleteDisk(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		logger.Warn(err)
	}
	if err != nil {
		logger.Warn(err)
	}
	helper.DeleteDisk(id)
}

func createRole(w http.ResponseWriter, r *http.Request) {
	dataSource := resources.ComputerRoleForm{}
	err := json.NewDecoder(r.Body).Decode(&dataSource)
	if err != nil {
		logger.Warn(err)
	}
	helper.CreateComputerRole(dataSource)
}

func deleteRole(w http.ResponseWriter, r *http.Request) {
	dataSource := resources.ComputerRoleForm{}
	err := json.NewDecoder(r.Body).Decode(&dataSource)
	if err != nil {
		logger.Warn(err)
	}
	helper.DeleteComputerRole(dataSource)
}
