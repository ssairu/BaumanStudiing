package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/mgutz/logxi/v1"
	"math"
	"net"
	"strconv"
)

import "sample/src/proto"

type Sec struct {
	A1 float64
	B1 float64
	A2 float64
	B2 float64
}

// Client - состояние клиента.
type Client struct {
	logger log.Logger    // Объект для печати логов
	conn   *net.TCPConn  // Объект TCP-соединения
	enc    *json.Encoder // Объект для кодирования и отправки сообщений
	secs   []Sec         // срез отрезков
}

// NewClient - конструктор клиента, принимает в качестве параметра
// объект TCP-соединения.
func NewClient(conn *net.TCPConn) *Client {
	return &Client{
		logger: log.New(fmt.Sprintf("client %s", conn.RemoteAddr().String())),
		conn:   conn,
		enc:    json.NewEncoder(conn),
		secs:   []Sec{}}
}

// serve - метод, в котором реализован цикл взаимодействия с клиентом.
// Подразумевается, что метод serve будет вызаваться в отдельной go-программе.
func (client *Client) serve() {
	defer client.conn.Close()
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

// handleRequest - метод обработки запроса от клиента. Он возвращает true,
// если клиент передал команду "quit" и хочет завершить общение.
func (client *Client) handleRequest(req *proto.Request) bool {
	switch req.Command {
	case "quit":
		client.respond("ok", nil)
		return true
	case "add":
		errorMsg := ""
		if req.Data == nil {
			errorMsg = "data field is absent"
		} else {
			var sec proto.Section
			if err := json.Unmarshal(*req.Data, &sec); err != nil {
				errorMsg = "malformed data field"
			} else {
				v1, ok1 := strconv.Atoi(sec.AX)
				v2, ok2 := strconv.Atoi(sec.AY)
				v3, ok3 := strconv.Atoi(sec.BX)
				v4, ok4 := strconv.Atoi(sec.BY)

				if ok1 != nil && ok2 != nil && ok3 != nil && ok4 != nil {
					errorMsg = "malformed data field"
				} else {
					info := "(" + sec.AX + "; " + sec.AY + ")   (" + sec.BX + "; " + sec.BY + ")"
					client.logger.Info("performing addition", "value", info)
					x := Sec{A1: float64(v1), A2: float64(v2), B1: float64(v3), B2: float64(v4)}
					client.secs = append(client.secs, x)
					if len(client.secs) > 2 {
						client.secs = client.secs[len(client.secs)-2 : len(client.secs)]
					}
				}
			}
		}
		if errorMsg == "" {
			client.respond("ok", nil)
		} else {
			client.logger.Error("addition failed", "reason", errorMsg)
			client.respond("failed", errorMsg)
		}
	case "cross":
		if len(client.secs) < 2 {
			client.logger.Error("calculation failed", "reason", "no args")
			client.respond("failed", "no two sections")
		} else {
			var x, y float64
			EPS := 1e-9
			x1, y1, x2, y2 := client.secs[0].A1, client.secs[0].A2, client.secs[0].B1, client.secs[0].B2
			x3, y3, x4, y4 := client.secs[1].A1, client.secs[1].A2, client.secs[1].B1, client.secs[1].B2

			det := func(a, b, c, d float64) float64 {
				return a*d - b*c
			}

			btw := func(a, b, c float64) bool {
				return math.Min(a, b) <= c+EPS &&
					c <= math.Max(a, b)+EPS
			}

			A1, B1 := y1-y2, x2-x1
			C1 := -A1*x1 - B1*y1
			A2, B2 := y3-y4, x4-x3
			C2 := -A2*x3 - B2*y3

			zn := det(A1, B1, A2, B2)
			mark := false
			if zn != 0 {
				x = -det(C1, B1, C2, B2) / zn
				y = -det(A1, C1, A2, C2) / zn
				mark = btw(x1, x2, x) && btw(y1, y2, y) &&
					btw(x3, x4, x) && btw(y3, y4, y)
			}

			if mark {
				client.respond("result", &proto.Pair{
					First:  strconv.FormatFloat(x, 'f', -1, 64),
					Second: strconv.FormatFloat(y, 'f', -1, 64),
				})
			} else {
				client.respond("result", &proto.Pair{
					First:  "not",
					Second: "crossed",
				})
			}
		}
	case "print":
		x, y := "no section", "no section"

		if len(client.secs) > 0 {
			x1, y1, x2, y2 := client.secs[0].A1, client.secs[0].A2, client.secs[0].B1, client.secs[0].B2
			x = "(" + strconv.FormatFloat(x1, 'f', -1, 64) + ", " +
				strconv.FormatFloat(y1, 'f', -1, 64) + ") -- (" +
				strconv.FormatFloat(x2, 'f', -1, 64) + ", " +
				strconv.FormatFloat(y2, 'f', -1, 64) + ")"
		}
		if len(client.secs) > 1 {
			x3, y3, x4, y4 := client.secs[1].A1, client.secs[1].A2, client.secs[1].B1, client.secs[1].B2
			y = "(" + strconv.FormatFloat(x3, 'f', -1, 64) + ", " +
				strconv.FormatFloat(y3, 'f', -1, 64) + ") -- (" +
				strconv.FormatFloat(x4, 'f', -1, 64) + ", " +
				strconv.FormatFloat(y4, 'f', -1, 64) + ")"
		}

		client.respond("print", &proto.Pair{
			First:  x,
			Second: y,
		})
	default:
		client.logger.Error("unknown command")
		client.respond("failed", "unknown command")
	}
	return false
}

// respond - вспомогательный метод для передачи ответа с указанным статусом
// и данными. Данные могут быть пустыми (data == nil).
func (client *Client) respond(status string, data interface{}) {
	var raw json.RawMessage
	raw, _ = json.Marshal(data)
	client.enc.Encode(&proto.Response{status, &raw})
}

func main() {
	// Работа с командной строкой, в которой может указываться необязательный ключ -addr.
	var addrStr string
	flag.StringVar(&addrStr, "addr", "127.0.0.1:8001", "specify ip address and port")
	flag.Parse()

	// Разбор адреса, строковое представление которого находится в переменной addrStr.
	if addr, err := net.ResolveTCPAddr("tcp", addrStr); err != nil {
		log.Error("address resolution failed", "address", addrStr)
	} else {
		log.Info("resolved TCP address", "address", addr.String())

		// Инициация слушания сети на заданном адресе.
		if listener, err := net.ListenTCP("tcp", addr); err != nil {
			log.Error("listening failed", "reason", err)
		} else {
			// Цикл приёма входящих соединений.
			for {
				if conn, err := listener.AcceptTCP(); err != nil {
					log.Error("cannot accept connection", "reason", err)
				} else {
					log.Info("accepted connection", "address", conn.RemoteAddr().String())

					// Запуск go-программы для обслуживания клиентов.
					go NewClient(conn).serve()
				}
			}
		}
	}
}
