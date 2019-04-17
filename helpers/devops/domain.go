package devops

import (
	"fmt"
	"time"

	"github.com/wecatch/devops-ui/db"
	model "github.com/wecatch/devops-ui/models/devops"
)

// CreateDomain create
func CreateDomain(name, host string, private uint, ip string) {
	domain := model.Domain{
		Name:    name,
		Host:    host,
		Private: private,
		IP:      ip,
	}
	now := time.Now()
	domain.UpdatedAt = now
	domain.CreatedAt = now
	fmt.Println(domain)
	db.DB.Create(&domain)
}

// UpdateDomain update
func UpdateDomain(id int, name, host string, private uint) {
	domain := model.Domain{
		BaseModel: model.BaseModel{
			ID: id,
		},
	}
	db.DB.Model(&domain).Updates(&model.Domain{Name: name, Host: host})
}

// QueryDomain domain list
func QueryDomain(page, limit int, name, host string) []model.Domain {
	if limit == 0 {
		limit = 10
	}
	domains := make([]model.Domain, limit)
	conditon := make([]interface{}, 0)
	query := ""
	if name != "" {
		query += "name = ?"
		conditon = append(conditon, name)
	}
	if host != "" {
		query += "and host = ?"
		conditon = append(conditon, host)
	}
	db.DB.Where(query, conditon...).Limit(limit).Offset((page - 1) * limit).Find(&domains)
	return domains
}

// DeleteDomain domain delete
func DeleteDomain(id int) {
	db.DB.Delete(&model.Domain{}, "id = ?", id)
}
