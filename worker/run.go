package worker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/wecatch/devops-ui/resources"
	"github.com/wecatch/devops-ui/utils/log"
)

// StartWorker to start this agent as worker
func StartWorker() {
	logger = log.Logger("worker")
	//轮询 job
	go pollJob()
	//执行 job
	go runJob()
	//执行 log
	go reportLog()
	go processRunningJob()
	go processSuccessJob()
	processFailedJob()
}

func pollJob() {
	for {
		resp := devopsClient.Post("/app/deploy/poll", io.Reader(nil))
		logger.Info("poll job running")
		if resp == nil {
			logger.Warn("poll job failed, retry after ", retryTimeout, " seconds")
			time.Sleep(time.Second * retryTimeout)
			continue
		}
		if resp.StatusCode != 200 {
			logger.Warn("poll job response status code is ", resp.StatusCode, " retry after", retryTimeout, "seconds")
			time.Sleep(time.Second * retryTimeout)
			continue
		}

		newJob := NewDeployJob(resp.Body)
		if newJob == nil {
			logger.Warn("polled job is invalid, retry ", retryTimeout, " seconds")
			time.Sleep(time.Second * retryTimeout)
			continue
		}
		logger.Debug(newJob)
		if _, ok := runningJopMap.m[newJob.AppID]; !ok && newJob.DeployID != 0 {
			runningJopMap.AddJob(newJob)
			readyJobChan <- newJob
		}

		time.Sleep(time.Second * retryTimeout)
	}
}

func runJob() {
	for job := range readyJobChan {
		go func(job *DeployJob) {
			// 正在执行的 job
			runningJobChan <- job
			// 保证每个 job 在对 host 进行服务更新时是串行的，即一个批次处理完成，然后更新下一批次
		forLabel:
			for {
				select {
				case currentCmd := <-job.updateCodeCmdChan:
					execCmd(currentCmd)
				case currentCmd := <-job.reloadServiceCmdChan:
					execCmd(currentCmd)
				case currentCmd := <-job.checkServiceCmdChan:
					execCmd(currentCmd)
					//如果服务未更新完毕，下一批停顿 interval 时间
					if !job.IsCheckSuccess() {
						time.Sleep(time.Duration(job.Interval) * time.Second)
					}
				default:
					if job.isTerminated {
						job.Status = FAILSTATUS
						break forLabel //结束 job 的执行
					}
					// 单次更新主机最多
					j := job.Max

					//在此判断保证 j 不能超过 capacity 的大小，防止越界
					//leftHosts 的实际大小如果小于 j，则 j 实际的 capacity 也是应该也比 max 小
					if len(job.LeftHosts) < j {
						j = len(job.LeftHosts)
					}

					// 由于 leftHosts capacity 越来越小，j 可能会超过 capacity
					hosts := job.LeftHosts[0:j]
					if len(hosts) == 0 {
						break forLabel //结束 job 的执行
					}
					//更新 leftHosts 为剩下的主机，此时 job.leftHosts 的 capacity = capacity - (j-0)
					job.LeftHosts = job.LeftHosts[j:]

					rd := RunningCmd{
						cmd:     NewCmd(job.CmdDir, job.CmdName, job.UpdateCodeCmd, job.AppName, hosts, job.BinaryURL),
						cmdType: UPDATECODECMD,
						Status:  NEWSTATUS,
						job:     job,
						hosts:   hosts,
					}
					job.updateCodeCmdChan <- &rd
				}
			}

		}(job)
	}
}

func reportLog() {
	// file, err := os.OpenFile("deploy.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0750)
	// defer file.Close()
	// if err != nil {
	// 	log.Println("open log file failed", err)
	// }
	filelog := &FileLogWriter{}
	httplog := &HTTPLogWriter{}

	for out := range cmdOutChan {
		writerList := []LogWriter{filelog, httplog, out.job.jobLog}
		for _, w := range writerList {
			//同步写 控制打开 http 连接的数量
			w.Write(out)
		}
	}
}

func processRunningJob() {
	for job := range runningJobChan {
		job.Status = DOINGSTATUS
		data := resources.DeployForm{}
		data.ID = job.DeployID
		logger.WithField("deployID", data.ID).Debug("processs running")
		data.Status = resources.DeployJobDoing
		jsonData, _ := json.Marshal(data)
		body := bytes.NewBuffer(jsonData)
		//更新任务状态
		resp := devopsClient.Put("/app/deploy/poll/update", body)
		if resp == nil {
			logger.WithField("deployId", job.DeployID).WithField("status", data.Status).Error("update status failed")
		}
	}
}

func processSuccessJob() {
	for job := range successJobChan {
		job.Status = SUCCESSSTATUS
		//上传日志
		job.jobLog.Close(job)
		//移除正在执行 job
		runningJopMap.RemoveJob(job)
		data := resources.DeployForm{}
		data.ID = job.DeployID
		logger.WithField("deployID", data.ID).Debug("processs success")
		data.Status = resources.DeployJobSuccess
		jsonData, _ := json.Marshal(data)
		body := bytes.NewBuffer(jsonData)
		//更新任务状态
		devopsClient.Put("/app/deploy/poll/update", body)
	}
}

func processFailedJob() {
	for job := range failedJobChan {
		job.Status = FAILSTATUS
		//上传日志
		job.jobLog.Close(job)
		//移除正在执行 job
		runningJopMap.RemoveJob(job)
		data := resources.DeployForm{}
		data.ID = job.DeployID
		logger.WithField("deployID", data.ID).Debug("processs fail")
		data.Status = resources.DeployJobFail
		jsonData, _ := json.Marshal(data)
		body := bytes.NewBuffer(jsonData)
		//更新任务状态
		devopsClient.Put("/app/deploy/poll/update", body)
	}
}

func execCmd(currentCmd *RunningCmd) {
	statusChan := currentCmd.cmd.Start()
	currentCmd.Status = DOINGSTATUS
	job := currentCmd.job

forLabel: // 命令执行的 label
	for {
		select {
		//命令执行状态
		case status := <-statusChan:
			cmdLogger := logger.WithField("app_name", job.AppName).WithField("cmd", status.Cmd).WithField("action", currentCmd.cmdType).WithField("hosts", currentCmd.hosts)
			if status.StartTs > 0 {
				cmdLogger.Debug("Doing")
			}
			// 执行结束
			if status.StopTs > 0 {
				//更新主机状态
				cmdLogger.Debug("Finished")
				var nextCmd RunningCmd
				var nextCmdChan chan *RunningCmd
				switch currentCmd.cmdType {
				case UPDATECODECMD:
					//execute next command reload service
					job.UpdateCodeStatus(currentCmd.hosts)
					nextCmd = RunningCmd{
						cmd:     NewCmd(job.CmdDir, job.CmdName, job.ReloadServiceCmd, job.AppName, currentCmd.hosts, job.BinaryURL),
						cmdType: RELOADSERVICECMD,
						Status:  NEWSTATUS,
						job:     job,
						hosts:   currentCmd.hosts,
					}
					nextCmdChan = job.reloadServiceCmdChan
				case RELOADSERVICECMD:
					job.UpdateReloadStatus(currentCmd.hosts)
					//execute next command check service
					nextCmd = RunningCmd{
						cmd:     NewCmd(job.CmdDir, job.CmdName, job.CheckServiceCmd, job.AppName, currentCmd.hosts, job.BinaryURL),
						cmdType: CHECKSERVICECMD,
						Status:  NEWSTATUS,
						job:     job,
						hosts:   currentCmd.hosts,
					}
					nextCmdChan = job.checkServiceCmdChan
				case CHECKSERVICECMD:
					job.UpdateCheckStatus(currentCmd.hosts)
					if job.IsCheckSuccess() {
						// 如果通过 chan 异步更改可能会导致，回传的日志状态来不及改变，所以在进入异步之前改变
						job.Status = SUCCESSSTATUS
						successJobChan <- job
					}
				}
				//命令正确退出，开始执行下一个命令
				if status.Exit == 0 && status.Complete && status.Error == nil {
					if nextCmdChan != nil {
						nextCmdChan <- &nextCmd
					}
					//命令执行成功
					currentCmd.Status = SUCCESSSTATUS
					cmdOutChan <- &CmdOut{message: "", job: job, cmd: currentCmd}
				} else {
					//报告执行失败的 job, 一旦一个 job 执行失败，后续命令应该及时终止
					//命令执行失败
					currentCmd.Status = FAILSTATUS
					job.isTerminated = true
					// 如果通过 chan 异步更改可能会导致，回传的日志状态来不及改变
					job.Status = FAILSTATUS
					failedJobChan <- job
					cmdOutChan <- &CmdOut{message: fmt.Sprint(status.Error), job: job, cmd: currentCmd}
				}

				break forLabel //退出当前命令的执行，准备执行下一个命令
			}
		//执行输出
		case outlog := <-currentCmd.cmd.Stdout:
			if outlog != "" {
				cmdOutChan <- &CmdOut{message: outlog, job: job, cmd: currentCmd}
			}

		//错误输出
		case errlog := <-currentCmd.cmd.Stderr:
			if errlog != "" {
				cmdOutChan <- &CmdOut{message: errlog, job: job, cmd: currentCmd}
			}
		}
	}
}
