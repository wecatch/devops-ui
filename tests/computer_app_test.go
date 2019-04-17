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

var computerClient = httputil.RequestClient{
	HOST:        "http://127.0.0.1:8000/v1",
	Prefix:      "/computer",
	ContentType: "application/json",
}

func TestComputerList(t *testing.T) {
	url := "/list"
	resp := computerClient.Get(url)
	if resp == nil {
		t.Fail()
	}
	data, _ := ioutil.ReadAll(resp.Body)
	var v []resources.ComputerForm
	json.Unmarshal(data, &v)
	for _, d := range v {
		if d.CPU == 0 || d.RAM == 0 {
			t.Fail()
		}
	}
}

func TestComputerCreate(t *testing.T) {
	form := resources.ComputerForm{
		Computer: model.Computer{
			CPU:       2,
			RAM:       4,
			PrivateIP: "10.12.292.1",
			PublicIP:  "10.12.292.1",
		},
	}
	data, _ := json.Marshal(form)
	body := bytes.NewBuffer(data)
	resp := computerClient.Post("/create", body)
	if resp == nil {
		t.Fail()
	}
	if resp.StatusCode != 200 {
		t.Fail()
	}
}

func TestComputerUpdate(t *testing.T) {
	form := resources.ComputerForm{
		Computer: model.Computer{
			CPU:       2,
			RAM:       4,
			PrivateIP: "10.12.292.1",
			PublicIP:  "10.12.292.1",
		},
	}
	data, _ := json.Marshal(form)
	body := bytes.NewBuffer(data)
	resp := computerClient.Put("/4", body)
	if resp == nil {
		t.Fail()
	}
	if resp.StatusCode != 200 {
		t.Fail()
	}
}

func TestComputerDelete(t *testing.T) {
	resp := computerClient.Delete("/1", io.Reader(nil))
	if resp == nil {
		t.Fail()
	}
	if resp.StatusCode != 200 {
		t.Fail()
	}
}
