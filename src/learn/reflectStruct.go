package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

/**
    *@Description TODO
    *@Date 2021/3/17 5:37 下午
**/


func main()  {
	/*js := `[{"c": "41", "create": "2021-03-17", "formula": "40"},{"c": "41", "create": "2021-03-17", "formula": "40"}]`
	a := "400"
	// reflect.New works kind of like the built-in function new
	intPtr := reflect.New(reflect.TypeOf(a))
	fmt.Println(intPtr)
	fmt.Println(intPtr.String())
	// Just to prove it
	b := intPtr.Elem().Interface().(string)
	fmt.Println(b)

	res := make([]map[string]interface{}, 0)
	json.Unmarshal([]byte(js),&res)

	for _, v := range res {
		fmt.Println(v["c"])
		fmt.Println(v["create"])
	}*/
	js := `[{"c": "41", "create": "2021-03-17", "formula": "40"},{"c": "41", "create": "2021-03-17", "formula": "40"}]`
	tmpDynamicStuct := []reflect.StructField{}
	tmpStruct := reflect.StructField{
		Name: "CREATE",
		Type: reflect.TypeOf("string"),
	}
	tmpDynamicStuct = append(tmpDynamicStuct, tmpStruct)
	tmpStruct1 := reflect.StructField{
		Name: "C",
		Type: reflect.TypeOf("string"),
	}
	tmpDynamicStuct = append(tmpDynamicStuct, tmpStruct1)
	typ := reflect.StructOf(tmpDynamicStuct)
	tSlice := reflect.MakeSlice(reflect.SliceOf(typ), 0, 0)

	tmp := reflect.New(tSlice.Type()).Elem().Addr().Interface()
	json.Unmarshal([]byte(js), tmp)
	fmt.Println(tmp)


	inputType := reflect.TypeOf(tmp)
	fmt.Println("Type of input is :", inputType.Name())
	inputValue := reflect.ValueOf(tmp)
	fmt.Println("Value of input is :", inputValue)

	st := reflect.ValueOf(tmp)

	if st.Kind() == reflect.Ptr {
		st = st.Elem()
		fmt.Println("Value input points to is :", st)
		len := st.Len()
		fmt.Println("Value input points to len :", len)
		aa := st.Index(0)
		fmt.Println(aa.NumField())
		fmt.Println(aa.Kind())

		if aa.Kind() == reflect.Struct {
			c := aa.FieldByName("C")
			CREATE := aa.FieldByName("CREATE")
			create := aa.FieldByName("create")
			fmt.Println(c)
			fmt.Println(CREATE)
			fmt.Println(create)
			/*for index := 0; index < aa.NumField(); index++ {
				fmt.Printf("%v ", aa.Field(index))
				fmt.Println(aa.Field(index).Type)
			}*/
			// 通过字段名, 找到字段类型信息
			/*if catType, ok := typeOfCat.FieldByName("C"); ok {
				// 从tag中取出需要的tag
				//fmt.Println(catType.Tag.Get("json"), catType.Tag.Get("id"))
				fmt.Println(catType)
			}*/
		}
	}

}