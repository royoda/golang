package main

import (
	"fmt"
	"github.com/golang-module/carbon"
	"regexp"
)

// 截取字符串
func main() {
	t := carbon.Now()
	t = carbon.Now().Yesterday().StartOfDay()
	// 10分钟后
	t.AddMinutes(10).ToDateTimeString()
	/*ms := map[string]string{
		"AP-JT-SSW28"	:	"ssw@xinyu100.com",
		"AP-YQ-DX3"		:	"HN_daxie3@xinyu100.com",
		"AP-YQ-DX91"	:	"h5_0_3",
		"AP-YQ-HZ3"		:	"HB_haoze3@xinyu100.com",
		"AP-YQ-HZ4"		:	"HB_haoze4@xinyu100.com",
		"AP-YQ-HZ5"		:	"HB_haoze5@xinyu100.com",
		"AP-YQ-KSMZ15"	:	"kashi3@xinyu100.com",
		"AP-YQ-KSMZ27"	:	"kashi2@xinyu100.com",
		"AP-YQ-KSMZ41"	:	"h5_0_9",
		"AP-YQ-KSMZ45"	:	"shuwei@mizhua.com",
		"AP-YQ-SSW28"	:	"h5_0_10",
		"AP-YQ-ST"		:	"HN_shuntong2@xinyu100.com",
		"AP-YQ-ST2"		:	"HN_shuntong2@xinyu100.com",
		"AP-YQ-XY08"	:	"xinyu3@xinyu100.com",
		"AP-YQ-XY21"	:	"ants123456@gmail.com",
		"AP-YQ-XY35"	:	"2220937808@qq.com",
		"AP-YQ-XY46"	:	"2672826400@qq.com",
		"AP-YQ-XY64"	:	"xinyu2@xinyu100.com",
		"AP-YQ-XY77"	:	"xinyu1@xinyu100.com",
		"AP-YQ059"		:	"SSW3",
		"AP-YQ062"		:	"HB_haoze@xinyu100.com",
		"AP-YQ172"		:	"SSW4",
		"AP-YQ199"		:	"SSW5",
		"AP-YQ244"		:	"SSW2",
		"AP-YQ452"		:	"mangguotang@aliyun.com",
		"AP-YQ561"		:	"HN_shuntong1@xinyu100.com",
		"AP-YQ643"		:	"SSW1",
		"AP-YQ679"		:	"HN_shuntong@mizhua.com",
		"AP-YQ755"		:	"HN_shuntong4@xinyu100.com",
		"AP-YQ880"		:	"HN_daxie4@xinyu100.com",
		"AP-YQ882"		:	"HN_shuntong5@xinyu100.com",
		"AP-YQ923"		:	"HN_daxie5@xinyu100.com",
		"AP-YQ980"		:	"HN_shuntong3@xinyu100.com",
		"WX-YQ-DX81"	:	"1555591481",
		"WX-YQ-GZMZ61"	:	"1512939161",
		"WX-YQ101"		:	"1277281101",
		"WX-YQ121"		:	"1573583121",
		"WX-YQ291"		:	"1550647291",
		"WX-YQ381"		:	"1510799381(焦糖头像)",
		"WX-YQ471"		:	"1555282471",
		"WX-YQ851"		:	"1571827851",
		"WX-YQ931"		:	"1528731931",
	}*/
	arr := []string{"ALIPAY-100007-5B73D3BF", "AP-YQ452-1006107-6006B2E1", "AP-JT-SSW28-1000010-5e31242e2"}
	for _, s := range arr {
		reg := `(.*?)-\d{1,8}-.*`
		//match , _:= regexp.MatchString(reg, s)
		r, _ := regexp.Compile(reg)
		s1 := r.FindStringSubmatch(s)
		fmt.Println(s1[1])
	}
}
