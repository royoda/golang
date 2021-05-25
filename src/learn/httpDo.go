package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

/**
    *@Description 用于测试http请求接口
    *@Date 2021/3/25 4:46 下午
**/
const (
	key          = "$bv%goKM&EybE^G6"
	beta         = "https://beta.2tianxin.com/beta/1111/platform_stat_innerapi"
	innerapiBeta = "https://beta.2tianxin.com/beta/1111/platform_stat_innerapi"
	localbeta    = "http://192.168.0.35:10003/beta/1111/platform_int/platform/api/"
	localpre     = "http://192.168.0.35:10003/pre/platform_int/platform/api/"
	pro          = "https://pub-stat.2tianxin.com/proxycommon/platform_int/platform/api/"
	pre          = "https://beta.2tianxin.com/pre/platform_int/platform/api/"
	localaut     = "http://192.168.0.35:10003/aut/platform_int/platform/api/"
	aut          = "https://beta.2tianxin.com/aut/platform_int/platform/api/"
)

//https://pub-stat.2tianxin.com/platform_stat/third_callback  地址
//https://pub-stat.2tianxin.com/proxycommon

func Md5Byte(inputByte []byte) string {
	m := md5.New()
	m.Write(inputByte)
	return hex.EncodeToString(m.Sum(nil))
}

func main() {
	client := &http.Client{}
	str := `{"Func":"GetOneEventAnalyzeDetail",
            "Param": {
                "applicationName":  "test_oxygen",
                "eventName":  "tracking_evt_consortia_member3",
                "groupByArr": [],
                "sumItemMap": {},
                "applicationId": 1,

                "page": 1,
                "pageSize": 10000,
                "timeGroupByDimension": "0",
                 "whereMap": {}
                }}
		`
	PostSendHttpRequest(client, str, localbeta)
}

func PostSendHttpRequest(client *http.Client, str, url string) {
	sign := Md5Byte([]byte(str + key))
	req, err := http.NewRequest("POST", url, strings.NewReader(str))
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("Security-Key", key)
	req.Header.Set("Md5-Key-Sign", sign)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(body))
}
