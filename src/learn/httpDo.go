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
	str := `{
				"Func": "FetchAnalyseData",
				"Param": {
					"applicationName": "test_oxygen",
					"page": 1,
					"pageSize": 30,
					"queryAnalyzeType": 1,
					"queryAnalyzeIds": "15",
					"timeGranularity": "3",
					 "startTime": "2021-04-22",
					"endTime": "2021-04-27",
					"day": 7
				}
			}`
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
