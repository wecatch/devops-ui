package ucloud

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"

	helper "github.com/wecatch/devops-ui/helpers/devops"
	"github.com/wecatch/devops-ui/utils/http"
	"github.com/wecatch/devops-ui/utils/log"
)

var logger = log.MakeLogger()

func init() {
	logger = log.Logger("utils.ucloud")
}

//UHost for uclound host
type UHost struct {
	CPU uint `json:"cpu"`
	RAM uint `json:"Memory"`
	// 云平台资源 id
	HostID string `json:"UHostId"`
	// 云平台 tag
	HostTag string `json:"Tag"`
	// 云平台 name
	HostName  string      `json:"Name"`
	IPSet     interface{} `json:"IPSet"`
	PublicIP  string
	PrivateIP string
}

//ExtractIP for uhost
func (host *UHost) ExtractIP() {
	value, ok := host.IPSet.([]interface{})
	if ok {
		for _, ip := range value {
			networkInfo, ok := ip.(map[string]interface{})
			if ok {
				if networkInfo["Type"].(string) == "Private" {
					host.PrivateIP = networkInfo["IP"].(string)
				} else {
					host.PublicIP = networkInfo["IP"].(string)
				}
			}
		}
	}
}

//ResponseHost for uclound api
type ResponseHost struct {
	RetCode  int      `json:"RetCode"`
	UHostSet []*UHost `json:"UHostSet"`
}

//APIClient ucloud api client
type APIClient struct {
	httpClient *httputil.RequestClient
	privateKey string
	publicKey  string
}

// UcloundClient for uclound api client
var UcloundClient *APIClient

//NewUcloundClient for uclound client
func NewUcloundClient(privateKey, publicKey, host string) *APIClient {
	client := &APIClient{
		privateKey: privateKey,
		publicKey:  publicKey,
	}

	client.httpClient = &httputil.RequestClient{
		HOST:        host,
		ContentType: "application/json",
	}
	return client
}

// signParams 签名请求参数
func (client *APIClient) signParams(privateKey string, params map[string]interface{}) string {
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}

	paramsData := ""
	sort.Strings(keys)
	for _, k := range keys {
		value := params[k]
		switch value := value.(type) {
		case string:
			paramsData = paramsData + k + value
		default:
			paramsData = paramsData + k + fmt.Sprint(value)
		}
	}

	paramsData += privateKey
	h := sha1.New()
	h.Write([]byte(paramsData))
	return fmt.Sprintf("%x", h.Sum(nil))
}

//Fetch fetch data from ucloud
func (client *APIClient) Fetch(uri string, params map[string]interface{}) []byte {
	if params != nil {
		params["PublicKey"] = client.publicKey
		params["Signature"] = client.signParams(client.privateKey, params)
	}

	resp := client.httpClient.Get(uri, params)
	data, _ := ioutil.ReadAll(resp.Body)
	return data
}

//SyncAllHosts sync host from uclound
func (client *APIClient) SyncAllHosts(regions []string) {
	limit := 1000000
	// delete old computers
	helper.DeleteAllComputers()
	for _, reg := range regions {
		params := map[string]interface{}{
			"Region": reg,
			"Action": "DescribeUHostInstance",
			"Limit":  limit,
			"Offset": 0,
		}
		for {
			data := client.Fetch("", params)
			var resp ResponseHost
			err := json.Unmarshal(data, &resp)
			if err != nil {
				logger.Warn(err)
				break
			}

			for _, host := range resp.UHostSet {
				host.ExtractIP()
				helper.CreateComputerWithField(
					host.HostID,
					host.HostName,
					host.HostTag,
					host.CPU,
					host.RAM,
					host.PrivateIP,
					host.PublicIP,
				)
			}
			if len(resp.UHostSet) < limit {
				break
			} else {
				params["Offset"] = params["Offset"].(int) + limit
			}
		}
	}
}
