package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	log "github.com/mgutz/logxi/v1"
	"net"
	"os"
	"strconv"
	"time"
)

import "sample/src/proto"

var addrStr = []string{"127.0.0.1:8001", "127.0.0.1:8002", "127.0.0.1:8003"}
var addrIn, addrOut = addrStr[2], addrStr[0]
var NAME = "node3"
var getMes []string

// interact - функция, содержащая цикл взаимодействия с сервером.
func interact(connIn, connOut *net.TCPConn) {
	defer connOut.Close()
	defer connIn.Close()
	encoder := json.NewEncoder(connOut)
	decoder := json.NewDecoder(connIn)

	go serve(decoder, encoder)

	for {
		// Чтение команды из стандартного потока ввода
		fmt.Printf("command = ")
		var command string
		fmt.Scanln(&command)

		// Отправка запроса.
		switch command {
		case "quit":
			send_request(encoder, "quit", nil)
			return
		case "send":
			var mes proto.Message
			address := ""
			fmt.Printf("Введите сообщение на одной строке:\n")
			reader := bufio.NewReader(os.Stdin)
			mes.Message, _ = reader.ReadString('\n')
			mes.Message = mes.Message[:len(mes.Message)-1]
			fmt.Printf("Сколько адресатов этого сообщения?\n")
			fmt.Scan(&address)
			fmt.Printf("Введите имена адресатов\n")
			n := 0
			for _, err := strconv.Atoi(address); err != nil; {
				fmt.Println("Это не число, введите число, пожалуйста")
				_, err = strconv.Atoi(address)
			}
			n, _ = strconv.Atoi(address)
			for i := 0; i < n; i++ {
				var buf string
				fmt.Scan(&buf)
				mes.Names = append(mes.Names, buf)
			}

			mes.InitName = NAME
			send_request(encoder, "send", &mes)
		case "print":
			if len(getMes) == 0 {
				fmt.Println("no messages")
			} else {
				for _, x := range getMes {
					fmt.Printf("\"%s\"\n", x)
				}
			}
		default:
			fmt.Printf("error: unknown command\n")
			continue
		}
	}
}

func serve(decoder *json.Decoder, encoder *json.Encoder) {
	for {
		var req proto.Request
		if err := decoder.Decode(&req); err != nil {
			fmt.Printf("cannot decode message: %s", err)
			break
		} else {
			//fmt.Printf("received command: %s", req.Command)
			switch req.Command {
			case "send":
				if req.Data == nil {
					fmt.Printf("error: data field is absent in response\n")
				} else {
					var mes proto.Message
					if err := json.Unmarshal(*req.Data, &mes); err != nil {
						fmt.Printf("error: malformed data field in response\n")
					} else {
						if mes.InitName != NAME {
							send_request(encoder, "send", &mes)
						}
						for _, node := range mes.Names {
							if node == NAME {
								getMes = append(getMes, mes.Message)
							}
						}
						//fmt.Printf("***    get message \"%s\"    ***", mes.Message)
					}
				}
			}
		}
	}
}

// send_request - вспомогательная функция для передачи запроса с указанной командой
// и данными. Данные могут быть пустыми (data == nil).
func send_request(encoder *json.Encoder, command string, data interface{}) {
	var raw json.RawMessage
	raw, _ = json.Marshal(data)
	encoder.Encode(&proto.Request{command, &raw})
}

func main() {
	// Работа с командной строкой, в которой может указываться необязательный ключ -addr.

	// Разбор адреса, установка соединения с сервером и
	// запуск цикла взаимодействия с сервером.

	if addrInTCP, err := net.ResolveTCPAddr("tcp", addrIn); err != nil {
		log.Error("address resolution failed", "address", addrStr)
	} else {
		log.Info("resolved TCP address", "address", addrInTCP.String())

		// Инициация слушания сети на заданном адресе.
		if listener, err := net.ListenTCP("tcp", addrInTCP); err != nil {
			log.Error("listening failed", "reason", err)
		} else {

			if addrOutTCP, err := net.ResolveTCPAddr("tcp", addrOut); err != nil {
				fmt.Printf("error: %v\n", err)
			} else {
				for {
					connOut, err := net.DialTCP("tcp", nil, addrOutTCP)
					for err != nil {
						fmt.Printf("error: %v\n", err)
						time.Sleep(5 * time.Second)
						connOut, err = net.DialTCP("tcp", nil, addrOutTCP)
					}
					if connIn, err := listener.AcceptTCP(); err != nil {
						log.Error("cannot accept connection", "reason", err)
					} else {
						log.Info("accepted connection", "address", connOut.RemoteAddr().String())
						interact(connIn, connOut)
					}
				}
			}
		}
	}
}
