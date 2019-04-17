package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	helper "github.com/wecatch/devops-ui/helpers/devops"
	"github.com/wecatch/devops-ui/logshow"
	"github.com/wecatch/devops-ui/resources"
	"github.com/wecatch/devops-ui/utils/format"
	"github.com/wecatch/devops-ui/utils/log"
)

var logger = log.MakeLogger()

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options

// InitRoutes init domain app route
func InitRoutes(router *mux.Router) *mux.Router {
	logger = log.Logger("apps.app")
	appRouter := router.PathPrefix("/v1/app").Subrouter()
	appRouter.Use(format.FormatResponseMiddleware)
	appRouter.HandleFunc("/list", getApps).Methods("GET")
	appRouter.HandleFunc("/tag", getTags).Methods("GET")
	appRouter.HandleFunc("/tag", createTag).Methods("POST")
	appRouter.HandleFunc("/create", createApp).Methods("POST")
	appRouter.HandleFunc("/{id:[0-9]+}", oneApp).Methods("GET")
	appRouter.HandleFunc("/deploy", oneAppDeploy).Methods("GET")
	appRouter.HandleFunc("/deploy", newDeploy).Methods("POST")
	appRouter.HandleFunc("/deploy/{id:[0-9]+}", deployDetail).Methods("GET")
	appRouter.HandleFunc("/deploy/poll", pollDeploy).Methods("POST")
	appRouter.HandleFunc("/deploy/poll/update", updateDeployStatus).Methods("PUT")
	appRouter.HandleFunc("/deploy/poll/log", echoDeployLog).Methods("GET")
	appRouter.HandleFunc("/{id:[0-9]+}", updateApp).Methods("PUT")
	appRouter.HandleFunc("/{id:[0-9]+}", deleteApp).Methods("DELETE")
	return router
}

// getApps service list handle
func getApps(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	var page, limit = 1, 1000
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

	ret := helper.QueryApp(page, limit, q.Get("name"), q.Get("url"))
	jsonData, _ := json.Marshal(ret)
	w.Write(jsonData)
}

// createApp handle
func createApp(w http.ResponseWriter, r *http.Request) {
	dataSource := resources.AppForm{}
	err := json.NewDecoder(r.Body).Decode(&dataSource)
	if err != nil {
		logger.Warn(err)
	}
	helper.CreateApp(dataSource)
}

// updateApp handle
func updateApp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dataSource := resources.AppForm{}
	err := json.NewDecoder(r.Body).Decode(&dataSource)
	if err != nil {
		logger.Warn(err)
	}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		logger.Warn("convert id to fails")
	}
	helper.UpdateApp(id, dataSource)
}

func oneApp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		logger.Warn("convert id to fails")
	}
	ret := helper.QueryOneApp(id, "", "")
	jsonData, _ := json.Marshal(ret)
	w.Write(jsonData)
}

// deleteApp handle
func deleteApp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		logger.Println("convert id to fails")
	}
	helper.DeleteApp(id)
}

// oneAppDeploy handle
func oneAppDeploy(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	var page, limit = 1, 10
	id, err := strconv.Atoi(q.Get("app_id"))
	if err != nil {
		logger.Println("convert app_id to fails")
	}
	if q.Get("page") != "" {
		page, err = strconv.Atoi(q.Get("page"))
		if err != nil {
			logger.Println("page parse error")
		}
	}

	if q.Get("limit") != "" {
		limit, err = strconv.Atoi(q.Get("limit"))
		if err != nil {
			logger.Println("limit parse error")
		}
	}

	ret := helper.QueryAppDeploy(id, page, limit)
	jsonData, _ := json.Marshal(ret)
	w.Write(jsonData)
}

// getTags for app group
func getTags(w http.ResponseWriter, r *http.Request) {
	ret := helper.QueryTag(1, 0)
	jsonData, _ := json.Marshal(ret)
	w.Write(jsonData)
}

func createTag(w http.ResponseWriter, r *http.Request) {
	dataSource := resources.TagForm{}
	err := json.NewDecoder(r.Body).Decode(&dataSource)
	if err != nil {
		logger.Warn(err)
	}
	if dataSource.Name == "" || dataSource.Kind == "" {
		logger.Warn("name or kind not be empty")
		return
	}
	helper.CreateTag(dataSource)
}

// deployDetail handle
func deployDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		logger.Warn("convert deploy id to fails")
	}

	ret := helper.QueryOneDeploy(id)
	jsonData, _ := json.Marshal(ret)
	w.Write(jsonData)
}

// newDeploy handle
func newDeploy(w http.ResponseWriter, r *http.Request) {
	dataSource := resources.DeployForm{}
	err := json.NewDecoder(r.Body).Decode(&dataSource)
	if err != nil {
		logger.Warn(err)
	}
	if dataSource.AppID == 0 {
		logger.Warn("app_id not be empty")
		return
	}
	ret := helper.CreateDeploy(dataSource)
	jsonData, _ := json.Marshal(ret)
	w.Write(jsonData)
}

func pollDeploy(w http.ResponseWriter, r *http.Request) {
	ret := helper.QueryNewDeploy()
	jsonData, _ := json.Marshal(ret)
	w.Write(jsonData)
}

func updateDeployStatus(w http.ResponseWriter, r *http.Request) {
	dataSource := resources.DeployForm{}
	err := json.NewDecoder(r.Body).Decode(&dataSource)
	if err != nil {
		logger.Warn(err)
	}
	helper.UpdateDeployStatus(dataSource.ID, dataSource.Status)

}

func echoDeployLog(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	deployID, err := strconv.Atoi(q.Get("deploy_id"))
	if err != nil {
		logger.Warn("deploy_id is not valid")
	}
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error("upgrade:", err)
		return
	}
	defer c.Close()
	key := fmt.Sprintf("%d", deployID)
	for {
		value, found := logshow.LogC.Get(key)
		if !found {
			time.Sleep(time.Second * 2)
			continue
		}
		lq := value.(*logshow.LogQueue)
		message := bytes.Join(lq.Pop(), []byte(","))
		//无消息需要传递
		if len(message) == 0 {
			time.Sleep(time.Second * 2)
			continue
		}
		message = bytes.Join([][]byte{[]byte("["), message, []byte("]")}, []byte(""))
		err = c.WriteMessage(1, message)
		if err != nil {
			logger.Warn("write:", err)
			break
		}
		logger.Debug(string(message))
	}
}
