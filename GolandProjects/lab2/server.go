package main

import (
	"github.com/mgutz/logxi/v1"
	"golang.org/x/net/html"
	"golang.org/x/text/encoding/charmap"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

const INDEX_HTML = `
    <!doctype html>
    <html lang="ru">
        <head>
            <meta charset="windows-1251">
            <title>статья из дневника за {{.Date}}</title>
        </head>
        <body>
            {{if .Data}}
				{{.Date}}
				<br/>
				<a href="https://lleo.me/dnevnik/{{.Path}}">{{.Data}}</a>
				<br/>
			{{else}}
                Не удалось загрузить записи из дневника за {{.Path}} !
            {{end}}
        </body>
    </html>
    `

var indexHtml = template.Must(template.New("index").Parse(INDEX_HTML))

type mes struct {
	Date, Data, Path string
}

func getAttr(node *html.Node, key string) string {
	for _, attr := range node.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}

func isElem(node *html.Node, tag string) bool {
	return node != nil && node.Type == html.ElementNode && node.Data == tag
}

func isText(node *html.Node) bool {
	return node != nil && node.Type == html.TextNode
}

func isDiv(node *html.Node, class string) bool {
	return isElem(node, "div") && getAttr(node, "class") == class
}

func search(node *html.Node, path string) *mes {
	if isDiv(node, "header") {
		c := node.LastChild.FirstChild
		date := node.FirstChild
		if isText(c) && isText(date) {

			decoder := charmap.Windows1251.NewDecoder()
			reader := decoder.Reader(strings.NewReader(date.Data))
			d, err := ioutil.ReadAll(reader)
			if err != nil {
				panic(err)
			}

			reader = decoder.Reader(strings.NewReader(c.Data))
			name, err := ioutil.ReadAll(reader)
			if err != nil {
				panic(err)
			}

			return &mes{string(d), string(name), path}
		}

	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if x := search(c, path); x != nil {
			return x
		}
	}
	return nil
}

func downloadNews(date string) *mes {
	log.Info("sending request to https://lleo.me/dnevnik/" + date)
	if response, err := http.Get("https://lleo.me/dnevnik/" + date); err != nil {
		log.Error("request to https://lleo.me/dnevnik/"+date+" failed", "error", err)
	} else {
		defer response.Body.Close()
		status := response.StatusCode
		log.Info("got response from https://lleo.me/dnevnik/"+date, "status", status)
		if status == http.StatusOK {
			if doc, err := html.Parse(response.Body); err != nil {
				log.Error("invalid HTML from https://lleo.me/dnevnik/"+date, "error", err)
			} else {
				log.Info("HTML from https://lleo.me/dnevnik/" + date + " parsed successfully")
				return search(doc, date)
			}
		}
	}
	return &mes{"", "", date}
}

func serveClient(response http.ResponseWriter, request *http.Request) {
	path := request.URL.Path
	year := request.URL.Query().Get("year")
	month := request.URL.Query().Get("month")
	day := request.URL.Query().Get("day")
	log.Info("got request", "Method", request.Method, "Path", path)
	if path != "/" && path != "/index.html" {
		log.Error("invalid path", "Path", path)
		response.WriteHeader(http.StatusNotFound)
	} else if err := indexHtml.Execute(response, downloadNews(year+"/"+month+"/"+day)); err != nil {
		log.Error("HTML creation failed", "error", err)
	} else {
		log.Info("response sent to client successfully")
	}
}

func main() {
	http.HandleFunc("/", serveClient)
	log.Info("starting listener")
	log.Error("listener failed", "error", http.ListenAndServe("127.0.0.1:80", nil))
}
