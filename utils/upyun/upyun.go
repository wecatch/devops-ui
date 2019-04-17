package upyun

import (
	"fmt"
	"os"

	"github.com/upyun/go-sdk/upyun"
	"github.com/wecatch/devops-ui/utils/log"
)

var logger = log.MakeLogger()

func init() {
	logger = log.Logger("utils.upyun")
}

//APIClient for upyun client
type APIClient struct {
	client *upyun.UpYun
	prefix string
}

//UpyunClient for global use
var UpyunClient *APIClient

//NewUpyunClient for making new upyun api client
func NewUpyunClient(prefix, bucket, op, passwd string) *APIClient {
	up := upyun.NewUpYun(&upyun.UpYunConfig{
		Bucket:   bucket,
		Operator: op,
		Password: passwd,
	})

	return &APIClient{
		client: up,
		prefix: prefix,
	}
}

//UploadFile for path
func (upClient *APIClient) UploadFile(filepath string) string {
	f, err := os.Stat(filepath)
	if err != nil {
		logger.Warn(filepath + " error " + fmt.Sprint(err))
		fmt.Println(err)
		return ""
	}

	fileURL := upClient.prefix + f.Name()
	upClient.client.Put(&upyun.PutObjectConfig{
		Path:      fileURL,
		LocalPath: filepath,
	})

	return fileURL
}
