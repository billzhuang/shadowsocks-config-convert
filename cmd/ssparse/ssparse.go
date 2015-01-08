package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	sc "github.com/shadowsocks-configfile/shadowsocksconfigfile"
	"io/ioutil"
)

type Configitem struct {
	Local_address string `json:"local_address"`
	Local_port    int    `json:"local_port"`
	Method        string `json:"method"`
	Password      string `json:"password"`
	Server        string `json:"server"`
	Server_port   int    `json:"server_port"`
	Timeout       int    `json:"timeout,omitempty"`
}

type GuiConfig struct {
	Configs []*Configitem `json:"configs"`
	Index   int           `json:"index"`
}

func main() {
	data, err := ioutil.ReadFile("C:\\Users\\bill.zhuang\\Dropbox\\Backup\\gui-config.json")
	sc.Check(err)

	var f = GuiConfig{}
	err = json.Unmarshal(data, &f)
	sc.Check(err)

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
}
