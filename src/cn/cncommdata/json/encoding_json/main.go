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

//序列化结构体
func SerializeStruct() {

	server := Server{
		ServerName: "json-for-struct",
		ServerIP:   "127.0.0.1",
		ServerPort: 8080,
	}

	marshal, err := json.Marshal(server)

	if err != nil {
		fmt.Printf("struct serialize error : %v\n", err.Error())
		return
	}

	fmt.Printf("server serialize success: %v\n", string(marshal))

}

//序列化map
func SerializeMap() {

	m := make(map[string]interface{})

	m["ServerName"] = "json-for-map"
	m["ServerIP"] = "192.168.0.100"
	m["ServerPort"] = "8888"

	marshal, err := json.Marshal(m)

	if err != nil {
		fmt.Printf("map serialize error : %v\n", err.Error())
		return
	}

	fmt.Printf("map serialize success: %v\n", string(marshal))

}

func main() {

	//序列化结构体
	SerializeStruct()
	//序列化map
	SerializeMap()
}
