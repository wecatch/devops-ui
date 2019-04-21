package devops

import (
	"time"

	"github.com/wecatch/devops-ui/db"
	model "github.com/wecatch/devops-ui/models/devops"
	"github.com/wecatch/devops-ui/resources"
)

//CreateRegion for create region
func CreateRegion(form resources.CloudAvailableRegionForm) {
	region := model.CloudAvailableRegion{
		Region: form.Region,
	}
	now := time.Now()
	region.UpdatedAt = now
	region.CreatedAt = now
	db.Session().Create(&region)
}

//QueryRegion for region query
func QueryRegion() []string {
	var ret []model.CloudAvailableRegion
	db.Session().Where(&model.CloudAvailableRegion{}).Find(&ret)
	var result []string
	for _, r := range ret {
		result = append(result, r.Region)
	}

	return result
}
