package main

import (
    "encoding/json"
    "fmt"
)


type EchartAnalyzeData struct {
    AnalyzeSql     string            `json:"analyze_sql"`
    AnalyzeParam   []interface{}     `json:"analyze_param"`
    FieldMap       map[string]string `json:"field_map"`
}

type EventModel struct {
    EventId   int64  `json:"event_id"`
    EventName string `json:"event_name"`
}
func main() {
    var objs = []EventModel{{EventId: 1, EventName: "11"},{EventId: 2, EventName: "2"}}
    fmt.Println(objs)
    for i, v := range objs {
        v.EventId = 3
        fmt.Println(i, v)
    }
    fmt.Println(objs)
   /* s := make([]string, 10)
    for i := 0; i < 3; i++ {
        s[i] = "tcy"
    }
    fmt.Println(strings.Join(s, "最棒")) //tcy最棒tcy最棒tcy最棒最棒最棒最棒最棒最棒最棒*/
    /*var b strings.Builder
    b.Grow(3)
    b.WriteString("1")
    for i := 0; i < 10; i++ {
        b.WriteString(strconv.Itoa(i))
    }
    fmt.Println(b.String())
    fmt.Println(b.Len())*/
    //ss := "[{\"analyze_param\": [\"2021-03-03 00:00:00\",\"2021-03-10 23:59:59\",\"2021-03-03 00:00:00\",\"2021-03-10 23:59:59\",\"2021-03-03 00:00:00\",\"2021-03-10 23:59:59\"],\"analyze_sql\": \"SELECT level, count() AS c FROM (select user_id, windowFunnel(1800)(time,event_type = 'ad_user_register2', event_type = 'ad_user_login1', event_type = 'ad_user_recharge1') AS level from ((Select * from (Select user_id, create_at as time, 'ad_user_register2' as event_type From ad_user_register2 where create_at between ? and ?  union all Select user_id, create_at as time, 'ad_user_login1' as event_type From ad_user_login1 where create_at between ? and ?  union all Select user_id, create_at as time, 'ad_user_recharge1' as event_type From ad_user_recharge1 where create_at between ? and ? ))) group by user_id)  where level > 0 GROUP BY level ORDER BY level;\",\"field_map\": null,\"event_model_list\": [{\"event_id\": 43,\"event_name\": \"渠道注册事件\"},{\"event_id\": 46,\"event_name\": \"渠道登录事件\"},{\"event_id\": 45,\"event_name\": \"渠道充值事件\"}]}]"
    ss := `[{"analyze_param":["2021-02-26 15:47:07","2021-03-05 15:47:07"],"analyze_sql":"select count(distinct ad_channel) ad_channelsum, %select_group_by_time_dimension, ad_channel ad_channelext from advertising_data2  where create_at between ? and ?  group by %group_by_time_dimension,ad_channel","field_map":{"ad_channel":"渠道","ad_channelsum":"投放数据2渠道的去重数","create":"日"}}]`

    tmpEchartAnalyzeData := EchartAnalyzeData{}
    fmt.Println(ss)
    err := json.Unmarshal([]byte(ss), &tmpEchartAnalyzeData)
    fmt.Println(err)
    s := "1111"
    fmt.Println(&s)

    a := []byte(s)
    fmt.Printf("%p", a)
}