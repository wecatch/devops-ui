package ucloud

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/wecatch/devops-ui/db"
)

func TestSignParams(t *testing.T) {
	// 'abcdefg', dict({"name": "name", "limit": 1})
	// b6fdb07f319928aec3436fd91ac4a3bbff405073
	client := &APIClient{}

	sign := client.signParams("abcdefg", map[string]interface{}{
		"name":  "name",
		"limit": 1,
	})
	if sign != "b6fdb07f319928aec3436fd91ac4a3bbff405073" {
		t.Fail()
	}
}

func TestFetch(t *testing.T) {
	privateKey := os.Getenv("ucloud_private_key")
	publickKey := os.Getenv("ucloud_publick_key")
	if privateKey == "" && publickKey == "" {
		t.SkipNow()
	}
	client := NewUcloundClient(
		privateKey,
		publickKey,
		"https://api.ucloud.cn")
	data := client.Fetch("", map[string]interface{}{
		"Region": "cn-north-02",
		"Action": "DescribeUHostInstance",
		"Limit":  100,
		"Offset": 0,
	})
	var respData interface{}
	json.Unmarshal(data, &respData)
	value := respData.(map[string]interface{})
	_, ok := value["RetCode"]
	if ok != true {
		t.Fail()
	}
}

func TestSyncAllHosts(t *testing.T) {
	privateKey := os.Getenv("ucloud_private_key")
	publickKey := os.Getenv("ucloud_publick_key")
	if privateKey == "" && publickKey == "" {
		t.SkipNow()
	}

	dbUser := os.Getenv("database_user")
	dbAddress := os.Getenv("database_address")
	dbPasswd := os.Getenv("database_paswd")
	dbName := os.Getenv("database_name")
	dbDebug := false
	if os.Getenv("database_debug") == "true" {
		dbDebug = true
	}
	db.NewDB(dbAddress, dbUser, dbPasswd, dbName, dbDebug)
	client := NewUcloundClient(
		privateKey,
		publickKey,
		"https://api.ucloud.cn")
	client.SyncAllHosts([]string{"cn-north-02", "cn-north-03"})
}

func TestHostExtractIP(t *testing.T) {
	host := UHost{
		IPSet: []interface{}{
			map[string]interface{}{
				"SubnetId": "subnet-jbmmv3",
				"IP":       "10.10.103.230",
				"Mac":      "52:54:00:68:93:4C",
				"VPCId":    "uvnet-rygcoh",
				"Type":     "Private",
			},
			map[string]interface{}{
				"IPId":      "eipev4152",
				"IP":        "106.75.7.104",
				"Bandwidth": 2,
				"Type":      "Bgp",
				"Weight":    50,
			},
		},
	}

	host.ExtractIP()
	if host.PrivateIP != "10.10.103.230" {
		t.Fail()
	}
	if host.PublicIP != "106.75.7.104" {
		t.Fail()
	}
}
