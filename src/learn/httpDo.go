package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
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

	localinnerapipre  = "http://192.168.0.35:10003/pre/platform_stat_innerapi"
	localinnerapibeta = "http://192.168.0.35:10003/beta/1111/platform_stat_innerapi"
	innerapibeta      = "https://beta.2tianxin.com/beta/1111/platform_stat_innerapi"
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
	str := `{"Func":"HttpGetOxygenUserDetail",
            "Param": {
                "applicationName":  "test_oxygen",
                "applicationId": 1,
				"startTime": "2021-05-27 00:00:00",
				"endTime": "2021-06-03 00:00:00",
                "page": 1,
                "pageSize": 100,
 
"reqData":{"user_id":"1303992"}

                }}
		`
	PostSendHttpRequest(client, str, innerapibeta)
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
	var xx ResData
	json.Unmarshal(body, &xx)
	ioutil.WriteFile("./new1.xlsx", xx.Bytes, 0777)

	fmt.Println(string(body))
}

type ResData struct {
	Code           int32             `json:"code"`
	Msg            string            `json:"msg"`
	Total          int32             `json:"total"`
	Datas          interface{}       `json:"datas"`
	OssCdnHost     string            `json:"ossCdnHost"`
	TrendVideoHost string            `json:"trendVideoHost"`
	ParamExplain   map[string]string `json:"param_explain"`
	TimeOut        int64
	Bytes          []byte
	Filename       string
}
