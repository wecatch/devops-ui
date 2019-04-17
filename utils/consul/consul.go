package consul

import (
	consulApi "github.com/hashicorp/consul/api"
	"github.com/wecatch/devops-ui/utils/log"
)

var logger = log.MakeLogger()

func init() {
	logger = log.Logger("utils.consul")
}

//ConsulClient for gloabl consul api client
var ConsulClient *consulApi.Client

//NewConsulClient return consul api client
func NewConsulClient(address string) *consulApi.Client {

	config := consulApi.DefaultConfig()
	config.Address = address
	var consulClient, err = consulApi.NewClient(config)
	if err != nil {
		logger.Warn(err)
	}

	return consulClient
}
