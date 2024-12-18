// UserData.go
package main

import (
	"sync"
	"time"
)

type UserData struct {
	Login    string   `json:"login"`
	Password string   `json:"password"`
	Signals  []Signal `json:"signal"`
}

type Signal struct {
	RSRP int       `json:"RSRP"`
	Lat  float64   `json:"lat"`
	Lon  float64   `json:"lon"`
	Time time.Time `json:"time"`
}

var (
	dataMutex sync.Mutex
)
