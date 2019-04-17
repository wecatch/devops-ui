package gitlab

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/wecatch/devops-ui/utils/format"
	"github.com/wecatch/devops-ui/utils/gitlab"
	"github.com/wecatch/devops-ui/utils/log"
)

var logger = log.MakeLogger()

// InitRoutes init domain app route
func InitRoutes(router *mux.Router) *mux.Router {
	logger = log.Logger("apps.gitlab")
	appRouter := router.PathPrefix("/v1/gitlab").Subrouter()
	appRouter.Use(format.FormatResponseMiddleware)
	appRouter.HandleFunc("/projects", getProjects).Methods("GET")
	appRouter.HandleFunc("/tags", getTags).Methods("GET")
	appRouter.HandleFunc("/groups", getGroups).Methods("GET")
	appRouter.HandleFunc("/commits", getCommits).Methods("GET")

	return router
}

// getProjects domain list handle
func getProjects(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	var gid = 0
	var err error
	if q.Get("gid") != "" {
		gid, err = strconv.Atoi(q.Get("gid"))
		if err != nil {
			logger.Warn("convert group id to int fails")
		}
	}
	var jsonData []byte
	if gid > 0 {
		jsonData, err = json.Marshal(gitlab.GitLabClient.ListGitlabGroupProjects(gid))
	} else {
		jsonData, err = json.Marshal(gitlab.GitLabClient.ListGitlabProjects())
	}
	w.Write(jsonData)
}

func getTags(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	var pid = 0
	var err error
	if q.Get("pid") != "" {
		pid, err = strconv.Atoi(q.Get("pid"))
		if err != nil {
			logger.Warn("convert id to fails")
		}
	}
	jsonData, _ := json.Marshal(gitlab.GitLabClient.ListGitlabProjectTags(pid))
	w.Write(jsonData)
}

// getGroups group list handle
func getGroups(w http.ResponseWriter, r *http.Request) {
	jsonData, _ := json.Marshal(gitlab.GitLabClient.ListGitlabGroups())
	w.Write(jsonData)
}

// getCommits commit list handle
func getCommits(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	var pid = 0
	var err error
	if q.Get("pid") != "" {
		pid, err = strconv.Atoi(q.Get("pid"))
		if err != nil {
			logger.Warn("convert id to fails")
		}
	}
	jsonData, _ := json.Marshal(gitlab.GitLabClient.ListGitlabProjectCommits(pid))
	w.Write(jsonData)
}
