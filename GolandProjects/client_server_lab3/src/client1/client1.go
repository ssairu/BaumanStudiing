package main

import (
	"encoding/json"
	"fmt"
	log "github.com/mgutz/logxi/v1"
	"net"
	"strconv"
	"time"
)

import "sample/src/proto"

var addrStr = []string{"127.0.0.1:8001", "127.0.0.1:8002", "127.0.0.1:8003"}
var addrIn, addrOut = addrStr[1], addrStr[2]
var NAME = "node1"
var getMes []string

// interact - функция, содержащая цикл взаимодействия с сервером.
func interact(connIn, connOut *net.TCPConn) {
	defer connOut.Close()
	defer connIn.Close()
	encoder := json.NewEncoder(connOut)
	decoder := json.NewDecoder(connIn)
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
			addres := ""
			fmt.Printf("Введите сообщение на одной строке:\n")
			fmt.Scanln(&mes.Message)
			fmt.Printf("Сколько адресатов этого сообщения?\n")
			fmt.Scan(&addres)
			fmt.Printf("Введите имена адресатов\n")
			n := 0
			for _, err := strconv.Atoi(addres); err != nil; {
				fmt.Println("Это не число, введите число, пожалуйста")
				_, err = strconv.Atoi(addres)
			}
			n, _ = strconv.Atoi(addres)
			for i := 0; i < n; i++ {
				var buf string
				fmt.Scan(&buf)
				mes.Names = append(mes.Names, buf)
			}

			mes.InitName = NAME
			send_request(encoder, "send", &mes)
		case "print":
			if len(getMes) == 0 {
				println("no messages")
			} else {
				for x := range getMes {
					println(x)
				}
			}
		default:
			fmt.Printf("error: unknown command\n")
			continue
		}

		// Получение ответа.
		var resp proto.Response
		if err := decoder.Decode(&resp); err != nil {
			fmt.Printf("error: %v\n", err)
			break
		}

		// Вывод ответа в стандартный поток вывода.
		switch resp.Status {
		case "ok":
			fmt.Printf("ok\n")
		case "failed":
			if resp.Data == nil {
				fmt.Printf("error: data field is absent in response\n")
			} else {
				var errorMsg string
				if err := json.Unmarshal(*resp.Data, &errorMsg); err != nil {
					fmt.Printf("error: malformed data field in response\n")
				} else {
					fmt.Printf("failed: %s\n", errorMsg)
				}
			}
		case "result":
			if resp.Data == nil {
				fmt.Printf("error: data field is absent in response\n")
			} else {
				var mes proto.Message
				if err := json.Unmarshal(*resp.Data, &mes); err != nil {
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
					fmt.Printf("*get message \"%s\"*", mes.Message)
				}
			}
		default:
			fmt.Printf("error: server reports unknown status %q\n", resp.Status)
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
