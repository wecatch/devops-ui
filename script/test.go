package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/wecatch/devops-ui/pbmessage"

	"github.com/StabbyCutyou/buffstreams"
	xlog "github.com/apex/log"
	"github.com/apex/log/handlers/text"
	scmd "github.com/go-cmd/cmd"
	"github.com/sirupsen/logrus"
	"github.com/xanzy/go-gitlab"
)

func gitlabMain() {
	client := gitlab.NewClient(nil, "Ay3qMksYxKZjLGXeWsHF")
	client.SetBaseURL("http://git.shequan.com/")
	opt := &gitlab.ListProjectsOptions{
		Sort:       gitlab.String("asc"),
		Membership: gitlab.Bool(true),
	}
	projects, _, err := client.Projects.ListProjects(opt)
	if err != nil {
		fmt.Println(err)
	}
	for _, project := range projects {
		fmt.Println(project.WebURL)
		fmt.Println(project.Name)
		fmt.Println("======")
	}
}

type writer int

func (*writer) Write(p []byte) (int, error) {
	fmt.Println("out", len(p)) //out 16384
	return len(p), nil
}

func main() {
	// cmdFileMain()
	// cmdStreamMain()
	// cmdPipeMain()
	// cmdGocmdMain()
	// logrusLog()
	// var hosts []string
	// var leftHosts []string
	// hosts = append(hosts, "1")
	// leftHosts = append(leftHosts)
	// fmt.Println(hosts)
	// copy(leftHosts, hosts)
	// fmt.Println(leftHosts)
	// selectMain()

	// mainEchoLog()

	var s []string
	for _, h := range []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"} {
		s = append(s, h)
	}
	fmt.Printf("%T", len(s))
	fmt.Println("s ", s)
	fmt.Println("s cap ", cap(s))
	fmt.Println("s length ", len(s))
	ss := s[0:5]
	fmt.Println("ss ", ss)
	fmt.Println("ss cap ", cap(ss))
	fmt.Println("ss length ", len(ss))
	s = s[5:]
	fmt.Println("s ", s)
	fmt.Println("s cap ", cap(s))
	fmt.Println("s length ", len(s))

	ss = s[0:5] // slice 的实际 length 是 j - i, 只要 j <= cap(s)
	fmt.Println("ss ", ss)
	fmt.Println("ss cap ", cap(ss))
	fmt.Println("ss length ", len(ss))

	//panic 实际 s 的 length 是 4
	//s = s[5:]
	s = s[4:]
	fmt.Println("s ", s)
	fmt.Println("s cap ", cap(s))
	fmt.Println("s length ", len(s))

	// fmt.Println(s)
	// ss = s[0:5]
	// fmt.Println(ss)
	// fmt.Println("ss cap ", cap(ss))
	// fmt.Println("ss length ", len(ss))
	// s = s[len(ss):]
	// fmt.Println(s)
	// fmt.Println(s)
	// fmt.Println(s[0:5])
	// fmt.Println("cap ", cap(s))
	// fmt.Println("length ", len(s))
	// // s = s[5:]
	// fmt.Println(s)
	// fmt.Println(s[0:6])
	// fmt.Println("cap ", cap(s))
	// fmt.Println("length ", len(s))

	var y []int
	for i := 0; i < 10; i++ {
		y = append(y, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
	}

}

// StartWorker to start this agent as worker
func cmdFileMain() {
	fmt.Println("start worker")
	cmd := exec.Command("python", "test.py")
	f, _ := os.OpenFile("log.log", os.O_APPEND|os.O_WRONLY, 750)
	cmd.Stdout = f
	cmd.Run()
}

func cmdStreamMain() {
	fmt.Println("start worker")
	cmd := exec.Command("python", "test.py")
	cmd.Stdout = new(writer)
	cmd.Run()
}

func cmdPipeMain() {
	fmt.Println("start worker")
	cmd := exec.Command("python", "test.py")
	stdout, err := cmd.StdoutPipe()
	go func() {
		var buf []byte
		io.ReadAtLeast(stdout, buf, 10)
		fmt.Println(string(buf))
	}()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}

func cmdGocmdMain() {
	cmdOptions := scmd.Options{
		Buffered:  false,
		Streaming: true,
	}
	findCmd := scmd.NewCmdOptions(cmdOptions, "python", "test.py")

	go func() {
		for {
			select {
			case line := <-findCmd.Stdout:
				fmt.Println(line)
			case line := <-findCmd.Stderr:
				fmt.Fprintln(os.Stderr, line)
			}
		}
	}()
	// Run and wait for Cmd to return, discard Status
	<-findCmd.Start()
}

func apexLog() {
	xlog.SetHandler(text.New(os.Stderr))

	ctx := xlog.WithFields(xlog.Fields{
		"file": "something.png",
		"type": "image/png",
		"user": "tobi",
	})

	for range time.Tick(time.Millisecond * 200) {
		ctx.Info("upload")
		ctx.Info("upload complete")
		ctx.Warn("upload retry")
		ctx.WithError(errors.New("unauthorized")).Error("upload failed")
		ctx.Errorf("failed to upload %s", "img.png")
	}
}

func logrusLog() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	// log := logrus.New()
	logrus.Info("hello")
	logrus.Error("hello")
	logrus.Warn("hello")
	logrus.Warning("hello")
	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}

func selectMain() {

	cmdOptions := scmd.Options{
		Buffered:  false,
		Streaming: true,
	}
	findCmd := scmd.NewCmdOptions(cmdOptions, "python", "test.py")
	statusChan := findCmd.Start()

	// go func() {
outfor:
	for {
		fmt.Println("for begin")
		select {
		case status := <-statusChan:
			fmt.Println(status.Complete)
			if status.Complete == true {
				break outfor
			}
		case line := <-findCmd.Stdout:
			fmt.Println(line)
		case line := <-findCmd.Stderr:
			fmt.Println(line)
			// default:
			// 	time.Sleep(time.Second * 2)
			// 	fmt.Println("default")
			// 	break

		}
		fmt.Println("for end")
	}
	// }()

	// time.Sleep(time.Second * 10)
}

func mainEchoLog() {
	cfg := buffstreams.TCPConnConfig{
		MaxMessageSize: 2048,                                                       // You want this to match the MaxMessageSize the server expects for messages on that socket
		Address:        buffstreams.FormatAddress("127.0.0.1", strconv.Itoa(8081)), // Any address with the pattern ip:port. The FormatAddress helper is here for convenience.
	}

	btc, err := buffstreams.DialTCP(&cfg)
	if err != nil {
		fmt.Println(err)
	}

	lm := pbmessage.LogMessage{
		Message:  "i am message",
		DeployId: 10,
		Time:     time.Now().Unix(),
	}
	msgByte, err := proto.Marshal(&lm)
	if err != nil {
		fmt.Println(err)
	}
	btc.Write(msgByte)
}
