package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Configitem struct {
	Local_address string `json:"local_address"`
	Local_port    int    `json:"local_port"`
	Method        string `json:"method"`
	Password      string `json:"password"`
	Server        string `json:"server"`
	Server_port   int    `json:"server_port"`
	Timeout       int    `json:"timeout"`
}

type GuiConfig struct {
	Configs []*Configitem `json:"configs"`
	Index   int           `json:"index"`
}

func main() {
	data, err := ioutil.ReadFile("C:\\Users\\bill.zhuang\\Dropbox\\Backup\\gui-config.json")
	check(err)

	var f = GuiConfig{}
	err = json.Unmarshal(data, &f)
	check(err)

	index := len(f.Configs)
	var buffer bytes.Buffer

	buffer.WriteString("{\n\"local_port\": 1081,\n \"server_password\": [\n")
	for k, v := range f.Configs {
		buffer.WriteString(fmt.Sprintf("[\"%v:%v\", \"%v\", \"%v\"]", v.Server, v.Server_port, v.Password, v.Method))
		if k < index-1 {
			buffer.WriteString(",\n")
		} else {
			buffer.WriteString("\n")
		}
	}

	buffer.WriteString("]\n}")

	ioutil.WriteFile("C:\\Users\\bill.zhuang\\Dropbox\\Backup\\config.json", []byte(buffer.String()), 0644)
	fmt.Println("done")
	/*for k, v := range f {
		switch vv := v.(type) {
		case string:
			//fmt.Println(k, "is string", vv)
		case int:
			//fmt.Println(k, "is int", vv)
		case []interface{}:
			//fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	} */

	//fmt.Println(f.Index)
	//fmt.Println(f.Configs[0].Server)

	/*
		item1 := &Configitem{
			Local_address: "eouoeu",
			Local_port:    111,
			Method:        "eoeoue",
			Password:      "ouoeu",
			Server:        "oeuoeu",
			Server_port:   "qoeueo",
			Timeout:       3333,
		}

		config1 := &GuiConfig{
			Index:   1,
			Configs: []*Configitem{item1}}

		fmt.Println(config1.Configs[0].Local_address)

		res2, err := json.Marshal(config1)
		check(err)
		fmt.Println(string(res2))*/
}

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}
