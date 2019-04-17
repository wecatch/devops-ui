package resources

import (
	"github.com/wecatch/devops-ui/models/devops"
)

// form model for form submit
type (
	// DomainForm for json data
	DomainForm struct {
		devops.Domain
	}

	// ComputerForm for json data
	ComputerForm struct {
		devops.Computer
	}

	// DiskForm for json data
	DiskForm struct {
		devops.Disk
	}

	// ServiceForm for json data
	ServiceForm struct {
		devops.Service
	}

	// AppForm for json data
	AppForm struct {
		devops.App
	}

	// TagForm for json data
	TagForm struct {
		devops.Tag
	}

	// ConsulServiceForm for consul register service
	ConsulServiceForm struct {
		HostID    string `json:"host_id"`
		AppID     int    `json:"app_id"`
		Name      string `json:"name"`
		Port      int    `json:"port"`
		PrivateIP string `json:"private_ip"`
		Tag       string `json:"tag"`
		ServiceID string `json:"service_id"`
	}

	//ComputerRoleForm for computer create role
	ComputerRoleForm struct {
		AppID  []int    `json:"app_id"`
		HostID []string `json:"host_id"`
	}

	//DeployForm for deploy create
	DeployForm struct {
		devops.Deploy
		Hosts []string `json:"hosts"`
	}

	CloudAvailableRegionForm struct {
		devops.CloudAvailableRegion
	}
)

const (
	//DeployJobDoing deploy doing status
	DeployJobDoing = "doing"
	//DeployJobFail deploy fail status
	DeployJobFail = "fail"
	//DeployJobSuccess deploy success status
	DeployJobSuccess = "success"
	//DeployJobRollback deploy rollback status
	DeployJobRollback = "rollback"
	//DeployJobNew deploy new status
	DeployJobNew = "new"
)
