package main

import (
	"fmt"
	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

/**
    *@Description TODO
    *@Date 2021/5/8 3:22 下午
**/

var (
	OxygenMysqlCon           *gorm.DB
	PlatformStatMysqlConBeta *gorm.DB
	PlatformStatMysqlConAut  *gorm.DB
	ConDataAccCon            *gorm.DB
)
var (
	ClickHouseConInfo      string
	DirectClickhouseSwitch = DirectClickhouseTypeOpen
	ClickHouseConClient    *sqlx.DB
)

type DirectClickhouseType int32

const (
	DirectClickhouseTypeDefault DirectClickhouseType = 0
	DirectClickhouseTypeOpen    DirectClickhouseType = 1
)

func InitClickHouse(host, user, password string, port int) {
	ClickHouseConInfo = fmt.Sprintf("tcp://%s:%d?username=%s&password=%s&read_timeout=60&write_timeout=60&debug=true&compress=true", host, port, user, url.QueryEscape(password))
	var err error
	ClickHouseConClient, err = sqlx.Open("clickhouse", ClickHouseConInfo)
	if err != nil {
		log.Fatalln(err)
	}
}

var (
	RedisCon *redis.Client // 8G
)

func InitRedis(host, pwd string, db int) {
	RedisCon = New(host, pwd, db)
	//log.Info("初始化redis完成")
}

func New(host string, pwd string, db int) *redis.Client {
	configs := strings.Split(host, ",")
	con := redis.NewClient(&redis.Options{
		PoolSize: 100,
		Addr:     strings.TrimSpace(configs[0]),
		Password: pwd,
		DB:       db,
	})
	_, err := con.Ping().Result()
	if err != nil {
		panic("redis初始化失败, err:" + err.Error())
	}
	return con
}

func InitBeta() {
	InitRedis("beta-hn1c-1.redis.rds.aliyuncs.com:6379", "beta-redis-1", 5)
	PlatformStatMysqlConBeta, _ = gorm.Open(mysql.Open("public:Erp3C6R_0uRl44ls@tcp(beta-hn1c-mysql0.mysql.rds.aliyuncs.com:3306)/data_center?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local"), &gorm.Config{})
	InitClickHouse("cc-wz9uj9t11f1b3evcw.ads.rds.aliyuncs.com", "mmorpg", "Erp23swwe44ls", 3306)
}

func InitAut() {
	InitRedis("beta-hn1c-1.redis.rds.aliyuncs.com:6379", "beta-redis-1", 10)
	PlatformStatMysqlConAut, _ = gorm.Open(mysql.Open("public:Erp3C6R_0uRl44ls@tcp(beta-hn1c-mysql0.mysql.rds.aliyuncs.com:3306)/data_center_aut?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local"))
	InitClickHouse("cc-wz9uj9t11f1b3evcw.ads.rds.aliyuncs.com", "mmorpg", "Erp23swwe44ls", 3306)
}

const (
	betaAppName = "test_oxygen"
	preAppName  = "oxygen"

	betaMysql = "data_center"
	preMysql  = "data_center_tmp"
)

func main() {
	InitBeta()
	var eventName = "tracking_evt_set_bind_phone_new"
	// 执行导入事件结构之后需要在心娱门户中赋权，才能点击权限
	//ImportEventStruct(eventName, preAppName, betaMysql, preMysql)
	// push 推送操作,此处需要注意的时按顺序的，有点坑，建议还是从页面点击
	/*client := &http.Client{}
	str := `{"Func":"PushEvent","Obj":"event","Param":{"accessPath":"/vue_admin_event","pageId":6,"accountId":222,"accountName":"罗有达","applicationId":1}}`
	ul := `https://beta.2tianxin.com/aut/platform_stat_admin/platform/`
	PostSendHttpRequest(client, str, ul)*/
	// -- 查询字段信息 行转列 需要推送之后在执行下列方法导入数据
	ImportEventData(eventName, "2020-05-20 00:00:00", "2021-06-03 23:59:59",
		betaMysql, preMysql, betaAppName, preAppName)
}
func ImportEventData(eventName, startTime, endTime, srcMysqlEnv, tarMysqlEnv, srcApp, tarApp string) {
	/* 情况所有数据
	_, err := ClickHouseConClient.Exec("truncate table oxygen." + eventName)
	if err != nil {
		log.Fatalln(err)
		return
	}*/
	// 获取结构题
	var srcResults []map[string]interface{}
	err := PlatformStatMysqlConBeta.Raw("select group_concat(english_name) cl from "+GetEnvCombination(srcMysqlEnv, "platform_stat_event_dimension")+" where "+
		"event_id = (select id from "+GetEnvCombination(srcMysqlEnv, "platform_stat_event")+" where english_name = ?"+
		" and application_id = (select id from "+GetEnvCombination(srcMysqlEnv, "platform_stat_application")+" where english_name = ?)) and english_name not like 'platform_key_%'",
		eventName, srcApp).Scan(&srcResults).Error
	if err != nil {
		log.Fatalln(err)
		return
	}
	if srcResults[0]["cl"] == nil {
		log.Fatalln("源事件不存在")
	}
	var tarResults []map[string]interface{}
	err = PlatformStatMysqlConBeta.Raw("select group_concat(english_name) cl from "+GetEnvCombination(tarMysqlEnv, "platform_stat_event_dimension")+" where "+
		"event_id = (select id from "+GetEnvCombination(tarMysqlEnv, "platform_stat_event")+" where english_name = ?"+
		" and application_id = (select id from "+GetEnvCombination(tarMysqlEnv, "platform_stat_application")+" where english_name = ?)) and english_name not like 'platform_key_%'",
		eventName, tarApp).Scan(&tarResults).Error
	if err != nil {
		log.Fatalln(err)
		return
	}
	if tarResults[0]["cl"] == nil {
		log.Fatalln("目标事件不存在")
	}
	if srcResults[0]["cl"] != tarResults[0]["cl"] {
		log.Fatalln("源表和目标事件字段不一致")
		return
	}
	tx, _ := ClickHouseConClient.Begin()
	_, err = tx.Exec("INSERT INTO " + GetEnvCombination(tarApp, eventName) + "(" + tarResults[0]["cl"].(string) + ") select " + tarResults[0]["cl"].(string) + " " +
		"from " + GetEnvCombination(srcApp, eventName) + " Where create_at BETWEEN '" + startTime + "' and '" + endTime + "'")
	if err != nil {
		log.Fatalln(err)
	}
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
}

func ImportEventStruct(eventName, tarApplicationName, srcMysqlEnv, tarMysqlEnv string) {
	// -- 插入事件信息
	err := PlatformStatMysqlConBeta.Exec(""+
		"insert into "+GetEnvCombination(tarMysqlEnv, "platform_stat_event")+"(application_id, event_group_id, event_type, name,`desc`, "+
		"english_name, event_join_switch,create_at, update_at, event_attr_type) (select application_id, event_group_id, event_type,name,`desc`, "+
		"english_name,event_join_switch,create_at,update_at,event_attr_type from "+GetEnvCombination(srcMysqlEnv, "platform_stat_event")+" where english_name = ?)",
		eventName).Error
	if err != nil {
		log.Fatalln(err)
		return
	}
	// -- 插入维度字段信息
	err = PlatformStatMysqlConBeta.Exec("INSERT INTO "+GetEnvCombination(tarMysqlEnv, "platform_stat_event_dimension")+
		"(event_id, dimension_type, used_switch, unique_switch, name, english_name, param_format,"+
		" create_at, update_at, enum_switch, enum_custom_switch, is_from_default_param, alias)  select 0 as event_id, dimension_type, used_switch,"+
		" unique_switch, name, english_name, param_format, create_at, update_at, enum_switch, enum_custom_switch, is_from_default_param, alias"+
		" from data_center.platform_stat_event_dimension  where event_id = (select id from "+GetEnvCombination(srcMysqlEnv, "platform_stat_event")+" where english_name = ?) "+
		"and english_name not like 'platform_key_%'", eventName).Error
	if err != nil {
		log.Fatalln(err)
		return
	}
	// -- 修改event所属id
	err = PlatformStatMysqlConBeta.Exec("update "+GetEnvCombination(tarMysqlEnv, "platform_stat_event_dimension")+" set event_id = (select id from "+
		"data_center_tmp.platform_stat_event where english_name = ?) "+
		"where event_id = 0", eventName, eventName).Error
	if err != nil {
		log.Fatalln(err)
		return
	}
	// -- 修改连通状态
	err = PlatformStatMysqlConBeta.Exec("update "+GetEnvCombination(tarMysqlEnv, "platform_stat_event")+"  set event_join_switch = 1  "+
		"where id in (select a.* from (select id from "+GetEnvCombination(tarMysqlEnv, "platform_stat_event")+" where english_name = ?) a)", eventName).Error
	if err != nil {
		log.Fatalln(err)
		return
	}
	// -- 修改AppId
	err = PlatformStatMysqlConBeta.Exec("update "+GetEnvCombination(tarMysqlEnv, "platform_stat_event")+"  set application_id = (select id from "+GetEnvCombination(tarMysqlEnv, "platform_stat_application")+
		" where english_name = ?)"+
		"where id in (select a.* from (select id from "+GetEnvCombination(tarMysqlEnv, "platform_stat_event")+" where english_name = ?) a)", tarApplicationName, eventName).Error
	if err != nil {
		log.Fatalln(err)
		return
	}
}

func PostSendHttpRequest(client *http.Client, str, url string) {
	req, err := http.NewRequest("POST", url, strings.NewReader(str))
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("cookie", "xymh-cookie=422f6abe0a9f83411e916a5d4ee28dc1de82978ddd2cc7745772ad7199032a12; table=1314cb2f262a50df47fa1d12f000e954a731fbf21cac0a8c817c8dcc90c03655")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(body))
}

// 删除aut事件数据
func DeleteEventAndDimension(eventName, mysqlEnv string) {
	err := PlatformStatMysqlConBeta.Exec("Delete "+GetEnvCombination(mysqlEnv, "platform_stat_event_dimension")+" Where event_id = ("+
		"select id from "+GetEnvCombination(mysqlEnv, "platform_stat_event")+" Where english_name = ? ) ",
		eventName).Error
	if err != nil {
		log.Fatalln(err)
		return
	}
	err = PlatformStatMysqlConBeta.Exec("Delete "+GetEnvCombination(mysqlEnv, "platform_stat_event")+" Where english_name = ? ",
		eventName).Error
	if err != nil {
		log.Fatalln(err)
		return
	}

}

func GetEnvCombination(env, tableName string) string {
	return env + "." + tableName
}
