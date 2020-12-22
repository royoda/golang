package main

import (
	"fmt"
	"net/url"
	"time"
)

type one struct {
	One int64
}
type two struct {
	Two int64
}
func main() {
	str := fmt.Sprintf("TRADE-%d-%d-%d-%s%d",
		1,
		8,
		1003,
		time.Now().Format("20060102150405"),
		time.Now().UnixNano()%10e6,
	)
	fmt.Println(time.Now().Format("20060102150405"))
	fmt.Println(time.Now().UnixNano()%10e6)
	fmt.Println(str)

	const path = "/api/BalanceQuery"
	base, _ := url.Parse("http://baidu.com")
	u, _ := base.Parse(path)
	fmt.Println(u)
	/*fmt.Println("hello golang github")
	s := "         "
	fmt.Printf("%d %q\n", len(s), s)
	t := strings.TrimSpace(s)
	fmt.Printf("%d %q\n", len(t), t)*/
	var o one
	o.One = 1
	fmt.Println(obj(o))
	var t two
	t.Two = 2
	fmt.Println(obj(t))

}
func obj(i interface{}) int64 {
	if o,ok := i.(one);ok {
		fmt.Println("我是你大爷",o.One)
	}
	switch v := i.(type) {
	case one:
		var o = v
		return o.One
	case two:
		var o = v
		return o.Two
	}
	return 0
}
