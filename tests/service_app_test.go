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

var serviceClient = httputil.RequestClient{
	HOST:        "http://127.0.0.1:8000/v1",
	Prefix:      "/service",
	ContentType: "application/json",
}

func TestServiceList(t *testing.T) {
	url := "/list"
	resp := serviceClient.Get(url)
	if resp == nil {
		t.Fail()
	}
	data, _ := ioutil.ReadAll(resp.Body)
	var v []resources.ServiceForm
	json.Unmarshal(data, &v)
	for _, d := range v {
		if d.URL == "" || d.Name == "" {
			t.Fail()
		}
	}
}

func TestServiceCreate(t *testing.T) {
	form := resources.ServiceForm{
		Service: model.Service{
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
	resp := serviceClient.Post("/create", body)
	if resp == nil {
		t.Fail()
	}
	if resp.StatusCode != 200 {
		t.Fail()
	}
}

func TestServiceUpdate(t *testing.T) {
	form := resources.ServiceForm{
		Service: model.Service{
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
	resp := serviceClient.Put("/1", body)
	if resp == nil {
		t.Fail()
	}
	if resp.StatusCode != 200 {
		t.Fail()
	}
}

func TestServiceDelete(t *testing.T) {
	resp := serviceClient.Delete("/10", io.Reader(nil))
	if resp == nil {
		t.Fail()
	}
	if resp.StatusCode != 200 {
		t.Fail()
	}
}
