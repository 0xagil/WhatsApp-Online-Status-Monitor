package main

import (
	"sync"
	"time"

	"go.mau.fi/whatsmeow"
)

var (
	client        *whatsmeow.Client
	userStatus    = make(map[string]string)
	userNames     = make(map[string]string)
	userStatusLog = make(map[string][]StatusLog)
	mu            sync.RWMutex
)

type StatusLog struct {
	Time   time.Time
	Status string
}

type ContactData struct {
	CurrentStatus string
	Log           []StatusLog
}

type OnlineRange struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

type StatusUpdate struct {
	JID          string        `json:"jid"`
	Username     string        `json:"username"`
	OnlineRanges []OnlineRange `json:"onlineRanges"`
	IsOnline     bool          `json:"isOnline"`
}
