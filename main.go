package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"

	"flag"

	"github.com/gobuffalo/packr"
	"github.com/wecatch/devops-ui/apps"
	"github.com/wecatch/devops-ui/apps/home"
	"github.com/wecatch/devops-ui/db"
	"github.com/wecatch/devops-ui/logshow"
	"github.com/wecatch/devops-ui/utils/consul"
	"github.com/wecatch/devops-ui/utils/gitlab"
	"github.com/wecatch/devops-ui/utils/log"
	"github.com/wecatch/devops-ui/utils/ucloud"
	"github.com/wecatch/devops-ui/utils/upyun"
	"github.com/wecatch/devops-ui/worker"
)

func startServer(
	port int,
	consulAddress string,
	gitlabToken string,
	gitlabBaseURL string,
	databaseAddress string,
	databaseUser string,
	databasePasswd string,
	databaseName string,
	databaseDebug bool,
	ucloudPrivateKey string,
	ucloudPublickKey string,
) {

	//初始化 consul
	consul.ConsulClient = consul.NewConsulClient(consulAddress)
	//初始化 gitlab
	gitlab.GitLabClient = gitlab.NewGitLabClient(gitlabToken, gitlabBaseURL)
	//初始化 database
	db.NewDB(databaseAddress, databaseUser, databasePasswd, databaseName, databaseDebug)
	//初始化 ucloud api
	ucloud.UcloundClient = ucloud.NewUcloundClient(ucloudPrivateKey, ucloudPublickKey, "https://api.ucloud.cn")

	logger := log.Logger("server")
	router := apps.InitRouter()
	box := packr.NewBox("./static")
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(box)))
	// mux 是严格按照路由的顺序进行匹配，这个和我在社区报的一个 bug 类似，就是如果发现有 error 出现，接下来的 route 就不会得到遍历
	// 所以把 not found 放到最后来匹配，防止把 static 给覆盖了
	router.NotFoundHandler = http.HandlerFunc(home.Home)
	handler := handlers.LoggingHandler(os.Stdout, router)
	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", port),
		Handler: handler,
	}
	server.SetKeepAlivesEnabled(false)
	logger.Println("Server start Listening...", port)
	// start run http server
	server.ListenAndServe()
}

func main() {
	serverMode := flag.Bool("server", false, "Start this agent as server")
	workerMode := flag.Bool("worker", false, "Start this agent as worker")
	consulAddress := flag.String("consul_address", "", "Consul API address")
	databaseAddress := flag.String("database_address", "", "Database address")
	databaseUser := flag.String("database_user", "", "Database user")
	databasePasswd := flag.String("database_paswd", "", "Database passwd")
	databaseName := flag.String("database_name", "devops", "Database name")
	databaseDebug := flag.Bool("database_debug", false, "Database mode")
	logServerPort := flag.Int("log_server_port", 8081, "Log server listening port")
	httpServerPort := flag.Int("http_server_port", 8000, "API Server port")
	gitlabToken := flag.String("gitlab_token", "", "Gitlab request token")
	gitlabBaseURL := flag.String("gitlab_base_url", "", "Gitlab request base url")
	logPath := flag.String("log_path", "devops.log", "App log path")
	logLevel := flag.String("log_level", "info", "App log level")
	ucloudPrivateKey := flag.String("ucloud_private_key", "", "ucloud private key")
	ucloudPublickKey := flag.String("ucloud_publick_key", "", "ucloud publick key")

	upyunBucket := flag.String("upyun_bucket", "", "Upyun bucket")
	upyunOp := flag.String("upyun_op", "", "Upyun operator")
	upyunPasswd := flag.String("upyun_passwd", "", "Upyun passwd")
	upyunPrefix := flag.String("upyun_prefix", "", "Upyun prefix")
	flag.Parse()
	//初始化日志
	log.NewLogConf(*logPath, *logLevel)
	upyun.UpyunClient = upyun.NewUpyunClient(*upyunPrefix, *upyunBucket, *upyunOp, *upyunPasswd)
	if *serverMode {
		go logshow.StartWorker(*logServerPort)
		startServer(
			*httpServerPort,
			*consulAddress,
			*gitlabToken,
			*gitlabBaseURL,
			*databaseAddress,
			*databaseUser,
			*databasePasswd,
			*databaseName,
			*databaseDebug,
			*ucloudPrivateKey,
			*ucloudPublickKey,
		)
		*workerMode = false
	}

	if *workerMode {
		worker.StartWorker()
	}
	// fmt.Println(*port)
}
