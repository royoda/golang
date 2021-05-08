package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"strings"
	"time"
)

/**
    *@Description TODO
    *@Date 2021/4/29 11:08 上午
**/

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

func main() {
	redisClient2 := New("beta-hn1c-1.redis.rds.aliyuncs.com:6379", "beta-redis-1", 1)

	c, err := redisClient2.Get("Key:Phone:18789916768111").Int()
	fmt.Println(err)
	if c >= 10 {
		fmt.Println("验证码验证次数已达上限")
		return
	}
	// 点击获取验证码
	redisClient2.Set("Key:Sms:Code:18789916768", "000000", time.Duration(10)*time.Second)
	// 获取验证码用于验证
	smsCode, _ := redisClient2.Get("Key:Sms:Code:18789916768").Bytes()

	var inputCode string
	fmt.Println("请输入验证码：")

	for i := 0; i < 1000; i++ {
		c, _ = redisClient2.Get("Key:Phone:18789916768").Int()
		if c >= 10 {
			fmt.Println("验证码验证次数已达上限")
			break
		}
		// 输入验证码
		fmt.Scanf("%s", &inputCode)
		smsCodeStr := string(smsCode)
		if smsCodeStr != inputCode {
			fmt.Println("验证码验证错误")
			// 这里应该是删除短信验证码缓存的
			redisClient2.Incr("Key:Phone:18789916768")
			redisClient2.Expire("Key:Phone:18789916768", time.Second*24)
			continue
		}
		fmt.Println("验证码验证成功～～～～～噗噗噗")
		redisClient2.Del("Key:Sms:Code:18789916768") // 删除验证码缓存
		redisClient2.Del("Key:Phone:18789916768")    // 删除记录次数缓存
		break
	}
}
