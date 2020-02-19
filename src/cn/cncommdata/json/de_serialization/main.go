package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	//Tag用法
	ServerName string `json:"name"`
	ServerIP   string `json:"ip"`
	ServerPort int    `json:"port"`
}

//反序列化结构体
func DeSerializeStruct(jsonStr string) {

	s := new(Server)
	err := json.Unmarshal([]byte(jsonStr), &s)
	if err != nil {
		fmt.Printf("deSerialize error : %v\n", err.Error())
		return
	}
	fmt.Printf("struct deserialize success: %v\n", s)

}

//反序列化map
func DeSerializeMap(jsonStr string) {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &m)

	if err != nil {
		fmt.Printf("deSerializeMap error:%v\n", err.Error())
		return
	}

	fmt.Printf("deSerializeMap success : %v\n", m)

}

func main() {

	jsonStr := `{"name":"json-for-struct", "ip":"127.0.0.1", "port":8080}`
	DeSerializeStruct(jsonStr)
	DeSerializeMap(jsonStr)

}
