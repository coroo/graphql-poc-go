package utils

import (
	"encoding/json"
)

func InterfaceToMap(res interface{}) (map[string]interface{}, error) {
    b, err := json.Marshal(res)
    if err != nil {
        panic(err)
    }
	var m map[string]interface{}
	e := json.Unmarshal(b, &m)
	if e != nil {
	   return nil, e
	}
	
	return m, nil
}