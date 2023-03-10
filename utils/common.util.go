package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
)

//전달된 구조체를 Json string으로 변환하는 함수
func ChangeStruct2JsonStr(x interface{}) string {
	e, err := json.Marshal(x)
	 if err != nil {
        fmt.Println("[ChangeStruct2JsonStr] err = ", err)
        return ""
    }
    return string(e)
}

// PrintJSON converts payload to JSON and prints it
func PrintJSON(payload interface{}) {
	response, _ := json.Marshal(payload)
	fmt.Printf("%s\n", response)
}



//FieldName / Value 조회 함수
func showStruct(data interface{}) {
	v:=reflect.ValueOf(data)
	typeOfS := v.Type()
    for i := 0; i< v.NumField(); i++ {
        fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
    }
}