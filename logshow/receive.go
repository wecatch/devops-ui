package logshow

import (
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/StabbyCutyou/buffstreams"
	"github.com/golang/protobuf/proto"
	"github.com/patrickmn/go-cache"
	"github.com/wecatch/devops-ui/pbmessage"
	"github.com/wecatch/devops-ui/utils/log"
)

var logger = log.MakeLogger()

//LogC log cache queue
var LogC = cache.New(1*time.Minute, 1*time.Minute)

const maxBufferSize = 1000

//LogQueue for deploy job show log 最大 100 个消息的队列
type LogQueue struct {
	sync.Mutex
	logBuffer [maxBufferSize][]byte
	used      int
}

//
func (lq *LogQueue) push(msg []byte) {
	lq.Mutex.Lock()
	defer lq.Mutex.Unlock()
	if len(msg) == 0 {
		return
	}
	if lq.used < maxBufferSize {
		lq.logBuffer[lq.used] = msg
		lq.used++
	}
}

//Pop as many as message
func (lq *LogQueue) Pop() [][]byte {
	lq.Mutex.Lock()
	defer lq.Mutex.Unlock()
	//没有消息
	if lq.used == 0 {
		return nil
	}
	var ret [][]byte
	for i := 0; i < lq.used; i++ {
		ret = append(ret, lq.logBuffer[i])
		lq.logBuffer[i] = []byte("")
	}
	lq.used = 0
	return ret
}

func logCallback(bts []byte) error {
	lm := pbmessage.LogMessage{}
	err := proto.Unmarshal(bts, &lm)
	if err != nil {
		logger.Warn("proto unmarshal err: ", err)
		return err
	}

	jsonData, err := json.Marshal(lm)
	if err != nil {
		logger.Warn("logMessage json encode err: ", err)
		return err
	}

	key := fmt.Sprintf("%d", lm.DeployId)
	retry := 0
	for {
		//最多尝试三次，如果还无法正常添加则丢掉日志消息
		if retry >= 3 {
			break
		}

		value, found := LogC.Get(key)
		if found {
			lq := value.(*LogQueue)
			lq.push(jsonData)
			break
		} else {
			// 每个 deploy 有一个 logQueue
			lq := LogQueue{}
			lq.push(jsonData)
			//如果
			err := LogC.Add(key, &lq, time.Minute*3)
			if err != nil {
				logger.WithField("deploy_id", lm.DeployId).Warn("deploy job push log failed: ", err)
				retry++
				time.Sleep(time.Millisecond * 100)
				continue
			}

			break
		}
	}
	// helper.UpdateDeployLog(int(lm.DeployId), lm.Message)
	return nil
}

//StartWorker log show
func StartWorker(port int) {
	logger = log.Logger("logshow")
	cfg := buffstreams.TCPListenerConfig{
		MaxMessageSize: 1024,
		EnableLogging:  true,
		Address:        buffstreams.FormatAddress("", strconv.Itoa(port)),
		Callback:       logCallback,
	}

	btl, err := buffstreams.ListenTCP(cfg)
	if err != nil {
		logger.Errorln("buffstreams listen ", err)
	}
	if err := btl.StartListening(); err != nil {
		logger.Errorln("buffstreams listening ", err)
	}
}
