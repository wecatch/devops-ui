package upyun

import (
	"os"
	"testing"
)

func TestUploadFile(t *testing.T) {
	// 'abcdefg', dict({"name": "name", "limit": 1})
	// b6fdb07f319928aec3436fd91ac4a3bbff405073

	upyunBucket := os.Getenv("upyun_bucket")
	upyunOp := os.Getenv("upyun_op")
	passwd := os.Getenv("upyun_passwd")
	prefix := os.Getenv("upyun_prefix")
	if upyunOp == "" && passwd == "" {
		t.SkipNow()
	}
	client := NewUpyunClient(prefix, upyunBucket, upyunOp, passwd)
	file, err := os.OpenFile("/tmp/test.log", os.O_RDWR|os.O_CREATE, 0644)
	if err == nil {
		file.Write([]byte("hello test"))
		file.Close()
	} else {
		t.Log(err)
		t.Fail()
	}
	if client.UploadFile("/tmp/test.log") == "" {
		t.Fail()
	}
}
