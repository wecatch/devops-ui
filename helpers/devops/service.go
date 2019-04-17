package devops

import (
	"fmt"
	"time"

	"github.com/wecatch/devops-ui/db"
	model "github.com/wecatch/devops-ui/models/devops"
	"github.com/wecatch/devops-ui/resources"
)

// CreateService create
func CreateService(form resources.ServiceForm) {
	service := model.Service{
		Name:          form.Name,
		URL:           form.URL,
		Desc:          form.Desc,
		DeployDir:     form.DeployDir,
		RepositoryURL: form.RepositoryURL,
		MonitorURL:    form.MonitorURL,
	}
	now := time.Now()
	service.UpdatedAt = now
	service.CreatedAt = now
	fmt.Println(service)
	db.DB.Create(&service)
}

// UpdateService update
func UpdateService(id int, form resources.ServiceForm) {
	service := model.Service{
		BaseModel: model.BaseModel{
			ID: id,
		},
	}
	db.DB.Model(&service).Updates(&model.Service{
		Name:          form.Name,
		URL:           form.URL,
		Desc:          form.Desc,
		DeployDir:     form.DeployDir,
		RepositoryURL: form.RepositoryURL,
		MonitorURL:    form.MonitorURL,
	})
}

// QueryService service list
func QueryService(page, limit int, name, url string) []model.Service {
	if limit == 0 {
		limit = 10
	}
	ret := make([]model.Service, limit)
	conditon := make([]interface{}, 0)
	query := ""
	if name != "" {
		query += "name = ?"
		conditon = append(conditon, name)
	}
	if url != "" {
		query += "and url = ?"
		conditon = append(conditon, url)
	}
	db.DB.Where(query, conditon...).Limit(limit).Offset((page - 1) * limit).Find(&ret)

	return ret
}

// DeleteService service delete
func DeleteService(id int) {
	db.DB.Delete(&model.Service{}, "id = ?", id)
}
