package main

import (
	"net/http"
	"sync"
)

var varsLock sync.Mutex
var vars map[*http.Request]map[string]interface{}

func OpenVars(r *http.Request) {
	varsLock.Lock()
	if vars == nil {
		vars = map[*http.Request]map[string]interface{}{}
	}
	vars[r] = map[string]interface{}{}
	varsLock.Unlock()
}

func CloseVars(r *http.Request) {
	varsLock.Lock()
	delete(vars, r)
	varsLock.Unlock()
}

func GetVars(r *http.Request, key string) interface{} {
	varsLock.Lock()
	value := vars[r][key]
	varsLock.Unlock()
	return value
}

func SetVars(r *http.Request, key string, value interface{}) {
	varsLock.Lock()
	vars[r][key] = value
	varsLock.Unlock()
}
