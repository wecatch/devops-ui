package devops

import (
	"strings"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/wecatch/devops-ui/db"
	model "github.com/wecatch/devops-ui/models/devops"
	"github.com/wecatch/devops-ui/resources"
	"github.com/wecatch/devops-ui/services"
	"github.com/wecatch/devops-ui/utils/log"
)

var logger = log.MakeLogger()

func init() {
	logger = log.Logger("helper")
}

// CreateApp create
func CreateApp(form resources.AppForm) {
	service := model.App{
		Name:             form.Name,
		URL:              form.URL,
		InternalURL:      form.InternalURL,
		Desc:             form.Desc,
		DeployDir:        form.DeployDir,
		RepositoryURL:    form.RepositoryURL,
		MonitorURL:       form.MonitorURL,
		RepositoryID:     form.RepositoryID,
		UpdateCodeCmd:    form.UpdateCodeCmd,
		ReloadServiceCmd: form.ReloadServiceCmd,
		CheckServiceCmd:  form.CheckServiceCmd,
		CmdName:          form.CmdName,
		CmdDir:           form.CmdDir,
		Port:             form.Port,
	}
	now := time.Now()
	service.UpdatedAt = now
	service.CreatedAt = now
	db.DB.Create(&service)
}

// UpdateApp update
func UpdateApp(id int, form resources.AppForm) {
	service := model.App{
		BaseModel: model.BaseModel{
			ID: id,
		},
	}
	logger.Debug(form.Port)
	db.DB.Model(&service).Updates(&model.App{
		Name:             form.Name,
		Tag:              form.Tag,
		URL:              form.URL,
		InternalURL:      form.InternalURL,
		Desc:             form.Desc,
		DeployDir:        form.DeployDir,
		RepositoryURL:    form.RepositoryURL,
		RepositoryID:     form.RepositoryID,
		MonitorURL:       form.MonitorURL,
		UpdateCodeCmd:    form.UpdateCodeCmd,
		ReloadServiceCmd: form.ReloadServiceCmd,
		CheckServiceCmd:  form.CheckServiceCmd,
		CmdName:          form.CmdName,
		CmdDir:           form.CmdDir,
		Port:             form.Port,
	})
}

// QueryApp app list
func QueryApp(page, limit int, name, url string) []model.App {
	if limit == 0 {
		limit = 10
	}
	ret := make([]model.App, limit)
	app := model.App{}
	if name != "" {
		app.Name = name
	}

	if url != "" {
		app.URL = url
	}
	db.DB.Where(&app).Limit(limit).Offset((page - 1) * limit).Find(&ret)

	return ret
}

// QueryOneApp for query one app
func QueryOneApp(id int, tag, name string) model.App {
	app := model.App{}
	//  When query with struct, GORM will only query with those fields has non-zero value, that means if your field’s value is 0, '', false or other zero values
	db.DB.Where(&model.App{BaseModel: model.BaseModel{ID: id}, Tag: tag, Name: name}).First(&app)
	return app
}

// DeleteApp app delete
func DeleteApp(id int) {
	db.DB.Delete(&model.App{}, "id = ?", id)
}

// QueryAppDeploy for deploy task
func QueryAppDeploy(appID int, page, limit int) []model.Deploy {
	if limit == 0 {
		limit = 20
	}
	ret := make([]model.Deploy, limit)
	db.DB.Where(&model.Deploy{AppID: appID}).Limit(limit).Order("created_at desc").Offset((page - 1) * limit).Find(&ret)

	return ret
}

// QueryOneDeploy for deploy task
func QueryOneDeploy(deployID int) model.Deploy {
	ret := model.Deploy{}
	db.DB.Where(&model.Deploy{BaseModel: model.BaseModel{ID: deployID}}).First(&ret)
	return ret
}

// QueryTag for app tag
func QueryTag(page, limit int) []model.Tag {

	if limit == 0 {
		limit = 100
	}
	var ret []model.Tag
	db.DB.Where("").Limit(limit).Offset((page - 1) * limit).Find(&ret)

	return ret
}

// CreateTag for app tag create
func CreateTag(form resources.TagForm) {
	tag := model.Tag{
		Name: form.Name,
		Kind: form.Kind,
	}
	now := time.Now()
	tag.UpdatedAt = now
	tag.CreatedAt = now
	db.DB.Create(&tag)
}

// CreateDeploy for app tag create
func CreateDeploy(form resources.DeployForm) model.Deploy {
	tag := model.Deploy{
		AppID:     form.AppID,
		CommitID:  form.CommitID,
		CommitTag: form.CommitTag,
		Interval:  form.Interval,
		Max:       form.Max,
		Desc:      form.Desc,
		Status:    form.Status,
		BinaryURL: form.BinaryURL,
		Hosts:     strings.Join(form.Hosts, ","),
	}
	now := time.Now()
	tag.UpdatedAt = now
	tag.CreatedAt = now
	tag.StartedAt = now
	tag.FinishedAt = now
	db.DB.Create(&tag)

	return tag

}

//QueryNewDeploy find new deploy job
func QueryNewDeploy() services.DeployData {
	var deploy model.Deploy
	var app model.App
	var job services.DeployData
	var hostIds []string
	var hosts []string
	var value string
	db.DB.Where(&model.Deploy{Status: resources.DeployJobNew}).Last(&deploy)
	db.DB.Where("id = ?", deploy.AppID).First(&app)
	rows, _ := db.DB.Table("computer_role").Select("host_id").Where("app_id = ?", app.ID).Rows()

	if deploy.Hosts == "" {
		for rows.Next() {
			err := rows.Scan(&value)
			if err != nil {
				logger.Warn(err)
			}
			hostIds = append(hostIds, value)
		}
		rows, _ = db.DB.Table("computer").Select("private_ip").Where("host_id in (?)", hostIds).Rows()
		for rows.Next() {
			err := rows.Scan(&value)
			if err != nil {
				logger.Warn(err)
			}
			hosts = append(hosts, value)
		}
	} else {
		hosts = strings.Split(deploy.Hosts, ",")
	}

	job.CommitID = deploy.CommitID
	job.CommitTag = deploy.CommitTag
	job.RollbackID = deploy.RollbackID
	job.RollbackTag = deploy.RollbackTag
	job.BinaryURL = deploy.BinaryURL
	job.ID = deploy.ID
	job.AppID = deploy.AppID
	job.AppName = app.Name
	job.Hosts = hosts
	job.Max = deploy.Max
	job.UpdateCodeCmd = app.UpdateCodeCmd
	job.ReloadServiceCmd = app.ReloadServiceCmd
	job.CheckServiceCmd = app.CheckServiceCmd
	job.CmdName = app.CmdName
	job.CmdDir = app.CmdDir

	return job
}

//UpdateDeployStatus update job status
func UpdateDeployStatus(id int, status string) {
	switch status {
	case resources.DeployJobDoing:
		db.DB.Model(&model.Deploy{BaseModel: model.BaseModel{ID: id}}).Update("status", status)
	case resources.DeployJobFail, resources.DeployJobRollback, resources.DeployJobSuccess:
		db.DB.Model(&model.Deploy{BaseModel: model.BaseModel{ID: id}}).Updates(&model.Deploy{Status: status, FinishedAt: time.Now()})
	}
}

//UpdateDeployLog update deploy log
func UpdateDeployLog(id int, message string) {
	logger.Warn(id, message)
	db.DB.Model(&model.Deploy{BaseModel: model.BaseModel{ID: id}}).Update("log", gorm.Expr("concat(log, ?)", message))
}
