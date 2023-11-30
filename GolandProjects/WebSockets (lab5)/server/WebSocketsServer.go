package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-ping/ping"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

type pingData struct {
	Addr        string        `json:"addr"`
	PacketsSent int           `json:"packets-sent"`
	PacketsRecv int           `json:"packets-recv"`
	PacketLoss  float64       `json:"packet-loss"`
	MinRtt      time.Duration `json:"min-rtt"`
	AvgRtt      time.Duration `json:"avg-rtt"`
	MaxRtt      time.Duration `json:"max-rtt"`
	StdDevRtt   time.Duration `json:"std-dev-rtt"`
	Status      int           `json:"status"`
}

func main() {

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		websocket, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Websocket Connected!")
		listen(websocket)
	})
	http.ListenAndServe("localhost:8081", nil)
}

func pingering(ip string) pingData {
	pinger, err := ping.NewPinger(ip)
	res := pingData{ip, 0, 0, 0,
		0, 0, 0, 0, -1}
	if err != nil {
		return res
	}
	pinger.Count = 3
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		res.Status = -2
		return res
	}
	stats := pinger.Statistics()

	res = pingData{stats.Addr, stats.PacketsSent, stats.PacketsRecv,
		stats.PacketLoss, stats.MinRtt, stats.AvgRtt, stats.MaxRtt,
		stats.StdDevRtt, 1}

	return res
}

func listen(conn *websocket.Conn) {
	for {
		// read a message
		messageType, messageContent, err := conn.ReadMessage()

		if err != nil {
			log.Println(err)
			return
		}

		// print out that message
		fmt.Println(string(messageContent))
		res := pingering(string(messageContent))

		fmt.Printf("Result is: %s", res)

		b, err := json.Marshal(res)
		if err != nil {
			log.Println(err)
			return
		}

		if err := conn.WriteMessage(messageType, b); err != nil {
			log.Println(err)
			return
		}
	}
}
