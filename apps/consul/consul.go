package consul

import (
	"encoding/json"
	"net/http"
	"strings"
	"fmt"

	"github.com/gorilla/mux"
	consulApi "github.com/hashicorp/consul/api"
	helper "github.com/wecatch/devops-ui/helpers/devops"
	"github.com/wecatch/devops-ui/resources"
	consulUtil "github.com/wecatch/devops-ui/utils/consul"
	"github.com/wecatch/devops-ui/utils/format"
	"github.com/wecatch/devops-ui/utils/log"
)

var logger = log.MakeLogger()

// InitRoutes init domain app route
func InitRoutes(router *mux.Router) *mux.Router {
	logger = log.Logger("apps.consul")
	appRouter := router.PathPrefix("/v1/consul").Subrouter()
	appRouter.Use(format.FormatResponseMiddleware)
	appRouter.HandleFunc("/service/all", getAllService).Methods("GET")
	appRouter.HandleFunc("/service/list", getService).Methods("GET")
	appRouter.HandleFunc("/kv/keys", getKeys).Methods("GET")
	appRouter.HandleFunc("/app/register", registerApp).Methods("PUT")
	appRouter.HandleFunc("/app/unregister", unregisterApp).Methods("PUT")
	return router
}

func getAllService(w http.ResponseWriter, r *http.Request) {
	catalog := consulUtil.ConsulClient.Catalog()
	services, _, err := catalog.Services(nil)
	if err != nil {
		logger.Warn(err)
	}

	jsonData, _ := json.Marshal(services)
	w.Write(jsonData)
}

// consul service list
func getService(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	tag := r.URL.Query().Get("tag")
	if name == "" && tag == "" {
		return
	}

	catalog := consulUtil.ConsulClient.Catalog()
	services, _, err := catalog.Service(name, tag, nil)
	if err != nil {
		logger.Warn(err)
	}

	list := make([]map[string]interface{}, 0)
	for _, s := range services {
		ret := make(map[string]interface{})
		ret["service_id"] = s.ServiceID
		ret["service_name"] = s.ServiceName
		ret["service_address"] = s.ServiceAddress
		ret["service_port"] = s.ServicePort
		ret["service_tags"] = s.ServiceTags
		ret["datacenter"] = s.Datacenter
		list = append(list, ret)
	}

	jsonData, _ := json.Marshal(list)
	w.Write(jsonData)
}

// consul kv keys
func getKeys(w http.ResponseWriter, r *http.Request) {
	dc := r.URL.Query().Get("dc")
	prefix := r.URL.Query().Get("prefix")
	separator := r.URL.Query().Get("separator")
	if dc == "" {
		dc = "dc1"
	}
	options := consulApi.QueryOptions{
		Datacenter: dc,
	}
	kv := consulUtil.ConsulClient.KV()
	keys, _, err := kv.Keys(prefix, separator, &options)
	if err != nil {
		logger.Warn(err)
	}

	jsonData, _ := json.Marshal(keys)
	w.Write(jsonData)
}

// consul register service
func registerApp(w http.ResponseWriter, r *http.Request) {
	form := resources.ConsulServiceForm{}
	err := json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		logger.Warn(err)
	}
	agent := consulUtil.ConsulClient.Agent()
	service := consulApi.AgentServiceRegistration{
		ID:      strings.Join([]string{form.Name, form.HostID}, "-"),
		Name:    form.Name,
		Tags:    []string{form.Tag},
		Port:    form.Port,
		Address: form.PrivateIP,
	}
	agent.ServiceRegister(&service)
	helper.UpdateComputerRole(form.AppID, form.HostID, 1, form.Name)
}

//unregiser app
func unregisterApp(w http.ResponseWriter, r *http.Request) {
	form := resources.ConsulServiceForm{}
	err := json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		logger.Warn(err)
	}
	agent := consulUtil.ConsulClient.Agent()
	if form.ServiceID == "" {
		agent.ServiceDeregister(strings.Join([]string{form.Name, form.HostID}, "-"))
	} else {
		fmt.Println(form.ServiceID)
		agent.ServiceDeregister(form.ServiceID)
	}

	helper.UpdateComputerRole(form.AppID, form.HostID, 0, form.Name)
}
