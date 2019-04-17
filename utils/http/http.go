package httputil

import (
	"fmt"
	"io"
	"net/http"
	netUrl "net/url"
	"strings"

	"github.com/wecatch/devops-ui/utils/log"
)

var logger = log.MakeLogger()

func init() {
	logger = log.Logger("utils.http")
}

// RequestClient client for test
type RequestClient struct {
	HOST        string
	ContentType string
	Client      http.Client
	Prefix      string
}

// JoinURL composie url and host
func (client RequestClient) JoinURL(url string) string {
	if strings.Contains(url, "http") || strings.Contains(url, "https") {
		return url
	}

	return strings.Join([]string{client.HOST, client.Prefix, url}, "")
}

//Get client get
func (client RequestClient) Get(url string, params map[string]interface{}) *http.Response {
	url = client.JoinURL(url)
	if params != nil {
		var paramsSlice []string
		for k, value := range params {
			paramsSlice = append(paramsSlice, k+"="+netUrl.QueryEscape(fmt.Sprint(value)))
		}
		url = strings.Join([]string{url, "?", strings.Join(paramsSlice, "&")}, "")
	}
	resp, err := http.Get(url)
	if err != nil {
		logger.Warn(err)
		return nil
	}

	return resp
}

//Post client post
func (client RequestClient) Post(url string, body io.Reader) *http.Response {
	resp, err := http.Post(client.JoinURL(url), client.ContentType, body)
	if err != nil {
		logger.Warn(err)
		return nil
	}

	return resp
}

//Put client pust
func (client RequestClient) Put(url string, body io.Reader) *http.Response {
	req, err := http.NewRequest(http.MethodPut, client.JoinURL(url), body)
	if err != nil {
		logger.Warn(err)
		return nil
	}

	resp, err := client.Client.Do(req)
	if err != nil {
		logger.Warn(err)
		return nil
	}

	return resp

}

//Delete client pust
func (client RequestClient) Delete(url string, body io.Reader) *http.Response {
	req, err := http.NewRequest(http.MethodDelete, client.JoinURL(url), body)
	if err != nil {
		logger.Warn(err)
		return nil
	}

	resp, err := client.Client.Do(req)
	if err != nil {
		logger.Warn(err)
		return nil
	}

	return resp

}
