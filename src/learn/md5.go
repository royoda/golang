package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
)

func MD5(str string) string {
	md5String := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5String
}

func main() {
	/*currentTime := time.Now()
	  startTime := currentTime.AddDate(0, 0, -1).Format("200601") + " 00:00:00" // 获取昨天的时间
	  endTime := currentTime.AddDate(0, 0, -1).Format("2006-01-02") + " 23:59:59" // 获取昨天的时间
	  fmt.Println(startTime)
	  fmt.Println(endTime)*/
	var ssss = "{\"app_id\":8,\"content\":\"心与科技\"}$bv%goKM&EybE^G6"
	fmt.Printf("%v\n", ssss)
	fmt.Println(MD5(ssss))
	var am = make(map[string]interface{}, 0)
	am["app_id"] = 8
	am["content"] = "心与科技"
	b, _ := json.Marshal(am)
	bstr := string(b) + "$bv%goKM&EybE^G6"
	fmt.Printf("%v\n", bstr)
	fmt.Println(MD5(bstr))

	/*signData := fmt.Sprintf("%d%s%s%s", 10014, "ea0696fe07302fd93158e8539ec124bec957090c", "ad_opt_admin_perform",
	      "register")
	  w := md5.New()
	  io.WriteString(w, signData)
	  sign := fmt.Sprintf("%x", w.Sum(nil))
	  fmt.Println(sign)

	  data := []byte(signData)
	  v := md5.Sum(data)
	  sign1 := fmt.Sprintf("%x", v)
	  fmt.Println(sign1)*/
	set := make(map[string]bool)
	set["1"] = true
	set["2"] = true
	set["3"] = true
	for v := range set {
		fmt.Println(v)
	}
}
