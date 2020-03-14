package controllers

import (
	
    "encoding/json"
)
//Autocheck asasd
func Autocheck(jsonFile []byte) {

	var j map[string]interface{}
    err := json.Unmarshal(jsonFile, &j)
    if (err != nil) {
        return
    }
    println(j)
}