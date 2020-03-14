package controllers

import (
	
    "encoding/json"
)

func autocheck(jsonFile []byte) {

	var j map[string]interface{}
    err := json.Unmarshal(jsonFile, &j)
    if (err != nil) {
        return
    }
    println(j)
}