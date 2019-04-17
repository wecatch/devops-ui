package worker

import (
	"fmt"
	"strings"

	"github.com/go-cmd/cmd"
)

// CmdType cmd 类型
type CmdType string

const (
	//UPDATECODECMD 变更代码 cmd
	UPDATECODECMD CmdType = "update"
	//RELOADSERVICECMD 重启服务 cmd
	RELOADSERVICECMD CmdType = "reload"
	//CHECKSERVICECMD 检查服务 cmd
	CHECKSERVICECMD CmdType = "check"
)

//StatusType job status
type StatusType string

const (
	//NEWSTATUS new status
	NEWSTATUS StatusType = "new"
	//DOINGSTATUS doing status
	DOINGSTATUS StatusType = "doing"
	//SUCCESSSTATUS success status
	SUCCESSSTATUS StatusType = "success"
	//FAILSTATUS fail status
	FAILSTATUS StatusType = "fail"
)

//RunningCmd 正在执行 cmd
type RunningCmd struct {
	cmd     *cmd.Cmd
	cmdType CmdType
	// cmdStatus cmd.Status
	job    *DeployJob
	hosts  []string
	Status StatusType
}

//NewCmd for new cmd statement
func NewCmd(cmdDir, cmdName, cmdStatement string, appName string, hosts []string, fileName string) *cmd.Cmd {
	cmdName = strings.Trim(cmdName, " ")
	statement := ""
	if strings.Count(cmdStatement, "%s") == 2 {
		if strings.Contains(cmdStatement, "file_name") {
			statement = fmt.Sprintf(cmdStatement, strings.Join(hosts, ","), fileName)
		} else {
			statement = fmt.Sprintf(cmdStatement, appName, strings.Join(hosts, ","))
		}
	} else if strings.Count(cmdStatement, "%s") == 1 {
		statement = fmt.Sprintf(cmdStatement, strings.Join(hosts, ","))
	} else {
		statement = cmdStatement
	}

	if !strings.HasPrefix(cmdName, "make") && !strings.HasPrefix(cmdName, "python") && !strings.HasPrefix(cmdName, "fab") {
		jobCmd := cmd.NewCmdOptions(cmdOptions, "echo", fmt.Sprintf("'%s not allowed'", cmdName))
		jobCmd.Dir = cmdDir
		return jobCmd
	}

	args := strings.Split(statement, " ")
	newArgs := make([]string, 0)
	for _, stat := range args {
		if strings.Trim(stat, " ") != "" {
			newArgs = append(newArgs, strings.Trim(stat, " "))
		}
	}
	// fmt.Println(newArgs)
	// fmt.Println(cmdName)

	jobCmd := cmd.NewCmdOptions(cmdOptions, cmdName, newArgs...)
	jobCmd.Dir = cmdDir
	return jobCmd
}
