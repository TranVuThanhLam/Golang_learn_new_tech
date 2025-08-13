package main

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(request *http.Request) bool {
		return true
	},
}

type Client struct {
	Conn *websocket.Conn
}

var (
	clients   = make(map[*Client]bool)
	broadcast = make(chan Coupon)
	mutex     sync.Mutex
)

type Coupon struct {
	
}
