package main

import (
	"encoding/json"
	"fmt"
	log "github.com/mgutz/logxi/v1"
	"net"
)

import "sample/src/proto"

var NAME = "node1"
var getMes []string

// interact - функция, содержащая цикл взаимодействия с сервером.
func interact(connIn, connOut *net.TCPConn) {
	defer connOut.Close()
	encoder := json.NewEncoder(connOut)
	decoder := json.NewEncoder(connIn)
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
			var n = 0
			fmt.Printf("Введите сообщение на одной строке:\n")
			fmt.Scanln(&mes.Message)
			fmt.Printf("Сколько адресатов этого сообщения?")
			fmt.Scanln(&n)
			fmt.Printf("Введите имена адресатов")
			for i := 0; i < n; i++ {
				var buf string
				fmt.Scanln(&buf)
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
				var pair proto.Pair
				if err := json.Unmarshal(*resp.Data, &pair); err != nil {
					fmt.Printf("error: malformed data field in response\n")
				} else {
					if pair.First == "not" {
						fmt.Printf("result: %s %s\n", pair.First, pair.Second)
					} else {
						fmt.Printf("result: (%s, %s)\n", pair.First, pair.Second)
					}
				}
			}
		case "print":
			if resp.Data == nil {
				fmt.Printf("error: data field is absent in response\n")
			} else {
				var pair proto.Pair
				if err := json.Unmarshal(*resp.Data, &pair); err != nil {
					fmt.Printf("error: malformed data field in response\n")
				} else {
					fmt.Printf("Sections:\n %s\n %s\n", pair.First, pair.Second)
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

// respond - вспомогательный метод для передачи ответа с указанным статусом
// и данными. Данные могут быть пустыми (data == nil).
func (client *Client) respond(status string, data interface{}) {
	var raw json.RawMessage
	raw, _ = json.Marshal(data)
	client.enc.Encode(&proto.Response{status, &raw})
}

// Client - состояние клиента.
type Client struct {
	logger  log.Logger    // Объект для печати логов
	connIn  *net.TCPConn  // Объект TCP-соединения
	connOut *net.TCPConn  // Объект TCP-соединения
	enc     *json.Encoder // Объект для кодирования и отправки сообщений
	dec     *json.Decoder // Объект для декодирования и получения сообщений
	//secs   []Sec         // срез отрезков
}

// NewClient - конструктор клиента, принимает в качестве параметра
// объект TCP-соединения.
func NewClient(connOut, connIn *net.TCPConn) *Client {
	return &Client{
		logger:  log.New(fmt.Sprintf("client %s", connOut.RemoteAddr().String())),
		connOut: connOut,
		connIn:  connIn,
		enc:     json.NewEncoder(connOut),
		//secs:   []Sec{}
	}
}

// serve - метод, в котором реализован цикл взаимодействия с клиентом.
// Подразумевается, что метод serve будет вызаваться в отдельной go-программе.
func (client *Client) serve() {
	defer client.connIn.Close()
	decoder := json.NewDecoder(client.conn)
	for {
		var req proto.Request
		if err := decoder.Decode(&req); err != nil {
			client.logger.Error("cannot decode message", "reason", err)
			break
		} else {
			client.logger.Info("received command", "command", req.Command)
			if client.handleRequest(&req) {
				client.logger.Info("shutting down connection")
				break
			}
		}
	}
}

func main() {
	// Работа с командной строкой, в которой может указываться необязательный ключ -addr.
	var addrStr = []string{"127.0.0.1:8001", "127.0.0.1:8002", "127.0.0.1:8003"}

	// Разбор адреса, установка соединения с сервером и
	// запуск цикла взаимодействия с сервером.

	if addr, err := net.ResolveTCPAddr("tcp", addrStr[1]); err != nil {
		log.Error("address resolution failed", "address", addrStr)
	} else {
		log.Info("resolved TCP address", "address", addr.String())

		// Инициация слушания сети на заданном адресе.
		if listener, err := net.ListenTCP("tcp", addr); err != nil {
			log.Error("listening failed", "reason", err)
		} else {
			if addrIn, err := net.ResolveTCPAddr("tcp", addrStr[0]); err != nil {
				fmt.Printf("error: %v\n", err)
			} else if connIn, err := net.DialTCP("tcp", nil, addrIn); err != nil {
				fmt.Printf("error: %v\n", err)
			} else {
				for {
					if connOut, err := listener.AcceptTCP(); err != nil {
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
