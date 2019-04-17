package services

import (
	"github.com/wecatch/devops-ui/models/devops"
)

type (
	// ComputerData for api
	ComputerData struct {
		devops.Computer
		AppID          int `json:"app_id,omitempty"`
		RegisterStatus int `json:"register_status"`
		// 应用名称
		Name string `json:"name"`
		// 应用分组
		Tag  string `json:"tag"`
		Port int    `json:"port"`
	}
	// DeployData model for deploy job
	DeployData struct {
		devops.Deploy
		AppName string `json:"app_name"`
		//上线主机
		Hosts []string `json:"hosts"`
		// 更新代码命令
		UpdateCodeCmd string `json:"update_code_cmd"`
		// 重启服务命令
		ReloadServiceCmd string `json:"reload_service_cmd"`
		// 检测服务命令
		CheckServiceCmd string `json:"check_service_cmd"`
		CmdName         string `json:"cmd_name"`
		CmdDir          string `json:"cmd_dir"`
	}
)
