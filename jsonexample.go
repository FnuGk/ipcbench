package ipcbench

import (
	"encoding/json"
)

type Example struct {
	A int    `json:"a,omitempty"`
	S string `json:"s,omitempty"`
}

func serializeExample() {
	example := Example{
		A: 42,
		S: "a string",
	}
	buf, _ := json.Marshal(&example)
	//fmt.Println(string(buf))

	json.Unmarshal(buf, &example)
}
