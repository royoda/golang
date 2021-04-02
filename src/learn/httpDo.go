package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

/**
    *@Description TODO
    *@Date 2021/3/25 4:46 下午
**/
const key = "$bv%goKM&EybE^G6"

func MD5Sum(str string) string {
	md5String := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5String
}

const beta = "https://beta.2tianxin.com/beta/1111/platform_int/platform/api/"
const pre = "https://beta.2tianxin.com/pre/platform_int/platform/api/"
const localbeta = "http://127.0.0.1:10003/beta/1111/platform_int/platform/api/"
const localpre = "http://192.168.0.35:10003/pre/platform_int/platform/api/"

func Md5Byte(inputByte []byte) string {
	m := md5.New()
	m.Write(inputByte)
	return hex.EncodeToString(m.Sum(nil))
}

func main() {
/*	sign := Md5Byte([]byte(`{"Func": "FetchAnalyseData", "Param": {"applicationName": "oxygen", "page": 1, "pageSize": 30, "queryAnalyzeIds": "6", "queryAnalyzeType": 1, "timeGranularity": "3", "startTime": "2021-03-12 14:47:36", "endTime": "2021-03-30 00:00:00"}}$bv%goKM&EybE^G6`))
	fmt.Println(sign)0c683c8e318b1a9d4a8f0aa5cddee0d6*/
	client := &http.Client{}
	/*str := `{ "Func":"FetchAnalyseData",
        "Param": {
        "applicationName": "oxygen",
        "page": 1,
        "pageSize": 30,
        "queryAnalyzeIds": "6",
        "queryAnalyzeType": 1,
        "timeGranularity": "3",
        "startTime": "2021-03-12 14:47:36",
        "endTime": "2021-03-19 14:47:36"
    }
      }`*/
	str :=`{"Func":"GetOneEventAnalyzeDetail","Param":{"eventName":"advertising_data","applicationName":"oxygen","pageSize":30,"page":1,"startTime":"2021-03-01 00:00:00","endTime":"2021-03-06 23:59:59","timeGroupByDimension":"","groupByArr":["creative_name","creative_id","third_api_type"],"sumItemMap":{"share":"","follow":"","like":"","show":""},"whereMap":{"follow":"","share":"","creative_id":"","creative_name":"","third_api_type":"","like":"","show":"","ad_id":"","platform_key_str_0":"","ad_name":"","ad_channel":"","platform_key_int_0":"","platform_key_sum_int_0":"","cost":"","video_id":"","campaign_id":"","campaign_name":"","comment":""}}}`
	//str := `{"Func": "FetchAnalyseData", "Param": {"applicationName": "test_oxygen", "page": 1, "pageSize": 30, "queryAnalyzeIds": "8", "queryAnalyzeType": 1, "timeGranularity": "3", "startTime": "2021-03-23 00:00:00", "endTime": "2021-03-31 00:00:00"}}`
	sign := Md5Byte([]byte(str+key))
	fmt.Println(sign)
	req, err := http.NewRequest("POST", pre, strings.NewReader(str))
	if err != nil {
		// handle error
	}
	req.Header.Set("Security-Key", key)
	req.Header.Set("Md5-Key-Sign", sign)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println(string(body))
}