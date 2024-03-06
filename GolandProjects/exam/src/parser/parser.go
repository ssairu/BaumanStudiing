package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Obj struct {
	Name    string `json:"name"`
	Gps     string `json:"gps"`
	Address string `json:"address"`
	Tel     string `json:"tel"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {

		response, err := http.Get("http://pstgu.yss.su/iu9/mobiledev/lab4_yandex_map/?x=var05")

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		defer response.Body.Close()

		jsonData, err := ioutil.ReadAll(response.Body)

		var data []Obj

		json.Unmarshal(jsonData, &data)

		fmt.Fprintln(w, "<!DOCTYPE html>\n<html lang=\"en\">\n"+
			"<head>\n<meta charset=\"UTF-8\">\n<title>Home page</title>\n"+
			"</head>\n<body>\n"+
			"\n<p>")
		for i := 0; i < len(data); i++ {
			fmt.Fprintf(w, "<p>[%d] item Name : %s\n <br>", i, data[i].Name)
			fmt.Fprintf(w, ". item Gps : %s\n<br>", data[i].Gps)
			fmt.Fprintf(w, ". item Tel : %s\n<br>", data[i].Tel)
			fmt.Fprintf(w, ". item Address : %s\n<br>\n</p>", data[i].Address)
		}
		fmt.Fprintln(w, "\n</p>\n</body>\n</html>")

	})

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe("185.139.70.64:8002", nil))
}
