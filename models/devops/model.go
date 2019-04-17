package devops

import (
	"bytes"
	"database/sql/driver"
	"errors"
	"time"
)

// JSON for mysql json field
type JSON []byte

// Value value
func (j JSON) Value() (driver.Value, error) {
	if j.IsNull() {
		return nil, nil
	}
	return string(j), nil
}

//Scan value
func (j *JSON) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		return errors.New("Invalid Scan Source")
	}
	*j = append((*j)[0:0], s...)
	return nil
}

// MarshalJSON value
func (j JSON) MarshalJSON() ([]byte, error) {
	if j == nil {
		return []byte("null"), nil
	}
	return j, nil
}

//UnmarshalJSON value
func (j *JSON) UnmarshalJSON(data []byte) error {
	if j == nil {
		return errors.New("null point exception")
	}
	*j = append((*j)[0:0], data...)
	return nil
}

//IsNull for json
func (j JSON) IsNull() bool {
	return len(j) == 0 || string(j) == "null"
}

// Equals for json
func (j JSON) Equals(j1 JSON) bool {
	return bytes.Equal([]byte(j), []byte(j1))
}

// BaseModel for all devops model
type BaseModel struct {
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;not null" json:"updated_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;not null" json:"created_at,omitempty"`
	ID        int       `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id,omitempty"`
}

// Domain model 域名 顶级或二级域名
type Domain struct {
	BaseModel
	Name    string `gorm:"column:name;type:varchar(256)" json:"name"`
	Host    string `gorm:"column:host;type:varchar(256)" json:"host"`
	Private uint   `gorm:"column:private;type:tinyint;not null;default:0" json:"private"`
	IP      string `gorm:"column:ip;type:varchar(256);" json:"ip"`
}

// TableName for domain in sql
func (Domain) TableName() string {
	return "domain"
}

// Computer 机器
type Computer struct {
	BaseModel
	CPU uint `gorm:"column:cpu;type:int(11);default:0" json:"cpu"`
	RAM uint `gorm:"column:ram;type:int(11);default:0" json:"ram"`
	// 云平台资源 id
	HostID string `gorm:"column:host_id;type:varchar(32):default:''" json:"host_id"`
	// 云平台 tag
	HostTag string `gorm:"column:host_tag;type:varchar(32):default:''" json:"host_tag"`
	// 云平台 name
	HostName  string `gorm:"column:host_name;type:varchar(32):default:''" json:"host_name"`
	PrivateIP string `gorm:"column:private_ip;type:varchar(32):default:''" json:"private_ip"`
	PublicIP  string `gorm:"column:public_ip;type:varchar(32):default:''" json:"public_ip"`
}

// ComputerRole 主机角色和app通过 tag 和 name 进行关联
type ComputerRole struct {
	BaseModel
	// 应用 id
	AppID  int    `gorm:"column:app_id;type:int(11);default:0" json:"app_id,omitempty"`
	HostID string `gorm:"column:host_id;type:varchar(32);default:''" json:"host_id"`
	// Tag    string `gorm:"column:tag;type:varchar(32):default:''" json:"tag"`
	// Name   string `gorm:"column:name;type:varchar(32):default:''" json:"name"`
	//0 未注册 1 已注册
	RegisterStatus int `gorm:"column:register_status;type:varchar(32):default:''" json:"register_status"`
}

//Tag 标签用以区分主机，应用分组提供的基础功能
type Tag struct {
	BaseModel
	Name string `gorm:"column:name;type:varchar(32):default:''" json:"name"`
	// 服务分类 base 和 business
	Kind string `gorm:"column:kind;type:varchar(32):default:''" json:"kind"`
}

// Disk 磁盘
type Disk struct {
	BaseModel
	// 容量
	Size int `gorm:"column:size;type:int(11);default:0" json:"size"`
	// 剩余容量
	Left int `gorm:"column:left;type:int(11):default:0" json:"left"`
	// 机器 id
	ComputerID int `gorm:"column:computer_id;type:int(11)" json:"computer_id"`
}

// Service 服务是具有单一职责的一个完整的应用 app，一般对应一个地址，可以外部地址或内部地址，地址可以是一个域名或一个路径。
// 服务包括基础服务和应用服务
// 基础服务 提供应用服务基础功能的服务，比如数据库服务、消息队列服务、缓存服务等等
// 应用服务 执行特定功能的服务，比如认证鉴权、商品列表、订单交易等，应用服务一般是对内服务
type Service struct {
	BaseModel
	// 服务名称
	Name string `gorm:"column:name;type:varchar(256)" json:"name"`
	// 服务地址(对内或对外)
	URL string `gorm:"column:url;type:varchar(256)" json:"url"`
	// 服务功能描述
	Desc string `gorm:"column:desc;type:varchar(512)" json:"desc"`
	// 是否核心链路
	IsImportant uint `gorm:"column:is_important;type:tinyint(11):default:0" json:"is_important"`
	// 仓库地址
	RepositoryURL string `gorm:"column:repository_url;type:varchar(256)" json:"repository_url"`
	// 部署目录
	DeployDir string `gorm:"column:deploy_dir;type:varchar(256)" json:"deploy_dir"`
	// 服务监控地址
	MonitorURL string `gorm:"column:monitor_url;type:varchar(256)" json:"monitor_url"`
}

// App for application 应用是具有完整业务逻辑的一个或一组服务
// 应用一般是对外的服务，比如合作方的一个页面就属于一个应用，为某个客户端直接提供服务的也属于一个应用
type App struct {
	BaseModel
	// 应用名称
	Name string `gorm:"column:name;type:varchar(32)" json:"name"`
	// 应用分组
	Tag string `gorm:"column:tag;type:varchar(32)" json:"tag"`
	// 应用调用地址 对外
	URL string `gorm:"column:url;type:varchar(256)" json:"url"`
	// 应用调用地址对内
	InternalURL string `gorm:"column:internal_url;type:varchar(256)" json:"internal_url"`
	// 应用调用端口(对内或对外)
	Port int `gorm:"column:port;type:int(11)" json:"port"`
	// 应用功能描述
	Desc string `gorm:"column:desc;type:varchar(512)" json:"desc"`
	// 是否核心链路
	IsImportant uint `gorm:"column:is_important;type:tinyint(11):default:0" json:"is_important"`
	// 仓库地址
	RepositoryURL string `gorm:"column:repository_url;type:varchar(256)" json:"repository_url"`
	// 仓库 id gitlab 使用仓库 id 进行数据筛选
	RepositoryID int `gorm:"column:repository_id;type:int(11)" json:"repository_id"`
	// 部署目录
	DeployDir string `gorm:"column:deploy_dir;type:varchar(256)" json:"deploy_dir"`
	// 应用监控地址
	MonitorURL string `gorm:"column:monitor_url;type:varchar(256)" json:"monitor_url"`
	// 更新代码命令
	UpdateCodeCmd string `gorm:"column:update_code_cmd;type:varchar(512)" json:"update_code_cmd"`
	// 重启服务命令
	ReloadServiceCmd string `gorm:"column:reload_service_cmd;type:varchar(512)" json:"reload_service_cmd"`
	// 检测服务命令
	CheckServiceCmd string `gorm:"column:check_service_cmd;type:varchar(512)" json:"check_service_cmd"`
	// 命令名称
	CmdName string `gorm:"column:cmd_name;type:varchar(512)" json:"cmd_name"`
	// 命令目录
	CmdDir string `gorm:"column:cmd_dir;type:varchar(512)" json:"cmd_dir"`
}

// Deploy task for deploy
type Deploy struct {
	BaseModel
	// 应用id
	AppID int `gorm:"column:app_id;type:int(11)" json:"app_id,omitempty"`
	// 上线commit 号
	CommitID string `gorm:"column:commit_id;type:varchar(64)" json:"commit_id"`
	// 描述
	Desc string `gorm:"column:desc;type:varchar(512)" json:"desc"`
	// 上线标签
	CommitTag string `gorm:"column:commit_tag;type:varchar(64)" json:"commit_tag,omitempty"`
	// 回滚标签
	RollbackTag string `gorm:"column:rollback_tag;type:varchar(64)" json:"rollback_tag,omitempty"`
	// 回滚 commit 号
	RollbackID string `gorm:"column:rollback_id;type:varchar(64)" json:"rollback_id,omitempty"`
	// 一次部署多少台
	Max int `gorm:"column:max;type:int(11);default:1" json:"max"`
	// 每次更新主机之后等待多久
	Interval int `gorm:"column:interval;type:int(11);default:0" json:"interval"`
	// 上线状态 new,success,fail,doing,rollback
	Status     string    `gorm:"column:status;type:varchar(12)" json:"status"`
	StartedAt  time.Time `gorm:"column:started_at;type:datetime;not null" json:"started_at"`
	FinishedAt time.Time `gorm:"column:finished_at;type:datetime;not null" json:"finished_at"`
	// 单次部署的主机列表
	Hosts string `gorm:"column:hosts;type:varchar(256)" json:"hosts"`
	// 二进制包
	BinaryURL string `gorm:"column:binary_url;type:varchar(256)" json:"binary_url"`
}

//CloudAvailableRegion for cloud region
type CloudAvailableRegion struct {
	BaseModel
	Region string `gorm:"column:region;type:varchar(64)" json:"region"`
}
