package worker

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"sync"
)

// DeployJob 部署任务
type DeployJob struct {
	DeployID    int    `json:"id"`
	AppID       int    `json:"app_id"`
	AppName     string `json:"app_name"`
	DeployDir   string `json:"deploy_dir"`
	CommitID    string `json:"commit_id"`
	CommitTag   string `json:"commit_tag"`
	RollbackTag string `json:"rollback_tag"`
	RollbackID  string `json:"rollback_id"`
	BinaryURL   string `json:"binary_url"`
	//一次部署多少主机
	Max int `json:"max"`
	//每个批次主机之间间隔多久
	Interval int `json:"interval"`
	//上线主机
	Hosts     []string `json:"hosts"`
	LeftHosts []string
	// 代码更新状态 0 未更新 1 更新
	codeStatus map[string]int
	// 服务重启状态 0 未重启 1 重启
	reloadStatus map[string]int
	// 服务检查状态 0 未重启 1 重启
	Status           StatusType
	checkStatus      map[string]int
	UpdateCodeCmd    string `json:"update_code_cmd"`
	ReloadServiceCmd string `json:"reload_service_cmd"`
	CheckServiceCmd  string `json:"check_service_cmd"`
	CmdName          string `json:"cmd_name"`
	CmdDir           string `json:"cmd_dir"`
	// 重启服务的 cmd Channel
	reloadServiceCmdChan chan *RunningCmd
	// 更新代码的 cmd Channel
	updateCodeCmdChan chan *RunningCmd
	// 检查服务 cmd Channel
	checkServiceCmdChan chan *RunningCmd
	// 是否需要终端
	isTerminated bool
	sync.Mutex
	jobLog         *JobLogWriter
	jobLogFileName string
}

//UpdateCodeStatus 更新主机代码更新状态
func (t *DeployJob) UpdateCodeStatus(hosts []string) {
	t.Mutex.Lock()
	defer t.Mutex.Unlock()
	for _, host := range hosts {
		if status, ok := t.codeStatus[host]; ok == true && status == 0 {
			t.codeStatus[host] = 1
		}
	}
}

//UpdateReloadStatus 更新主机服务变化状态
func (t *DeployJob) UpdateReloadStatus(hosts []string) {
	t.Mutex.Lock()
	defer t.Mutex.Unlock()
	for _, host := range hosts {
		if status, ok := t.reloadStatus[host]; ok == true && status == 0 {
			t.reloadStatus[host] = 1
		}
	}

}

//UpdateCheckStatus 更新主机服务变化状态
func (t *DeployJob) UpdateCheckStatus(hosts []string) {
	t.Mutex.Lock()
	defer t.Mutex.Unlock()
	for _, host := range hosts {
		if status, ok := t.checkStatus[host]; ok == true && status == 0 {
			t.checkStatus[host] = 1
		}
	}
}

//IsCodeSuccess deployJob 包含的全部主机代码是否都变更成功
func (t *DeployJob) IsCodeSuccess() bool {
	for _, host := range t.Hosts {
		if status, ok := t.codeStatus[host]; ok && status == 0 {
			return false
		}
	}

	return true
}

//IsReloadSuccess deployJob 包含的全部主机服务是否都变更成功
func (t *DeployJob) IsReloadSuccess() bool {
	for _, host := range t.Hosts {
		if status, ok := t.reloadStatus[host]; ok && status == 0 {
			return false
		}
	}

	return true
}

//IsCheckSuccess deployJob 包含的全部主机服务是否都变更成功
func (t *DeployJob) IsCheckSuccess() bool {
	for _, host := range t.Hosts {
		if status, ok := t.checkStatus[host]; ok && status == 0 {
			return false
		}
	}

	return true
}

type responseJSON struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Resp    *DeployJob `json:"resp"`
}

// NewDeployJob 生成新的部署任务
func NewDeployJob(body io.Reader) *DeployJob {
	resp := responseJSON{}
	err := json.NewDecoder(body).Decode(&resp)
	if err != nil {
		logger.Warn("decode job failed", err)
		return nil
	}
	job := resp.Resp
	job.reloadStatus = make(map[string]int)
	job.codeStatus = make(map[string]int)
	job.checkStatus = make(map[string]int)
	job.updateCodeCmdChan = make(chan *RunningCmd, 1)
	job.reloadServiceCmdChan = make(chan *RunningCmd, 1)
	job.checkServiceCmdChan = make(chan *RunningCmd, 1)
	job.Status = NEWSTATUS
	job.jobLogFileName = strings.Join([]string{"/tmp/", job.AppName, fmt.Sprintf("_%d", job.DeployID)}, "")
	job.jobLog = newJobLogWriter(job)
	for _, host := range job.Hosts {
		job.reloadStatus[host] = 0
		job.codeStatus[host] = 0
		job.checkStatus[host] = 0
		job.LeftHosts = append(job.LeftHosts, host)
	}

	return job
}
