package main

import (
	"fmt"
	"net/http"
	"reflect"
	"time"
	"github.com/tomwright/queryparam/v4"
)

func searchUsersHandler(r *http.Request, rw http.ResponseWriter) {
	req := struct {
		UserIDs      []string  `queryparam:"id"`
		TeamIDs      []string  `queryparam:"team-id"`
		MustBeActive bool      `queryparam:"must-be-active"`
		CreatedAfter time.Time `queryparam:"created-after"`
	}{}

	err := queryparam.Parse(r.URL.Query(), &req)
	switch err {
	case nil:
		break
	case queryparam.ErrInvalidBoolValue: // only necessary if the request contains a bool value
		// they have entered a non-bool value.
		// this can be handled this as a 400 or ignored to default to false.
		return
	default:
		// something went wrong when parsing a value.
		// send a 500.
		return
	}

}


func main()  {
	type target struct{
		ID int `intsss:"iii"`
	}
	var ss = &target{
		ID:0,
	}
	targetValue := reflect.ValueOf(ss)
	fmt.Println(targetValue)
	fmt.Println(targetValue.Kind())
	fmt.Println(targetValue.IsNil())

	t1 := targetValue.Elem()
	t2 := t1.Type()

	for i := 0; i < t2.NumField(); i++ {
		fmt.Println(t2.Field(i))
		fmt.Println(t2.Field(i))
	}
	if targetValue.Kind() != reflect.Ptr || targetValue.IsNil() {
	}
}
