package utils

import (
	"encoding/json"
	"fmt"
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