package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"testing"

	model "github.com/wecatch/devops-ui/models/devops"
	"github.com/wecatch/devops-ui/resources"
	"github.com/wecatch/devops-ui/utils/http"
)

var appClient = httputil.RequestClient{
	HOST:        "http://127.0.0.1:8000/v1",
	Prefix:      "/app",
	ContentType: "application/json",
}

func TestAppList(t *testing.T) {
	url := "/list"
	resp := appClient.Get(url)
	if resp == nil {
		t.Fail()
	}
	data, _ := ioutil.ReadAll(resp.Body)
	var v []resources.AppForm
	json.Unmarshal(data, &v)
	for _, d := range v {
		if d.URL == "" || d.Name == "" {
			t.Fail()
		}
	}
}

func TestAppCreate(t *testing.T) {
	form := resources.AppForm{
		App: model.App{
			Name:          "i am name",
			URL:           "i am url",
			Desc:          "desc",
			DeployDir:     "deploy_url",
			RepositoryURL: "repository_url",
			MonitorURL:    "monitor_url",
		},
	}
	data, _ := json.Marshal(form)
	body := bytes.NewBuffer(data)
	resp := appClient.Post("/create", body)
	if resp == nil {
		t.Fail()
	}
	if resp.StatusCode != 200 {
		t.Fail()
	}
}

func TestAppUpdate(t *testing.T) {
	form := resources.AppForm{
		App: model.App{
			Name:          "i am name",
			URL:           "i am url",
			Desc:          "desc",
			DeployDir:     "deploy_url",
			RepositoryURL: "repository_url",
			MonitorURL:    "monitor_url",
		},
	}
	data, _ := json.Marshal(form)
	body := bytes.NewBuffer(data)
	resp := appClient.Put("/1", body)
	if resp == nil {
		t.Fail()
	}
	if resp.StatusCode != 200 {
		t.Fail()
	}
}

func TestAppDelete(t *testing.T) {
	resp := appClient.Delete("/10", io.Reader(nil))
	if resp == nil {
		t.Fail()
	}
	if resp.StatusCode != 200 {
		t.Fail()
	}
}
