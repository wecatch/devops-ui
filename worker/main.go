package worker

import (
	"fmt"
	"os"
	"sync"

	"github.com/StabbyCutyou/buffstreams"
	"github.com/go-cmd/cmd"
	"github.com/wecatch/devops-ui/utils/http"
	"github.com/wecatch/devops-ui/utils/log"
)

var logCfg = buffstreams.TCPConnConfig{
	MaxMessageSize: 2048,                                                                                 // You want this to match the MaxMessageSize the server expects for messages on that socket
	Address:        buffstreams.FormatAddress(os.Getenv("log_service_ip"), os.Getenv("log_server_port")), // Any address with the pattern ip:port. The FormatAddress helper is here for convenience.
}

var echoLogClient *buffstreams.TCPConn

var logger = log.MakeLogger()

// 准备就绪的 job
var readyJobChan = make(chan *DeployJob, 3)

// 部署失败的 job
var failedJobChan = make(chan *DeployJob, 3)

// 执行成功的 job
var successJobChan = make(chan *DeployJob, 3)

// 正在执行的 job
var runningJobChan = make(chan *DeployJob, 1)

// cmd 输出日志
var cmdOutChan = make(chan *CmdOut, 3)

//同样的 app 一次只能有一个部署任务在执行
var runningJopMap = RunningJobMap{
	m: make(map[int]int),
}

const retryTimeout = 5

var cmdOptions = cmd.Options{
	Buffered:  false,
	Streaming: true,
}

var devopsClient = httputil.RequestClient{
	HOST:        fmt.Sprintf("http://%s:%s", os.Getenv("log_service_ip"), os.Getenv("http_server_port")),
	Prefix:      "/v1",
	ContentType: "application/json",
}

//RunningJobMap 控制同一个 app 同时只能有一个 job 在执行
type RunningJobMap struct {
	sync.Mutex
	m map[int]int
}

//AddJob add deploy
func (t *RunningJobMap) AddJob(job *DeployJob) {
	t.Mutex.Lock()
	defer t.Mutex.Unlock()
	t.m[job.AppID] = 1
}

//RemoveJob remove deploy job
func (t *RunningJobMap) RemoveJob(job *DeployJob) {
	t.Mutex.Lock()
	defer t.Mutex.Unlock()
	delete(t.m, job.AppID)
}

func newEchoLogClient() *buffstreams.TCPConn {
	client, err := buffstreams.DialTCP(&logCfg)
	if err != nil {
		logger.Error("dial log server faied", err)
	}

	return client
}
