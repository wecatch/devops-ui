package worker

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/wecatch/devops-ui/pbmessage"
	// "github.com/wecatch/devops-ui/utils/http"
	"github.com/wecatch/devops-ui/utils/upyun"
)

//LogWriter job execute log wirter
type LogWriter interface {
	Write(*CmdOut) (int, error)
}

//CmdOut cmd execute out
type CmdOut struct {
	job     *DeployJob
	message string
	cmd     *RunningCmd
}

func (c *CmdOut) newLogMessage() *pbmessage.LogMessage {
	lm := pbmessage.LogMessage{
		Message:      c.message,
		DeployId:     int32(c.job.DeployID),
		Time:         time.Now().Unix(),
		CmdType:      string(c.cmd.cmdType),
		CmdStatus:    string(c.cmd.Status),
		DeployStatus: string(c.job.Status),
		Hosts:        strings.Join(c.cmd.hosts, ""),
	}

	return &lm
}

//FileLogWriter file log support
type FileLogWriter struct {
	file *os.File
}

// Write for file writer
func (w *FileLogWriter) Write(out *CmdOut) (int, error) {
	job := out.job
	logger.WithField("appName", job.AppName).WithField("deployID", job.DeployID).WithField("hosts", out.cmd.hosts).Info(out.message)
	return len(out.message), nil
}

//HTTPLogWriter http log support
type HTTPLogWriter struct {
}

// Write for http writer
func (w *HTTPLogWriter) Write(out *CmdOut) (int, error) {

	outlogger := logger.WithField("deploy_id", out.job.DeployID)

	if echoLogClient == nil {
		echoLogClient = newEchoLogClient()
	}

	if echoLogClient == nil {
		return 0, errors.New("echo log failed")
	}

	msgByte, err := proto.Marshal(out.newLogMessage())
	if err != nil {
		outlogger.Warn("message log marshal failed", err)
	}

	n, err := echoLogClient.Write(msgByte)
	if err != nil {
		outlogger.Warn("write log failed ", err)
		// 关闭旧的连接，无论是否有错
		echoLogClient.Close()
		// 创建新的连接
		echoLogClient = newEchoLogClient()
	} else {
		outlogger.Debug("write http log message", out.message)
	}
	return n, err
}

//JobLogWriter for single deploy job
type JobLogWriter struct {
	file *os.File
}

// Write for JobLogWriter
func (w *JobLogWriter) Write(out *CmdOut) (int, error) {
	outlogger := logger.WithField("deploy_id", out.job.DeployID)
	n, err := w.file.Write([]byte(out.message + "\n"))
	if err != nil {
		outlogger.Warn("write file log failed", err)
		n = 0
	}

	return n, err
}

//Close for JobLogWriter file
func (w *JobLogWriter) Close(job *DeployJob) {
	go func(job *DeployJob) {
		time.Sleep(time.Second * 10)
		job.jobLog.file.Close()
		upyun.UpyunClient.UploadFile(job.jobLogFileName)
		os.Remove(job.jobLogFileName)
	}(job)
}

func newJobLogWriter(job *DeployJob) *JobLogWriter {
	outlogger := logger.WithField("deploy_id", job.DeployID)
	logWriter := JobLogWriter{}
	var err error
	logWriter.file, err = os.OpenFile(job.jobLogFileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		outlogger.Warn("create deploy log file ", err)
		return nil
	}
	return &logWriter
}
