package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/websocket"
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

var SERVER = "localhost:8081"
var PATH = "/"
var TIMESWAIT = 0
var TIMESWAITMAX = 20
var in = bufio.NewReader(os.Stdin)

func getInput(input chan string) {
	result, err := in.ReadString('\n')
	if err != nil {
		log.Println(err)
		return
	}
	input <- result[:len(result)-1]
}

func main() {
	fmt.Println("Connecting to:", SERVER, "at", PATH)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	input := make(chan string, 1)
	go getInput(input)
	URL := url.URL{Scheme: "ws", Host: SERVER, Path: PATH}
	c, _, err := websocket.DefaultDialer.Dial(URL.String(), nil)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	defer c.Close()
	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("ReadMessage() error:", err)
				return
			}
			var stats pingData
			json.Unmarshal(message, &stats)

			switch stats.Status {
			case 1:
				fmt.Printf("\n--- %s ping statistics ---\n", stats.Addr)
				fmt.Printf("%d packets transmitted, %d packets received, %v%% packet loss\n",
					stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
				fmt.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
					stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
			case -1:
				fmt.Println("No such host!")
			case -2:
				fmt.Println("CAN'T RUN pinger on the server")
			}
		}
	}()

	for {
		select {
		case <-time.After(20 * time.Second):
			log.Println("Please give me input!", TIMESWAIT)
			TIMESWAIT++
			if TIMESWAIT > TIMESWAITMAX {
				syscall.Kill(syscall.Getpid(), syscall.SIGINT)
			}
		case <-done:
			return
		case t := <-input:
			err := c.WriteMessage(websocket.TextMessage, []byte(t))
			if err != nil {
				log.Println("Write error:", err)
				return
			}
			TIMESWAIT = 0
			go getInput(input)
		case <-interrupt:
			log.Println("Caught interrupt signal - quitting!")
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))

			if err != nil {
				log.Println("Write close error:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(2 * time.Second):
			}
			return
		}
	}
}
