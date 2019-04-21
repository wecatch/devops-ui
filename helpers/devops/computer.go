package devops

import (
	"time"

	"github.com/wecatch/devops-ui/db"
	model "github.com/wecatch/devops-ui/models/devops"
	"github.com/wecatch/devops-ui/resources"
	"github.com/wecatch/devops-ui/services"
)

// CreateComputer create computer
func CreateComputer(form resources.ComputerForm) {
	computer := model.Computer{
		CPU:       form.CPU,
		RAM:       form.RAM,
		PrivateIP: form.PrivateIP,
		PublicIP:  form.PublicIP,
	}
	now := time.Now()
	computer.UpdatedAt = now
	computer.CreatedAt = now
	db.Session().Create(&computer)
}

//CreateComputerWithField for insert host
func CreateComputerWithField(hostID, hostName, hostTag string, CPU, RAM uint, privateIP, publicIP string) {
	computer := model.Computer{
		HostID:    hostID,
		HostName:  hostName,
		HostTag:   hostTag,
		CPU:       CPU,
		RAM:       RAM,
		PrivateIP: privateIP,
		PublicIP:  publicIP,
	}
	now := time.Now()
	computer.UpdatedAt = now
	computer.CreatedAt = now
	db.Session().Create(&computer)
}

// UpdateComputer update
func UpdateComputer(id int, form resources.ComputerForm) {
	computer := model.Computer{
		CPU:       form.CPU,
		RAM:       form.RAM,
		PrivateIP: form.PrivateIP,
		PublicIP:  form.PublicIP,
	}
	db.Session().Model(&computer).Where("id=?", id).Updates(computer)
}

// QueryAppComputer according to tag and name query computer list
func QueryAppComputer(appID interface{}, page, limit int, appName string) (result []services.ComputerData) {
	if limit == 0 {
		limit = 50
	}

	var oneAppID = 0
	var appIDS []int

	oneAppID, ok := appID.(int)
	if !ok {
		appIDS, ok = appID.([]int)
	}

	// select app.id as app_id, app.name, app.tag, computer.`cpu`, computer.ram,computer.private_ip, computer.public_ip,computer.host_id,computer_role.register_status
	// from app inner join `computer_role` on app.id = computer_role.app_id inner join computer on `computer_role`.host_id = computer.host_id;
	sql := db.Session().Table("app").Select(
		`app.id as app_id, app.name, app.tag, app.port, computer.cpu, computer.ram,computer.private_ip,
		computer.public_ip,computer.host_id,computer_role.register_status`).Joins(
		`inner join computer_role on app.id = computer_role.app_id inner join computer on computer_role.host_id = computer.host_id`).Order(
		"name desc").Limit(limit).Offset((page - 1) * limit)

	if oneAppID != 0 {
		sql.Where("app.id = ?", appID).Scan(&result)
	} else if len(appIDS) != 0 {
		sql.Where("app.id in (?)", appIDS).Scan(&result)
	} else if appName != "" {
		sql.Where("app.name = ?", appName).Scan(&result)
	} else {
		sql.Scan(&result)
	}
	return
}

// QueryComputer computer list
func QueryComputer(tag, name, searchValue string, page, limit int) []model.Computer {
	if limit == 0 {
		limit = 20
	}
	var ret []model.Computer

	db.Session().Where("host_tag = ?", searchValue).Or("host_name = ?", searchValue).Or("private_ip like ?", "%"+searchValue+"%").Limit(limit).Offset((page - 1) * limit).Find(&ret)
	// db.Session().Where(&model.Computer{HostTag: tag, HostName: name}).Limit(limit).Offset((page - 1) * limit).Find(&ret)
	return ret
}

//QueryOneComputer for one computer according to id
func QueryOneComputer(id int, hostID string) model.Computer {
	computer := model.Computer{}
	db.Session().Where(&model.Computer{HostID: hostID, BaseModel: model.BaseModel{ID: id}}).First(&computer)
	return computer
}

// DeleteComputer for computer delete
func DeleteComputer(id int) {
	db.Session().Delete(&model.Computer{}, "id = ?", id)
}

// DeleteAllComputers
func DeleteAllComputers() {
	db.Session().Delete(&model.Computer{}, "id > ?", 0)
}

// CreateDisk for disk create
func CreateDisk(form resources.DiskForm) {
	disk := model.Disk{
		Size:       form.Size,
		Left:       form.Left,
		ComputerID: form.ComputerID,
	}
	now := time.Now()
	disk.UpdatedAt = now
	disk.CreatedAt = now
	db.Session().Create(&disk)
}

// UpdateDisk for disk update
func UpdateDisk(id int, form resources.DiskForm) {
	disk := model.Disk{
		Size:       form.Size,
		Left:       form.Left,
		ComputerID: form.ComputerID,
	}
	disk.UpdatedAt = time.Now()
	db.Session().Model(&disk).Where("id=?", id).Updates(disk)
}

// DeleteDisk for disk delete
func DeleteDisk(id int) {
	db.Session().Delete(&model.Disk{}, "id = ?", id)
}

// UpdateComputerRole for register status chanee
func UpdateComputerRole(appID int, HostID string, registerStatus int, appName string) {
	if registerStatus != 0 && registerStatus != 1 {
		return
	}
	if appID == 0 || HostID == "" {
		return
	}
	//Model 方法在接受 struct 时只会映射 id 到 where 子句中
	var app model.App
	if appID == 0 {
		// 根据 name 查找 app
		db.Session().Where(&model.App{Name: appName}).Find(&app)
		appID = app.ID
	}

	db.Session().Model(&model.ComputerRole{}).Where(&model.ComputerRole{AppID: appID, HostID: HostID}).Update("register_status", registerStatus)
}

//CreateComputerRole for computer and app role create
func CreateComputerRole(form resources.ComputerRoleForm) {
	now := time.Now()
	for _, appID := range form.AppID {
		for _, hostID := range form.HostID {
			role := model.ComputerRole{
				AppID:  appID,
				HostID: hostID,
			}
			role.CreatedAt = now
			role.UpdatedAt = now
			db.Session().Create(&role)
		}
	}
}

//DeleteComputerRole for computer and app role delete
func DeleteComputerRole(form resources.ComputerRoleForm) {
	db.Session().Delete(&model.ComputerRole{}, "app_id in (?) and host_id in (?)", form.AppID, form.HostID)
}
