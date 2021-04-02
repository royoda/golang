package csv

import (
	"fmt"
	"math"
	"os"
	"testing"
	"time"
)

// 导出Excel
func TestWriterCSV(t *testing.T) {
	type User struct {
		Id   int
		Name string
		Age  int
	}
	header := []string{"Id", "ID", "Name", "姓名", "Age", "年龄"}
	users := make([]*User, 0)
	for i := 0; i < 100000; i++ {
		users = append(users, &User{Id: i, Name: "安迪", Age: 18})
	}
	start := time.Now()
	buf := WriterCSV(header, users)
	file, err := os.OpenFile("test.csv", os.O_CREATE|os.O_RDWR, 777)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	file.Write(buf)
	end := time.Now()
	fmt.Println("execute time=", end.UnixNano()-start.UnixNano())
	//fmt.Println(string(buf))

	lm := make([]map[string]interface{}, 0)
	for i := 0; i < 100000; i++ {
		lm = append(lm, map[string]interface{}{"Id": i, "Name": "安迪", "Age": 20})
	}
	start = time.Now()
	WriterCSV(header, lm)
	end = time.Now()
	fmt.Println("execute time=", end.UnixNano()-start.UnixNano())
	//fmt.Println(string(buf))
}

func TestName(t *testing.T) {
	//2147483647
	fmt.Println(math.MaxInt64)
	fmt.Println(math.MaxInt32)
	fmt.Println(math.MaxInt8)

	fmt.Println(^uint32(0) >> 1)
}
